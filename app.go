package app

import (
	"fmt"
	"sort"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/decentrandom/decentrandom/x/rand"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	// tmtypes "github.com/tendermint/tendermint/types"
)

const (
	appName = "rand"
)

type randApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	keyMain          *sdk.KVStoreKey
	keyAccount       *sdk.KVStoreKey
	keyRounds        *sdk.KVStoreKey
	keyFeeCollection *sdk.KVStoreKey
	keyStaking       *sdk.KVStoreKey
	tkeyStaking      *sdk.TransientStoreKey
	keySlashing      *sdk.KVStoreKey
	keyDistr         *sdk.KVStoreKey
	tkeyDistr        *sdk.TransientStoreKey
	keyParams        *sdk.KVStoreKey
	tkeyParams       *sdk.TransientStoreKey

	accountKeeper       auth.AccountKeeper
	bankKeeper          bank.Keeper
	stakingKeeper       staking.Keeper
	slashingKeeper      slashing.Keeper
	distrKeeper         distr.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	paramsKeeper        params.Keeper
	randKeeper          rand.Keeper
}

// NewRandApp - 앱 생성
func NewRandApp(logger log.Logger, db dbm.DB) *randApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &randApp{
		BaseApp: bApp,
		cdc:     cdc,

		keyMain:          sdk.NewKVStoreKey("main"),
		keyAccount:       sdk.NewKVStoreKey("acc"),
		keyRounds:        sdk.NewKVStoreKey("rounds"),
		keyFeeCollection: sdk.NewKVStoreKey("fee_collection"),
		keyParams:        sdk.NewKVStoreKey("params"),
		keyStaking:       sdk.NewKVStoreKey("staking"),
		keySlashing:      sdk.NewKVStoreKey("slashsing"),
		tkeyStaking:      sdk.NewTransientStoreKey("transient_stake"),
		keyDistr:         sdk.NewKVStoreKey(distr.StoreKey),
		tkeyDistr:        sdk.NewTransientStoreKey(distr.TStoreKey),

		tkeyParams: sdk.NewTransientStoreKey("transient_params"),
	}

	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams)

	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		app.paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace,
	)

	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(cdc, app.keyFeeCollection)

	stakingKeeper := staking.NewKeeper(
		app.cdc,
		app.keyStaking, app.tkeyStaking,
		app.bankKeeper, app.paramsKeeper.Subspace(staking.DefaultParamspace),
		staking.DefaultCodespace,
	)

	app.stakingKeeper = *stakingKeeper.SetHooks(NewStakingHooks(app.distrKeeper.Hooks(), app.slashingKeeper.Hooks()))

	app.slashingKeeper = slashing.NewKeeper(
		app.cdc,
		app.keySlashing,
		app.stakingKeeper.GetValidatorSet(),
		app.paramsKeeper.Subspace(slashing.DefaultParamspace),
		slashing.DefaultCodespace,
	)

	app.randKeeper = rand.NewKeeper(
		app.bankKeeper,
		app.keyRounds,
		app.cdc,
	)

	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))

	app.Router().
		AddRoute(bank.RouterKey, bank.NewHandler(app.bankKeeper)).
		AddRoute(staking.RouterKey, staking.NewHandler(app.stakingKeeper)).
		AddRoute(distr.RouterKey, distr.NewHandler(app.distrKeeper)).
		AddRoute(slashing.RouterKey, slashing.NewHandler(app.slashingKeeper)).
		AddRoute("rand", rand.NewHandler(app.randKeeper))

	app.QueryRouter().
		AddRoute(distr.QuerierRoute, distr.NewQuerier(app.distrKeeper)).
		AddRoute(slashing.QuerierRoute, slashing.NewQuerier(app.slashingKeeper, app.cdc)).
		AddRoute(staking.QuerierRoute, staking.NewQuerier(app.stakingKeeper, app.cdc)).
		AddRoute("rand", rand.NewQuerier(app.randKeeper))

	app.SetInitChainer(app.initChainer)

	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keyStaking,
		app.keyDistr,
		app.keySlashing,
		app.keyRounds,
		app.keyFeeCollection,
		app.keyParams,
		app.tkeyStaking,
		app.tkeyDistr,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

// initFromGenesisState -
func (app *randApp) initFromGenesisState(ctx sdk.Context, genesisState GenesisState) []abci.ValidatorUpdate {
	genesisState.Sanitize()

	// load the accounts
	for _, gacc := range genesisState.Accounts {
		acc := gacc.ToAccount()
		acc = app.accountKeeper.NewAccount(ctx, acc) // set account number
		app.accountKeeper.SetAccount(ctx, acc)
	}

	// initialize distribution (must happen before staking)
	distr.InitGenesis(ctx, app.distrKeeper, genesisState.DistrData)

	// load the initial staking information
	validators, err := staking.InitGenesis(ctx, app.stakingKeeper, genesisState.StakingData)
	if err != nil {
		panic(err) // TODO find a way to do this w/o panics
	}

	// initialize module-specific stores
	auth.InitGenesis(ctx, app.accountKeeper, app.feeCollectionKeeper, genesisState.AuthData)
	bank.InitGenesis(ctx, app.bankKeeper, genesisState.BankData)
	slashing.InitGenesis(ctx, app.slashingKeeper, genesisState.SlashingData, genesisState.StakingData.Validators.ToSDKValidators())

	// validate genesis state
	if err := RandValidateGenesisState(genesisState); err != nil {
		panic(err) // TODO find a way to do this w/o panics
	}

	if len(genesisState.GenTxs) > 0 {
		for _, genTx := range genesisState.GenTxs {
			var tx auth.StdTx
			err = app.cdc.UnmarshalJSON(genTx, &tx)
			if err != nil {
				panic(err)
			}

			bz := app.cdc.MustMarshalBinaryLengthPrefixed(tx)
			res := app.BaseApp.DeliverTx(bz)
			if !res.IsOK() {
				panic(res.Log)
			}
		}

		validators = app.stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	}
	return validators
}

// initChainer -
func (app *randApp) initChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	stateJSON := req.AppStateBytes

	var genesisState GenesisState
	err := app.cdc.UnmarshalJSON(stateJSON, &genesisState)
	if err != nil {
		panic(err)
	}

	validators := app.initFromGenesisState(ctx, genesisState)

	// sanity check
	if len(req.Validators) > 0 {
		if len(req.Validators) != len(validators) {
			panic(fmt.Errorf("len(RequestInitChain.Validators) != len(validators) (%d != %d)",
				len(req.Validators), len(validators)))
		}
		sort.Sort(abci.ValidatorUpdates(req.Validators))
		sort.Sort(abci.ValidatorUpdates(validators))
		for i, val := range validators {
			if !val.Equal(req.Validators[i]) {
				panic(fmt.Errorf("validators[%d] != req.Validators[%d] ", i, i))
			}
		}
	}

	app.assertRuntimeInvariants()

	return abci.ResponseInitChain{}
}

// MakeCodec - Amino를 사용하기 위한 codec 생성
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	rand.RegisterCodec(cdc)
	staking.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

// BeginBlocker -
func (app *randApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {

	// distribute rewards for the previous block
	distr.BeginBlocker(ctx, req, app.distrKeeper)

	// 감액 처리
	tags := slashing.BeginBlocker(ctx, req, app.slashingKeeper)

	return abci.ResponseBeginBlock{
		Tags: tags.ToKVPairs(),
	}
}

// EndBlocker -
func (app *randApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	validatorUpdates, tags := staking.EndBlocker(ctx, app.stakingKeeper)

	return abci.ResponseEndBlock{
		ValidatorUpdates: validatorUpdates,
		Tags:             tags,
	}
}

var _ sdk.StakingHooks = StakingHooks{}

// StakingHooks -
type StakingHooks struct {
	dh distr.Hooks
	sh slashing.Hooks
}

// NewStakingHooks -
func NewStakingHooks(dh distr.Hooks, sh slashing.Hooks) StakingHooks {
	return StakingHooks{dh, sh}
}

// AfterValidatorCreated -
func (h StakingHooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	h.dh.AfterValidatorCreated(ctx, valAddr)
	h.sh.AfterValidatorCreated(ctx, valAddr)
}

// BeforeValidatorModified -
func (h StakingHooks) BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) {
	h.dh.BeforeValidatorModified(ctx, valAddr)
	h.sh.BeforeValidatorModified(ctx, valAddr)
}

// AfterValidatorRemoved -
func (h StakingHooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.AfterValidatorRemoved(ctx, consAddr, valAddr)
	h.sh.AfterValidatorRemoved(ctx, consAddr, valAddr)
}

// AfterValidatorBonded -
func (h StakingHooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.AfterValidatorBonded(ctx, consAddr, valAddr)
	h.sh.AfterValidatorBonded(ctx, consAddr, valAddr)
}

// AfterValidatorBeginUnbonding -
func (h StakingHooks) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
	h.sh.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
}

// BeforeDelegationCreated -
func (h StakingHooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.BeforeDelegationCreated(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationCreated(ctx, delAddr, valAddr)
}

// BeforeDelegationSharesModified -
func (h StakingHooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.BeforeDelegationSharesModified(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationSharesModified(ctx, delAddr, valAddr)
}

// BeforeDelegationRemoved -
func (h StakingHooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.BeforeDelegationRemoved(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationRemoved(ctx, delAddr, valAddr)
}

// AfterDelegationModified -
func (h StakingHooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.AfterDelegationModified(ctx, delAddr, valAddr)
	h.sh.AfterDelegationModified(ctx, delAddr, valAddr)
}

// BeforeValidatorSlashed -
func (h StakingHooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) {
	h.dh.BeforeValidatorSlashed(ctx, valAddr, fraction)
	h.sh.BeforeValidatorSlashed(ctx, valAddr, fraction)
}

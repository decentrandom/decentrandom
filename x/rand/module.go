package rand

import (
	"encoding/json"
	//"fmt"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/decentrandom/decentrandom/x/rand/client/cli"
	"github.com/decentrandom/decentrandom/x/rand/client/rest"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic -
type AppModuleBasic struct{}

// Name -
func (AppModuleBasic) Name() string {
	return ModuleName
}

// RegisterCodec -
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// DefaultGenesis -
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
	//return nil
}

// ValidateGenesis -
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState

	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	// Once json successfully marshalled, passes along to genesis.go
	return ValidateGenesis(data)
}

// RegisterRESTRoutes -
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr, StoreKey)
}

// GetQueryCmd -
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(StoreKey, cdc)
}

// GetTxCmd -
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(StoreKey, cdc)
}

// AppModule -
type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

// NewAppModule -
func NewAppModule(k Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

// Name -
func (AppModule) Name() string {
	return ModuleName
}

// RegisterInvariants -
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// Route -
func (am AppModule) Route() string {
	return RouterKey
}

// NewHandler -
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

// QuerierRoute -
func (am AppModule) QuerierRoute() string {
	return ModuleName
}

// NewQuerierHandler -
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

// BeginBlock -
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock -
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// InitGenesis -
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	//InitGenesis(ctx, am.keeper, genesisState)
	return InitGenesis(ctx, am.keeper, genesisState)
}

// ExportGenesis -
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return ModuleCdc.MustMarshalJSON(gs)
}

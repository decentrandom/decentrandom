package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey -
const RouterKey = ModuleName

/*
NewRound
*/

// MsgNewRound -
type MsgNewRound struct {
	ID            string
	Difficulty    uint8
	Owner         sdk.AccAddress
	Nonce         string
	NonceHash     string
	Targets       []string
	DepositCoin   sdk.Coin
	ScheduledTime time.Time
}

// NewMsgNewRound -
func NewMsgNewRound(id string, difficulty uint8, owner sdk.AccAddress, nonceHash string, targets []string, depositCoin sdk.Coin, scheduledTime time.Time) MsgNewRound {
	return MsgNewRound{
		ID:            id,
		Difficulty:    difficulty,
		Owner:         owner,
		Nonce:         "",
		NonceHash:     nonceHash,
		Targets:       targets,
		DepositCoin:   depositCoin,
		ScheduledTime: scheduledTime,
	}
}

// Route -
func (msg MsgNewRound) Route() string {
	return RouterKey
}

// Type -
func (msg MsgNewRound) Type() string {
	return "new_round"
}

// ValidateBasic -
func (msg MsgNewRound) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address can not be empty")
	}

	if len(msg.ID) == 0 || len(msg.NonceHash) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id, noncehash can not be empty")
	}

	// Validate round hash

	if msg.Difficulty < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "difficulty must be greater than 0.")
	}

	return nil
}

// GetSignBytes -
func (msg MsgNewRound) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners -
func (msg MsgNewRound) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

/*
DeployNonce
*/

// MsgDeployNonce -
type MsgDeployNonce struct {
	ID    string
	Owner sdk.AccAddress
	Nonce string
}

// NewMsgDeployNonce -
func NewMsgDeployNonce(id string, owner sdk.AccAddress, nonce string) MsgDeployNonce {
	return MsgDeployNonce{
		ID:    id,
		Owner: owner,
		Nonce: nonce,
	}
}

// Route -
func (msg MsgDeployNonce) Route() string {
	return "rand"
}

// Type -
func (msg MsgDeployNonce) Type() string {
	return "deploy_round"
}

// ValidateBasic -
func (msg MsgDeployNonce) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address can not be empty.")
	}

	if len(msg.ID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id, noncehash can not be empty.")
	}

	return nil
}

// GetSignBytes -
func (msg MsgDeployNonce) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners -
func (msg MsgDeployNonce) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

/*
AddTargets
*/

// MsgAddTargets -
type MsgAddTargets struct {
	ID      string
	Owner   sdk.AccAddress
	Targets []string
}

// NewMsgAddTargets -
func NewMsgAddTargets(id string, owner sdk.AccAddress, targets []string) MsgAddTargets {
	return MsgAddTargets{
		ID:      id,
		Owner:   owner,
		Targets: targets,
	}
}

// Route -
func (msg MsgAddTargets) Route() string {
	return "rand"
}

// Type -
func (msg MsgAddTargets) Type() string {
	return "add_targets"
}

// ValidateBasic - 모집단 추가 ValidateBasic
func (msg MsgAddTargets) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address can not be empty.")
	}
	if len(msg.ID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id can not be empty.")
	}

	if msg.Targets == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "target can not be empty.")
	}

	return nil
}

// GetSignBytes -
func (msg MsgAddTargets) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners -
func (msg MsgAddTargets) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

/*
MsgUpdateTargets
*/

// MsgUpdateTargets -
type MsgUpdateTargets struct {
	ID      string
	Owner   sdk.AccAddress
	Targets []string
}

// NewMsgUpdateTargets -
func NewMsgUpdateTargets(id string, owner sdk.AccAddress, targets []string) MsgUpdateTargets {
	return MsgUpdateTargets{
		ID:      id,
		Owner:   owner,
		Targets: targets,
	}
}

// Route -
func (msg MsgUpdateTargets) Route() string {
	return "rand"
}

// Type -
func (msg MsgUpdateTargets) Type() string {
	return "update_targets"
}

// ValidateBasic -
func (msg MsgUpdateTargets) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address can not be empty.")
	}
	if len(msg.ID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "id can not be empty.")
	}

	if msg.Targets == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "target can not be empty")
	}

	return nil
}

// GetSignBytes -
func (msg MsgUpdateTargets) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners -
func (msg MsgUpdateTargets) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

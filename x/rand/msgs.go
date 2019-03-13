package rand

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgNewRound is a struct format for a message
type MsgNewRound struct {
	ID            string
	Difficulty    int16
	Owner         sdk.AccAddress
	Nonce         int16
	NonceHash     string
	Targets       []string
	ScheduledTime time.Time
	SeedHeights   []int64
}

// NewMsgNewRound - Nonce must be 0, SeedHeights must be nil
func NewMsgNewRound(id string, difficulty int16, owner sdk.AccAddress, nonceHash string, targets []string, scheduledTime time.Time) MsgNewRound {
	return MsgNewRound{
		ID:            id,
		Difficulty:    difficulty,
		Owner:         owner,
		Nonce:         0,
		NonceHash:     nonceHash,
		Targets:       targets,
		ScheduledTime: scheduledTime,
		SeedHeights:   nil,
	}
}

// Route -
func (msg MsgNewRound) Route() string {
	return "rand"
}

// Type -
func (msg MsgNewRound) Type() string {
	return "new_round"
}

// ValidateBasic -
func (msg MsgNewRound) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.ID) == 0 || len(msg.NonceHash) == 0 {
		return sdk.ErrUnknownRequest("Id and/or NonceHash cannot be empty.")
	}

	if msg.Difficulty < 1 {
		return sdk.ErrUnknownRequest("Difficulty must greater than zero(0).")
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

// MsgAddTargets -
type MsgAddTargets struct {
	ID      string
	Owner   sdk.AccAddress
	Targets []string
}

//NewMsgAddTargets -
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

// ValidateBasic -
func (msg MsgAddTargets) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.ID) == 0 {
		return sdk.ErrUnknownRequest("Round ID cannot be empty.")
	}

	if msg.Targets == nil {
		return sdk.ErrUnknownRequest("Targets cannot be empty.")
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

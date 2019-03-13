package rand

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgNewRound struct {
	Id            string
	Difficulty    int16
	Owner         sdk.AccAddress
	Nonce         int16
	NonceHash     string
	Targets       []string
	ScheduledTime time.Time
	SeedHeights   []string
}

// Nonce must be 0, SeedHeights must be nil
func NewMsgNewRound(id string, difficulty int16, owner sdk.AccAddress, nonceHash string, targets []string, scheduledTime time.Time) MsgNewRound {
	return MsgNewRound{
		Id:            id,
		Difficulty:    difficulty,
		Owner:         owner,
		Nonce:         0,
		NonceHash:     nonceHash,
		Targets:       targets,
		ScheduledTime: scheduledTime,
		SeedHeights:   nil,
	}
}

func (msg MsgNewRound) Route() string {
	return "rand"
}

func (msg MsgNewRound) Type() string {
	return "new_round"
}

// Validate
func (msg MsgNewRound) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.Id) == 0 || len(msg.NonceHash) == 0 {
		return sdk.ErrUnknownRequest("Id and/or NonceHash cannot be empty.")
	}

	if msg.Difficulty < 1 {
		return sdk.ErrUnknownRequest("Difficulty must greater than zero(0).")
	}

	return nil
}

func (msg MsgNewRound) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgNewRound) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgAddTargets struct {
	Id      string
	Owner   sdk.AccAddress
	Targets []string
}

func NewMsgAddTargets(id string, owner sdk.AccAddress, targets []string) MsgAddTargets {
	return MsgAddTargets{
		Id:      id,
		Owner:   owner,
		Targets: targets,
	}
}

func (msg MsgAddTargets) Route() string {
	return "rand"
}

func (msg MsgAddTargets) Type() string {
	return "add_targets"
}

func (msg MsgAddTargets) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Id) == 0 {
		return sdk.ErrUnknownRequest("Round ID cannot be empty.")
	}

	if targets == nil {
		return sdk.ErrUnknownRequest("Targets cannot be empty.")
	}

	return nil
}

func (msg MsgAddTargets) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgAddTargets) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

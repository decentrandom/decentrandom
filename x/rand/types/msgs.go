package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey -
const RouterKey = ModuleName // this was defined in your key.go file

/*
NewRound
: 라운드 신규 생성
*/

// MsgNewRound - 신규 라운드 생성을 위한 msg 구조체
type MsgNewRound struct {
	ID            string
	Difficulty    uint8
	Owner         sdk.AccAddress
	Nonce         string
	NonceHash     string
	Targets       []string
	ScheduledTime time.Time
}

// NewMsgNewRound - 초기이므로 Nonce는 0이고, SeedHeights는 빈 값
func NewMsgNewRound(id string, difficulty uint8, owner sdk.AccAddress, nonceHash string, targets []string, scheduledTime time.Time) MsgNewRound {
	return MsgNewRound{
		ID:            id,
		Difficulty:    difficulty,
		Owner:         owner,
		Nonce:         "",
		NonceHash:     nonceHash,
		Targets:       targets,
		ScheduledTime: scheduledTime,
	}
}

// Route - 라운드 신규 생성 Route
func (msg MsgNewRound) Route() string {
	return RouterKey
}

// Type - 라운드 신규 생성 Type
func (msg MsgNewRound) Type() string {
	return "new_round"
}

// ValidateBasic - 라운드 신규 생성 ValidateBasic
func (msg MsgNewRound) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.ID) == 0 || len(msg.NonceHash) == 0 {
		return sdk.ErrUnknownRequest("ID와 NonceHash는 필수 항목 입니다.")
	}

	// 그럴리는 없지만 ID hash 값이 중복되는 경우는 어떻게?
	// important ****** to-do

	if msg.Difficulty < 1 {
		return sdk.ErrUnknownRequest("난이도는 0보다 커야합니다.")
	}

	return nil
}

// GetSignBytes - 라운드 신규 생성 GetSignBytes
func (msg MsgNewRound) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners - 라운드 신규 생성 GetSigners
func (msg MsgNewRound) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

/*
DeployNonce
: 라운드 Nonce 배포
*/

// MsgDeployNonce - 라운드 논스 배포 msg의 구조체
type MsgDeployNonce struct {
	ID    string
	Owner sdk.AccAddress
	Nonce string
}

// NewMsgDeployNonce - 라운드 논스 배포
func NewMsgDeployNonce(id string, owner sdk.AccAddress, nonce string) MsgDeployNonce {
	return MsgDeployNonce{
		ID:    id,
		Owner: owner,
		Nonce: nonce,
	}
}

// Route - 라운드 논스 배포 Route
func (msg MsgDeployNonce) Route() string {
	return "rand"
}

// Type - 라운드 논스 배포 Type
func (msg MsgDeployNonce) Type() string {
	return "deploy_round"
}

// ValidateBasic - 라운드 논스 배포 ValidateBasic
func (msg MsgDeployNonce) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.ID) == 0 {
		return sdk.ErrUnknownRequest("ID와 NonceHash는 필수 항목 입니다.")
	}

	return nil
}

// GetSignBytes - 라운드 논스 배포 GetSignBytes
func (msg MsgDeployNonce) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners - 라운드 논스 배포 GetSigners
func (msg MsgDeployNonce) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

/*
AddTargets
: 라운드에 모집단 추가, 현재는 동일한 값을 중복 등록할 수 없음
*/

// MsgAddTargets - 모집단 추가 msg 구조체
type MsgAddTargets struct {
	ID      string
	Owner   sdk.AccAddress
	Targets []string
}

// NewMsgAddTargets - 모집단 추가
func NewMsgAddTargets(id string, owner sdk.AccAddress, targets []string) MsgAddTargets {
	return MsgAddTargets{
		ID:      id,
		Owner:   owner,
		Targets: targets,
	}
}

// Route - 모집단 추가 Route
func (msg MsgAddTargets) Route() string {
	return "rand"
}

// Type - 모집단 추가 Type
func (msg MsgAddTargets) Type() string {
	return "add_targets"
}

// ValidateBasic - 모집단 추가 ValidateBasic
func (msg MsgAddTargets) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.ID) == 0 {
		return sdk.ErrUnknownRequest("라운드 ID는 필수 항목 입니다.")
	}

	if msg.Targets == nil {
		return sdk.ErrUnknownRequest("Targets은 필수 항목 입니다.")
	}

	return nil
}

// GetSignBytes - 모집단 추가 GetSignBytes
func (msg MsgAddTargets) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners - 모집단 추가 GetSigners
func (msg MsgAddTargets) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

/*
MsgRemoveTargets
: 라운드 모집단 삭제
*/

// MsgRemoveTargets - 라운드 모집단 삭제 msg 구조체
type MsgRemoveTargets struct {
	ID      string
	Owner   sdk.AccAddress
	Targets []string
}

// NewMsgRemoveTargets - 라운드 모집단 삭제
func NewMsgRemoveTargets(id string, owner sdk.AccAddress, targets []string) MsgRemoveTargets {
	return MsgRemoveTargets{
		ID:      id,
		Owner:   owner,
		Targets: targets,
	}
}

// Route - 라운드 모집단 삭제 Route
func (msg MsgRemoveTargets) Route() string {
	return "rand"
}

// Type - 라운드 모집단 삭제 Type
func (msg MsgRemoveTargets) Type() string {
	return "remove_targets"
}

// ValidateBasic - 라운드 모집단 삭제 ValidateBasic
func (msg MsgRemoveTargets) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.ID) == 0 {
		return sdk.ErrUnknownRequest("라운드 ID는 필수 항목 입니다.")
	}

	if msg.Targets == nil {
		return sdk.ErrUnknownRequest("Targets는 필수 항목 입니다.")
	}

	return nil
}

// GetSignBytes - 라운드 모집단 삭제 GetSignBytes
func (msg MsgRemoveTargets) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners - 라운드 모집단 삭제 GetSigners
func (msg MsgRemoveTargets) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

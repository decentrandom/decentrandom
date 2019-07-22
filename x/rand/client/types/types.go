package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Round - 라운드 기본 구조체
type Round struct {
	Difficulty    uint8          `json:"difficulty"`
	Owner         sdk.AccAddress `json:"owner"`
	Nonce         string         `json:"nonce"`
	NonceHash     string         `json:"nonce_hash"`
	Targets       []string       `json:"targets"`
	ScheduledTime time.Time      `json:"scheduled_time"`
}

func (a Artist) String() string {
	return fmt.Sprintf(`Song %d:
		  Artist ID:	%s
		  Image:		%s
		  Name:			%s`, a.ArtistID, a.Image, a.Name)
}

// Artists is an array of song
// To FIX with new fields
type Artists []*Artist

func (artists Artists) String() string {
	out := fmt.Sprintf("%10s - (%15s) - (%40s) - [%10s] - Create Time\n", "ID", "Title", "Owner", "CreateTime")
	for _, artist := range artists {
		out += fmt.Sprintf("%10d - (%15s) - (%40s) - [%10s]\n",
			artist.ArtistID, artist.Image, artist.Name)
	}

	return strings.TrimSpace(out)
}
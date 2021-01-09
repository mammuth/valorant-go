package api

import (
	"strings"
	"testing"
)

func TestVerboseMapName(t *testing.T) {
	match := Match{MapId: "/Game/Maps/Triad/Triad"}
	if name := strings.ToLower(match.VerboseMapName()); name != "haven" {
		t.Errorf("Map verbose name doesn't match. It returend %s", name)
	}
}

func TestEloChange(t *testing.T) {
	match := Match{TierProgressBeforeUpdate: 50, TierProgressAfterUpdate: 80}
	if eloChange := match.EloChange(); eloChange != 30 {
		t.Errorf("Elo change reported %d instead of 30", eloChange)
	}

	match2 := Match{TierProgressBeforeUpdate: 50, TierProgressAfterUpdate: 30}
	if eloChange := match2.EloChange(); eloChange != -20 {
		t.Errorf("Elo change reported %d instead of -20", eloChange)
	}
}

package mptoken

import (
	"strings"
	"testing"
)

func TestBuildMPTokenIssuanceIDAndParseMPTokenIssuanceID(t *testing.T) {
	// Example issuer address (classic XRPL address)
	issuer := "raZ3wTTKiMHn3BiStvz4ET9rbCHfU1DMak"
	sequence := uint32(6560006)
	mptoken := "006419063CEBEB49FC20032206CE0F203138BFC59F1AC578"

	id, err := BuildMPTokenIssuanceID(sequence, issuer)
	if err != nil {
		t.Fatalf("BuildMPTokenIssuanceID failed: %v", err)
	}
	if len(id) != 48 {
		t.Errorf("Expected ID length 48, got %d", len(id))
	}
	if id[:8] != "00641906" { // 6560006 in hex is 0x00641906
		t.Errorf("Expected sequence hex prefix, got %s", id[:8])
	}
	if id != mptoken {
		t.Errorf("Expected ID %s, got %s", mptoken, id)
	}

	seq2, issuer2, err := ParseMPTokenIssuanceID(id)
	if err != nil {
		t.Fatalf("ParseMPTokenIssuanceID failed: %v", err)
	}
	if seq2 != sequence {
		t.Errorf("Expected sequence %d, got %d", sequence, seq2)
	}
	if issuer2 != issuer {
		t.Errorf("Expected issuer %s, got %s", issuer, issuer2)
	}
}

func TestParseMPTokenIssuanceID_InvalidLength(t *testing.T) {
	_, _, err := ParseMPTokenIssuanceID("short")
	if err == nil {
		t.Error("Expected error for invalid length, got nil")
	}
}

func TestParseMPTokenIssuanceID_InvalidHex(t *testing.T) {
	// 8 chars for sequence + 40 invalid hex chars
	badID := "00BC614E" + strings.Repeat("ZZ", 20)
	_, _, err := ParseMPTokenIssuanceID(badID)
	if err == nil {
		t.Error("Expected error for invalid hex, got nil")
	}
}

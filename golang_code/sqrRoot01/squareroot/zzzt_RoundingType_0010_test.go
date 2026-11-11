package squareroot

import "testing"

// TestRoundingTypeStringAwayFromZero
//
// Verifies that the String() method returns "AwayFromZero" for the
// AwayFromZero enum constant.
//
// Input Parameters
// ----------------
//
// testing.T
//  Standard Go testing fixture.
//
// Return Parameters
// -----------------
//
// None
func TestRoundingTypeStringAwayFromZero(t *testing.T) {

	expected := "AwayFromZero"
	actual := AwayFromZero.String()

	if actual != expected {
		t.Errorf("AwayFromZero.String() = %q, want %q", actual, expected)
	}
}

// TestRoundingTypeStringBankers
//
// Verifies that the String() method returns "Bankers" for the Bankers
// enum constant.
//
// Input Parameters
// ----------------
//
// testing.T
//  Standard Go testing fixture.
//
// Return Parameters
// -----------------
//
// None
func TestRoundingTypeStringBankers(t *testing.T) {

	expected := "Bankers"
	actual := Bankers.String()

	if actual != expected {
		t.Errorf("Bankers.String() = %q, want %q", actual, expected)
	}
}

// TestRoundingTypeStringTruncation
//
// Verifies that the String() method returns "Truncation" for the
// Truncation enum constant.
//
// Input Parameters
// ----------------
//
// testing.T
//  Standard Go testing fixture.
//
// Return Parameters
// -----------------
//
// None
func TestRoundingTypeStringTruncation(t *testing.T) {

	expected := "Truncation"
	actual := Truncation.String()

	if actual != expected {
		t.Errorf("Truncation.String() = %q, want %q", actual, expected)
	}
}

// TestRoundingTypeStringZeroValue
//
// Verifies that the String() method returns "INVALID" for the zero
// (uninitialized) value of RoundingType.
//
// Input Parameters
// ----------------
//
// testing.T
//  Standard Go testing fixture.
//
// Return Parameters
// -----------------
//
// None
func TestRoundingTypeStringZeroValue(t *testing.T) {

	var rt RoundingType
	expected := "INVALID"
	actual := rt.String()

	if actual != expected {
		t.Errorf("zero RoundingType.String() = %q, want %q", actual, expected)
	}
}

// TestRoundingTypeStringOutOfRange
//
// Verifies that the String() method returns "INVALID" for out-of-range
// RoundingType values.
//
// Input Parameters
// ----------------
//
// testing.T
//  Standard Go testing fixture.
//
// Return Parameters
// -----------------
//
// None
func TestRoundingTypeStringOutOfRange(t *testing.T) {

	rt := RoundingType(99)
	expected := "INVALID"
	actual := rt.String()

	if actual != expected {
		t.Errorf("out-of-range RoundingType.String() = %q, want %q", actual, expected)
	}
}

// TestRoundingTypeEnumValues
//
// Verifies that the enum constants are numbered starting at 1 using
// iota + 1: AwayFromZero=1, Bankers=2, Truncation=3.
//
// Input Parameters
// ----------------
//
// testing.T
//  Standard Go testing fixture.
//
// Return Parameters
// -----------------
//
// None
func TestRoundingTypeEnumValues(t *testing.T) {

	if AwayFromZero != 1 {
		t.Errorf("AwayFromZero = %d, want 1", AwayFromZero)
	}

	if Bankers != 2 {
		t.Errorf("Bankers = %d, want 2", Bankers)
	}

	if Truncation != 3 {
		t.Errorf("Truncation = %d, want 3", Truncation)
	}
}

// TestRoundingTypeNamesArray
//
// Verifies that the RoundingTypeNames array has "INVALID" at index 0
// and contains the correct string representations for each enum value.
//
// Input Parameters
// ----------------
//
// testing.T
//  Standard Go testing fixture.
//
// Return Parameters
// -----------------
//
// None
func TestRoundingTypeNamesArray(t *testing.T) {

	if RoundingTypeNames[0] != "INVALID" {
		t.Errorf("RoundingTypeNames[0] = %q, want %q", RoundingTypeNames[0], "INVALID")
	}

	if len(RoundingTypeNames) != 4 {
		t.Errorf("len(RoundingTypeNames) = %d, want 4", len(RoundingTypeNames))
	}

	if RoundingTypeNames[1] != "AwayFromZero" {
		t.Errorf("RoundingTypeNames[1] = %q, want %q", RoundingTypeNames[1], "AwayFromZero")
	}

	if RoundingTypeNames[2] != "Bankers" {
		t.Errorf("RoundingTypeNames[2] = %q, want %q", RoundingTypeNames[2], "Bankers")
	}

	if RoundingTypeNames[3] != "Truncation" {
		t.Errorf("RoundingTypeNames[3] = %q, want %q", RoundingTypeNames[3], "Truncation")
	}
}

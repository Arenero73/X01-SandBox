package squareroot

// RoundingType defines the strategy used for numeric rounding.
type RoundingType int

// Exported RoundingType enum values starting at 1.
const (
	AwayFromZero RoundingType = iota + 1
	Bankers
	Truncation
)

// RoundingTypeNames maps RoundingType values to their string representations.
// Index 0 represents the uninitialized / invalid state.
var RoundingTypeNames = [...]string{
	"INVALID",
	"AwayFromZero",
	"Bankers",
	"Truncation",
}

// String implements the fmt.Stringer interface for RoundingType.
func (rt RoundingType) String() string {
	if int(rt) < 1 || int(rt) >= len(RoundingTypeNames) {
		return RoundingTypeNames[0]
	}

	return RoundingTypeNames[rt]
}

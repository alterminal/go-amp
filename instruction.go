package goamp

// Instruction
const (
	END byte = iota // End
	INT             // Int64
	FLO             // Float64
	U8              // UTF-8
	ATM             // Atom ascii string
	TRU             // True
	FAL             // False
	MAP             // Map
	LIS             // List
	TUP             // Tuple
)

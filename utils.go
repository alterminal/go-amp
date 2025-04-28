package goamp

func PrintBinary(data byte) {
	// Print the binary representation of the byte
	for i := 7; i >= 0; i-- {
		if (data>>i)&1 == 1 {
			print("1")
		} else {
			print("0")
		}
	}
	println()
}

func ByteToBinaryString(data byte) string {
	// Convert the byte to a binary string
	binaryString := ""
	for i := 7; i >= 0; i-- {
		if (data>>i)&1 == 1 {
			binaryString += "1"
		} else {
			binaryString += "0"
		}
	}
	return binaryString
}

func IntToBinaryString(data int) string {
	// Convert the integer to a binary string
	binaryString := ""
	for i := 63; i >= 0; i-- {
		if (data>>i)&1 == 1 {
			binaryString += "1"
		} else {
			binaryString += "0"
		}
	}
	return binaryString
}

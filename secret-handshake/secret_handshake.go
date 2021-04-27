package secret

var handshakes = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

func Handshake(input uint) []string {
	var result []string
	var checklist []uint

	if (input & 16) == 16 {
		checklist = []uint{8, 4, 2, 1}
	} else {
		checklist = []uint{1, 2, 4, 8}
	}

	for _, v := range checklist {
		if (input & v) == v {
			result = append(result, handshakes[v])
		}
	}

	return result
}

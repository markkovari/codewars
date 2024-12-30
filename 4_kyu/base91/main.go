package main

var (
	enctab = [...]byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
		'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '#',
		'$', '%', '&', '(', ')', '*', '+', ',', '.', '/', ':', ';', '<', '=', '>', '?',
		'@', '[', ']', '^', '_', '`', '{', '|', '}', '~', '"',
	}
	dectab = [...]byte{
		91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
		91, 91, 91, 91, 91, 91, 91, 91, 91, 62, 90, 63, 64, 65, 66, 91, 67, 68, 69, 70, 71, 91, 72, 73,
		52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 74, 75, 76, 77, 78, 79, 80, 0, 1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 81, 91, 82, 83, 84, 85, 26,
		27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		51, 86, 87, 88, 89, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
		91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
		91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
		91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
		91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
		91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	}
)

func decode(src []byte) []byte {
	// before and next must be at least uint32
	var before, next uint32 = 0, 0
	value := -1
	decoded := []byte{}

	for i := 0; i < len(src); i++ {
		p := dectab[src[i]]
		if p > 90 {
			// skip invalid character silently
			continue
		}
		if value < 0 {
			value = int(p)
			continue
		}
		value += int(p) * 91
		before |= uint32(value) << next
		if value&0x1fff > 88 {
			next += 13
		} else {
			next += 14
		}
		for {
			decoded = append(decoded, uint8(before))
			before >>= 8
			next -= 8
			if next <= 7 {
				break
			}
		}
		value = -1
	}

	if value > -1 {
		decoded = append(decoded, uint8(before|uint32(value)<<next))
	}

	return decoded
}

func encode(src []byte) []byte {
	var before, next uint32 = 0, 0
	encoded := []byte{}

	for i := 0; i < len(src); i++ {
		before |= uint32(src[i]) << next
		next += 8

		if next <= 13 {
			continue
		}

		value := before & 0x1fff
		if value > 88 {
			before >>= 13
			next -= 13
		} else {
			value = before & 0x3fff
			before >>= 14
			next -= 14
		}

		encoded = append(encoded, enctab[value%91], enctab[value/91])
	}

	if next != 0 {
		encoded = append(encoded, enctab[before%91])
		if next > 7 || before > 90 {
			encoded = append(encoded, enctab[before/91])
		}
	}

	return encoded
}

func Encode(d []byte) []byte {
	return encode(d)
}

func Decode(d []byte) []byte {
	return decode(d)
}

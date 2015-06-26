package teletext

// Generate parity bit for 7 bit data
func Parity(data []byte) []byte {
	for i, b := range data {
		hweight := 0
		for b != 0 {
			hweight += int(b & 1)
			b >>= 1
		}
		if (hweight & 1) == 0 {
			data[i] |= 0x80
		}
	}
	return data
}

// 8:4 hamming encoding
func Ham(data []byte) []byte {
	for i, d := range data {
		d1 := d & 1
		d2 := (d >> 1) & 1
		d3 := (d >> 2) & 1
		d4 := (d >> 3) & 1

		p1 := (1 + d1 + d3 + d4) & 1
		p2 := (1 + d1 + d2 + d4) & 1
		p3 := (1 + d1 + d2 + d3) & 1
		p4 := (1 + p1 + d1 + p2 + d2 + p3 + d3 + d4) & 1

		data[i] = (p1 | (d1 << 1) | (p2 << 2) | (d2 << 3) | (p3 << 4) | (d3 << 5) | (p4 << 6) | (d4 << 7))
	}
	return data
}

// Unescape .tti line data
func unescape(line []byte) []byte {
	out := []byte{}
	escaped := false
	for _, c := range line {
		if escaped {
			out = append(out, c-0x40)
			escaped = false
		} else if c == 0x1b {
			escaped = true
		} else if c == 0x10 {
			out = append(out, 0x0d)
		} else if c > 0x80 {
			out = append(out, c-0x80)
		} else {
			out = append(out, c)
		}
	}
	return out
}

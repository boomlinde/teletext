package teletext

// Generates parity bit for 7 bit data
func Parity(data []byte) []byte {
	out := []byte{}
	for _, b := range data {
		orig := b
		hweight := 0
		for b != 0 {
			hweight += int(b & 1)
			b >>= 1
		}
		if (hweight & 1) == 0 {
			orig |= 0x80
		}
		out = append(out, orig)
	}
	return out
}

// 8:4 hamming encodes the input bytes
func Ham(data []byte) []byte {
	out := []byte{}
	for _, b := range data {
		d1 := b & 1
		d2 := (b >> 1) & 1
		d3 := (b >> 2) & 1
		d4 := (b >> 3) & 1

		p1 := (1 + d1 + d3 + d4) & 1
		p2 := (1 + d1 + d2 + d4) & 1
		p3 := (1 + d1 + d2 + d3) & 1
		p4 := (1 + p1 + d1 + p2 + d2 + p3 + d3 + d4) & 1

		b := (p1 | (d1 << 1) | (p2 << 2) | (d2 << 3) | (p3 << 4) | (d3 << 5) | (p4 << 6) | (d4 << 7))
		out = append(out, b)
	}
	return out
}

// Unescapes .tti line data
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

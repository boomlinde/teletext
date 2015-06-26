package teletext

import (
	"bytes"
	"strconv"
)

// Converts .tti file data to the internal page format
func ConvertTTI(title string, data []byte) Page {
	out := Page{}
	lines := bytes.Split(data, []byte{0xd, 0xa})
	pagenumber := 100
	for _, l := range lines {
		line := unescape(l)
		if bytes.HasPrefix(line, []byte("PN,")) {
			lsplit := bytes.SplitN(line, []byte(","), 2)
			pagenumber, _ = strconv.Atoi(string(lsplit[1]))
			pagenumber /= 100
			out = append(out, PageHeader{Header{Page: pagenumber}, title})
		} else if bytes.HasPrefix(line, []byte("OL,")) {
			lsplit := bytes.SplitN(line, []byte(","), 3)
			row, _ := strconv.Atoi(string(lsplit[1]))
			data := lsplit[2]
			if len(data) > 40 {
				data = data[:40]
			}
			out = append(out, OutputLine{Header{pagenumber, row}, data})
		}
	}
	return out
}

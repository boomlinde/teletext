package teletext

// Converts 971 format .ttv file data to the internal page format
func ConvertTTV(title string, pagenumber int, data []byte) Page {
	out := Page{PageHeader{Header{Page: pagenumber}, title}}
	for y := 1; y < 24; y++ {
		ldata := data[y*40 : y*40+40]
		out = append(out, OutputLine{Header{pagenumber, y}, ldata})
	}
	return out
}

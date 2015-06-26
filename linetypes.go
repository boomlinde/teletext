package teletext

// Page header line type
type PageHeader struct {
	Header
	Title string // Page title
}

func (p PageHeader) Serialize() []byte {
	title := p.Title
	for len(title) != 32 {
		title += " "
	}
	p.Header.Row = 0
	number := p.Header.Page % 100
	data := p.Header.Serialize()
	data = append(data, Ham([]byte{byte(number % 10), byte(number / 10), 0, 0, 0, 0, 0, 0})...)
	data = append(data, Parity([]byte(title))...)
	return data
}

// Output line type
type OutputLine struct {
	Header
	Data []byte // Line data
}

func (o OutputLine) Serialize() []byte {
	for len(o.Data) < 40 {
		o.Data = append(o.Data, byte(' '))
	}
	return append(o.Header.Serialize(), Parity(o.Data)...)
}

// A Go library for generating teletext data suitable for
// [raspi-teletext](https://github.com/ali1234/raspi-teletext). Contains data
// structures and routines suitable for building pages from the ground up or
// generating them using `.tti` files.
package teletext

import (
	"sort"
)

// Teletext line header.
type Header struct {
	Page int // Page number
	Row  int // Row number
}

// Interface for line types
type Line interface {
	Serialize() []byte
	GetHeader() *Header
	SetPage(int)
}

type Page []Line

// Returns a pointer to the line header
func (h Header) GetHeader() *Header {
	return &h
}

// Sets the page number of a line
func (h Header) SetPage(page int) {
	h.Page = page
}

func (p Page) Len() int           { return len(p) }
func (p Page) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Page) Less(i, j int) bool { return p[i].GetHeader().Row < p[j].GetHeader().Row }

func (h Header) Serialize() []byte {
	magazine := h.Page / 100
	return Ham([]byte{
		byte(magazine | ((h.Row & 1) << 3)),
		byte(h.Row >> 1),
	})
}

func (p Page) Serialize() []byte {
	data := []byte{}
	sort.Sort(p)
	for _, line := range p {
		data = append(data, line.Serialize()...)
	}
	return data
}

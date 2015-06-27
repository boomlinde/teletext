# teletext
--
    import "github.com/boomlinde/teletext"

A Go library for generating teletext data suitable for
[raspi-teletext](https://github.com/ali1234/raspi-teletext). Contains data
structures and routines suitable for building pages from the ground up or
generating them using `.tti` files.

## Usage

### func  Ham

    func Ham(data []byte) []byte

8:4 hamming encodes the input bytes

### func  Parity

    func Parity(data []byte) []byte

Generates parity bit for 7 bit data

### type Header

    type Header struct {
    	Page int // Page number
    	Row  int // Row number
    }


Teletext line header.

### func (Header) GetHeader

    func (h Header) GetHeader() *Header

Returns a pointer to the line header

### func (Header) Serialize

    func (h Header) Serialize() []byte


### func (Header) SetPage

    func (h Header) SetPage(page int)

Sets the page number of a line

### type Line

    type Line interface {
    	Serialize() []byte
    	GetHeader() *Header
    	SetPage(int)
    }


Interface for line types

### type OutputLine

    type OutputLine struct {
    	Header
    	Data []byte // Line data
    }


Output line type

### func (OutputLine) Serialize

    func (o OutputLine) Serialize() []byte


### type Page

    type Page []Line



### func  ConvertTTI

    func ConvertTTI(title string, data []byte) Page

Converts .tti file data to the internal page format

### func (Page) Len

    func (p Page) Len() int

Satisfies sorting interface

### func (Page) Less

    func (p Page) Less(i, j int) bool

Satisfies sorting interface

### func (Page) Serialize

    func (p Page) Serialize() []byte

Serializes a teletext page

### func (Page) Swap

    func (p Page) Swap(i, j int)

Satisfies sorting interface

### type PageHeader

    type PageHeader struct {
    	Header
    	Title string // Page title
    }


Page header line type

### func (PageHeader) Serialize

    func (p PageHeader) Serialize() []byte

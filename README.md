# teletext
--
    import "github.com/boomlinde/teletext"

A Go library for generating teletext data suitable for
[raspi-teletext](https://github.com/ali1234/raspi-teletext). Contains data
structures and routines suitable for building pages from the ground up or
generating them using `.tti` files.

## Usage

#### func  Ham

```go
func Ham(data []byte) []byte
```
8:4 hamming encodes the input bytes

#### func  Parity

```go
func Parity(data []byte) []byte
```
Generates parity bit for 7 bit data

#### type Header

```go
type Header struct {
	Page int // Page number
	Row  int // Row number
}
```

Teletext line header.

#### func (Header) GetHeader

```go
func (h Header) GetHeader() *Header
```
Returns a pointer to the line header

#### func (Header) Serialize

```go
func (h Header) Serialize() []byte
```

#### func (Header) SetPage

```go
func (h Header) SetPage(page int)
```
Sets the page number of a line

#### type Line

```go
type Line interface {
	Serialize() []byte
	GetHeader() *Header
	SetPage(int)
}
```

Interface for line types

#### type OutputLine

```go
type OutputLine struct {
	Header
	Data []byte // Line data
}
```

Output line type

#### func (OutputLine) Serialize

```go
func (o OutputLine) Serialize() []byte
```

#### type Page

```go
type Page []Line
```


#### func  ConvertTTI

```go
func ConvertTTI(title string, data []byte) Page
```
Converts .tti file data to the internal page format

#### func (Page) Len

```go
func (p Page) Len() int
```
Satisfies sorting interface

#### func (Page) Less

```go
func (p Page) Less(i, j int) bool
```
Satisfies sorting interface

#### func (Page) Serialize

```go
func (p Page) Serialize() []byte
```
Serializes a teletext page

#### func (Page) Swap

```go
func (p Page) Swap(i, j int)
```
Satisfies sorting interface

#### type PageHeader

```go
type PageHeader struct {
	Header
	Title string // Page title
}
```

Page header line type

#### func (PageHeader) Serialize

```go
func (p PageHeader) Serialize() []byte
```

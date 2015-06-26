package main

import (
	"github.com/boomlinde/teletext"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %s [file.tti ...]\n", os.Args[0])
		os.Exit(1)
	}
	out := []byte{}
	for _, arg := range os.Args[1:] {
		log.Printf("Converting \"%s\"", arg)
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			panic(err)
		}
		out = append(out, teletext.ConvertTTI("teletext", data).Serialize()...)
	}
	if len(out) != 0 {
		log.Println("Starting carousel")
		for {
			os.Stdout.Write(out)
		}
	}
}

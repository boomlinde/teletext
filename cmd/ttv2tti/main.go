package main

import (
	"flag"
	"fmt"
	"github.com/boomlinde/teletext"
	"io/ioutil"
	"log"
)

func main() {
	var pn int
	var sc string
	var ps string
	var fl string
	var pageInc int

	flag.IntVar(&pn, "PN", 10000, "The page number (including carousel frame)")
	flag.StringVar(&sc, "SC", "0000", "Page subcode")
	flag.StringVar(&ps, "PS", "80C0", "Page status")
	flag.StringVar(&fl, "FL", "", "Fastext line")
	flag.IntVar(&pageInc, "pageinc", 100, "Page increment per input file")
	flag.Parse()

	for _, fname := range flag.Args() {
		log.Println("Converting", fname)
		bytes, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatal(err)
		}

		page := teletext.ConvertTTV("", 0, bytes)
		for i := range page {
			page[i].GetHeader().Page = pn
		}

		ttiFname := fname + ".tti"

		headers := map[string]string{
			"PN": fmt.Sprintf("%d", pn),
			"SC": sc,
			"PS": ps,
		}
		if fl != "" {
			headers["FL"] = fl
		}
		outdata := page.BuildTTI(headers)
		if err := ioutil.WriteFile(ttiFname, outdata, 0666); err != nil {
			panic(err)
		}
		pn += pageInc
	}
	log.Println("Done")
}

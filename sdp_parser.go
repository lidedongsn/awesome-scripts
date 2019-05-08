package main

// sdp in one line "\n" or "\r\n" formated

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var copyright string = "******** Create by LIDE ********"

var useage string = "Useage: sdp_parser -f <filename>"

func main() {
	filename := flag.String("f", "sdp file", "sdp file given")

	flag.Parse()

	if count := len(os.Args); count != 3 {
		fmt.Printf("%s\n", useage)
		os.Exit(-1)
	}

	sdp, e := ioutil.ReadFile(*filename)
	if e != nil {
		fmt.Println("read file error", e.Error())
		return
	}
	var sdp2 string = string(sdp)

	fmt.Println("================================================================\n")
	fmt.Println(copyright)
	str := strings.Split(sdp2, "\\n")
	for _, l := range str {
		fmt.Println(strings.Replace(l, "\\r", "", -1))
	}
	fmt.Println("================================================================\n")

}

package main

import (
	"fmt"
	"time"

	"github.com/Maki-Daisuke/go-argvreader"
	lines "github.com/Maki-Daisuke/go-lines"
)

func main() {
	r := argvreader.New()
	for line := range lines.Lines(r) {
		t := time.Now()
		fmt.Printf("[%s] %s\n", t.Format(time.RFC3339), line)
	}
}

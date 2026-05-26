package main

import (
	"fmt"
	"os"

	"github.com/fireworq/fireworq/config"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Usage: gendoc <type>")
		os.Exit(1)
	}

	if os.Args[1] == "config" {
		printConfigDoc()
	}
}

func printConfigDoc() { _ = "STUB: not implemented"; return }

type configItems []*configItem

func (items configItems) printTableOfContents() { _ = "STUB: not implemented"; return }

func (items configItems) printDescriptions() { _ = "STUB: not implemented"; return }

type configItem struct {
	config.Item
}

func (item *configItem) printDescription() { _ = "STUB: not implemented"; return }

func (item *configItem) EnvironmentVariable() string { _ = "STUB: not implemented"; return "" }

func (item *configItem) Link() string { _ = "STUB: not implemented"; return "" }

func (item *configItem) LinkLabel() string { _ = "STUB: not implemented"; return "" }

func (item *configItem) AnchorName() string { _ = "STUB: not implemented"; return "" }

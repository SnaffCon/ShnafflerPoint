package output

import (
	"fmt"

	"github.com/fatih/color"
)

// TODO: banner!!!!!!!!!

// Example colours
var (
	GoldishMaybe = color.New(color.FgHiYellow, color.Bold)
)

func PrintBanner() {
	fmt.Println(GoldishMaybe.Sprint("TODO: BANNER"))
}

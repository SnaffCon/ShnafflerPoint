package output

import (
	"fmt"

	"github.com/fatih/color"
)

var banner = [11]string{
	" _______ __                 ___  ___ __            ",
	"|   _   |  |--.-----.---.-.'  _.'  _|  .-----.----.",
	"|   1___|     |     |  _  |   _|   _|  |  -__|   _|",
	"|____   |__|__|__|__|___._|__| |__| |__|_____|__|  ",
	"|:  1   |              _______       __       __   ",
	"|::.. . |             |   _   .-----|__.-----|  |_ ",
	"`-------'             |.  1   |  _  |  |     |   _|",
	"                      |.  ____|_____|__|__|__|____|",
	" by @sh3r4_hax        |:  |                        ",
	"  & @mikeloss         |::.|                        ",
	"                      `---'                        ",
}

// Example colours
var (
	GoldishMaybe = color.New(color.FgHiYellow, color.Bold)
	ReddishMaybe = color.New(color.FgRed, color.Bold)
)

func PrintBanner() {
	for index, bannerline := range banner {
		if index <= 3 {
			fmt.Println(GoldishMaybe.Sprint(bannerline))
		} else {
			fmt.Print(GoldishMaybe.Sprint(bannerline[0:20]))
			fmt.Println(ReddishMaybe.Sprint(bannerline[21:51]))
		}
	}
}

package spsnaffler

import (
	"fmt"
	"github.com/SnaffCon/ShnafflerPoint/internal/pkg/logging"
	"os"

	"github.com/SnaffCon/ShnafflerPoint/internal/pkg/config"
	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/api"
	strategy "github.com/koltyakov/gosip/auth/addin"
)

// example reference code for when I forget what I was supposed to be doing
// https://github.com/koltyakov/gosip
func main() {
	auth := &strategy.AuthCnfg{
		SiteURL:      os.Getenv("SPAUTH_SITEURL"),
		ClientID:     os.Getenv("SPAUTH_CLIENTID"),
		ClientSecret: os.Getenv("SPAUTH_CLIENTSECRET"),
	}

	client := &gosip.SPClient{AuthCnfg: auth}

	sp := api.NewSP(client)

	res, err := sp.Web().Select("Title").Get()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", res.Data().Title)
}

// Run the whole fucken thing (except for UI)
func Run(cfg *config.Config, c config.State) {
	// TODO
	logging.Warn("NOT IMPLEMENTED")
}

package main

import (
	"context"

	"github.com/SnaffCon/ShnafflerPoint/internal/pkg/config"
	"github.com/SnaffCon/ShnafflerPoint/internal/pkg/output"
	"github.com/SnaffCon/ShnafflerPoint/pkg/spsnaffler"
)

func main() {

	// the most important part
	output.PrintBanner()

	// init logging and output
	ctx := context.Context(context.Background())
	uiChan := make(chan *output.Message)
	output.Init(ctx, uiChan, output.DEGUB, output.GREEN)

	// Parse config
	cfg := config.Parse()
	state := config.State{
		Output: uiChan,
		Ctx:    ctx,
	}

	// Do the stuff
	spsnaffler.Run(cfg, state)
}

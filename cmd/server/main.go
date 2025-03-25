package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/gotemplate"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"github.com/osamikoyo/publics/internal/modules/event"
	"github.com/osamikoyo/publics/internal/modules/user"
)

func main() {
	flamingo.App(
		[]dingo.Module{
			new(user.UserModule),
			new(gotemplate.Module),
			new(event.EventModule),
			new(requestlogger.Module),
		},
	)
}

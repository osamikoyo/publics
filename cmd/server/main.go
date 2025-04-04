package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"github.com/osamikoyo/publics/internal/modules/event"
	"github.com/osamikoyo/publics/internal/modules/member"
	"github.com/osamikoyo/publics/internal/modules/user"
)

func main() {
	flamingo.App(
		[]dingo.Module{
			new(member.MemberModule),
			new(user.UserModule),
			new(event.EventModule),
			new(requestlogger.Module),
		},
	)
}

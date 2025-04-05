package entity

import "github.com/osamikoyo/publics/internal/modules/event/entity"

type Element struct {
	Parents []entity.Event
	Self    entity.Event
}

package entity

import "github.com/osamikoyo/publics/internal/modules/event/entity"

type UID uint64

type Element struct {
	ID      UID
	Parents []entity.Event
	Self    entity.Event
}

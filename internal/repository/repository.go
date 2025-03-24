package repository

import "github.com/osamikoyo/publics/internal/repository/models"

type Repo interface {
	CreateEvent(*models.Event) error
	GetEventBy(key string, value string) []models.Event
	UpdateEvent(id uint, newEvent *models.Event)
	DeleteEvent(*models.Event) error 
}



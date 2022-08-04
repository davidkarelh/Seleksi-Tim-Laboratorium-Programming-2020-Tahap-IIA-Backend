package repository

import (
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/entity"
	"gorm.io/gorm"
)

type IHistoryRepository interface {
	InsertHistory(history entity.History) entity.History
	UpdateHistory(history entity.History) entity.History
	FindByHistoryname(Historyname string) entity.History
}

type HistoryConnection struct {
	connection *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) IHistoryRepository {
	return &HistoryConnection{
		connection: db,
	}
}

func (db *HistoryConnection) InsertHistory(history entity.History) entity.History {
	db.connection.Save(&history)
	return history
}

func (db *HistoryConnection) UpdateHistory(history entity.History) entity.History {
	db.connection.Save(&history)
	return history
}

func (db *HistoryConnection) FindByHistoryname(Historyname string) entity.History {
	var History entity.History
	db.connection.Where("Historyname = ?", Historyname).Take(&History)
	return History
}

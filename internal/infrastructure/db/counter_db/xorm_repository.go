package counter_db

import (
	"github.com/zethange/goodsocd/internal/domain/counter"
	"xorm.io/xorm"
)

type CommandCounterDB struct {
	Id       int64 `xorm:"pk autoincr"`
	ChatId   int64 `xorm:"not null index"`
	UserId   int64 `xorm:"not null index"`
	UserName string
	FullName string
	Count    int `xorm:"not null default 0"`
}

type XORMRepository struct {
	engine *xorm.Engine
}

func NewXORMRepository(engine *xorm.Engine) *XORMRepository {
	return &XORMRepository{engine: engine}
}

func (r *XORMRepository) FindByChatAndUser(chatID, userID int64) (*counter.CommandCounter, error) {
	var model CommandCounterDB
	has, err := r.engine.Where("chat_id = ? AND user_id = ?", chatID, userID).Get(&model)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return toDomainEntity(model), nil
}

func (r *XORMRepository) Save(entity *counter.CommandCounter) error {
	model := fromDomainEntity(entity)
	if model.Id == 0 {
		_, err := r.engine.Insert(model)
		return err
	}
	_, err := r.engine.ID(model.Id).Update(model)
	return err
}

func (r *XORMRepository) GetTopByChat(chatID int64, limit int) ([]*counter.CommandCounter, error) {
	var models []CommandCounterDB
	err := r.engine.Where("chat_id = ?", chatID).Desc("count").Limit(limit).Find(&models)
	if err != nil {
		return nil, err
	}

	entities := make([]*counter.CommandCounter, len(models))
	for i, m := range models {
		entities[i] = toDomainEntity(m)
	}
	return entities, nil
}

func toDomainEntity(model CommandCounterDB) *counter.CommandCounter {
	return &counter.CommandCounter{
		ID:       model.Id,
		ChatID:   model.ChatId,
		UserID:   model.UserId,
		UserName: model.UserName,
		Count:    model.Count,
	}
}

func fromDomainEntity(entity *counter.CommandCounter) *CommandCounterDB {
	return &CommandCounterDB{
		Id:       entity.ID,
		ChatId:   entity.ChatID,
		UserId:   entity.UserID,
		UserName: entity.UserName,
		Count:    entity.Count,
	}
}

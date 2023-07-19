package models

import (
	"github.com/google/uuid"
)

type ItemState struct {
	Amount   uint   `json:"amount"`
	Uuid     string `json:"uuid"`
	UserUuid string `json:"user_uuid"`
	ItemUuid string `json:"item_uuid"`
}

func NewItemState(
	amount uint,
	userUuid string,
	itemUuid string,
) *ItemState {
	return &ItemState{
		Amount:   amount,
		Uuid:     uuid.NewString(),
		UserUuid: userUuid,
		ItemUuid: itemUuid,
	}
}

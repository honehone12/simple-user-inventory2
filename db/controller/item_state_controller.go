package controller

import (
	"context"
	"simple-user-inventory2/db/models"
	"time"

	"github.com/redis/rueidis"
)

type ItemStateController struct {
	kv rueidis.Client
}

func NewItemStateController(kv rueidis.Client) ItemStateController {
	return ItemStateController{kv: kv}
}

func (c ItemStateController) Create(
	userUuid string,
	itemUuid string,
	amount uint,
) error {
	item := models.NewItemState(amount, userUuid, itemUuid)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	push := c.kv.B().Rpush().Key(userUuid).Element(item.Uuid).Build()
	create := c.kv.B().JsonSet().Key(item.Uuid).Path("$").Value(
		rueidis.JSON(item),
	).Build()
	ress := c.kv.DoMulti(ctx, push, create)
	for i := 0; i < 2; i++ {
		if err := ress[i].Error(); err != nil {
			return err
		}
	}
	return nil
}

func (c ItemStateController) ReadAllUserItems(
	userUuid string,
	ptr *[]*[]models.ItemState,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	pull := c.kv.B().Lrange().Key(userUuid).Start(0).Stop(-1).Build()
	res := c.kv.Do(ctx, pull)
	cancel()
	if err := res.Error(); err != nil {
		return err
	}
	uuids, err := res.AsStrSlice()
	if err != nil {
		return err
	}

	len := len(uuids)
	cmds := make(rueidis.Commands, 0, len)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	for i := 0; i < len; i++ {
		cmds = append(cmds, c.kv.B().JsonGet().Key(uuids[i]).Path("$").Build())
	}
	ress := c.kv.DoMulti(ctx, cmds...)
	cancel()
	for i := 0; i < len; i++ {
		if err := ress[i].Error(); err != nil {
			return err
		}

		tmp := []models.ItemState{}
		err := ress[i].DecodeJSON(&tmp)
		if err != nil {
			return err
		}
		*ptr = append(*ptr, &tmp)
	}

	return nil
}

func (c ItemStateController) UpdateAmount(
	uuid string,
	amount uint,
	ptr *[]uint,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cmd := c.kv.B().JsonNumincrby().Key(uuid).Path("$.amount").Value(
		float64(amount),
	).Build()
	res := c.kv.Do(ctx, cmd)
	if err := res.Error(); err != nil {
		return err
	}

	err := res.DecodeJSON(&ptr)
	return err
}

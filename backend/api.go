package main

import (
	"github.com/gofiber/fiber"
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
)

func ListFloorsInAHole(c *fiber.Ctx) error {
	// validate
	holeID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	var query ListModel
	err = common.ValidateQuery(c, &query)
	if err != nil {
		return err
	}

	// get floors
	var floors Floors
	// use ranking field to locate faster
	querySet, err := floors.MakeQuerySet(&holeID, &query.Offset, &query.Size, c)
	if err != nil {
		return err
	}
	result := querySet.Order(query.OrderBy + " " + query.Sort).
		Find(&floors)
	if result.Error != nil {
		return result.Error
	}

	return floors.Preprocess(c)
}

func (floors Floors) MakeQuerySet(holeID *int, offset, size *int, c *fiber.Ctx) (*gorm.DB, error) {
	if holeID != nil {
		querySet = DB.Preload("Mention").Where("hole_id = ?", holeID)
	}

	if offset != nil {
		querySet = querySet.Offset(*offset)
	}
	if size != nil {
		querySet = querySet.Limit(*size)
	}
	return querySet, nil
}

func (floors Floors) Preprocess(c *fiber.Ctx) (err error) {
	userID, err := common.GetUserID(c)
	if err != nil {
		return
	}

	// get floors' like
	err = floors.loadFloorLikes(c)
	if err != nil {
		return
	}

	// set floors IsMe
	for _, floor := range floors {
		floor.IsMe = userID == floor.UserID
	}

	// set some default values
	for _, floor := range floors {
		err = floor.SetDefaults(c)
		if err != nil {
			return
		}
	}
	return
}

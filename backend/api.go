package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"slices"
	"strconv"
)

func ListFloorsInASpecialHole(c *gin.Context) {
	var err error
	idString := c.Param("holeID")
	holeID, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid hole holeID: %s", idString)
		return
	}
	if !slices.Contains(Config.SpecialHoleIDs, holeID) {
		c.String(http.StatusForbidden, "hole %d is not a special hole", holeID)
		return
	}

	var query ListModel
	if err := c.ShouldBindQuery(&query); err != nil {
		c.String(http.StatusBadRequest, "invalid query: %v", err)
		return
	}
	// get floors
	var floors Floors
	// use ranking field to locate faster
	querySet, err := floors.MakeQuerySet(&holeID, &query.Offset, &query.Size)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to make query set: %v", err)
		return
	}

	validSorts := map[string]bool{"asc": true, "desc": true, "": true}
	validOrderBy := map[string]bool{"id": true, "like": true, "": true}

	if !validSorts[query.Sort] || !validOrderBy[query.OrderBy] {
		c.String(http.StatusBadRequest, "invalid sort or order_by")
		return
	}

	result := querySet.Order(fmt.Sprintf("`%s` %s", query.OrderBy, query.Sort)).
		Find(&floors)
	if result.Error != nil {
		c.String(http.StatusInternalServerError, "failed to query floors: %v", result.Error)
		return
	}

	c.JSON(http.StatusOK, floors)
}

func (floors Floors) MakeQuerySet(holeID *int, offset, size *int) (*gorm.DB, error) {
	var querySet *gorm.DB
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

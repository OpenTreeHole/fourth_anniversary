package main

import (
	"time"
)

type ListModel struct {
	Size    int    `json:"size" query:"size" form:"size" default:"30" validate:"min=0,max=50"`              // length of object array
	Offset  int    `json:"offset" query:"offset" form:"offset" default:"0" validate:"min=0"`                // offset of object array
	Sort    string `json:"sort" query:"sort" form:"sort" default:"asc" validate:"oneof=asc desc"`           // Sort order
	OrderBy string `json:"order_by" query:"order_by" form:"order_by" default:"id" validate:"oneof=id like"` // SQL ORDER BY field
}
type Floor struct {
	/// saved fields
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"time_created"`
	UpdatedAt time.Time `json:"time_updated"`

	/// base info

	// content of the floor, no more than 15000, should be sensitive checked, no more than 10000 in frontend
	Content string `json:"content"`

	// a random username
	Anonyname string `json:"anonyname"`

	// the ranking of this floor in the hole
	Ranking int `json:"ranking"`

	// floor_id that it replies to, for dialog mode, in the same hole
	ReplyTo int `json:"reply_to"`

	// like number
	Like int `json:"like"`

	// dislike number
	Dislike int `json:"dislike"`

	// whether the floor is deleted
	Deleted bool `json:"deleted"`

	// the modification times of floor.content
	Modified int `json:"modified"`

	// fold reason
	Fold string `json:"fold_v2"`

	// additional info, like "树洞管理团队"
	SpecialTag string `json:"special_tag"`

	// the user who wrote it
	UserID int `json:"-"`

	// the hole it belongs to
	HoleID int `json:"hole_id"`

	// many to many mentions
	Mention Floors `json:"mention" gorm:"many2many:floor_mention;"`
}

func (Floor) TableName() string {
	return "floor"
}

type Floors []*Floor

type FloorMention struct {
	FloorID   int `json:"floor_id"`
	MentionID int `json:"mention_id"`
}

func (FloorMention) TableName() string {
	return "floor_mention"
}

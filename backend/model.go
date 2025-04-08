package main

type ListModel struct {
	Size    int    `json:"size" query:"size" default:"30" validate:"min=0,max=50"`          // length of object array
	Offset  int    `json:"offset" query:"offset" default:"0" validate:"min=0"`              // offset of object array
	Sort    string `json:"sort" query:"sort" default:"asc" validate:"oneof=asc desc"`       // Sort order
	OrderBy string `json:"order_by" query:"order_by" default:"id" validate:"oneof=id like"` // SQL ORDER BY field
}
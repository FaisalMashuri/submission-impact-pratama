package request

type ProductRequest struct {
	CodeProduct string `json:"codeProduct" form:"codeProduct"`
	NameProduct string `json:"nameProduct" form:"nameProduct"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	UOM         string `json:"uom" form:"uom"`
}

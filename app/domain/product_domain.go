package domain

import (
	"github.com/google/uuid"
)

type ProductDomain struct {
	Id          uuid.UUID
	CodeProduct string
	NameProduct string
	Description string
	Price       int
	UOM         string
}

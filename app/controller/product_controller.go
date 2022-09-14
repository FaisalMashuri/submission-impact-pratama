package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/FaisalMashuri/submission-golang/app/domain"
	"github.com/FaisalMashuri/submission-golang/app/request"
	"github.com/FaisalMashuri/submission-golang/app/response"
	"github.com/FaisalMashuri/submission-golang/app/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	UseCase usecase.ProductUsecase
}

func NewProductController(u usecase.ProductUsecase) *ProductController {
	return &ProductController{
		UseCase: u,
	}
}

func (p *ProductController) CreateProduct(c echo.Context) error {
	ctx := c.Request().Context()
	createdProduct := request.ProductRequest{}
	c.Bind(&createdProduct)
	uom := strings.ToUpper(createdProduct.UOM)
	switch uom {
	case "PCS":
		createdProduct.UOM = uom
	case "ROLL":
		createdProduct.UOM = uom
	case "SHEET":
		createdProduct.UOM = uom
	default:
		defer c.Logger().Error(fmt.Errorf("UOM tidak sesuai"))
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("UOM tidak ada"))
	}
	id := uuid.New()
	productDomain := domain.ProductDomain{
		Id:          id,
		CodeProduct: createdProduct.CodeProduct,
		NameProduct: createdProduct.NameProduct,
		Description: createdProduct.Description,
		Price:       createdProduct.Price,
		UOM:         createdProduct.UOM,
	}

	product, err := p.UseCase.CreateProduct(ctx, productDomain)
	if err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, product)
}

func (p *ProductController) GetAllProduct(c echo.Context) error {
	//TODO implement me
	ctx := c.Request().Context()

	product, err := p.UseCase.GetAllProduct(ctx)
	if err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, product)
}

func (p *ProductController) GetByID(c echo.Context) error {
	//TODO implement me
	ctx := c.Request().Context()
	codeProduct := c.Param("codeProduct")
	product, err := p.UseCase.GetProductByCodeProduct(ctx, codeProduct)
	if err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, product)
}

func (p *ProductController) UpdateProduct(c echo.Context) error {
	//TODO implement me
	ctx := c.Request().Context()
	id := c.Param("id")
	updatedProduct := request.ProductRequest{}
	c.Bind(&updatedProduct)
	productDomain := domain.ProductDomain{
		CodeProduct: updatedProduct.CodeProduct,
		NameProduct: updatedProduct.NameProduct,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		UOM:         updatedProduct.UOM,
	}
	product, err := p.UseCase.UpdateProduct(ctx, id, productDomain)
	if err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, product)
}

func (p *ProductController) DeleteProduct(c echo.Context) error {
	//TODO implement me
	ctx := c.Request().Context()
	id := c.Param("id")
	if err := p.UseCase.DeleteProduct(ctx, id); err != nil {
		defer c.Logger().Error(fmt.Errorf(err.Error()))
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, fmt.Sprintf("Data dengan id %s berhasil dihapus", id))
}

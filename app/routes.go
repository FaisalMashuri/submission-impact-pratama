package app

import (
	"net/http"

	"github.com/FaisalMashuri/submission-golang/app/controller"
	"github.com/labstack/echo/v4"
)

type RouteParams struct {
	Productcontroller *controller.ProductController
}

func (p RouteParams) Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hai")
	})
	e.POST("/product", p.Productcontroller.CreateProduct)
	e.GET("/product", p.Productcontroller.GetAllProduct)
	e.GET("/product/:codeProduct", p.Productcontroller.GetByID)
	e.PUT("/product/:id", p.Productcontroller.UpdateProduct)
	e.DELETE("/product/:id", p.Productcontroller.DeleteProduct)
}

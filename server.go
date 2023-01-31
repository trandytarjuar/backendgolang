config.ConnectDB()
route := echo.New()

route.GET("product/search_product", func(c echo.Context) error {
	response := new(Response)
	product, err := model.GetAll(c.QueryParam("keywords")) // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal melihat data product"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses melihat data product"
		response.Data = product
	}
	return c.JSON(http.StatusOK, response)
})
route.Start(":9000")
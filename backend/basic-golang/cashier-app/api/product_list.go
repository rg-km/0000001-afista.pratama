package api

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
>>>>>>> 709389bab4a4c744b6239109076774a155e4eb5a
	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
<<<<<<< HEAD
	encoder := json.NewEncoder(w)

	response := ProductListSuccessResponse{}
	response.Products = make([]Product, 0)

	products, err := api.productsRepo.SelectAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(DashboardErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	fmt.Println(products)

	encoder.Encode(ProductListSuccessResponse{Products: []Product{}}) // TODO: replace this
=======
	// encoder.Encode(ProductListSuccessResponse{Products: []Product{}}) // TODO: replace this
>>>>>>> 709389bab4a4c744b6239109076774a155e4eb5a
}

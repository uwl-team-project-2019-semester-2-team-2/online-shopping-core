package services

import (
	"fmt"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
	"image/draw"
	"log"
	"strconv"
)

/**
Structures
*/
type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Brand       Brand  `json:"brand"`
	Colour      string `json:"colour"`
	Description string `json:"description"`
	Variants    []struct {
		Id       string `json:"id"`
		Title    string `json:"title"`
		Quantity uint32 `json:"quantity"`
		Stock    []struct {
			Id string `json:"id"`
		} `json:"stock"`
	} `json:"variants"`
	Images  []draw.Image `json:"images"`
	Related []Product    `json:"related"`
}

/**
REST Methods
*/
func GetProduct(r typhon.Request) typhon.Response {
	response := typhon.NewResponse(r)

	productIdStr, ok := typhon.RouterForRequest(r).Params(r)["productId"]
	productId, err := strconv.ParseInt(productIdStr, 10, 32)

	if err != nil {
		log.Print("error parsing productId: " + err.Error())
		response.Error = terrors.InternalService("", "error parsing product id", map[string]string{
			"error": err.Error(),
		})
		return response
	}

	if !ok {
		response.Error = terrors.InternalService("missing_parameter", "ProductID parameter missing in request", nil)
		return response
	}

	log.Print(fmt.Sprintf("processing get request for product %d", productId))

	return typhon.NewResponse(r)
}

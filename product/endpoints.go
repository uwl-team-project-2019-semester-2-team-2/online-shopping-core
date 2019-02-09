package product

import (
	"encoding/json"
	"fmt"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
	"log"
	"strconv"
)

func Routes(router *typhon.Router) {
	router.GET("/product/:productId", Get)
	router.PUT("/product", Put)
}

func Get(r typhon.Request) typhon.Response {
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

func Put(r typhon.Request) typhon.Response {
	productBytes, err := r.BodyBytes(true)

	if err != nil {
		log.Println("error reading body bytes: " + err.Error())
		return r.Response(err)
	}

	newProduct := &Product{}
	err = json.Unmarshal(productBytes, newProduct)

	if err != nil {
		log.Println("error unmarshalling json body: " + err.Error())
		return r.Response(err)
	}

	// TODO: Save new product?

	return r.Response(newProduct)
}

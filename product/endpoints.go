package product

import (
	"fmt"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"log"
	"strconv"
)

func Routes(router *typhon.Router) {
	router.GET("/product/:productId", Get)
	router.PUT("/product", Put)
}

func Get(r typhon.Request) typhon.Response {
	productIdStr, ok := typhon.RouterForRequest(r).Params(r)["productId"]
	productId, err := strconv.ParseInt(productIdStr, 10, 32)

	if err != nil {
		log.Print("error parsing productId: " + err.Error())
		return r.Response(terrors.InternalService("", "error parsing product id", map[string]string{
			"error": err.Error(),
		}))
	}

	if !ok {
		return r.Response(terrors.InternalService("missing_parameter", "ProductID parameter missing in request", nil))
	}

	log.Print(fmt.Sprintf("processing get request for product %d", productId))

	result := database.DB.QueryRow("SELECT product.id, line.name, line.description FROM product LEFT JOIN product_line line ON (product.product_line_id = line.id) WHERE product.id=?", productId)

	product := Product{}
	err = result.Scan(&product.Id, &product.Name, &product.Description)

	if err != nil {
		return r.Response(terrors.NotFound("product_not_found", "The requested product could not be found.", map[string]string{
			"error": err.Error(),
		}))
	}

	return r.Response(product)
}

func Put(r typhon.Request) typhon.Response {
	newProduct := &Product{}
	err := r.Decode(newProduct)

	if err != nil {
		return r.Response(err)
	}

	// TODO: Save new product?

	return r.Response(newProduct)
}

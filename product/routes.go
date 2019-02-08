package product

import "github.com/monzo/typhon"

func Routes(router *typhon.Router) {
	router.GET("/product/:productId", Get)
}

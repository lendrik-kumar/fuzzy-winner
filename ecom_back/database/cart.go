package database

import "errors"

var (
	ErrCantFindProduct = errors.New("Can't find product")
	ErrCantDecodeProduct = errors.New("Can't decode product")
	ErrUserIdIsNotValid = errors.New("User ID is not valid")
	ErrCantUpdateProduct = errors.New("Can't update product")
	ErrCantRemoveItemCart = errors.New("Can't remove item from cart")
	ErrCantGetItem = errors.New("Can't get item")
	ErrCantBuyCartItem = errors.New("Can't buy cart item")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}

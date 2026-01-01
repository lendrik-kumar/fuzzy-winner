package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id"           bson:"_id"`
	FirstName      *string            `json:"first_name"    bson:"first_name"     validate:"required min=2 max=30"`
	LastName       *string            `json:"last_name"     bson:"last_name"      validate:"required min=2 max=30"`
	Phone          *string            `json:"phone"         bson:"phone"          validate:"required,min=10,max=15"`
	Email          *string            `json:"email"         bson:"email"          validate:"email,required"`
	Password       *string            `json:"password"      bson:"password"       validate:"required,min=6"`
	Token          *string            `json:"token"         bson:"token"          validate:"required"`
	RefreshToken   *string            `json:"refresh_token" bson:"refresh_token"`
	CreatedAt      time.Time          `json:"created_at"    bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"    bson:"updated_at"`
	UserId         string             `json:"user_id"       bson:"user_id"`
	UserCart       []ProductUser      `json:"usercart"      bson:"usercart"`
	AddressDetails []Address          `json:"address"       bson:"addresss"`
	OrderStatus    []Order            `json:"orders"        bson:"orders"`
}

type Product struct {
	ProductId   primitive.ObjectID `json:"_id"          bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name"`
	Price       *uint64            `json:"price"        bson:"price"`
	Rating      *uint8             `json:"rating"       bson:"rating"`
	Image       *string            `json:"image"        bson:"image"`
} 

type ProductUser struct {
	ProductId   primitive.ObjectID `json:"_id"          bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name"`
	Price       *uint64            `json:"price"        bson:"price"`
	Rating      *uint8             `json:"rating"       bson:"rating"`
	Image       *string            `json:"image"        bson:"image"`
}

type Address struct {
	AddressId primitive.ObjectID `json:"_id"          bson:"_id"`
	House     *string            `json:"house_name"   bson:"house_name"`
	Street    *string            `json:"street_name"  bson:"street_name"`
	City      *string            `json:"city_name"    bson:"city_name"`
	Pincode   *string            `json:"pin_code"     bson:"pin_code"`
}

type Order struct {
	OrderId       primitive.ObjectID `json:"_id"            bson:"_id"`
	OrderCart     []ProductUser      `json:"order_list"     bson:"order_list"`
	OrderedAt     time.Time          `json:"ordered_at"     bson:"ordered_at"`
	Price         uint64             `json:"total_price"    bson:"total_price"`
	Discount      uint64             `json:"discount"       bson:"discount"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}

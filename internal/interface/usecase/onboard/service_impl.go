package onboard

import (
	"context"
	"time"

	"gokomodo-test/internal/domain/buyer"
	"gokomodo-test/internal/domain/seller"
	"gokomodo-test/pkg/config"
	"gokomodo-test/pkg/response"
	"gokomodo-test/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	buyerRepo  buyer.BuyerRepository
	sellerRepo seller.SellerRepository
}

func NewService(
	buyerRepo buyer.BuyerRepository,
	sellerRepo seller.SellerRepository,
) *service {
	// to validate that repository is properly injected
	if buyerRepo == nil {
		panic("buyer repository is nil")
	}

	if sellerRepo == nil {
		panic("seller repository is nil")
	}

	return &service{
		buyerRepo:  buyerRepo,
		sellerRepo: sellerRepo,
	}
}

func (s *service) Login(ctx context.Context, loginData Login) (resp response.DefaultResponse, err error) {
	if loginData.Type != "buyer" && loginData.Type != "seller" {
		resp = response.ErrorResponse(response.CODE_BAD_REQUEST, "must be either buyer or seller")
		return
	}
	var userId string
	var password string
	if loginData.Type == "buyer" {
		var buyer buyer.Buyer
		buyer, err = s.buyerRepo.FindByEmail(ctx, loginData.Email)
		userId = buyer.ID
		password = buyer.Password
	} else if loginData.Type == "seller" {
		var seller seller.Seller
		seller, err = s.sellerRepo.FindByEmail(ctx, loginData.Email)
		userId = seller.ID
		password = seller.Password
	}

	if err != nil && err.Error() != "no rows in result set" {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}
	if userId == "" {
		resp = response.ErrorResponse(response.CODE_BAD_REQUEST, "user not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(loginData.Password))

	if err != nil {
		resp = response.ErrorResponse(response.CODE_BAD_REQUEST, "invalid password")
		return
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Set expiration time to 24 hours from now
		Id:        userId,
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(config.GetEnvString("JWT_SECRET_KEY")))

	if err != nil {
		utils.PrintError(err)
		resp = response.ErrorServerResponse()
		return
	}

	respData := Loginresp{
		Token: tokenString,
	}
	resp = response.CreateResponse(response.CODE_SUCCESS, response.MESSAGE_SUCCESS, respData)
	return
}

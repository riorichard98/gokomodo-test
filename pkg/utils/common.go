package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func PrintError(err error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	errorMessage := strings.ToUpper(err.Error())
	fmt.Printf("[%s] %s\n", timestamp, errorMessage)
}

func Validate(c echo.Context, body interface{}) (err error) {
	// Bind the request body to the struct
	if err = c.Bind(body); err != nil {
		PrintError(err)
		return
	}

	// Validate the request body using the Echo validator middleware
	if err = c.Validate(body); err != nil {
		PrintError(err)
		return
	}

	return nil
}

// Define a struct to match the structure of the token JSON
type TokenInfo struct {
	Raw    string `json:"Raw"`
	Method struct {
		Name string `json:"Name"`
		Hash int    `json:"Hash"`
	} `json:"Method"`
	Header    map[string]interface{} `json:"Header"`
	Claims    jwt.StandardClaims     `json:"Claims"`
	Signature string                 `json:"Signature"`
	Valid     bool                   `json:"Valid"`
}

func ClaimJWT(jwtData interface{}) (userId string) {
	var tokenInfo TokenInfo

	strJwtData, _ := json.Marshal(jwtData)

	// Unmarshal the JSON string into the struct
	err := json.Unmarshal([]byte(strJwtData), &tokenInfo)
	if err != nil {
		fmt.Println("error getting jwt data")
	}

	userId = tokenInfo.Claims.Id
	return
}

func IsValidUUID(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}

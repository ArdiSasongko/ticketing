package helper

import (
	"errors"
	"github.com/ArdiSasongko/ticketing_app/cfg"
	"github.com/ArdiSasongko/ticketing_app/model/enum"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func GetAuthId(c echo.Context) (int, error) {
	claims := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	authID, err := strconv.Atoi(claims.ID)

	if err != nil {
		return -1, err
	}

	return authID, nil
}

func Login(userType enum.UserType, userId int, userEmail string) (map[string]interface{}, error) {
	var tokenDuration time.Duration
	if userType == enum.UserTypeBuyer {
		tokenDuration = time.Duration(cfg.BuyerTokenExpiryHour) * time.Hour
	} else if userType == enum.UserTypeSeller {
		tokenDuration = time.Duration(cfg.SellerTokenExpiryHour) * time.Hour
	} else {
		tokenDuration = time.Duration(cfg.AdminTokenExpiryHour) * time.Hour
	}
	expiredTime := time.Now().Local().Add(tokenDuration)

	claims := JwtCustomClaims{
		ID:    strconv.Itoa(userId),
		Email: userEmail,
		Role:  string(userType),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.AppName,
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, errToken := NewTokenUseCase().GenerateAccessToken(claims)
	if errToken != nil {
		return nil, errors.New("internal server error")
	}

	return map[string]interface{}{
		"token":      token,
		"expires_at": expiredTime.Format(time.RFC3339),
	}, nil
}

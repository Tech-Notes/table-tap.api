package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/table-tap/api/internal/types"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponseData struct {
	Token string `json:"token"`
}

type SignInSuccessResponse struct {
	*types.ResponseBase
	Data *SignInResponseData `json:"data"`
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	payload := SignInRequest{}
	err := readJSON(r, &payload)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()
	businessUser, err := DBConn.GetLastActiveBusinessUserByEmail(ctx, payload.Email)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if businessUser != nil && businessUser.ID == 0 {
		writeError(w, http.StatusUnauthorized, ErrUserNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(businessUser.Password), []byte(payload.Password))
	if err != nil {
		writeError(w, http.StatusUnauthorized, ErrInvalidCredentials)
		return
	}

	token, err := generateToken(businessUser)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, SignInSuccessResponse{
		ResponseBase: &types.SuccessResponse,
		Data: &SignInResponseData{
			Token: token,
		},
	})
}

func generateToken(businessUser *types.BusinessUser) (string, error) {
	// create claims
	claims := jwt.MapClaims{
		"user_id":     businessUser.ID,
		"user_email":  businessUser.Email,
		"business_id": businessUser.BusinessID,
		"role":        businessUser.Role,
		"role_id":     businessUser.RoleID,
	}

	tokenSecretKey := os.Getenv("TOKEN_SECRET_KEY")

	// generate JWT token with tihs claims using jtw-go
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(tokenSecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

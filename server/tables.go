package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	internal "github.com/table-tap/api/internal/types"
	types "github.com/table-tap/api/internal/types"
	utils "github.com/table-tap/api/internal/utils"
)

func CreateTableHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	businessID := BusinessIDFromContext(ctx)
	secureToken := uuid.New().String()

	// Create order URL for QR
	orderURL := "https://ordertap.com/order/" + secureToken

	// Generate QR code image
	qrPNG, err := qrcode.Encode(orderURL, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Failed to generate QR", http.StatusInternalServerError)
		return
	}

	// Upload QR image to S3
	fileName := secureToken + ".png"
	imageURL, err := utils.UploadToS3(r.Context(), businessID, qrPNG, fileName, "table_qr")
	if err != nil {
		http.Error(w, "Failed to upload QR to S3", http.StatusInternalServerError)
		return
	}

	// Store in database (replace this with your actual DB logic)
	table := &types.Table{
		BusinessID: businessID,
		QrCodeURL:  imageURL,
		Status:     "active",
	}

	_, err = DBConn.CreateTable(ctx, table)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, internal.SuccessResponse)

}

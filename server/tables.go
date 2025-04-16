package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	types "github.com/table-tap/api/internal/types"
	utils "github.com/table-tap/api/internal/utils"
)

func CreateTableHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	businessID := utils.BusinessIDFromContext(ctx)
	secureToken := uuid.New().String()

	qrURL := "https://ordertap.com/order/" + secureToken

	// Generate QR code image
	qrPNG, err := qrcode.Encode(qrURL, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Failed to generate QR", http.StatusInternalServerError)
		return
	}

	// Upload QR image to S3
	fileName := secureToken + ".png"
	imageURL, err := utils.UploadToS3(ctx, businessID, qrPNG, fileName, "table_qrs")
	if err != nil {
		http.Error(w, "Failed to upload QR to S3", http.StatusInternalServerError)
		return
	}

	// Store in database (replace this with your actual DB logic)
	table := &types.Table{
		BusinessID: businessID,
		QrCodeURL:  imageURL,
		Status:     "active",
		Token:      secureToken,
	}

	id, err := DBConn.CreateTable(ctx, table)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusCreated, types.ActionSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &types.ActionSuccessResponseData{
			ID: id,
		},
	})

}

type GetTablesResponse struct {
	Tables []*types.Table `json:"tables"`
}

type GetTablesSuccessResponse struct {
	types.ResponseBase
	Data *GetTablesResponse `json:"data"`
}

func GetTablesHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	businessID := utils.BusinessIDFromContext(ctx)

	tables, err := DBConn.GetTableList(ctx, businessID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, GetTablesSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &GetTablesResponse{
			Tables: tables,
		},
	})
}

func GetTableByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	businessID := utils.BusinessIDFromContext(ctx)

	tableIDSring := chi.URLParam(r, "id")
	tableID, err := strconv.ParseInt(tableIDSring, 10, 64)

	if err != nil {
		writeError(w, http.StatusBadRequest, ErrRequiredTableID)
		return
	}

	table, err := DBConn.GetTableByID(ctx, businessID, tableID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, types.TableDetailSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &types.TableDetailResponse{
			Table: table,
		},
	})
}

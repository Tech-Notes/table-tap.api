package shopper

import "net/http"

type ValidateTableRequest struct {
	TableToken string `json:"tableToken"`
}

type ValidateTableResponse struct {
	Valid      bool   `json:"valid"`
	ShopID     int64  `json:"shopId,omitempty"`
	TableID    int64  `json:"tableId,omitempty"`
	TableToken string `json:"tableToken,omitempty"`
}

func ValidateTableHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := &ValidateTableRequest{}
	err := readJSON(r, data)
	if err != nil {
		writeError(w, http.StatusInternalServerError, ErrFailedRequestBody)
		return
	}

	table, err := DBConn.GetTableByToken(ctx, data.TableToken)
	if err != nil {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"valid": false,
		})
		return
	}

	writeJSON(w, http.StatusOK, ValidateTableResponse{
		Valid:      true,
		ShopID:     table.BusinessID,
		TableID:    table.ID,
		TableToken: table.Token,
	})
}

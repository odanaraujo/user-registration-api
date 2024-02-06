package userhandler

import (
	"encoding/json"
	"fmt"
	"github.com/odanaraujo/user-api/internal/dto"
	"github.com/odanaraujo/user-api/internal/handler/httperr"
	"github.com/odanaraujo/user-api/internal/handler/validation"
	"log/slog"
	"net/http"
)

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := dto.CreateUserDto{}
	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userHandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is empty")
		json.NewEncoder(w).Encode(msg)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	if httperr := validation.ValidateHttpData(req); httperr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httperr), slog.String("package", "handler_user"))
		w.WriteHeader(httperr.Code)
		json.NewEncoder(w).Encode(httperr)
		return
	}

	if err := h.service.CreateUser(r.Context(), req); err != nil {
		slog.Error(fmt.Sprintf("error to create user: %v", err), slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to create user")
		json.NewEncoder(w).Encode(msg)
		return
	}
}

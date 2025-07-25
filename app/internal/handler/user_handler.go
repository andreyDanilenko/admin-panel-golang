package handler

import (
	"encoding/json"
	"net/http"

	"admin/panel/internal/contract"
	"admin/panel/internal/middleware"
	"admin/panel/internal/model"
	"admin/panel/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	authService  *service.UserService
	validate     *validator.Validate
	errorWriter  contract.ErrorWriter
	responseJSON contract.ResponseWriter
}

func NewUserHandler(
	authService *service.UserService,
	ew contract.ErrorWriter,
	rw contract.ResponseWriter,
) *UserHandler {
	return &UserHandler{
		authService:  authService,
		validate:     validator.New(),
		errorWriter:  ew,
		responseJSON: rw,
	}
}

func (h *UserHandler) StartAuthFlow(w http.ResponseWriter, r *http.Request) {
	var input model.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorWriter.WriteWithCode(w, http.StatusBadRequest, "bad_request", "Неверный формат", nil)
		return
	}

	if err := h.validate.Struct(input); err != nil {
		h.errorWriter.WriteWithCode(w, http.StatusBadRequest, "validation_failed", "Ошибка валидации", err)
		return
	}

	user, err := h.authService.StartAuthFlow(r.Context(), input)
	if err != nil {
		h.errorWriter.WriteWithCode(w, http.StatusUnauthorized, "unauthorized", err.Error(), nil)
		return
	}

	h.responseJSON.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Код отправлен на email",
		"user":    user,
	})
}

func (h *UserHandler) ConfirmCode(w http.ResponseWriter, r *http.Request) {
	var input *model.ConfirmCodeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorWriter.WriteWithCode(w, http.StatusBadRequest, "bad_request", "Неверный формат", nil)
		return
	}

	if err := h.validate.Struct(input); err != nil {
		h.errorWriter.WriteWithCode(w, http.StatusBadRequest, "validation_failed", "Ошибка валидации", err)
		return
	}

	user, token, err := h.authService.ConfirmCode(r.Context(), input.Email, input.Code)
	if err != nil {
		h.errorWriter.WriteWithCode(w, http.StatusUnauthorized, "unauthorized", err.Error(), nil)
		return
	}

	h.responseJSON.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.authService.GetUsers(r.Context())
	if err != nil {
		h.errorWriter.WriteError(w, http.StatusInternalServerError, "Не удалось получить пользователей")
		return
	}

	h.responseJSON.WriteJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.authService.GetUserByID(r.Context(), id)
	if err != nil {
		h.errorWriter.WriteWithCode(w, http.StatusNotFound, "not_found", "Пользователь не найден", nil)
		return
	}

	h.responseJSON.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	authUserID, _ := r.Context().Value(middleware.UserIDKey).(string)
	role, _ := r.Context().Value(middleware.RoleKey).(model.UserRole)

	if authUserID != userID && role != model.RoleAdmin {
		h.errorWriter.WriteWithCode(w, http.StatusForbidden, "forbidden", "Другого пользователя может редактировать только администратор", nil)
		return
	}

	var input model.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorWriter.WriteWithCode(w, http.StatusBadRequest, "invalid_body", "Неверный формат запроса", nil)
		return
	}

	user, err := h.authService.UpdateUser(r.Context(), userID, input)
	if err != nil {
		h.errorWriter.WriteError(w, http.StatusInternalServerError, "Не удалось обновить пользователя")
		return
	}

	h.responseJSON.WriteJSON(w, http.StatusOK, user)
}

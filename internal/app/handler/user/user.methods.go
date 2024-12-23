package user

import (
	"encoding/json"
	"net/http"
	"siap_app/internal/app/entity"
	"siap_app/internal/app/entity/user"
	"siap_app/internal/app/helpers"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input user.RegisterRequest
	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failled umarshal data", err.Error())
		return
	}

	err = h.userUC.CreateUser(ctx, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "User created successfully", http.StatusCreated)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input user.LoginRequest

	ipAddress := r.Header.Get("X-Forwarded-For")
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}

	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failled umarshal data", err.Error())
		return
	}

	resp, err := h.userUC.LoginUser(ctx, ipAddress, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, resp, "login user successfully", http.StatusOK)
}

func (h *Handler) CreateUserByAdmin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input user.RegisterByAdminRequest
	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failled umarshal data", err.Error())
		return
	}

	err = h.userUC.CreateUserByAdmin(ctx, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "User created successfully", http.StatusCreated)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email, ok := r.Context().Value(entity.EmailKey).(string)
	if !ok || email == "" {
		helpers.SendUnauthorizedResponse(w)
		return
	}

	err := h.userUC.LogoutUser(ctx, email)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "Logout user successfully", http.StatusOK)
}

func (h *Handler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input user.UpdateRoleRequest
	tokenData, ok := ctx.Value("user").(entity.TokenData)
	if !ok {
		helpers.SendUnauthorizedResponse(w)
		return
	}

	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failled umarshal data", err.Error())
		return
	}

	err = h.userUC.UpdateRoleUser(ctx, tokenData.UserId, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "role has been updated", http.StatusOK)
}

func (h *Handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input user.UpdatePaswordeRequest
	tokenData, ok := ctx.Value("user").(entity.TokenData)
	if !ok {
		helpers.SendUnauthorizedResponse(w)
		return
	}

	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failled umarshal data", err.Error())
		return
	}

	err = h.userUC.UpdatePasswordUser(ctx, tokenData.UserId, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	helpers.SendSuccessResponse(w, nil, "role has been updated", http.StatusOK)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input entity.Pagination

	dataReq, err := helpers.GetInputDataRequest(r)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "error encrypt data", err.Error())
		return
	}

	err = json.Unmarshal(dataReq, &input)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "failled umarshal data", err.Error())
		return
	}

	res, total, err := h.userUC.GetUsers(ctx, input)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	resp := entity.ResponsePagination{
		Total: total,
		Data:  res,
	}

	helpers.SendSuccessResponse(w, resp, "success retrieve user data", http.StatusOK)
}

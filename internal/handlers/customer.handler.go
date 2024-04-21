package handlers

import (
	"customer-data-api/internal/models"
	"customer-data-api/internal/services"
	"customer-data-api/pkg"
	"net/http"
)

type customerHandler struct {
	serv services.CustomerService
}

func NewCustomerHandler(serv services.CustomerService) *customerHandler {
	return &customerHandler{serv: serv}
}

func (h *customerHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := models.CreateCustomerReq{}
	if err := pkg.GetJsonBody(r, &req); err != nil {
		pkg.WriteJsonError(w, err)
		return
	}

	res, err := h.serv.Create(&req)
	if err != nil {
		pkg.WriteJsonError(w, err)
		return
	}

	pkg.WriteJson(w, http.StatusOK, res)
}

func (h *customerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	custRes, err := h.serv.GetAllCustomer()
	if err != nil {
		pkg.WriteJsonError(w, err)
		return
	}

	pkg.WriteJson(w, http.StatusOK, custRes)
}

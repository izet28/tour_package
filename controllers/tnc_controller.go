package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/izet28/tour_package/models"
	"github.com/izet28/tour_package/services"
	"github.com/izet28/tour_package/utils"
	"gorm.io/gorm"
)

// TncController handles tour package related HTTP requests
type TncController struct {
	Service services.TncService
}

// NewTncController creates a new TncController
func NewTncController(db *gorm.DB) *TncController {
	return &TncController{
		Service: services.NewTncService(db),
	}
}

// GetAll handles the request to get all tour packages
func (c *TncController) GetAll(w http.ResponseWriter, r *http.Request) {
	packages, err := c.Service.GetAll()
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, packages)
}

// GetByID handles the request to get a tour package by ID
func (c *TncController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid package ID")
		return
	}

	paket, err := c.Service.GetByID(id)
	if err != nil {
		utils.RespondJSON(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, paket)
}

// Create handles the request to create a new tour package
func (c *TncController) Create(w http.ResponseWriter, r *http.Request) {
	var paket models.Tnc
	if err := json.NewDecoder(r.Body).Decode(&paket); err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate input
	if err := utils.ValidateTnc(paket); err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	newPackage, err := c.Service.Create(paket)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, newPackage)
}

// Update handles the request to update a tour package
func (c *TncController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid package ID")
		return
	}

	var paket models.Tnc
	if err := json.NewDecoder(r.Body).Decode(&paket); err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate input
	if err := utils.ValidateTnc(paket); err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedPackage, err := c.Service.Update(id, paket)
	if err != nil {
		utils.RespondJSON(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, updatedPackage)
}

// Delete handles the request to delete a tour package
func (c *TncController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid package ID")
		return
	}

	err = c.Service.Delete(id)
	if err != nil {
		utils.RespondJSON(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusNoContent, nil)
}

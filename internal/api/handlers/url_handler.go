package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"urlshortener/internal/pkg/service"
	"urlshortener/internal/pkg/validator"
	"urlshortener/pkg/logger"
)

// URLHandler handles HTTP requests for URL operations
type URLHandler struct {
	service   *service.URLService
	validator *validator.URLValidator
	logger    *zap.Logger
}

// NewURLHandler creates a new instance of URLHandler
func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{
		service:   service,
		validator: validator.NewURLValidator(),
		logger:    logger.GetLogger(),
	}
}

// createURLRequest represents the payload for creating a short URL
type createURLRequest struct {
	URL string `json:"url"`
}

// CreateShortURL handles the creation of a new short URL
func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req createURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("failed to decoded request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validator.ValidateURL(req.URL); err != nil {
		h.logger.Warn("url validation failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url, err := h.service.CreateShortURL(r.Context(), req.URL)
	if err != nil {
		h.logger.Error("failed to create short url", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.logger.Info("short url created", zap.String("original_url", req.URL), zap.String("short_code", url.ShortCode))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(url)
}

// GetURL handles retrieving a URL by its short code
func (h *URLHandler) GetURL(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["shortCode"]

	url, err := h.service.GetURL(r.Context(), shortCode)
	if err != nil {
		h.logger.Warn("url not found", zap.String("short_code", shortCode), zap.Error(err))
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	h.logger.Info("url retrieved", zap.String("short_code", shortCode), zap.String("original_url", url.OriginalURL))

	json.NewEncoder(w).Encode(url)
}

// UpdateURL handles updating the original URL of an existing short code
func (h *URLHandler) UpdateURL(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["shortCode"]

	var req createURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validator.ValidateURL(req.URL); err != nil {
		h.logger.Warn("url validation failed", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url, err := h.service.UpdateURL(r.Context(), shortCode, req.URL)
	if err != nil {
		h.logger.Warn("failed to update url", zap.String("short_code", shortCode), zap.Error(err))
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	h.logger.Info("url updated", zap.String("short_code", shortCode), zap.String("new_url", req.URL))

	json.NewEncoder(w).Encode(url)
}

// DeleteURL handles deleting a URL by its short code
func (h *URLHandler) DeleteURL(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["shortCode"]

	if err := h.service.DeleteURL(r.Context(), shortCode); err != nil {
		h.logger.Warn("failed to delete url", zap.String("short_code", shortCode), zap.Error(err))
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	h.logger.Info("url deleted", zap.String("short_code", shortCode))

	w.WriteHeader(http.StatusNoContent)
}

// GetStats handles retrieving the statistics of a URL by its short code
func (h *URLHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["shortCode"]

	url, err := h.service.GetStats(r.Context(), shortCode)
	if err != nil {
		h.logger.Warn("failed to get url stats", zap.String("short_code", shortCode), zap.Error(err))
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	h.logger.Info("url stats retrieved", zap.String("short_code", shortCode), zap.Int("access_count", url.AccessCount))

	json.NewEncoder(w).Encode(url)
}

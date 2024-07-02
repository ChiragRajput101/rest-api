package cart

import (
	"fmt"
	"net/http"

	"github.com/ChiragRajput101/rest-api/services/auth"
	"github.com/ChiragRajput101/rest-api/types"
	"github.com/ChiragRajput101/rest-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
	orderStore types.OrderStore
	userStore types.UserStore
}

func NewHandler(store types.ProductStore, orderStore types.OrderStore, userStore types.UserStore) *Handler {
	return &Handler{
		store: store,
		orderStore: orderStore,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods("POST")
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {

	userID := auth.GetUserIDFromContext(r.Context())
	fmt.Println(userID)

	var payload types.CartCheckoutPayload

	// check for malformated json payload

	if err := utils.ParseJSON(r,&payload); err != nil {
		utils.WriteError(w, http.StatusBadGateway, err)
		return
	}

	// validate the payload

	if err := utils.Validate.Struct(payload); err != nil {
		errs := err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadGateway, fmt.Errorf("invalid payload %v", errs))
		return
	}

	// getting out the IDs of the products in the cart
	productIds, err := getCartItemsIDs(payload.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// finding the products against the recieved IDs from cart items
	products, err := h.store.GetProductsByID(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, payload.Items, 7)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}


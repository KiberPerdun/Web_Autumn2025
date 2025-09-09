package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"rip/internal/app/repository"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}

func (h *Handler) GetOrders(ctx *gin.Context) {
	orders, err := h.Repository.GetOrders()
	if err != nil {
		logrus.Error(err)
	}

	searchQuery := ctx.Query("query")

	var authors []repository.Author

	if searchQuery != "" {
		authors, err = h.Repository.SearchAuthors(searchQuery)
	} else {
		authors, err = h.Repository.GetAuthors()
	}

	if err != nil {
		logrus.Error(err)
	}

	cartCount := len(getCartFromSession(ctx))

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"time":      time.Now().Format("15:04:05"),
		"orders":    orders,
		"authors":   authors,
		"query":     searchQuery,
		"cartCount": cartCount,
	})
}

func (h *Handler) GetAuthorByID(ctx *gin.Context) {
	id := ctx.Param("id")

	author, err := h.Repository.GetAuthorByID(id)
	if err != nil {
		logrus.Error(err)
		ctx.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "Автор не найден",
		})
		return
	}

	cartCount := len(getCartFromSession(ctx))

	ctx.HTML(http.StatusOK, "author.html", gin.H{
		"author":    author,
		"cartCount": cartCount,
	})
}

func (h *Handler) GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")

	order, err := h.Repository.GetOrderByID(id)
	if err != nil {
		logrus.Error(err)
		ctx.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "Заказ не найден",
		})
		return
	}

	ctx.HTML(http.StatusOK, "order.html", gin.H{
		"order": order,
	})
}

func (h *Handler) GetOrderForm(ctx *gin.Context) {
	cartCount := len(getCartFromSession(ctx))

	ctx.HTML(http.StatusOK, "order.html", gin.H{
		"time":      time.Now().Format("15:04:05"),
		"cartCount": cartCount,
	})
}

func (h *Handler) SubmitOrder(ctx *gin.Context) {
	ctx.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) GetCart(ctx *gin.Context) {
	cart := getCartFromSession(ctx)
	var authors []repository.Author
	for _, id := range cart {
		author, err := h.Repository.GetAuthorByID(strconv.Itoa(id))
		if err == nil {
			authors = append(authors, *author)
		}
	}
	ctx.HTML(http.StatusOK, "cart.html", gin.H{
		"authors":   authors,
		"cartCount": len(cart),
	})
}

func (h *Handler) AddToCart(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Некорректный ID")
		return
	}
	cart := getCartFromSession(ctx)
	for _, cid := range cart {
		if cid == id {
			ctx.Redirect(http.StatusSeeOther, "/cart")
			return
		}
	}
	cart = append(cart, id)
	setCartToSession(ctx, cart)
	ctx.Redirect(http.StatusSeeOther, "/cart")
}

func (h *Handler) RemoveFromCart(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Некорректный ID")
		return
	}
	cart := getCartFromSession(ctx)
	newCart := []int{}
	for _, cid := range cart {
		if cid != id {
			newCart = append(newCart, cid)
		}
	}
	setCartToSession(ctx, newCart)
	ctx.Redirect(http.StatusSeeOther, "/cart")
}

func getCartFromSession(ctx *gin.Context) []int {
	cart := []int{}
	cookie, err := ctx.Cookie("cart")
	if err != nil || cookie == "" {
		return cart
	}
	for _, idStr := range splitAndTrim(cookie, ",") {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			cart = append(cart, id)
		}
	}
	return cart
}

func setCartToSession(ctx *gin.Context, cart []int) {
	strs := []string{}
	for _, id := range cart {
		strs = append(strs, strconv.Itoa(id))
	}
	ctx.SetCookie("cart", join(strs, ","), 3600, "/", "", false, true)
}

func splitAndTrim(s, sep string) []string {
	parts := []string{}
	for _, p := range strings.Split(s, sep) {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			parts = append(parts, trimmed)
		}
	}
	return parts
}

func split(s, sep string) []string {
	return strings.Split(s, sep)
}

func join(strs []string, sep string) string {
	return strings.Join(strs, sep)
}

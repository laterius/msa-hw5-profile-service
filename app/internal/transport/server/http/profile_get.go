package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laterius/service_architecture_hw3/app/internal/domain"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

func NewGetProfile(r service.UserReader) *getProfileHandler {
	return &getProfileHandler{

		readerUser: r,
	}
}

type getProfileHandler struct {
	//reader service.UserRememberReader
	readerUser service.UserReader
}

func (h *getProfileHandler) Handle() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := ctx.ParamsInt(UserIdFieldName, 0)
		if err != nil {
			return fail(ctx, err)
		}

		rememberToken, ok := ctx.GetReqHeaders()["remember_token"]
		if ok != true {
			if err == http.ErrNoCookie {
				return ctx.SendStatus(http.StatusUnauthorized)
			}
		}

		user, err := h.readerUser.Get(domain.UserId(userId))
		if err != nil {
			return fail(ctx, err)
		}

		if user.Remember == rememberToken {
			return ctx.Render("profile", fiber.Map{
				"FirstName": user.FirstName,
				"LastName":  user.LastName,
				"Username":  user.Username,
				"Phone":     user.Phone,
				"Email":     user.Email,
				"Token":     user.Remember,
			})
		}

		return ctx.SendStatus(http.StatusForbidden)

		//return json(ctx, (&service.User{}).FromDomain(user))
	}
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nas03/scholar-ai/backend/internal/models"
	"github.com/nas03/scholar-ai/backend/internal/services"
	"github.com/nas03/scholar-ai/backend/pkg/response"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Ping godoc
// @Summary      Health check endpoint
// @Description  Returns a simple pong message to verify the API is running
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string  "Success response with pong message"
// @Router       /users/ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Register a new user account. Returns success with OTP requirement if registration is successful.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      models.CreateUserRequest  true  "User registration data"
// @Success      200      {object}  response.ResponseData     "User created successfully"
// @Failure      200      {object}  response.ResponseData     "Error response (user already exists, invalid input, etc.)"
// @Router       /users/create [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var payload models.CreateUserRequest

	// Validate JSON binding
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response.ErrorResponse(ctx, response.CodeInvalidInput, err.Error())
		return
	}

	// Call service to create user
	code := c.userService.CreateUser(ctx, payload.Username, payload.Password, payload.Email)

	// Handle response based on service result
	if code == response.CodeSuccess {
		data := map[string]any{"requiresOtp": true}
		response.SuccessResponse(ctx, code, data)
	} else {
		response.ErrorResponse(ctx, code, "")
	}
}

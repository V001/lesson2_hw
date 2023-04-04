package handlers

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/storage"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	storage     *storage.Storage
	userManager *service.Service
}

func NewUserHandler(storage *storage.Storage, userManager *service.Service) *UserHandler {
	return &UserHandler{storage: storage, userManager: userManager}
}

func (h *UserHandler) Create(c echo.Context) error {
 // зачем эта структура ? 
	var input struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := readJSON(c.Response(), c.Request(), &input)
	if err != nil {
		return err
	}
	
	user := &model.User{
		ID:      input.ID,
		Name:    input.Name,
		Surname: input.Surname,
		Email:   input.Email,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		return err
	}

	err = h.userManager.User.Create(*user)
	//err = h.userManager.User.Delete("1")
	if err != nil {
		return err
	}

	err = writeJSON(c.Response(), http.StatusAccepted, envelope{"user": user}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (h *UserHandler) Get(c echo.Context) {

}

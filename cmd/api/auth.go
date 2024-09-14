package main

import (
	"net/http"

	"github.com/pradytpk/go-blog/internal/store"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,max=255"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

// registerUserHandler godoc
//
//	@Summary		Register a user
//	@Description	Registers a user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	RegisterUserPayload	true	"User credentials"
//	@Succedd		201 {object} store.User "User registered"
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router			/authentication/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJson(w, r, payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// hash the user password
	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
	}

	ctx := r.Context()

	// store the user
	if err := app.store.UsersIF.CreateAndInvite(ctx, user, "uuidv4"); err != nil {
		app.internalServerError(w, r, err)
	}

	// return the response
	if err := app.jsonResponse(w, http.StatusOK, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

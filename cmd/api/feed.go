package main

import "net/http"

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	feeds, err := app.store.PostsIF.GetUserFeed(ctx, int64(58))
	if err != nil {
		app.internalServerError(w, r, err)
	}
	if err := app.jsonResponse(w, http.StatusOK, feeds); err != nil {
		app.internalServerError(w, r, err)
	}
}

package handlers

import (
	"context"
	"net/http"
)

func contextProvider(ctx context.Context) ContextProvider {
	return func(_ *http.Request) context.Context { return ctx }
}

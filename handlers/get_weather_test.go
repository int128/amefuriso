package handlers

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/int128/amefurisobot/domain"
	"github.com/int128/amefurisobot/usecases/mock_usecases"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWeather_ServeHTTP(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock_usecases.NewMockIGetWeather(ctrl)
	usecase.EXPECT().
		Do(ctx, domain.UserID("USER1"), domain.SubscriptionID("SUBSCRIPTION1")).
		Return(&domain.Weather{}, nil)

	req := httptest.NewRequest("GET", "/USER1/SUBSCRIPTION1/weather", nil)
	w := httptest.NewRecorder()
	h := Handlers{
		GetWeather: GetWeather{
			ContextProvider: func(_ *http.Request) context.Context {
				return ctx
			},
			Usecase: usecase,
		},
	}
	h.NewRouter().ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Code wants 200 but %v", w.Code)
	}
	contentType := w.Header().Get("content-type")
	if contentType != "image/png" {
		t.Errorf("content-type wants image/png but %v", contentType)
	}
}

func TestGetWeather_ServeHTTP_NotFound(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock_usecases.NewMockIGetWeather(ctrl)
	usecase.EXPECT().
		Do(ctx, domain.UserID("USER1"), domain.SubscriptionID("SUBSCRIPTION1")).
		Return(nil, domain.ErrNoSuchUser{})

	req := httptest.NewRequest("GET", "/USER1/SUBSCRIPTION1/weather", nil)
	w := httptest.NewRecorder()
	h := Handlers{
		GetWeather: GetWeather{
			ContextProvider: func(_ *http.Request) context.Context {
				return ctx
			},
			Usecase: usecase,
		},
	}
	h.NewRouter().ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Code wants 404 but %v", w.Code)
	}
}

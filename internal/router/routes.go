package router

import (
	"net/http"
	"urlshortner/internal/constant"
	"urlshortner/internal/controller"
)

var urlShortner = Routes{
	Route{"Url Shortner Service", http.MethodPost, constant.UrlShortnerPath, controller.ShortTheUrl},
	Route{"Redirect to url", http.MethodGet, constant.RedirectUrlPath, controller.RedirectURL},
}

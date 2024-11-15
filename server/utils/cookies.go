package utils

import (
	"net/http"
	"time"
)

const defaultPath string = "/"

func getAccessCookieOpts(ac string) *http.Cookie {
	return &http.Cookie{
		Name: "accessToken",
		HttpOnly: true,
		Value: ac,
		Path: defaultPath,
		SameSite: http.SameSiteStrictMode,
		Secure: false,
		Expires: time.Now().Add(15 * time.Minute),
	}
}

func getRefreshCookieOpts(rf string) *http.Cookie {
	return &http.Cookie{
		Name:     "refreshToken",
		HttpOnly: true,
		Value: rf,
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		Expires:  time.Now().AddDate(0,0,7),
		Path:     defaultPath,
	}
}

func SetAuthCookies(w http.ResponseWriter, accessToken string, refreshToken string){
	http.SetCookie(w, getAccessCookieOpts(accessToken))

	if refreshToken != "" {
		http.SetCookie(w, getRefreshCookieOpts(refreshToken))
	}
}

func ClearAuthCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "accessToken",
		Value:  "",
		Path:   defaultPath,
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "refreshToken",
		Value:  "",
		Path:   defaultPath,
		MaxAge: -1,
	})
}
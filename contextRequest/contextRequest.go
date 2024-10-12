package ContextRequest

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/bytedance/sonic"
)

type ContextRequest struct {
	W http.ResponseWriter
	R *http.Request
}

type Cookie struct {
	*http.Cookie
}

func NewCookie(name, value string, maxAge int, path string, secure, httpOnly bool, sameSite http.SameSite) *Cookie {
	return &Cookie{
		Cookie: &http.Cookie{
			Name:     name,
			Value:    value,
			MaxAge:   maxAge,
			Path:     path,
			SameSite: sameSite,
			Secure:   secure,
			HttpOnly: httpOnly,
		},
	}
}

func (c Cookie) ToHTTPCookie() *http.Cookie {
	return c.Cookie
}

var NotAPointerError = errors.New("Not a pointer")

func (ctx *ContextRequest) ParseBody(v interface{}) error {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Pointer {
		return NotAPointerError
	}

	if err := sonic.ConfigDefault.NewDecoder(ctx.R.Body).Decode(v); err != nil {
		return err
	}

	return nil
}

func (ctx *ContextRequest) JSON(status int, v interface{}) error {
	ctx.W.WriteHeader(status)
	return sonic.ConfigDefault.NewEncoder(ctx.W).Encode(v)
}

func (ctx *ContextRequest) Cookie(cookie *Cookie) {
	http.SetCookie(ctx.W, cookie.ToHTTPCookie())
}

package misugo

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/bytedance/sonic"
)

type ContextRequest struct {
	w       http.ResponseWriter
	r       *http.Request
	encoder sonic.Encoder
	decoder sonic.Decoder
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

	return ctx.decoder.Decode(v)
}

func (ctx *ContextRequest) JSON(status int, v interface{}) error {
	ctx.w.WriteHeader(status)
	return ctx.encoder.Encode(v)
}

func (ctx *ContextRequest) Cookie(cookie *Cookie) {
	http.SetCookie(ctx.w, cookie.ToHTTPCookie())
}

package http_util

import (
	"net/http"
	"strconv"
	"veric-backend-mvp/logic/eth"
)

type HTTPContext struct {
	*http.Request

	w http.ResponseWriter
}

func NewHTTPContext(r *http.Request, w http.ResponseWriter) *HTTPContext {
	return &HTTPContext{Request: r, w: w}
}

func (c *HTTPContext) GetHeaderPublicKey() (*eth.PublicKey, error) {
	signData := c.Header.Get("X-Token")
	if signData == "" {
		return nil, NewHttpError(0xE000001, "token not exists")
	}

	return GetUserPublicKeyFromSign(signData)
}

func (c *HTTPContext) QueryWithDefault(key, def string) string {
	query := c.URL.Query()
	if query.Has(key) {
		return query.Get(key)
	}

	return def
}

func (c *HTTPContext) QueryWithDefaultInt(key string, def int) int {
	query := c.URL.Query()
	if query.Has(key) {
		queryNum, err := strconv.Atoi(query.Get(key))
		if err != nil {
			return def
		}

		return queryNum
	}

	return def
}

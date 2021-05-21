package transport_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swiftcarrot/dashi/gqlgen/graphql/handler/testserver"
	"github.com/swiftcarrot/dashi/gqlgen/graphql/handler/transport"
)

func TestOptions(t *testing.T) {
	h := testserver.New()
	h.AddTransport(transport.Options{})

	t.Run("responds to options requests", func(t *testing.T) {
		resp := doRequest(h, "OPTIONS", "/graphql?query={me{name}}", ``)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "OPTIONS, GET, POST", resp.Header().Get("Allow"))
	})

	t.Run("responds to head requests", func(t *testing.T) {
		resp := doRequest(h, "HEAD", "/graphql?query={me{name}}", ``)
		assert.Equal(t, http.StatusMethodNotAllowed, resp.Code)
	})
}

package syncspec

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWellknown_MarshalJSON(t *testing.T) {
	w := Wellknown{
		TokenEndpoint:            "https://www.example.com/sync/v1/token",
		SearchUserEndpoint:       "https://www.example.com/sync/v1/user",
		SearchDepartmentEndpoint: "https://www.example.com/sync/v1/dept",
		TokenEndpointAuthMethods: []string{TokenEndpointAuthMethodBasic, TokenEndpointAuthMethodPost},
	}

	b, err := json.Marshal(&w)
	assert.NoError(t, err)
	assert.Contains(t, string(b), `"spec":"v1"`)
}
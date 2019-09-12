package anchore_client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stackrox/anchore-client/client"
)

func TestClient(t *testing.T) {
	config := client.NewConfiguration()
	config.BasePath = fmt.Sprintf("%s/v1", "http://localhost:8228")
	config.AddDefaultHeader("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("admin:foobar"))))

	cl := client.NewAPIClient(config)

	vulnResponse, _, err := cl.VulnerabilitiesApi.GetImageVulnerabilitiesByType(context.Background(),
		"sha256:f265c36f0ac53a26aaa4b5b760c19e0fe6feb64b4d610599f9cbdbee5b8f0773", "all", nil)
	if err != nil {
		t.Error(err)
		return
	}

	s, err := json.MarshalIndent(vulnResponse, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v\n", string(s))
}

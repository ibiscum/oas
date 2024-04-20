/*
Callback Example

Testing ChickenNuggetsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package callbacks

import (
	"context"
	"testing"

	openapiclient "github.com/ibiscum/oas/callbacks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_circular_ChickenNuggetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChickenNuggetsAPIService PostStreams", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ChickenNuggetsAPI.PostStreams(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
/*
API definition with a circular $ref

Testing ChickenNuggetsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package circular

import (
	"context"
	"testing"

	openapiclient "github.com/ibiscum/oas/circular"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_circular_ChickenNuggetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChickenNuggetsAPIService GetAnything", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ChickenNuggetsAPI.GetAnything(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

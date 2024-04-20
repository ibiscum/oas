/*
Polymorphism support

Testing PetsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package polymorphism

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func Test_polymorphism_PetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PetsAPIService AnyOfWithType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnyOfWithType(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService AnythingAllOfObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnythingAllOfObject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService AnythingAnyOfObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnythingAnyOfObject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService AnythingAnyOfPrimitive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnythingAnyOfPrimitive(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService AnythingNestedOneOfObjectRef", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnythingNestedOneOfObjectRef(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService AnythingNestedOneOfObjectWithNestedOneOf", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnythingNestedOneOfObjectWithNestedOneOf(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService AnythingNestedOneOfRef", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnythingNestedOneOfRef(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService AnythingOneOfObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnythingOneOfObject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService AnythingOneOfObjectRef", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnythingOneOfObjectRef(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService AnythingOneOfPrimitive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.AnythingOneOfPrimitive(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService GetPets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.GetPets(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService OneOfComplex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.OneOfComplex(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService OneOfWithType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.OneOfWithType(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
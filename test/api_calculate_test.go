/*
Infrahub-API

Testing CalculateAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package hyperstack

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/szczad/hyperstack-api-go"
)

func Test_hyperstack_CalculateAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CalculateAPIService GetCalculate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var resourceType string
		var id int32

		resp, httpRes, err := apiClient.CalculateAPI.GetCalculate(context.Background(), resourceType, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

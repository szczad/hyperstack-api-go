/*
Infrahub-API

Testing VirtualMachineAPIService

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

func Test_hyperstack_VirtualMachineAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VirtualMachineAPIService AddFirewallRuleToVirtualMachine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.AddFirewallRuleToVirtualMachine(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService AttachFirewallsToAVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vmId int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.AttachFirewallsToAVM(context.Background(), vmId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService CreateVirtualMachine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VirtualMachineAPI.CreateVirtualMachine(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService DeleteFirewallRuleFromVirtualMachine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var virtualMachineId int32
		var sgRuleId int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.DeleteFirewallRuleFromVirtualMachine(context.Background(), virtualMachineId, sgRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService DeleteVirtualMachine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.DeleteVirtualMachine(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService EditLabelsOfAnExistingVM", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var virtualMachineId int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.EditLabelsOfAnExistingVM(context.Background(), virtualMachineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService HardRebootVirtualMachine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.HardRebootVirtualMachine(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService HibernateVirtualMachine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var virtualMachineId int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.HibernateVirtualMachine(context.Background(), virtualMachineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService ListVirtualMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VirtualMachineAPI.ListVirtualMachines(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService ResizeVirtualMachine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var virtualMachineId int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.ResizeVirtualMachine(context.Background(), virtualMachineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService RestoreVirtualMachineFromHibernation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var virtualMachineId int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.RestoreVirtualMachineFromHibernation(context.Background(), virtualMachineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService RetrieveVirtualMachineDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.RetrieveVirtualMachineDetails(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService RetrieveVirtualMachinePerformanceMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var virtualMachineId int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.RetrieveVirtualMachinePerformanceMetrics(context.Background(), virtualMachineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService RetrieveVirtualMachinesAssociatedWithAContract", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contractId int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.RetrieveVirtualMachinesAssociatedWithAContract(context.Background(), contractId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService StartVirtualMachine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.StartVirtualMachine(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualMachineAPIService StopVirtualMachine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.VirtualMachineAPI.StopVirtualMachine(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

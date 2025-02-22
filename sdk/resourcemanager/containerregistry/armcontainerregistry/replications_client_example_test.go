//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ReplicationList.json
func ExampleReplicationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewReplicationsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("myResourceGroup", "myRegistry", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ReplicationListResult = armcontainerregistry.ReplicationListResult{
		// 	Value: []*armcontainerregistry.Replication{
		// 		{
		// 			Name: to.Ptr("myReplication"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/replications"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/replications/myReplication"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armcontainerregistry.ReplicationProperties{
		// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				RegionEndpointEnabled: to.Ptr(true),
		// 				Status: &armcontainerregistry.Status{
		// 					DisplayStatus: to.Ptr("Ready"),
		// 					Message: to.Ptr("The replication is ready."),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:15:37.0707808Z"); return t}()),
		// 				},
		// 				ZoneRedundancy: to.Ptr(armcontainerregistry.ZoneRedundancyDisabled),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ReplicationGet.json
func ExampleReplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewReplicationsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myResourceGroup", "myRegistry", "myReplication", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Replication = armcontainerregistry.Replication{
	// 	Name: to.Ptr("myReplication"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/replications"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/replications/myReplication"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.ReplicationProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RegionEndpointEnabled: to.Ptr(true),
	// 		Status: &armcontainerregistry.Status{
	// 			DisplayStatus: to.Ptr("Ready"),
	// 			Message: to.Ptr("The replication is ready."),
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:15:37.0707808Z"); return t}()),
	// 		},
	// 		ZoneRedundancy: to.Ptr(armcontainerregistry.ZoneRedundancyDisabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ReplicationCreate.json
func ExampleReplicationsClient_BeginCreate_replicationCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewReplicationsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "myResourceGroup", "myRegistry", "myReplication", armcontainerregistry.Replication{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Replication = armcontainerregistry.Replication{
	// 	Name: to.Ptr("myReplication"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/replications"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/replications/myReplication"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.ReplicationProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RegionEndpointEnabled: to.Ptr(true),
	// 		Status: &armcontainerregistry.Status{
	// 			DisplayStatus: to.Ptr("Ready"),
	// 			Message: to.Ptr("The replication is ready."),
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:15:37.0707808Z"); return t}()),
	// 		},
	// 		ZoneRedundancy: to.Ptr(armcontainerregistry.ZoneRedundancyDisabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ReplicationCreateZoneRedundant.json
func ExampleReplicationsClient_BeginCreate_replicationCreateZoneRedundant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewReplicationsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "myResourceGroup", "myRegistry", "myReplication", armcontainerregistry.Replication{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
		Properties: &armcontainerregistry.ReplicationProperties{
			RegionEndpointEnabled: to.Ptr(true),
			ZoneRedundancy:        to.Ptr(armcontainerregistry.ZoneRedundancyEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Replication = armcontainerregistry.Replication{
	// 	Name: to.Ptr("myReplication"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/replications"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/replications/myReplication"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.ReplicationProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RegionEndpointEnabled: to.Ptr(true),
	// 		Status: &armcontainerregistry.Status{
	// 			DisplayStatus: to.Ptr("Ready"),
	// 			Message: to.Ptr("The replication is ready."),
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:15:37.0707808Z"); return t}()),
	// 		},
	// 		ZoneRedundancy: to.Ptr(armcontainerregistry.ZoneRedundancyEnabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ReplicationDelete.json
func ExampleReplicationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewReplicationsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "myResourceGroup", "myRegistry", "myReplication", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/270d3cd664cca3ddc8511f92d3851a715e2c61db/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ReplicationUpdate.json
func ExampleReplicationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewReplicationsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myReplication", armcontainerregistry.ReplicationUpdateParameters{
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Replication = armcontainerregistry.Replication{
	// 	Name: to.Ptr("myReplication"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/replications"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/replications/myReplication"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.ReplicationProperties{
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RegionEndpointEnabled: to.Ptr(true),
	// 		Status: &armcontainerregistry.Status{
	// 			DisplayStatus: to.Ptr("Ready"),
	// 			Message: to.Ptr("The replication is ready."),
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:15:37.0707808Z"); return t}()),
	// 		},
	// 		ZoneRedundancy: to.Ptr(armcontainerregistry.ZoneRedundancyDisabled),
	// 	},
	// }
}

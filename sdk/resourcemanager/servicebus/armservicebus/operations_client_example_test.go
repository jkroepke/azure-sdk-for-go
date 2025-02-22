//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/SBOperations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewOperationsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
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
		// page.OperationListResult = armservicebus.OperationListResult{
		// 	Value: []*armservicebus.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/checkNameAvailability/action"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get namespace availability."),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Non Resource Operation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/register/action"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Registers the ServiceBus Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("ServiceBus Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/write"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Create Or Update Namespace "),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Resource"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/Delete"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Delete Namespace"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Namespace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/authorizationRules/write"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/authorizationRules/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/authorizationRules/delete"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Delete Namespace Authorization Rule"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/authorizationRules/listkeys/action"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace Listkeys"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/authorizationRules/regenerateKeys/action"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Resource Regeneratekeys"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/queues/write"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Queue"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Queue"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/queues/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get Queue"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Queue"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/queues/Delete"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Delete Queue"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Queue"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/queues/authorizationRules/write"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Queue Authorization Rule"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Queue AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/queues/authorizationRules/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr(" Get Queue Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Queue AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/queues/authorizationRules/delete"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Delete Queue Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Queue AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/queues/authorizationRules/listkeys/action"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("List Queue keys"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Queue AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/queues/authorizationRules/regenerateKeys/action"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Resource Regeneratekeys"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Queue AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/write"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Topic"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Topic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get Topic"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Topic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/Delete"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Delete Topic"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Topic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/authorizationRules/write"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Topic Authorization Rule"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Topic AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/authorizationRules/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr(" Get Topic Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Topic AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/authorizationRules/delete"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Delete Topic Authorization Rules"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Topic AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/authorizationRules/listkeys/action"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("List Topic keys"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Topic AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/authorizationRules/regenerateKeys/action"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Resource Regeneratekeys"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Topic AuthorizationRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/subscriptions/write"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update TopicSubscription"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("TopicSubscription"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/subscriptions/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get TopicSubscription"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("TopicSubscription"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/subscriptions/Delete"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Delete TopicSubscription"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("TopicSubscription"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/subscriptions/rules/write"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Rule"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/subscriptions/rules/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get Rule"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/topics/subscriptions/rules/Delete"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Delete Rule"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Rule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/metricDefinitions/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace metrics"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Namespace metrics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/diagnosticSettings/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace diagnostic settings"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Namespace diagnostic settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/diagnosticSettings/write"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Namespace diagnostic settings"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Namespace diagnostic settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ServiceBus/namespaces/logDefinitions/read"),
		// 			Display: &armservicebus.OperationDisplay{
		// 				Operation: to.Ptr("Get Namespace logs"),
		// 				Provider: to.Ptr("Microsoft Azure ServiceBus"),
		// 				Resource: to.Ptr("Namespace logs"),
		// 			},
		// 	}},
		// }
	}
}

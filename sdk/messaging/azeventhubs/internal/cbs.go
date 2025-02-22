// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	azlog "github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/auth"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/go-amqp"
)

const (
	cbsAddress           = "$cbs"
	cbsOperationKey      = "operation"
	cbsOperationPutToken = "put-token"
	cbsTokenTypeKey      = "type"
	cbsAudienceKey       = "name"
	cbsExpirationKey     = "expiration"
)

// NegotiateClaim attempts to put a token to the $cbs management endpoint to negotiate auth for the given audience
//
// contextWithTimeoutFn is intended to be context.WithTimeout in production code, but can be stubbed out when writing
// unit tests to keep timeouts reasonable.
func NegotiateClaim(ctx context.Context, audience string, conn amqpwrap.AMQPClient, provider auth.TokenProvider, contextWithTimeoutFn contextWithTimeoutFn) error {
	link, err := NewRPCLink(ctx, RPCLinkArgs{
		Client:   conn,
		Address:  cbsAddress,
		LogEvent: exported.EventAuth,
	})

	if err != nil {
		// In some circumstances we can end up in a situation where the link closing was cancelled
		// or interrupted, leaving $cbs still open by some dangling receiver or sender. The only way
		// to fix this is to restart the connection.
		if IsNotAllowedError(err) {
			log.Writef(exported.EventAuth, "Not allowed to open, connection will be reset: %s", err)
			return errConnResetNeeded
		}

		return err
	}

	closeLink := func(ctx context.Context, origErr error) error {
		ctx, cancel := contextWithTimeoutFn(ctx, defaultCloseTimeout)
		defer cancel()

		if err := link.Close(ctx); err != nil {
			if IsCancelError(err) {
				azlog.Writef(exported.EventAuth, "Failed closing claim link because it was cancelled. Connection will need to be reset")
				return errConnResetNeeded
			}

			azlog.Writef(exported.EventAuth, "Failed closing claim link: %s", err.Error())
			return err
		}

		return origErr
	}

	token, err := provider.GetToken(audience)
	if err != nil {
		azlog.Writef(exported.EventAuth, "Failed to get token from provider")
		return closeLink(ctx, err)
	}

	azlog.Writef(exported.EventAuth, "negotiating claim for audience %s with token type %s and expiry of %s", audience, token.TokenType, token.Expiry)

	msg := &amqp.Message{
		Value: token.Token,
		ApplicationProperties: map[string]interface{}{
			cbsOperationKey:  cbsOperationPutToken,
			cbsTokenTypeKey:  string(token.TokenType),
			cbsAudienceKey:   audience,
			cbsExpirationKey: token.Expiry,
		},
	}

	if _, err := link.RPC(ctx, msg); err != nil {
		azlog.Writef(exported.EventAuth, "Failed to send/receive RPC message")
		return closeLink(ctx, err)
	}

	return closeLink(ctx, nil)
}

var errConnResetNeeded = errors.New("connection must be reset, link/connection state may be inconsistent")

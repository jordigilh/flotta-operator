// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"

	"github.com/project-flotta/flotta-operator/restapi/operations"
)

type contextKey string

const AuthKey contextKey = "Auth"

//go:generate mockery -name OperationsAPI -inpkg

/* OperationsAPI  */
type OperationsAPI interface {
	/* EnrolDevice Initiates the process of enrolling the device */
	EnrolDevice(ctx context.Context, params operations.EnrolDeviceParams) middleware.Responder

	/* FinalizeDeviceRegistration Updates the status of the device as registered and its hardware configuration */
	FinalizeDeviceRegistration(ctx context.Context, params operations.FinalizeDeviceRegistrationParams) middleware.Responder

	/* GetDeviceConfiguration Returns the device configuration */
	GetDeviceConfiguration(ctx context.Context, params operations.GetDeviceConfigurationParams) middleware.Responder

	/* InitializeDeviceRegistration Initializes the process of registering the device */
	InitializeDeviceRegistration(ctx context.Context, params operations.InitializeDeviceRegistrationParams) middleware.Responder

	/* IsDeviceDeregistrable Confirms if a device can be deregistered */
	IsDeviceDeregistrable(ctx context.Context, params operations.IsDeviceDeregistrableParams) middleware.Responder
}

// Config is configuration for Handler
type Config struct {
	OperationsAPI
	Logger func(string, ...interface{})
	// InnerMiddleware is for the handler executors. These do not apply to the swagger.json document.
	// The middleware executes after routing but before authentication, binding and validation
	InnerMiddleware func(http.Handler) http.Handler

	// Authorizer is used to authorize a request after the Auth function was called using the "Auth*" functions
	// and the principal was stored in the context in the "AuthKey" context value.
	Authorizer func(*http.Request) error

	// Authenticator to use for all APIKey authentication
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// Authenticator to use for all Bearer authentication
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// Authenticator to use for all Basic authentication
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator
}

// Handler returns an http.Handler given the handler configuration
// It mounts all the business logic implementers in the right routing.
func Handler(c Config) (http.Handler, error) {
	h, _, err := HandlerAPI(c)
	return h, err
}

// HandlerAPI returns an http.Handler given the handler configuration
// and the corresponding *FlottaBackendAPI instance.
// It mounts all the business logic implementers in the right routing.
func HandlerAPI(c Config) (http.Handler, *operations.FlottaBackendAPIAPI, error) {
	spec, err := loads.Analyzed(swaggerCopy(SwaggerJSON), "")
	if err != nil {
		return nil, nil, fmt.Errorf("analyze swagger: %v", err)
	}
	api := operations.NewFlottaBackendAPIAPI(spec)
	api.ServeError = errors.ServeError
	api.Logger = c.Logger

	if c.APIKeyAuthenticator != nil {
		api.APIKeyAuthenticator = c.APIKeyAuthenticator
	}
	if c.BasicAuthenticator != nil {
		api.BasicAuthenticator = c.BasicAuthenticator
	}
	if c.BearerAuthenticator != nil {
		api.BearerAuthenticator = c.BearerAuthenticator
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.EnrolDeviceHandler = operations.EnrolDeviceHandlerFunc(func(params operations.EnrolDeviceParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.OperationsAPI.EnrolDevice(ctx, params)
	})
	api.FinalizeDeviceRegistrationHandler = operations.FinalizeDeviceRegistrationHandlerFunc(func(params operations.FinalizeDeviceRegistrationParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.OperationsAPI.FinalizeDeviceRegistration(ctx, params)
	})
	api.GetDeviceConfigurationHandler = operations.GetDeviceConfigurationHandlerFunc(func(params operations.GetDeviceConfigurationParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.OperationsAPI.GetDeviceConfiguration(ctx, params)
	})
	api.InitializeDeviceRegistrationHandler = operations.InitializeDeviceRegistrationHandlerFunc(func(params operations.InitializeDeviceRegistrationParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.OperationsAPI.InitializeDeviceRegistration(ctx, params)
	})
	api.IsDeviceDeregistrableHandler = operations.IsDeviceDeregistrableHandlerFunc(func(params operations.IsDeviceDeregistrableParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.OperationsAPI.IsDeviceDeregistrable(ctx, params)
	})
	api.ServerShutdown = func() {}
	return api.Serve(c.InnerMiddleware), api, nil
}

// swaggerCopy copies the swagger json to prevent data races in runtime
func swaggerCopy(orig json.RawMessage) json.RawMessage {
	c := make(json.RawMessage, len(orig))
	copy(c, orig)
	return c
}

// authorizer is a helper function to implement the runtime.Authorizer interface.
type authorizer func(*http.Request) error

func (a authorizer) Authorize(req *http.Request, principal interface{}) error {
	if a == nil {
		return nil
	}
	ctx := storeAuth(req.Context(), principal)
	return a(req.WithContext(ctx))
}

func storeAuth(ctx context.Context, principal interface{}) context.Context {
	return context.WithValue(ctx, AuthKey, principal)
}

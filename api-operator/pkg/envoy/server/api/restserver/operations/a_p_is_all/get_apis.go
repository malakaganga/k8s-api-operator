// Code generated by go-swagger; DO NOT EDIT.

package a_p_is_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/wso2/k8s-api-operator/api-operator/pkg/envoy/server/api/models"
)

// GetApisHandlerFunc turns a function with the right signature into a get apis handler
type GetApisHandlerFunc func(GetApisParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetApisHandlerFunc) Handle(params GetApisParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetApisHandler interface for that can handle valid get apis params
type GetApisHandler interface {
	Handle(GetApisParams, *models.Principal) middleware.Responder
}

// NewGetApis creates a new http.Handler for the get apis operation
func NewGetApis(ctx *middleware.Context, handler GetApisHandler) *GetApis {
	return &GetApis{Context: ctx, Handler: handler}
}

/*GetApis swagger:route GET /apis APIs (All) getApis

Get all apis in a zip file

This operation can be used to get all the APIs deployed in Kubernetes.


*/
type GetApis struct {
	Context *middleware.Context
	Handler GetApisHandler
}

func (o *GetApis) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetApisParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
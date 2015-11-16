
// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// GetEventsHandlerFunc turns a function with the right signature into a get events handler
type GetEventsHandlerFunc func() middleware.Responder

// Handle executing the request and returning a response
func (fn GetEventsHandlerFunc) Handle() middleware.Responder {
	return fn()
}

// GetEventsHandler interface for that can handle valid get events params
type GetEventsHandler interface {
	Handle() middleware.Responder
}

// NewGetEvents creates a new http.Handler for the get events operation
func NewGetEvents(ctx *middleware.Context, handler GetEventsHandler) *GetEvents {
	return &GetEvents{Context: ctx, Handler: handler}
}

/*
Get events.
*/
type GetEvents struct {
	Context *middleware.Context
	Handler GetEventsHandler
}

func (o *GetEvents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, nil); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle() // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetJourneys(params *GetJourneysParams, opts ...ClientOption) (*GetJourneysOK, error)

	GetLocations(params *GetLocationsParams, opts ...ClientOption) (*GetLocationsOK, error)

	GetStationsID(params *GetStationsIDParams, opts ...ClientOption) (*GetStationsIDOK, error)

	GetStopsIDArrivals(params *GetStopsIDArrivalsParams, opts ...ClientOption) (*GetStopsIDArrivalsOK, error)

	GetStopsIDDepartures(params *GetStopsIDDeparturesParams, opts ...ClientOption) (*GetStopsIDDeparturesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetJourneys finds journeys from a to b

  Uses [`hafasClient.journeys()`](https://github.com/public-transport/hafas-client/blob/5/docs/journeys.md) to **find journeys from A (`from`) to B (`to`)**.
*/
func (a *Client) GetJourneys(params *GetJourneysParams, opts ...ClientOption) (*GetJourneysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJourneysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetJourneys",
		Method:             "GET",
		PathPattern:        "/journeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetJourneysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJourneysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetJourneys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLocations finds stops stations p o is and addresses matching a query

  Uses [`hafasClient.locations()`](https://github.com/public-transport/hafas-client/blob/5/docs/locations.md) to **find stops/stations, POIs and addresses matching `query`**.
*/
func (a *Client) GetLocations(params *GetLocationsParams, opts ...ClientOption) (*GetLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLocationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLocations",
		Method:             "GET",
		PathPattern:        "/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLocationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLocations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStationsID returns a stop station from db stations

  Returns a stop/station from [db-stations](https://npmjs.com/package/db-stations).
*/
func (a *Client) GetStationsID(params *GetStationsIDParams, opts ...ClientOption) (*GetStationsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStationsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStationsID",
		Method:             "GET",
		PathPattern:        "/stations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStationsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStationsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStationsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStopsIDArrivals fetches arrivals at a stop station

  Works like `/stops/{id}/departures`, except that it uses [`hafasClient.arrivals()`](https://github.com/public-transport/hafas-client/blob/5/docs/arrivals.md) to **query arrivals at a stop/station**.
*/
func (a *Client) GetStopsIDArrivals(params *GetStopsIDArrivalsParams, opts ...ClientOption) (*GetStopsIDArrivalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStopsIDArrivalsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStopsIDArrivals",
		Method:             "GET",
		PathPattern:        "/stops/{id}/arrivals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStopsIDArrivalsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStopsIDArrivalsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStopsIDArrivals: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStopsIDDepartures fetches departures at a stop station

  Uses [`hafasClient.departures()`](https://github.com/public-transport/hafas-client/blob/5/docs/departures.md) to **query departures at a stop/station**.
*/
func (a *Client) GetStopsIDDepartures(params *GetStopsIDDeparturesParams, opts ...ClientOption) (*GetStopsIDDeparturesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStopsIDDeparturesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStopsIDDepartures",
		Method:             "GET",
		PathPattern:        "/stops/{id}/departures",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStopsIDDeparturesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStopsIDDeparturesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStopsIDDepartures: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

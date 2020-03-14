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

// ClientService is the interface for Client methods
type ClientService interface {
	GetStationPattern(params *GetStationPatternParams) (*GetStationPatternOK, error)

	GetTimetableFchgEvaNo(params *GetTimetableFchgEvaNoParams) (*GetTimetableFchgEvaNoOK, error)

	GetTimetablePlanEvaNoDateHour(params *GetTimetablePlanEvaNoDateHourParams) (*GetTimetablePlanEvaNoDateHourOK, error)

	GetTimetableRchgEvaNo(params *GetTimetableRchgEvaNoParams) (*GetTimetableRchgEvaNoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetStationPattern This public interface allows access to information about a station.
*/
func (a *Client) GetStationPattern(params *GetStationPatternParams) (*GetStationPatternOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStationPatternParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStationPattern",
		Method:             "GET",
		PathPattern:        "/station/{pattern}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xhtml+xml", "application/xml", "text/html"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStationPatternReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStationPatternOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStationPattern: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTimetableFchgEvaNo Returns a Timetable object (see Timetable) that contains all known changes for the station given by evaNo.

The data includes all known changes from now on until ndefinitely into the future. Once changes become obsolete (because their trip departs from the station) they are removed from this resource.

Changes may include messages. On event level, they usually contain one or more of the 'changed' attributes ct, cp, cs or cpth. Changes may also include 'planned' attributes if there is no associated planned data for the change (e.g. an unplanned stop or trip).

Full changes are updated every 30s and should be cached for that period by web caches.

*/
func (a *Client) GetTimetableFchgEvaNo(params *GetTimetableFchgEvaNoParams) (*GetTimetableFchgEvaNoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimetableFchgEvaNoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTimetableFchgEvaNo",
		Method:             "GET",
		PathPattern:        "/timetable/fchg/{evaNo}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xhtml+xml", "application/xml", "text/html"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTimetableFchgEvaNoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTimetableFchgEvaNoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTimetableFchgEvaNo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTimetablePlanEvaNoDateHour Returns a Timetable object (see Timetable) that contains planned data for the specified station (evaNo) within the hourly time slice given by date (format YYMMDD) and hour (format HH). The data includes stops for all trips that arrive or depart within that slice. There is a small overlap between slices since some trips arrive in one slice and depart in another.

Planned data does never contain messages. On event level, planned data contains the 'plannned' attributes pt, pp, ps and ppth while the 'changed' attributes ct, cp, cs and cpth are absent.

Planned data is generated many hours in advance and is static, i.e. it does never change. It should be cached by web caches.public interface allows access to information about a station.

*/
func (a *Client) GetTimetablePlanEvaNoDateHour(params *GetTimetablePlanEvaNoDateHourParams) (*GetTimetablePlanEvaNoDateHourOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimetablePlanEvaNoDateHourParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTimetablePlanEvaNoDateHour",
		Method:             "GET",
		PathPattern:        "/timetable/plan/{evaNo}/{date}{hour}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xhtml+xml", "application/xml", "text/html"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTimetablePlanEvaNoDateHourReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTimetablePlanEvaNoDateHourOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTimetablePlanEvaNoDateHour: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTimetableRchgEvaNo Returns a Timetable object (see Timetable) that contains all recent changes for the station given by evaNo. Recent changes are always a subset of the full changes. They may equal full changes but are typically much smaller. Data includes only those changes that became known within the last 2 minutes.

A client that updates its state in intervals of less than 2 minutes should load full changes initially and then proceed to periodically load only the recent changes in order to save bandwidth.

Recent changes are updated every 30s as well and should be cached for that period by web caches.

*/
func (a *Client) GetTimetableRchgEvaNo(params *GetTimetableRchgEvaNoParams) (*GetTimetableRchgEvaNoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimetableRchgEvaNoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTimetableRchgEvaNo",
		Method:             "GET",
		PathPattern:        "/timetable/rchg/{evaNo}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xhtml+xml", "application/xml", "text/html"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTimetableRchgEvaNoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTimetableRchgEvaNoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTimetableRchgEvaNo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

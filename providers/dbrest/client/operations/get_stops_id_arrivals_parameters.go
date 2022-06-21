// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetStopsIDArrivalsParams creates a new GetStopsIDArrivalsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStopsIDArrivalsParams() *GetStopsIDArrivalsParams {
	return &GetStopsIDArrivalsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStopsIDArrivalsParamsWithTimeout creates a new GetStopsIDArrivalsParams object
// with the ability to set a timeout on a request.
func NewGetStopsIDArrivalsParamsWithTimeout(timeout time.Duration) *GetStopsIDArrivalsParams {
	return &GetStopsIDArrivalsParams{
		timeout: timeout,
	}
}

// NewGetStopsIDArrivalsParamsWithContext creates a new GetStopsIDArrivalsParams object
// with the ability to set a context for a request.
func NewGetStopsIDArrivalsParamsWithContext(ctx context.Context) *GetStopsIDArrivalsParams {
	return &GetStopsIDArrivalsParams{
		Context: ctx,
	}
}

// NewGetStopsIDArrivalsParamsWithHTTPClient creates a new GetStopsIDArrivalsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStopsIDArrivalsParamsWithHTTPClient(client *http.Client) *GetStopsIDArrivalsParams {
	return &GetStopsIDArrivalsParams{
		HTTPClient: client,
	}
}

/* GetStopsIDArrivalsParams contains all the parameters to send to the API endpoint
   for the get stops ID arrivals operation.

   Typically these are written to a http.Request.
*/
type GetStopsIDArrivalsParams struct {

	// Bus.
	//
	// Default: true
	Bus *bool

	/* Direction.

	   Filter departures by direction.
	*/
	Direction *string

	/* Duration.

	   Show departures for how many minutes?

	   Default: 10
	*/
	Duration *int64

	// Ferry.
	//
	// Default: true
	Ferry *bool

	/* ID.

	   stop/station ID to show arrivals for
	*/
	ID string

	/* Language.

	   Language of the results.

	   Default: "en"
	*/
	Language *string

	/* LinesOfStops.

	   Parse & return lines of each stop/station?
	*/
	LinesOfStops *bool

	// National.
	//
	// Default: true
	National *bool

	// NationalExpress.
	//
	// Default: true
	NationalExpress *bool

	// Regional.
	//
	// Default: true
	Regional *bool

	// RegionalExp.
	//
	// Default: true
	RegionalExp *bool

	/* Remarks.

	   Parse & return hints & warnings?

	   Default: true
	*/
	Remarks *bool

	/* Results.

	   Max. number of departures. – Default: *whatever HAFAS wants*
	*/
	Results *int64

	// Suburban.
	//
	// Default: true
	Suburban *bool

	// Subway.
	//
	// Default: true
	Subway *bool

	// Taxi.
	//
	// Default: true
	Taxi *bool

	// Tram.
	//
	// Default: true
	Tram *bool

	/* When.

	   Date & time to get departures for. – Default: *now*

	   Format: date-time
	*/
	When *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get stops ID arrivals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStopsIDArrivalsParams) WithDefaults() *GetStopsIDArrivalsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get stops ID arrivals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStopsIDArrivalsParams) SetDefaults() {
	var (
		busDefault = bool(true)

		durationDefault = int64(10)

		ferryDefault = bool(true)

		languageDefault = string("en")

		linesOfStopsDefault = bool(false)

		nationalDefault = bool(true)

		nationalExpressDefault = bool(true)

		regionalDefault = bool(true)

		regionalExpDefault = bool(true)

		remarksDefault = bool(true)

		suburbanDefault = bool(true)

		subwayDefault = bool(true)

		taxiDefault = bool(true)

		tramDefault = bool(true)
	)

	val := GetStopsIDArrivalsParams{
		Bus:             &busDefault,
		Duration:        &durationDefault,
		Ferry:           &ferryDefault,
		Language:        &languageDefault,
		LinesOfStops:    &linesOfStopsDefault,
		National:        &nationalDefault,
		NationalExpress: &nationalExpressDefault,
		Regional:        &regionalDefault,
		RegionalExp:     &regionalExpDefault,
		Remarks:         &remarksDefault,
		Suburban:        &suburbanDefault,
		Subway:          &subwayDefault,
		Taxi:            &taxiDefault,
		Tram:            &tramDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithTimeout(timeout time.Duration) *GetStopsIDArrivalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithContext(ctx context.Context) *GetStopsIDArrivalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithHTTPClient(client *http.Client) *GetStopsIDArrivalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBus adds the bus to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithBus(bus *bool) *GetStopsIDArrivalsParams {
	o.SetBus(bus)
	return o
}

// SetBus adds the bus to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetBus(bus *bool) {
	o.Bus = bus
}

// WithDirection adds the direction to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithDirection(direction *string) *GetStopsIDArrivalsParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithDuration adds the duration to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithDuration(duration *int64) *GetStopsIDArrivalsParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetDuration(duration *int64) {
	o.Duration = duration
}

// WithFerry adds the ferry to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithFerry(ferry *bool) *GetStopsIDArrivalsParams {
	o.SetFerry(ferry)
	return o
}

// SetFerry adds the ferry to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetFerry(ferry *bool) {
	o.Ferry = ferry
}

// WithID adds the id to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithID(id string) *GetStopsIDArrivalsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetID(id string) {
	o.ID = id
}

// WithLanguage adds the language to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithLanguage(language *string) *GetStopsIDArrivalsParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetLanguage(language *string) {
	o.Language = language
}

// WithLinesOfStops adds the linesOfStops to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithLinesOfStops(linesOfStops *bool) *GetStopsIDArrivalsParams {
	o.SetLinesOfStops(linesOfStops)
	return o
}

// SetLinesOfStops adds the linesOfStops to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetLinesOfStops(linesOfStops *bool) {
	o.LinesOfStops = linesOfStops
}

// WithNational adds the national to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithNational(national *bool) *GetStopsIDArrivalsParams {
	o.SetNational(national)
	return o
}

// SetNational adds the national to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetNational(national *bool) {
	o.National = national
}

// WithNationalExpress adds the nationalExpress to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithNationalExpress(nationalExpress *bool) *GetStopsIDArrivalsParams {
	o.SetNationalExpress(nationalExpress)
	return o
}

// SetNationalExpress adds the nationalExpress to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetNationalExpress(nationalExpress *bool) {
	o.NationalExpress = nationalExpress
}

// WithRegional adds the regional to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithRegional(regional *bool) *GetStopsIDArrivalsParams {
	o.SetRegional(regional)
	return o
}

// SetRegional adds the regional to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetRegional(regional *bool) {
	o.Regional = regional
}

// WithRegionalExp adds the regionalExp to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithRegionalExp(regionalExp *bool) *GetStopsIDArrivalsParams {
	o.SetRegionalExp(regionalExp)
	return o
}

// SetRegionalExp adds the regionalExp to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetRegionalExp(regionalExp *bool) {
	o.RegionalExp = regionalExp
}

// WithRemarks adds the remarks to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithRemarks(remarks *bool) *GetStopsIDArrivalsParams {
	o.SetRemarks(remarks)
	return o
}

// SetRemarks adds the remarks to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetRemarks(remarks *bool) {
	o.Remarks = remarks
}

// WithResults adds the results to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithResults(results *int64) *GetStopsIDArrivalsParams {
	o.SetResults(results)
	return o
}

// SetResults adds the results to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetResults(results *int64) {
	o.Results = results
}

// WithSuburban adds the suburban to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithSuburban(suburban *bool) *GetStopsIDArrivalsParams {
	o.SetSuburban(suburban)
	return o
}

// SetSuburban adds the suburban to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetSuburban(suburban *bool) {
	o.Suburban = suburban
}

// WithSubway adds the subway to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithSubway(subway *bool) *GetStopsIDArrivalsParams {
	o.SetSubway(subway)
	return o
}

// SetSubway adds the subway to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetSubway(subway *bool) {
	o.Subway = subway
}

// WithTaxi adds the taxi to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithTaxi(taxi *bool) *GetStopsIDArrivalsParams {
	o.SetTaxi(taxi)
	return o
}

// SetTaxi adds the taxi to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetTaxi(taxi *bool) {
	o.Taxi = taxi
}

// WithTram adds the tram to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithTram(tram *bool) *GetStopsIDArrivalsParams {
	o.SetTram(tram)
	return o
}

// SetTram adds the tram to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetTram(tram *bool) {
	o.Tram = tram
}

// WithWhen adds the when to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) WithWhen(when *strfmt.DateTime) *GetStopsIDArrivalsParams {
	o.SetWhen(when)
	return o
}

// SetWhen adds the when to the get stops ID arrivals params
func (o *GetStopsIDArrivalsParams) SetWhen(when *strfmt.DateTime) {
	o.When = when
}

// WriteToRequest writes these params to a swagger request
func (o *GetStopsIDArrivalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Bus != nil {

		// query param bus
		var qrBus bool

		if o.Bus != nil {
			qrBus = *o.Bus
		}
		qBus := swag.FormatBool(qrBus)
		if qBus != "" {

			if err := r.SetQueryParam("bus", qBus); err != nil {
				return err
			}
		}
	}

	if o.Direction != nil {

		// query param direction
		var qrDirection string

		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {

			if err := r.SetQueryParam("direction", qDirection); err != nil {
				return err
			}
		}
	}

	if o.Duration != nil {

		// query param duration
		var qrDuration int64

		if o.Duration != nil {
			qrDuration = *o.Duration
		}
		qDuration := swag.FormatInt64(qrDuration)
		if qDuration != "" {

			if err := r.SetQueryParam("duration", qDuration); err != nil {
				return err
			}
		}
	}

	if o.Ferry != nil {

		// query param ferry
		var qrFerry bool

		if o.Ferry != nil {
			qrFerry = *o.Ferry
		}
		qFerry := swag.FormatBool(qrFerry)
		if qFerry != "" {

			if err := r.SetQueryParam("ferry", qFerry); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Language != nil {

		// query param language
		var qrLanguage string

		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {

			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}
	}

	if o.LinesOfStops != nil {

		// query param linesOfStops
		var qrLinesOfStops bool

		if o.LinesOfStops != nil {
			qrLinesOfStops = *o.LinesOfStops
		}
		qLinesOfStops := swag.FormatBool(qrLinesOfStops)
		if qLinesOfStops != "" {

			if err := r.SetQueryParam("linesOfStops", qLinesOfStops); err != nil {
				return err
			}
		}
	}

	if o.National != nil {

		// query param national
		var qrNational bool

		if o.National != nil {
			qrNational = *o.National
		}
		qNational := swag.FormatBool(qrNational)
		if qNational != "" {

			if err := r.SetQueryParam("national", qNational); err != nil {
				return err
			}
		}
	}

	if o.NationalExpress != nil {

		// query param nationalExpress
		var qrNationalExpress bool

		if o.NationalExpress != nil {
			qrNationalExpress = *o.NationalExpress
		}
		qNationalExpress := swag.FormatBool(qrNationalExpress)
		if qNationalExpress != "" {

			if err := r.SetQueryParam("nationalExpress", qNationalExpress); err != nil {
				return err
			}
		}
	}

	if o.Regional != nil {

		// query param regional
		var qrRegional bool

		if o.Regional != nil {
			qrRegional = *o.Regional
		}
		qRegional := swag.FormatBool(qrRegional)
		if qRegional != "" {

			if err := r.SetQueryParam("regional", qRegional); err != nil {
				return err
			}
		}
	}

	if o.RegionalExp != nil {

		// query param regionalExp
		var qrRegionalExp bool

		if o.RegionalExp != nil {
			qrRegionalExp = *o.RegionalExp
		}
		qRegionalExp := swag.FormatBool(qrRegionalExp)
		if qRegionalExp != "" {

			if err := r.SetQueryParam("regionalExp", qRegionalExp); err != nil {
				return err
			}
		}
	}

	if o.Remarks != nil {

		// query param remarks
		var qrRemarks bool

		if o.Remarks != nil {
			qrRemarks = *o.Remarks
		}
		qRemarks := swag.FormatBool(qrRemarks)
		if qRemarks != "" {

			if err := r.SetQueryParam("remarks", qRemarks); err != nil {
				return err
			}
		}
	}

	if o.Results != nil {

		// query param results
		var qrResults int64

		if o.Results != nil {
			qrResults = *o.Results
		}
		qResults := swag.FormatInt64(qrResults)
		if qResults != "" {

			if err := r.SetQueryParam("results", qResults); err != nil {
				return err
			}
		}
	}

	if o.Suburban != nil {

		// query param suburban
		var qrSuburban bool

		if o.Suburban != nil {
			qrSuburban = *o.Suburban
		}
		qSuburban := swag.FormatBool(qrSuburban)
		if qSuburban != "" {

			if err := r.SetQueryParam("suburban", qSuburban); err != nil {
				return err
			}
		}
	}

	if o.Subway != nil {

		// query param subway
		var qrSubway bool

		if o.Subway != nil {
			qrSubway = *o.Subway
		}
		qSubway := swag.FormatBool(qrSubway)
		if qSubway != "" {

			if err := r.SetQueryParam("subway", qSubway); err != nil {
				return err
			}
		}
	}

	if o.Taxi != nil {

		// query param taxi
		var qrTaxi bool

		if o.Taxi != nil {
			qrTaxi = *o.Taxi
		}
		qTaxi := swag.FormatBool(qrTaxi)
		if qTaxi != "" {

			if err := r.SetQueryParam("taxi", qTaxi); err != nil {
				return err
			}
		}
	}

	if o.Tram != nil {

		// query param tram
		var qrTram bool

		if o.Tram != nil {
			qrTram = *o.Tram
		}
		qTram := swag.FormatBool(qrTram)
		if qTram != "" {

			if err := r.SetQueryParam("tram", qTram); err != nil {
				return err
			}
		}
	}

	if o.When != nil {

		// query param when
		var qrWhen strfmt.DateTime

		if o.When != nil {
			qrWhen = *o.When
		}
		qWhen := qrWhen.String()
		if qWhen != "" {

			if err := r.SetQueryParam("when", qWhen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

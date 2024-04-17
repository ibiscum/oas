// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// AnythingAllOfObject invokes anything-all-of-object operation.
	//
	// AllOf with listed objects.
	//
	// POST /anything/all-of-object
	AnythingAllOfObject(ctx context.Context, request OptAnythingAllOfObjectReq) error
	// AnythingAnyOfPrimitive invokes anything-any-of-primitive operation.
	//
	// Description.
	//
	// POST /anything/any-of-primitive
	AnythingAnyOfPrimitive(ctx context.Context, request OptAnythingAnyOfPrimitiveReq) error
	// AnythingNestedOneOfObjectRef invokes anything-nested-one-of-object-ref operation.
	//
	// Description.
	//
	// POST /anything/nested-one-of-object-ref
	AnythingNestedOneOfObjectRef(ctx context.Context, request OptAnythingNestedOneOfObjectRefReq) error
	// AnythingOneOfObject invokes anything-one-of-object operation.
	//
	// Description.
	//
	// POST /anything/one-of-object
	AnythingOneOfObject(ctx context.Context, request OptAnythingOneOfObjectReq) error
	// AnythingOneOfObjectRef invokes anything-one-of-object-ref operation.
	//
	// Description.
	//
	// POST /anything/one-of-object-ref
	AnythingOneOfObjectRef(ctx context.Context, request OptAnythingOneOfObjectRefReq) error
	// AnythingOneOfPrimitive invokes anything-one-of-primitive operation.
	//
	// Description.
	//
	// POST /anything/one-of-primitive
	AnythingOneOfPrimitive(ctx context.Context, request OptAnythingOneOfPrimitiveReq) error
	// GetPets invokes get_pets operation.
	//
	// OneOf request with a nested allOf.
	//
	// PATCH /pets
	GetPets(ctx context.Context, request OptGetPetsReq) error
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// AnythingAllOfObject invokes anything-all-of-object operation.
//
// AllOf with listed objects.
//
// POST /anything/all-of-object
func (c *Client) AnythingAllOfObject(ctx context.Context, request OptAnythingAllOfObjectReq) error {
	_, err := c.sendAnythingAllOfObject(ctx, request)
	return err
}

func (c *Client) sendAnythingAllOfObject(ctx context.Context, request OptAnythingAllOfObjectReq) (res *AnythingAllOfObjectOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("anything-all-of-object"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/anything/all-of-object"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AnythingAllOfObject",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/anything/all-of-object"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAnythingAllOfObjectRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAnythingAllOfObjectResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AnythingAnyOfPrimitive invokes anything-any-of-primitive operation.
//
// Description.
//
// POST /anything/any-of-primitive
func (c *Client) AnythingAnyOfPrimitive(ctx context.Context, request OptAnythingAnyOfPrimitiveReq) error {
	_, err := c.sendAnythingAnyOfPrimitive(ctx, request)
	return err
}

func (c *Client) sendAnythingAnyOfPrimitive(ctx context.Context, request OptAnythingAnyOfPrimitiveReq) (res *AnythingAnyOfPrimitiveOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("anything-any-of-primitive"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/anything/any-of-primitive"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AnythingAnyOfPrimitive",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/anything/any-of-primitive"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAnythingAnyOfPrimitiveRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAnythingAnyOfPrimitiveResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AnythingNestedOneOfObjectRef invokes anything-nested-one-of-object-ref operation.
//
// Description.
//
// POST /anything/nested-one-of-object-ref
func (c *Client) AnythingNestedOneOfObjectRef(ctx context.Context, request OptAnythingNestedOneOfObjectRefReq) error {
	_, err := c.sendAnythingNestedOneOfObjectRef(ctx, request)
	return err
}

func (c *Client) sendAnythingNestedOneOfObjectRef(ctx context.Context, request OptAnythingNestedOneOfObjectRefReq) (res *AnythingNestedOneOfObjectRefOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("anything-nested-one-of-object-ref"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/anything/nested-one-of-object-ref"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AnythingNestedOneOfObjectRef",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/anything/nested-one-of-object-ref"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAnythingNestedOneOfObjectRefRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAnythingNestedOneOfObjectRefResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AnythingOneOfObject invokes anything-one-of-object operation.
//
// Description.
//
// POST /anything/one-of-object
func (c *Client) AnythingOneOfObject(ctx context.Context, request OptAnythingOneOfObjectReq) error {
	_, err := c.sendAnythingOneOfObject(ctx, request)
	return err
}

func (c *Client) sendAnythingOneOfObject(ctx context.Context, request OptAnythingOneOfObjectReq) (res *AnythingOneOfObjectOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("anything-one-of-object"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/anything/one-of-object"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AnythingOneOfObject",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/anything/one-of-object"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAnythingOneOfObjectRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAnythingOneOfObjectResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AnythingOneOfObjectRef invokes anything-one-of-object-ref operation.
//
// Description.
//
// POST /anything/one-of-object-ref
func (c *Client) AnythingOneOfObjectRef(ctx context.Context, request OptAnythingOneOfObjectRefReq) error {
	_, err := c.sendAnythingOneOfObjectRef(ctx, request)
	return err
}

func (c *Client) sendAnythingOneOfObjectRef(ctx context.Context, request OptAnythingOneOfObjectRefReq) (res *AnythingOneOfObjectRefOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("anything-one-of-object-ref"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/anything/one-of-object-ref"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AnythingOneOfObjectRef",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/anything/one-of-object-ref"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAnythingOneOfObjectRefRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAnythingOneOfObjectRefResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AnythingOneOfPrimitive invokes anything-one-of-primitive operation.
//
// Description.
//
// POST /anything/one-of-primitive
func (c *Client) AnythingOneOfPrimitive(ctx context.Context, request OptAnythingOneOfPrimitiveReq) error {
	_, err := c.sendAnythingOneOfPrimitive(ctx, request)
	return err
}

func (c *Client) sendAnythingOneOfPrimitive(ctx context.Context, request OptAnythingOneOfPrimitiveReq) (res *AnythingOneOfPrimitiveOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("anything-one-of-primitive"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/anything/one-of-primitive"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AnythingOneOfPrimitive",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/anything/one-of-primitive"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAnythingOneOfPrimitiveRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAnythingOneOfPrimitiveResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPets invokes get_pets operation.
//
// OneOf request with a nested allOf.
//
// PATCH /pets
func (c *Client) GetPets(ctx context.Context, request OptGetPetsReq) error {
	_, err := c.sendGetPets(ctx, request)
	return err
}

func (c *Client) sendGetPets(ctx context.Context, request OptGetPetsReq) (res *GetPetsOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("get_pets"),
		semconv.HTTPMethodKey.String("PATCH"),
		semconv.HTTPRouteKey.String("/pets"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "GetPets",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/pets"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PATCH", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeGetPetsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPetsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// Package schemas provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package schemas

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// N5StartsWithNumber defines model for 5StartsWithNumber.
type N5StartsWithNumber map[string]interface{}

// AnyType1 defines model for AnyType1.
type AnyType1 interface{}

// AnyType2 defines model for AnyType2.
type AnyType2 interface{}

// CustomStringType defines model for CustomStringType.
type CustomStringType string

// GenericObject defines model for GenericObject.
type GenericObject map[string]interface{}

// Issue9JSONBody defines parameters for Issue9.
type Issue9JSONBody interface{}

// Issue9Params defines parameters for Issue9.
type Issue9Params struct {
	Foo string `json:"foo"`
}

// Issue9RequestBody defines body for Issue9 for application/json ContentType.
type Issue9JSONRequestBody Issue9JSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// Issue127 request
	Issue127(ctx context.Context) (*http.Response, error)

	// Issue30 request
	Issue30(ctx context.Context, pFallthrough string) (*http.Response, error)

	// Issue41 request
	Issue41(ctx context.Context, n1param N5StartsWithNumber) (*http.Response, error)

	// Issue9 request  with any body
	Issue9WithBody(ctx context.Context, params *Issue9Params, contentType string, body io.Reader) (*http.Response, error)

	Issue9(ctx context.Context, params *Issue9Params, body Issue9JSONRequestBody) (*http.Response, error)
}

func (c *Client) Issue127(ctx context.Context) (*http.Response, error) {
	req, err := NewIssue127Request(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue30(ctx context.Context, pFallthrough string) (*http.Response, error) {
	req, err := NewIssue30Request(c.Server, pFallthrough)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue41(ctx context.Context, n1param N5StartsWithNumber) (*http.Response, error) {
	req, err := NewIssue41Request(c.Server, n1param)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue9WithBody(ctx context.Context, params *Issue9Params, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewIssue9RequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Issue9(ctx context.Context, params *Issue9Params, body Issue9JSONRequestBody) (*http.Response, error) {
	req, err := NewIssue9Request(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewIssue127Request generates requests for Issue127
func NewIssue127Request(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/127")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewIssue30Request generates requests for Issue30
func NewIssue30Request(server string, pFallthrough string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "fallthrough", pFallthrough)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/30/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewIssue41Request generates requests for Issue41
func NewIssue41Request(server string, n1param N5StartsWithNumber) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "1param", n1param)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/41/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewIssue9Request calls the generic Issue9 builder with application/json body
func NewIssue9Request(server string, params *Issue9Params, body Issue9JSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewIssue9RequestWithBody(server, params, "application/json", bodyReader)
}

// NewIssue9RequestWithBody generates requests for Issue9 with any type of body
func NewIssue9RequestWithBody(server string, params *Issue9Params, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/issues/9")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "foo", params.Foo); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		if !strings.HasSuffix(baseURL, "/") {
			baseURL += "/"
		}
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

type issue127Response struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GenericObject
	XML200       *GenericObject
	YAML200      *GenericObject
	JSONDefault  *GenericObject
}

// Status returns HTTPResponse.Status
func (r issue127Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r issue127Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type issue30Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r issue30Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r issue30Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type issue41Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r issue41Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r issue41Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type issue9Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r issue9Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r issue9Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// Issue127WithResponse request returning *Issue127Response
func (c *ClientWithResponses) Issue127WithResponse(ctx context.Context) (*issue127Response, error) {
	rsp, err := c.Issue127(ctx)
	if err != nil {
		return nil, err
	}
	return ParseIssue127Response(rsp)
}

// Issue30WithResponse request returning *Issue30Response
func (c *ClientWithResponses) Issue30WithResponse(ctx context.Context, pFallthrough string) (*issue30Response, error) {
	rsp, err := c.Issue30(ctx, pFallthrough)
	if err != nil {
		return nil, err
	}
	return ParseIssue30Response(rsp)
}

// Issue41WithResponse request returning *Issue41Response
func (c *ClientWithResponses) Issue41WithResponse(ctx context.Context, n1param N5StartsWithNumber) (*issue41Response, error) {
	rsp, err := c.Issue41(ctx, n1param)
	if err != nil {
		return nil, err
	}
	return ParseIssue41Response(rsp)
}

// Issue9WithBodyWithResponse request with arbitrary body returning *Issue9Response
func (c *ClientWithResponses) Issue9WithBodyWithResponse(ctx context.Context, params *Issue9Params, contentType string, body io.Reader) (*issue9Response, error) {
	rsp, err := c.Issue9WithBody(ctx, params, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseIssue9Response(rsp)
}

func (c *ClientWithResponses) Issue9WithResponse(ctx context.Context, params *Issue9Params, body Issue9JSONRequestBody) (*issue9Response, error) {
	rsp, err := c.Issue9(ctx, params, body)
	if err != nil {
		return nil, err
	}
	return ParseIssue9Response(rsp)
}

// ParseIssue127Response parses an HTTP response from a Issue127WithResponse call
func ParseIssue127Response(rsp *http.Response) (*issue127Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &issue127Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GenericObject
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.YAML200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		var dest GenericObject
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "xml") && rsp.StatusCode == 200:
		var dest GenericObject
		if err := xml.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.YAML200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "yaml") && rsp.StatusCode == 200:
		var dest GenericObject
		if err := yaml.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.YAML200 = &dest

	case rsp.StatusCode == 200:
	// Content-type (text/markdown) unsupported

	default:
		// Content-type (text/markdown) unsupported

	}

	return response, nil
}

// ParseIssue30Response parses an HTTP response from a Issue30WithResponse call
func ParseIssue30Response(rsp *http.Response) (*issue30Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &issue30Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseIssue41Response parses an HTTP response from a Issue41WithResponse call
func ParseIssue41Response(rsp *http.Response) (*issue41Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &issue41Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseIssue9Response parses an HTTP response from a Issue9WithResponse call
func ParseIssue9Response(rsp *http.Response) (*issue9Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &issue9Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /issues/127)
	Issue127(ctx echo.Context) error

	// (GET /issues/30/{fallthrough})
	Issue30(ctx echo.Context, pFallthrough string) error

	// (GET /issues/41/{1param})
	Issue41(ctx echo.Context, n1param N5StartsWithNumber) error

	// (GET /issues/9)
	Issue9(ctx echo.Context, params Issue9Params) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Issue127 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue127(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Issue127(ctx)
	return err
}

// Issue30 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue30(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "fallthrough" -------------
	var pFallthrough string

	err = runtime.BindStyledParameter("simple", false, "fallthrough", ctx.Param("fallthrough"), &pFallthrough)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fallthrough: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Issue30(ctx, pFallthrough)
	return err
}

// Issue41 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue41(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "1param" -------------
	var n1param N5StartsWithNumber

	err = runtime.BindStyledParameter("simple", false, "1param", ctx.Param("1param"), &n1param)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter 1param: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Issue41(ctx, n1param)
	return err
}

// Issue9 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue9(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params Issue9Params
	// ------------- Required query parameter "foo" -------------

	err = runtime.BindQueryParameter("form", true, true, "foo", ctx.QueryParams(), &params.Foo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter foo: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Issue9(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/issues/127", wrapper.Issue127)
	router.GET("/issues/30/:fallthrough", wrapper.Issue30)
	router.GET("/issues/41/:1param", wrapper.Issue41)
	router.GET("/issues/9", wrapper.Issue9)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RVTW/bOBD9K4PZBfZiSHaSRRHd2hyKHNoETYAemhxocWQxkUiGHDoRBP33gpRdKx9O",
	"EaQ9mZI4M+/NvHnusTStNZo0eyx69GVNrUjH/y9YOPbfFddfQ7skF19K8qVTlpXRWOBlrTyMIaBFS+BT",
	"CNwrrkGAHsNmyJ0lLNAsb6hkHGb4UXeXnaUFFv3u6WBfgdqERsKSQGhQmslVoqR+iIlOgmfTXrBTenWZ",
	"qvRYGdcKxgLL9HFX36drMewzaXKqPBsBFf1ThMMwQ6Ur8wIi8gyl8OShMg7WwikTPCjvQ3oVtASzJges",
	"WsrgvCHhCYSUIIC3sTH0SgvdwTKsoFIPJLMrHYEqbmhb5YLcOrVvTc6P1RfZPJtHAsaSFlZhgYfZPFvg",
	"DK3gOs0tH7Hki4MP8XFF/JzFF3FL4IMjCNoHa41jklAazfTAEJvhQRr9H4N1RK1l2N1KX0e4xpITMeWp",
	"xAJPY91YdYaOvDXaUwJ0MJ/Hn5RdJzDC2kaVKTK/8RHRVnnx9K+jCgv8J99JM9/oMn88uDimaa6HtnlP",
	"qkg+b4W7leZevztRJ96DJqaRVInQ8F9s3h9iPIx4t8o7nOd9JZqGa2fCqh6e6+8b+ShuCbfU3Rsnp0tu",
	"HaWNAKXTeollQ8ldNqLbCPoF7R2m1bDCiZaYnMfiR48q1ovLgTOMWbDACbak1bugHEks2AWaTVrwxDaG",
	"6wnHo0XeL1KpYe+WnW+RTMxR6dVoj7/M8QUiR+NC/47HWP9VCq9N8bnBR4oTjsd7mZ00ijRDqu+TI4DS",
	"pXGOSm66eG6CJJksMoKLdpZYL43sQGh5pXf09nrJ8Z4u3AVy3WScxrxtjONl8vzJyO4t27XZymknztJB",
	"NInZIxiVaDyNi5HEvmEQXIMF1sy2yPONjcc/hkwS2VbYTCgcroefAQAA//9q8HMvmQcAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

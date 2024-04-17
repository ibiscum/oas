// Package discriminators provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package discriminators

import (
	"net/http"
	"reflect"
	"testing"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestOneOfWithTopLevelDiscriminatorAndMappingResponse_Status(t *testing.T) {
	type fields struct {
		Body         []byte
		HTTPResponse *http.Response
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := OneOfWithTopLevelDiscriminatorAndMappingResponse{
				Body:         tt.fields.Body,
				HTTPResponse: tt.fields.HTTPResponse,
			}
			if got := r.Status(); got != tt.want {
				t.Errorf("OneOfWithTopLevelDiscriminatorAndMappingResponse.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		server string
		opts   []ClientOption
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.server, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
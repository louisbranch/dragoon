package parser

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tcs := map[string]struct {
		input    string
		services []Service
		err      error
	}{
		"ok": {
			input: `
				service Balancer {
					// dragoon:ignore
					rpc Check(HealthCheckRequest) returns (HealthCheckResponse) {}

					rpc CreditAccount(CreditAccountRequest) returns (CreditAccountResponse) {}

					rpc DebitAccount(DebitAccountRequest) returns (DebitAccountResponse) {}
				}`,
			services: []Service{
				{
					Name: "Balancer",
					RPCs: []RPC{
						{Name: "CreditAccount"},
						{Name: "DebitAccount"},
					},
				},
			},
		},
		"empty file": {},
		"invalid input": {
			input: "Service{",
			err:   errors.New(`<input>:1:1: found "Service" but expected [.proto element {comment|option|import|syntax|enum|service|package|message}]`),
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {

			tmpfile, err := ioutil.TempFile("", "proto")
			if err != nil {
				t.Fatal(err)
			}

			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tc.input)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			services, err := Parse(tmpfile.Name())

			if !reflect.DeepEqual(tc.services, services) {
				t.Errorf("expected %v, got %v", tc.services, services)
			}

			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected %q, got %q", tc.err.Error(), err.Error())
			}

		})
	}
}

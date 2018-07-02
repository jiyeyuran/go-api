// Package grpc resolves using http path
package grpc

import (
	"errors"
	"net/http"
	"strings"
)

type Resolver struct{}

func (r *Resolver) Resolve(req *http.Request) (string, error) {
	// /foo.Bar/Service
	if req.URL.Path == "/" {
		return "", errors.New("unknown name")
	}
	// [foo.Bar, Service]
	parts := strings.Split(req.URL.Path[1:], "/")
	// [foo, Bar]
	name := strings.Split(parts[0], ".")
	// foo
	return strings.Join(name[:len(name)-1], "."), nil
}

func (r *Resolver) String() string {
	return "grpc"
}

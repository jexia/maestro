package protocol

import (
	"context"
	"io"

	"github.com/jexia/maestro/codec"
	"github.com/jexia/maestro/header"
	"github.com/jexia/maestro/refs"
	"github.com/jexia/maestro/schema"
	"github.com/jexia/maestro/specs"
)

// ResponseWriter specifies the response writer implementation which could be used to both proxy forward a request or used to call a service
type ResponseWriter interface {
	Header() header.Store
	Write([]byte) (int, error)
	WriteHeader(int)
}

// Request represents the request object given to a caller implementation used to make calls
type Request struct {
	Header  header.Store
	Body    io.Reader
	Context context.Context
}

// Callers represents a collection of callers
type Callers []Caller

// Get attempts to return a caller with the given name
func (collection Callers) Get(name string) Caller {
	for _, caller := range collection {
		if caller.Name() == name {
			return caller
		}
	}

	return nil
}

// Caller constructs new calls which could be used to call services
type Caller interface {
	Name() string
	New(schema schema.Service, method string, functions specs.CustomDefinedFunctions, options schema.Options) (Call, error)
}

// Call is a preconfigured interface for a single service
type Call interface {
	Call(writer ResponseWriter, request *Request, refs *refs.Store) error
	References() []*specs.Property
	Close() error
}

// Listeners represents a collection of listeners
type Listeners []Listener

// Get attempts to return a listener with the given name
func (collection Listeners) Get(name string) Listener {
	for _, listener := range collection {
		if listener.Name() == name {
			return listener
		}
	}

	return nil
}

// Flow represents a flow which could be called by a protocol
type Flow interface {
	NewStore() *refs.Store
	GetName() string
	Call(ctx context.Context, refs *refs.Store) error
	Wait()
}

// Endpoint represents a protocol listener endpoint
type Endpoint struct {
	Listener string
	Flow     Flow
	Request  *specs.ParameterMap
	Response *specs.ParameterMap
	Forward  Call
	Options  specs.Options
}

// Listener specifies the listener implementation
type Listener interface {
	Name() string
	Serve() error
	Close() error
	Handle([]*Endpoint, map[string]codec.Constructor) error
}

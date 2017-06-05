package kit

import (
	"net/http"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// HTTPEndpoint encapsulates everything required to build
// an endpoint hosted on an HTTP server.
type HTTPEndpoint struct {
	Endpoint endpoint.Endpoint
	Decoder  httptransport.DecodeRequestFunc
	Encoder  httptransport.EncodeResponseFunc
	Options  []httptransport.ServerOption
}

// Service is the interface of mixed HTTP/gRPC that can be registered and
// hosted by a gizmo/server/kit server. Services provide hooks for service-wide options
// and middlewares and can be used as a means of dependency injection.
// In general, a Service should just contain the logic for deserializing HTTP
// requests, passing the request to a business logic interface abstraction,
// handling errors and serializing the apprioriate response.
//
// In other words, each Endpoint is similar to a 'controller' and the Service
// a container for injecting depedencies (business services, repositories, etc.)
// into each request handler.
//
type Service interface {
	// HTTPMiddleware is for service-wide http specific middlewares
	// for easy integration with 3rd party http.Handlers.
	HTTPMiddleware(http.Handler) http.Handler

	// HTTPOptions are service-wide go-kit HTTP server options
	HTTPOptions() []httptransport.ServerOption

	// Middleware is for any service-wide go-kit middlewares
	Middleware(endpoint.Endpoint) endpoint.Endpoint

	// HTTPRouterOptions allows users to override the default
	// behavior and use of the GorillaRouter.
	HTTPRouterOptions() []RouterOption

	// HTTPEndpoints default to using a JSON serializer if no encoder is provided.
	// For example:
	//
	//    return map[string]map[string]kithttp.HTTPEndpoint{
	//        "/cat/{id}": {
	//            "GET": {
	//                Endpoint: s.GetCatByID,
	//                Decoder:  decodeGetCatRequest,
	//            },
	//        },
	//        "/cats": {
	//            "PUT": {
	//                Endpoint: s.PutCats,
	//                HTTPDecoder:  decodePutCatsProtoRequest,
	//            },
	//            "GET": {
	//                Endpoint: s.GetCats,
	//                HTTPDecoder:  decodeGetCatsRequest,
	//            },
	//        },
	//  }
	HTTPEndpoints() map[string]map[string]HTTPEndpoint

	// RPCServiceDesc allows services to declare an alternate gRPC
	// representation of themselves ot be hosted on the RPC_PORT.
	RPCServiceDesc() *grpc.ServiceDesc

	// RPCOptions are for service-wide gRPC server options.
	RPCOptions() []grpc.ServerOption
}

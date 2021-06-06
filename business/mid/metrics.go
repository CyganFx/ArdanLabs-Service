package mid

import (
	"context"
	"expvar"
	"github.com/CyganFx/ArdanLabs-Service/foundation/web"
	"net/http"
	"runtime"
)

const curExpectedMaxNumOfRequestsPerSec = 100

// m contains the global program counters for the application.
var m = struct {
	gr  *expvar.Int
	req *expvar.Int
	err *expvar.Int
}{
	gr:  expvar.NewInt("goroutines"),
	req: expvar.NewInt("requests"),
	err: expvar.NewInt("errors"),
}

// Metrics updates program counters.
func Metrics() web.Middleware {

	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// Call the next handler.
			err := handler(ctx, w, r)

			m.req.Add(1)

			// Update the count for the number of active goroutines every 100 requests.
			if m.req.Value()%curExpectedMaxNumOfRequestsPerSec == 0 {
				m.gr.Set(int64(runtime.NumGoroutine()))
			}

			// Increment the errors counter if an error occurred on this request.
			if err != nil {
				m.err.Add(1)
			}

			// Return the error so it can be handled further up the chain.
			return err
		}

		return h
	}

	return m
}

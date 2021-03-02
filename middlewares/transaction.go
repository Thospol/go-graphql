package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	c "context"

	"github.com/go-chi/chi/middleware"
	"github.com/thospol/go-graphql/core/context"
	"github.com/thospol/go-graphql/core/postgres"
)

// Transaction to do transaction my sql
func Transaction(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		database := postgres.GetDatabase().Begin()
		timeout, cancel := c.WithTimeout(c.Background(), time.Second*5)
		context.SetDatabase(r, database.WithContext(timeout))
		next.ServeHTTP(w, r)

		if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
			_ = database.Rollback()

			logEntry := middleware.GetLogEntry(r)
			if logEntry != nil {
				logEntry.Panic(rvr, debug.Stack())
			} else {
				fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)

			}

			w.WriteHeader(http.StatusInternalServerError)
		}

		if context.GetErrMsg(r) != "" || database.Commit().Error != nil {
			_ = database.Rollback()
		}

		func() {
			defer cancel()
		}()
	}

	return http.HandlerFunc(fn)
}

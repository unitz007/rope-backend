package middlewares

import (
	"net/http"
	"ropc-service/handlers"
	"ropc-service/logger"
	"ropc-service/model"
)

func PanicRecovery(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				errorMsg := "Something went wrong"
				if e, ok := err.(error); ok {
					errorMsg = e.Error()
				}
				logger.Error(errorMsg, false)
				_ = handlers.PrintResponse(http.StatusBadRequest, w, model.NewResponse[any](errorMsg, nil))
			}
		}()

		h(w, r)
	}
}

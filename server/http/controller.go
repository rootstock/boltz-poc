package http

import (
	"encoding/json"
	"net/http"

	"github.com/patogallaiov/boltz-poc/storage"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func Controller(s *Server) *mux.Router {
	r := mux.NewRouter()
	corsMw := mux.CORSMethodMiddleware(r)

	// ROUTES

	// GET
	r.Path("/swapinfo").Methods(http.MethodGet).HandlerFunc(wrap(getSwapInfo, s))
	r.Path("/getpairs").Methods(http.MethodGet).HandlerFunc(wrap(getPair, s))
	r.Path("/health").Methods(http.MethodGet).HandlerFunc(wrap(checkHealthHandler, s))
	r.Path("/payment/").Methods(http.MethodGet, http.MethodOptions).HandlerFunc(wrap(getPayments, s))
	r.Path("/payment/{id}").Methods(http.MethodGet, http.MethodOptions).HandlerFunc(wrap(getPayments, s))
	r.Path("/config/").Methods(http.MethodGet, http.MethodOptions).HandlerFunc(wrap(getConfigs, s))
	r.Path("/config/{id}").Methods(http.MethodGet, http.MethodOptions).HandlerFunc(wrap(getConfigs, s))

	// POST
	r.Path("/payment").Methods(http.MethodPost, http.MethodOptions).HandlerFunc(wrap(createPayment, s))
	r.Path("/config").Methods(http.MethodPost, http.MethodOptions).HandlerFunc(wrap(saveConfig, s))

	r.Use(corsMw)
	return r
}

func getSwapInfo(s *Server, r *http.Request) (any, error) {
	return s.boltz.GetReverseSwapInfo()
}

func getPair(s *Server, r *http.Request) (any, error) {
	return s.boltz.GetPair()
}

func getPayments(s *Server, r *http.Request) (any, error) {
	vars := mux.Vars(r)
	key := vars["id"]
	if len(key) > 0 {
		return s.checkout.GetInvoice(key)
	}
	return s.checkout.GetInvoices()
}

func createPayment(s *Server, r *http.Request) (any, error) {
	request := storage.PaymentRequest{}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&request); err != nil {
		return nil, err
	}
	return s.checkout.CreateInvoice(request)
}

func saveConfig(s *Server, r *http.Request) (any, error) {
	request := storage.ConfigRequest{}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&request); err != nil {
		return nil, err
	}
	return s.config.SaveConfig(&storage.Config{
		Key:   request.Key,
		Value: request.Value,
	})
}

func getConfigs(s *Server, r *http.Request) (any, error) {
	vars := mux.Vars(r)
	key := vars["id"]
	if len(key) > 0 {
		return s.config.GetConfig(key)
	}
	return s.config.GetConfigs()
}

func checkHealthHandler(s *Server, r *http.Request) (any, error) {
	dbSvcStatus := svcStatusOk
	rskSvcStatus := svcStatusOk

	if err := s.db.CheckConnection(); err != nil {
		log.Error("error checking db connection status: ", err.Error())
		dbSvcStatus = svcStatusUnreachable
	}

	if err := s.rsk.CheckConnection(); err != nil {
		log.Error("error checking rsk connection status: ", err.Error())
		rskSvcStatus = svcStatusUnreachable
	}

	return healthRes{
		Services: servicesType{
			Db:  dbSvcStatus,
			Rsk: rskSvcStatus,
		},
	}, nil
}

func wrap(action func(s *Server, r *http.Request) (any, error), server *Server) (f func(http.ResponseWriter, *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := action(server, r)
		if err != nil {
			log.Error("error response: ", err.Error())
			response = &struct {
				message string
			}{
				message: "internal server error:" + err.Error(),
			}
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		err = enc.Encode(response)
		if err != nil {
			log.Error("error encoding response: ", err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}

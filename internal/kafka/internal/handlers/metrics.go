package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/public"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/metrics"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/presenters"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/services"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/handlers"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/shared"
	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/client/observatorium"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/gorilla/mux"
)

type metricsHandler struct {
	service services.ObservatoriumService
}

func NewMetricsHandler(service services.ObservatoriumService) *metricsHandler {
	return &metricsHandler{
		service: service,
	}
}

func (h metricsHandler) FederateMetrics(w http.ResponseWriter, r *http.Request) {
	kafkaId := strings.TrimSpace(mux.Vars(r)["id"])
	if kafkaId == "" {
		shared.HandleError(r, w, &errors.ServiceError{
			Code:     errors.ErrorBadRequest,
			Reason:   "missing path parameter: kafka id",
			HttpCode: http.StatusBadRequest,
		})
		return
	}

	kafkaMetrics := &observatorium.KafkaMetrics{}
	params := observatorium.MetricsReqParams{
		ResultType: observatorium.Query,
	}

	_, err := h.service.GetMetricsByKafkaId(r.Context(), kafkaMetrics, kafkaId, params)
	if err != nil {
		if err.Code == errors.ErrorNotFound {
			shared.HandleError(r, w, err)
		} else {
			glog.Errorf("Error getting metrics: %v", err)
			sentry.CaptureException(err)
			shared.HandleError(r, w, &errors.ServiceError{
				Code:     err.Code,
				Reason:   "error getting metrics",
				HttpCode: http.StatusInternalServerError,
			})
		}
		return
	}

	// Define metric collector
	collector := metrics.NewFederatedUserMetricsCollector(kafkaMetrics)
	registry := prometheus.NewPedanticRegistry()
	registry.MustRegister(collector)

	promHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		ErrorHandling: promhttp.HTTPErrorOnError,
	})
	promHandler.ServeHTTP(w, r)
}

func (h metricsHandler) GetMetricsByRangeQuery(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	params := observatorium.MetricsReqParams{}
	query := r.URL.Query()
	cfg := &handlers.HandlerConfig{
		Validate: []handlers.Validate{
			handlers.ValidatQueryParam(query, "duration"),
			handlers.ValidatQueryParam(query, "interval"),
		},
		Action: func() (i interface{}, serviceError *errors.ServiceError) {
			ctx := r.Context()
			params.ResultType = observatorium.RangeQuery
			extractMetricsQueryParams(r, &params)
			kafkaMetrics := &observatorium.KafkaMetrics{}
			foundKafkaId, err := h.service.GetMetricsByKafkaId(ctx, kafkaMetrics, id, params)
			if err != nil {
				return nil, err
			}
			metricList := public.MetricsRangeQueryList{
				Kind: "MetricsRangeQueryList",
				Id:   foundKafkaId,
			}
			metricsResult, err := presenters.PresentMetricsByRangeQuery(kafkaMetrics)
			if err != nil {
				return nil, err
			}
			metricList.Items = metricsResult

			return metricList, nil
		},
	}
	handlers.HandleGet(w, r, cfg)
}

func (h metricsHandler) GetMetricsByInstantQuery(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	params := observatorium.MetricsReqParams{}
	cfg := &handlers.HandlerConfig{
		Action: func() (i interface{}, serviceError *errors.ServiceError) {
			ctx := r.Context()
			params.ResultType = observatorium.Query
			extractMetricsQueryParams(r, &params)
			kafkaMetrics := &observatorium.KafkaMetrics{}
			foundKafkaId, err := h.service.GetMetricsByKafkaId(ctx, kafkaMetrics, id, params)
			if err != nil {
				return nil, err
			}
			metricList := public.MetricsInstantQueryList{
				Kind: "MetricsInstantQueryList",
				Id:   foundKafkaId,
			}
			metricsResult, err := presenters.PresentMetricsByInstantQuery(kafkaMetrics)
			if err != nil {
				return nil, err
			}
			metricList.Items = metricsResult

			return metricList, nil
		},
	}
	handlers.HandleGet(w, r, cfg)
}

func extractMetricsQueryParams(r *http.Request, q *observatorium.MetricsReqParams) {
	q.FillDefaults()
	queryParams := r.URL.Query()
	if dur := queryParams.Get("duration"); dur != "" {
		if num, err := strconv.ParseInt(dur, 10, 64); err == nil {
			duration := time.Duration(num) * time.Minute
			q.Start = q.End.Add(-duration)
		}
	}
	if step := queryParams.Get("interval"); step != "" {
		if num, err := strconv.Atoi(step); err == nil {
			q.Step = time.Duration(num) * time.Second
		}
	}
	if filters, ok := queryParams["filters"]; ok && len(filters) > 0 {
		q.Filters = filters
	}

}

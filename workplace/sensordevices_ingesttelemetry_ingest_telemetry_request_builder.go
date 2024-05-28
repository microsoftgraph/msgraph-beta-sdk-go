package workplace

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SensordevicesIngesttelemetryIngestTelemetryRequestBuilder provides operations to call the ingestTelemetry method.
type SensordevicesIngesttelemetryIngestTelemetryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SensordevicesIngesttelemetryIngestTelemetryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SensordevicesIngesttelemetryIngestTelemetryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSensordevicesIngesttelemetryIngestTelemetryRequestBuilderInternal instantiates a new SensordevicesIngesttelemetryIngestTelemetryRequestBuilder and sets the default values.
func NewSensordevicesIngesttelemetryIngestTelemetryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SensordevicesIngesttelemetryIngestTelemetryRequestBuilder) {
    m := &SensordevicesIngesttelemetryIngestTelemetryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/workplace/sensorDevices/ingestTelemetry", pathParameters),
    }
    return m
}
// NewSensordevicesIngesttelemetryIngestTelemetryRequestBuilder instantiates a new SensordevicesIngesttelemetryIngestTelemetryRequestBuilder and sets the default values.
func NewSensordevicesIngesttelemetryIngestTelemetryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SensordevicesIngesttelemetryIngestTelemetryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSensordevicesIngesttelemetryIngestTelemetryRequestBuilderInternal(urlParams, requestAdapter)
}
// Post ingest sensor telemetry for a workplace sensor device.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/workplacesensordevice-ingesttelemetry?view=graph-rest-beta
func (m *SensordevicesIngesttelemetryIngestTelemetryRequestBuilder) Post(ctx context.Context, body SensordevicesIngesttelemetryIngestTelemetryPostRequestBodyable, requestConfiguration *SensordevicesIngesttelemetryIngestTelemetryRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation ingest sensor telemetry for a workplace sensor device.
// returns a *RequestInformation when successful
func (m *SensordevicesIngesttelemetryIngestTelemetryRequestBuilder) ToPostRequestInformation(ctx context.Context, body SensordevicesIngesttelemetryIngestTelemetryPostRequestBodyable, requestConfiguration *SensordevicesIngesttelemetryIngestTelemetryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SensordevicesIngesttelemetryIngestTelemetryRequestBuilder when successful
func (m *SensordevicesIngesttelemetryIngestTelemetryRequestBuilder) WithUrl(rawUrl string)(*SensordevicesIngesttelemetryIngestTelemetryRequestBuilder) {
    return NewSensordevicesIngesttelemetryIngestTelemetryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

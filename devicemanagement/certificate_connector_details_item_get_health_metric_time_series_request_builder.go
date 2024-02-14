package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder provides operations to call the getHealthMetricTimeSeries method.
type CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderInternal instantiates a new CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder) {
    m := &CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails%2Did}/getHealthMetricTimeSeries", pathParameters),
    }
    return m
}
// NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder instantiates a new CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getHealthMetricTimeSeries
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a CertificateConnectorDetailsItemGetHealthMetricTimeSeriesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder) Post(ctx context.Context, body CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration)(CertificateConnectorDetailsItemGetHealthMetricTimeSeriesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCertificateConnectorDetailsItemGetHealthMetricTimeSeriesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CertificateConnectorDetailsItemGetHealthMetricTimeSeriesResponseable), nil
}
// PostAsGetHealthMetricTimeSeriesPostResponse invoke action getHealthMetricTimeSeries
// returns a CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder) PostAsGetHealthMetricTimeSeriesPostResponse(ctx context.Context, body CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration)(CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostResponseable), nil
}
// ToPostRequestInformation invoke action getHealthMetricTimeSeries
// returns a *RequestInformation when successful
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder) ToPostRequestInformation(ctx context.Context, body CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder when successful
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder) WithUrl(rawUrl string)(*CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder) {
    return NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

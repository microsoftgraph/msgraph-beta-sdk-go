package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder provides operations to call the getHealthMetricTimeSeries method.
type CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderInternal instantiates a new CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder and sets the default values.
func NewCertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder) {
    m := &CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails%2Did}/getHealthMetricTimeSeries", pathParameters),
    }
    return m
}
// NewCertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder instantiates a new CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder and sets the default values.
func NewCertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getHealthMetricTimeSeries
// Deprecated: This method is obsolete. Use PostAsGetHealthMetricTimeSeriesPostResponse instead.
// returns a CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder) Post(ctx context.Context, body CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration)(CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesResponseable), nil
}
// PostAsGetHealthMetricTimeSeriesPostResponse invoke action getHealthMetricTimeSeries
// returns a CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder) PostAsGetHealthMetricTimeSeriesPostResponse(ctx context.Context, body CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration)(CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesPostResponseable), nil
}
// ToPostRequestInformation invoke action getHealthMetricTimeSeries
// returns a *RequestInformation when successful
func (m *CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder) ToPostRequestInformation(ctx context.Context, body CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder when successful
func (m *CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder) WithUrl(rawUrl string)(*CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder) {
    return NewCertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

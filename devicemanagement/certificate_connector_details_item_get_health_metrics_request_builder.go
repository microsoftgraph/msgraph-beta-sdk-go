package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder provides operations to call the getHealthMetrics method.
type CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateConnectorDetailsItemGetHealthMetricsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateConnectorDetailsItemGetHealthMetricsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilderInternal instantiates a new GetHealthMetricsRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) {
    m := &CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails%2Did}/getHealthMetrics", pathParameters),
    }
    return m
}
// NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilder instantiates a new GetHealthMetricsRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getHealthMetrics
// Deprecated: This method is obsolete. Use PostAsGetHealthMetricsPostResponse instead.
func (m *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) Post(ctx context.Context, body CertificateConnectorDetailsItemGetHealthMetricsPostRequestBodyable, requestConfiguration *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilderPostRequestConfiguration)(CertificateConnectorDetailsItemGetHealthMetricsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCertificateConnectorDetailsItemGetHealthMetricsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CertificateConnectorDetailsItemGetHealthMetricsResponseable), nil
}
// PostAsGetHealthMetricsPostResponse invoke action getHealthMetrics
func (m *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) PostAsGetHealthMetricsPostResponse(ctx context.Context, body CertificateConnectorDetailsItemGetHealthMetricsPostRequestBodyable, requestConfiguration *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilderPostRequestConfiguration)(CertificateConnectorDetailsItemGetHealthMetricsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCertificateConnectorDetailsItemGetHealthMetricsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CertificateConnectorDetailsItemGetHealthMetricsPostResponseable), nil
}
// ToPostRequestInformation invoke action getHealthMetrics
func (m *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CertificateConnectorDetailsItemGetHealthMetricsPostRequestBodyable, requestConfiguration *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) WithUrl(rawUrl string)(*CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) {
    return NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
type CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderGetQueryParameters collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
type CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderGetQueryParameters
}
// CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderInternal instantiates a new CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder and sets the default values.
func NewCertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) {
    m := &CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder instantiates a new CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder and sets the default values.
func NewCertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property certificateConnectorDetails for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
// returns a CertificateConnectorDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateConnectorDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable), nil
}
// GetHealthMetrics provides operations to call the getHealthMetrics method.
// returns a *CertificateconnectordetailsItemGethealthmetricsGetHealthMetricsRequestBuilder when successful
func (m *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) GetHealthMetrics()(*CertificateconnectordetailsItemGethealthmetricsGetHealthMetricsRequestBuilder) {
    return NewCertificateconnectordetailsItemGethealthmetricsGetHealthMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetHealthMetricTimeSeries provides operations to call the getHealthMetricTimeSeries method.
// returns a *CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder when successful
func (m *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) GetHealthMetricTimeSeries()(*CertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder) {
    return NewCertificateconnectordetailsItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property certificateConnectorDetails in deviceManagement
// returns a CertificateConnectorDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, requestConfiguration *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateConnectorDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable), nil
}
// ToDeleteRequestInformation delete navigation property certificateConnectorDetails for deviceManagement
// returns a *RequestInformation when successful
func (m *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
// returns a *RequestInformation when successful
func (m *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property certificateConnectorDetails in deviceManagement
// returns a *RequestInformation when successful
func (m *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, requestConfiguration *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder when successful
func (m *CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) WithUrl(rawUrl string)(*CertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder) {
    return NewCertificateconnectordetailsCertificateConnectorDetailsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

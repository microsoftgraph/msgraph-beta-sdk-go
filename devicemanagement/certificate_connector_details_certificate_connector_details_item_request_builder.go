package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
type CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetQueryParameters collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
type CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetQueryParameters
}
// CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderInternal instantiates a new CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) {
    m := &CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder instantiates a new CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property certificateConnectorDetails for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, error) {
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
// returns a *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder when successful
func (m *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) GetHealthMetrics()(*CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) {
    return NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetHealthMetricTimeSeries provides operations to call the getHealthMetricTimeSeries method.
// returns a *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder when successful
func (m *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) GetHealthMetricTimeSeries()(*CertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder) {
    return NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property certificateConnectorDetails in deviceManagement
// returns a CertificateConnectorDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, requestConfiguration *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, error) {
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
func (m *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
// returns a *RequestInformation when successful
func (m *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, requestConfiguration *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder when successful
func (m *CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) WithUrl(rawUrl string)(*CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) {
    return NewCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

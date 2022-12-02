package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
type DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetQueryParameters collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
type DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetQueryParameters
}
// DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderInternal instantiates a new CertificateConnectorDetailsItemRequestBuilder and sets the default values.
func NewDeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) {
    m := &DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder instantiates a new CertificateConnectorDetailsItemRequestBuilder and sets the default values.
func NewDeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property certificateConnectorDetails for deviceManagement
func (m *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property certificateConnectorDetails in deviceManagement
func (m *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, requestConfiguration *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property certificateConnectorDetails for deviceManagement
func (m *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateConnectorDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable), nil
}
// GetHealthMetrics provides operations to call the getHealthMetrics method.
func (m *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) GetHealthMetrics()(*DeviceManagementCertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) {
    return NewDeviceManagementCertificateConnectorDetailsItemGetHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetHealthMetricTimeSeries provides operations to call the getHealthMetricTimeSeries method.
func (m *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) GetHealthMetricTimeSeries()(*DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilder) {
    return NewDeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property certificateConnectorDetails in deviceManagement
func (m *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, requestConfiguration *DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateConnectorDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateConnectorDetailsable), nil
}

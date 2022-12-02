package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder provides operations to manage the microsoftTunnelServers property of the microsoft.graph.microsoftTunnelSite entity.
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderGetQueryParameters a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderGetQueryParameters
}
// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderInternal instantiates a new MicrosoftTunnelServerItemRequestBuilder and sets the default values.
func NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) {
    m := &DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers/{microsoftTunnelServer%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder instantiates a new MicrosoftTunnelServerItemRequestBuilder and sets the default values.
func NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property microsoftTunnelServers for deviceManagement
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property microsoftTunnelServers in deviceManagement
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, requestConfiguration *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateServerLogCollectionRequest provides operations to call the createServerLogCollectionRequest method.
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) CreateServerLogCollectionRequest()(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemCreateServerLogCollectionRequestRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemCreateServerLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property microsoftTunnelServers for deviceManagement
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// GenerateServerLogCollectionRequest provides operations to call the generateServerLogCollectionRequest method.
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) GenerateServerLogCollectionRequest()(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGenerateServerLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable), nil
}
// GetHealthMetrics provides operations to call the getHealthMetrics method.
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) GetHealthMetrics()(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetHealthMetricTimeSeries provides operations to call the getHealthMetricTimeSeries method.
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) GetHealthMetricTimeSeries()(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property microsoftTunnelServers in deviceManagement
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, requestConfiguration *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable), nil
}

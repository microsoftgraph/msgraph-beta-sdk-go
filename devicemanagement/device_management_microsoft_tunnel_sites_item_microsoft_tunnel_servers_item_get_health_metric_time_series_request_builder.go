package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder provides operations to call the getHealthMetricTimeSeries method.
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilderInternal instantiates a new GetHealthMetricTimeSeriesRequestBuilder and sets the default values.
func NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder) {
    m := &DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers/{microsoftTunnelServer%2Did}/microsoft.graph.getHealthMetricTimeSeries";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder instantiates a new GetHealthMetricTimeSeriesRequestBuilder and sets the default values.
func NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getHealthMetricTimeSeries
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getHealthMetricTimeSeries
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBodyable, requestConfiguration *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesResponseable), nil
}

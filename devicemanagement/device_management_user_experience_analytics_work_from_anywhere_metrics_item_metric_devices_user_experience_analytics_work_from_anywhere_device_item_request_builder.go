package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder provides operations to manage the metricDevices property of the microsoft.graph.userExperienceAnalyticsWorkFromAnywhereMetric entity.
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetQueryParameters the work from anywhere metric devices.
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetQueryParameters
}
// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) {
    m := &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/{userExperienceAnalyticsWorkFromAnywhereMetric%2Did}/metricDevices/{userExperienceAnalyticsWorkFromAnywhereDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder instantiates a new UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property metricDevices for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the work from anywhere metric devices.
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property metricDevices in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property metricDevices for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the work from anywhere metric devices.
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable), nil
}
// Patch update the navigation property metricDevices in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDeviceable), nil
}

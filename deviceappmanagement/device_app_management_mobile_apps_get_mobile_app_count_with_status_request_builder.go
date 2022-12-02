package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilder provides operations to call the getMobileAppCount method.
type DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilderInternal instantiates a new GetMobileAppCountWithStatusRequestBuilder and sets the default values.
func NewDeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, status *string)(*DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilder) {
    m := &DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/microsoft.graph.getMobileAppCount(status='{status}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if status != nil {
        urlTplParams["status"] = *status
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilder instantiates a new GetMobileAppCountWithStatusRequestBuilder and sets the default values.
func NewDeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getMobileAppCount
func (m *DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getMobileAppCount
func (m *DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsGetMobileAppCountWithStatusRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementMobileAppsGetMobileAppCountWithStatusResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAppManagementMobileAppsGetMobileAppCountWithStatusResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementMobileAppsGetMobileAppCountWithStatusResponseable), nil
}

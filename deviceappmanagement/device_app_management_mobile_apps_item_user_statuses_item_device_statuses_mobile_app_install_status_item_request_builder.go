package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder provides operations to manage the deviceStatuses property of the microsoft.graph.userAppInstallStatus entity.
type DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderGetQueryParameters the install state of the app on devices.
type DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// App provides operations to manage the app property of the microsoft.graph.mobileAppInstallStatus entity.
func (m *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) App()(*DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesItemAppRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesItemAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderInternal instantiates a new MobileAppInstallStatusItemRequestBuilder and sets the default values.
func NewDeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) {
    m := &DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/userStatuses/{userAppInstallStatus%2Did}/deviceStatuses/{mobileAppInstallStatus%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder instantiates a new MobileAppInstallStatusItemRequestBuilder and sets the default values.
func NewDeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceStatuses for deviceAppManagement
func (m *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the install state of the app on devices.
func (m *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceStatuses in deviceAppManagement
func (m *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppInstallStatusable, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceStatuses for deviceAppManagement
func (m *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the install state of the app on devices.
func (m *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppInstallStatusable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppInstallStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppInstallStatusable), nil
}
// Patch update the navigation property deviceStatuses in deviceAppManagement
func (m *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppInstallStatusable, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppInstallStatusable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppInstallStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppInstallStatusable), nil
}

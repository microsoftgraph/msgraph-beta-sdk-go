package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder provides operations to manage the userStatuses property of the microsoft.graph.mobileApp entity.
type DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderGetQueryParameters the list of installation states for this mobile app.
type DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// App provides operations to manage the app property of the microsoft.graph.userAppInstallStatus entity.
func (m *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) App()(*DeviceAppManagementMobileAppsItemUserStatusesItemAppRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemUserStatusesItemAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderInternal instantiates a new UserAppInstallStatusItemRequestBuilder and sets the default values.
func NewDeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) {
    m := &DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/userStatuses/{userAppInstallStatus%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder instantiates a new UserAppInstallStatusItemRequestBuilder and sets the default values.
func NewDeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userStatuses for deviceAppManagement
func (m *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of installation states for this mobile app.
func (m *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userStatuses in deviceAppManagement
func (m *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserAppInstallStatusable, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property userStatuses for deviceAppManagement
func (m *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.userAppInstallStatus entity.
func (m *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) DeviceStatuses()(*DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.userAppInstallStatus entity.
func (m *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) DeviceStatusesById(id string)(*DeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppInstallStatus%2Did"] = id
    }
    return NewDeviceAppManagementMobileAppsItemUserStatusesItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of installation states for this mobile app.
func (m *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserAppInstallStatusable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserAppInstallStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserAppInstallStatusable), nil
}
// Patch update the navigation property userStatuses in deviceAppManagement
func (m *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserAppInstallStatusable, requestConfiguration *DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserAppInstallStatusable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserAppInstallStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserAppInstallStatusable), nil
}

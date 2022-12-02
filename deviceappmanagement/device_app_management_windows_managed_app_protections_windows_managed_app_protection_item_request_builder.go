package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder provides operations to manage the windowsManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetQueryParameters windows managed app policies.
type DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.windowsManagedAppProtection entity.
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Apps()(*DeviceAppManagementWindowsManagedAppProtectionsItemAppsRequestBuilder) {
    return NewDeviceAppManagementWindowsManagedAppProtectionsItemAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById provides operations to manage the apps property of the microsoft.graph.windowsManagedAppProtection entity.
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) AppsById(id string)(*DeviceAppManagementWindowsManagedAppProtectionsItemAppsManagedMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp%2Did"] = id
    }
    return NewDeviceAppManagementWindowsManagedAppProtectionsItemAppsManagedMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assign provides operations to call the assign method.
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Assign()(*DeviceAppManagementWindowsManagedAppProtectionsItemAssignRequestBuilder) {
    return NewDeviceAppManagementWindowsManagedAppProtectionsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsManagedAppProtection entity.
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Assignments()(*DeviceAppManagementWindowsManagedAppProtectionsItemAssignmentsRequestBuilder) {
    return NewDeviceAppManagementWindowsManagedAppProtectionsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.windowsManagedAppProtection entity.
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) AssignmentsById(id string)(*DeviceAppManagementWindowsManagedAppProtectionsItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment%2Did"] = id
    }
    return NewDeviceAppManagementWindowsManagedAppProtectionsItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderInternal instantiates a new WindowsManagedAppProtectionItemRequestBuilder and sets the default values.
func NewDeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) {
    m := &DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/windowsManagedAppProtections/{windowsManagedAppProtection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder instantiates a new WindowsManagedAppProtectionItemRequestBuilder and sets the default values.
func NewDeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property windowsManagedAppProtections for deviceAppManagement
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation windows managed app policies.
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property windowsManagedAppProtections in deviceAppManagement
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, requestConfiguration *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property windowsManagedAppProtections for deviceAppManagement
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get windows managed app policies.
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable), nil
}
// Patch update the navigation property windowsManagedAppProtections in deviceAppManagement
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, requestConfiguration *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable), nil
}
// TargetApps provides operations to call the targetApps method.
func (m *DeviceAppManagementWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) TargetApps()(*DeviceAppManagementWindowsManagedAppProtectionsItemTargetAppsRequestBuilder) {
    return NewDeviceAppManagementWindowsManagedAppProtectionsItemTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

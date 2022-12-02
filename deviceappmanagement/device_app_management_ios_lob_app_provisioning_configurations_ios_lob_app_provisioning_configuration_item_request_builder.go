package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder provides operations to manage the iosLobAppProvisioningConfigurations property of the microsoft.graph.deviceAppManagement entity.
type DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters the IOS Lob App Provisioning Configurations.
type DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Assign()(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignRequestBuilder) {
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Assignments()(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsRequestBuilder) {
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) AssignmentsById(id string)(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosLobAppProvisioningConfigurationAssignment%2Did"] = id
    }
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal instantiates a new IosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewDeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    m := &DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder instantiates a new IosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewDeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the IOS Lob App Provisioning Configurations.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatuses()(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemDeviceStatusesRequestBuilder) {
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemDeviceStatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus%2Did"] = id
    }
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemDeviceStatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the IOS Lob App Provisioning Configurations.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignments()(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemGroupAssignmentsRequestBuilder) {
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById provides operations to manage the groupAssignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignmentsById(id string)(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemGroupAssignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppProvisioningConfigGroupAssignment%2Did"] = id
    }
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemGroupAssignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable), nil
}
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) UserStatuses()(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesRequestBuilder) {
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) UserStatusesById(id string)(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus%2Did"] = id
    }
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

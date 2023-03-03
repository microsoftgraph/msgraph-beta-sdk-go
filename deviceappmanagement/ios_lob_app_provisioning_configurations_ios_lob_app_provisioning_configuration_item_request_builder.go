package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder provides operations to manage the iosLobAppProvisioningConfigurations property of the microsoft.graph.deviceAppManagement entity.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters the IOS Lob App Provisioning Configurations.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters
}
// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Assign()(*IosLobAppProvisioningConfigurationsItemAssignRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Assignments()(*IosLobAppProvisioningConfigurationsItemAssignmentsRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) AssignmentsById(id string)(*IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosLobAppProvisioningConfigurationAssignment%2Did"] = id
    }
    return NewIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal instantiates a new IosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    m := &IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder instantiates a new IosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatuses()(*IosLobAppProvisioningConfigurationsItemDeviceStatusesRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*IosLobAppProvisioningConfigurationsItemDeviceStatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus%2Did"] = id
    }
    return NewIosLobAppProvisioningConfigurationsItemDeviceStatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignments()(*IosLobAppProvisioningConfigurationsItemGroupAssignmentsRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupAssignmentsById provides operations to manage the groupAssignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignmentsById(id string)(*IosLobAppProvisioningConfigurationsItemGroupAssignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppProvisioningConfigGroupAssignment%2Did"] = id
    }
    return NewIosLobAppProvisioningConfigurationsItemGroupAssignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) UserStatuses()(*IosLobAppProvisioningConfigurationsItemUserStatusesRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) UserStatusesById(id string)(*IosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus%2Did"] = id
    }
    return NewIosLobAppProvisioningConfigurationsItemUserStatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

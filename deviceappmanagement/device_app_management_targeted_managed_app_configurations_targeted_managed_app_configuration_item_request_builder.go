package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters targeted managed app configurations.
type DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Apps()(*DeviceAppManagementTargetedManagedAppConfigurationsItemAppsRequestBuilder) {
    return NewDeviceAppManagementTargetedManagedAppConfigurationsItemAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById provides operations to manage the apps property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) AppsById(id string)(*DeviceAppManagementTargetedManagedAppConfigurationsItemAppsManagedMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp%2Did"] = id
    }
    return NewDeviceAppManagementTargetedManagedAppConfigurationsItemAppsManagedMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assign provides operations to call the assign method.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Assign()(*DeviceAppManagementTargetedManagedAppConfigurationsItemAssignRequestBuilder) {
    return NewDeviceAppManagementTargetedManagedAppConfigurationsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Assignments()(*DeviceAppManagementTargetedManagedAppConfigurationsItemAssignmentsRequestBuilder) {
    return NewDeviceAppManagementTargetedManagedAppConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) AssignmentsById(id string)(*DeviceAppManagementTargetedManagedAppConfigurationsItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment%2Did"] = id
    }
    return NewDeviceAppManagementTargetedManagedAppConfigurationsItemAssignmentsTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewDeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) {
    m := &DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewDeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation targeted managed app configurations.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppConfigurationable, requestConfiguration *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploymentSummary provides operations to manage the deploymentSummary property of the microsoft.graph.targetedManagedAppConfiguration entity.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) DeploymentSummary()(*DeviceAppManagementTargetedManagedAppConfigurationsItemDeploymentSummaryRequestBuilder) {
    return NewDeviceAppManagementTargetedManagedAppConfigurationsItemDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get targeted managed app configurations.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTargetedManagedAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppConfigurationable), nil
}
// Patch update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppConfigurationable, requestConfiguration *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppConfigurationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTargetedManagedAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetedManagedAppConfigurationable), nil
}
// TargetApps provides operations to call the targetApps method.
func (m *DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) TargetApps()(*DeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsRequestBuilder) {
    return NewDeviceAppManagementTargetedManagedAppConfigurationsItemTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

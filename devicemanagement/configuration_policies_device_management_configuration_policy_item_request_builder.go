package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
type ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters list of all Configuration policies
type ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters
}
// ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Assignments()(*ConfigurationPoliciesItemAssignmentsRequestBuilder) {
    return NewConfigurationPoliciesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) AssignmentsById(id string)(*ConfigurationPoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment%2Did"] = id
    }
    return NewConfigurationPoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal instantiates a new DeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    m := &ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder instantiates a new DeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property configurationPolicies for deviceManagement
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of all Configuration policies
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable), nil
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) MicrosoftGraphAssign()(*ConfigurationPoliciesItemMicrosoftGraphAssignRequestBuilder) {
    return NewConfigurationPoliciesItemMicrosoftGraphAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateCopy provides operations to call the createCopy method.
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) MicrosoftGraphCreateCopy()(*ConfigurationPoliciesItemMicrosoftGraphCreateCopyRequestBuilder) {
    return NewConfigurationPoliciesItemMicrosoftGraphCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReorder provides operations to call the reorder method.
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) MicrosoftGraphReorder()(*ConfigurationPoliciesItemMicrosoftGraphReorderRequestBuilder) {
    return NewConfigurationPoliciesItemMicrosoftGraphReorderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property configurationPolicies in deviceManagement
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Settings()(*ConfigurationPoliciesItemSettingsRequestBuilder) {
    return NewConfigurationPoliciesItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) SettingsById(id string)(*ConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting%2Did"] = id
    }
    return NewConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property configurationPolicies for deviceManagement
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation list of all Configuration policies
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property configurationPolicies in deviceManagement
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

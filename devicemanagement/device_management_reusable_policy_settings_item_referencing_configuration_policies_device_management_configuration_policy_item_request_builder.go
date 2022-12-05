package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder provides operations to manage the referencingConfigurationPolicies property of the microsoft.graph.deviceManagementReusablePolicySetting entity.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters configuration policies referencing the current reusable setting. This property is read-only.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters
}
// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Assign()(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder) {
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Assignments()(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignmentsRequestBuilder) {
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment%2Did"] = id
    }
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal instantiates a new DeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    m := &DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder instantiates a new DeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateCopy provides operations to call the createCopy method.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) CreateCopy()(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyRequestBuilder) {
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property referencingConfigurationPolicies for deviceManagement
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property referencingConfigurationPolicies in deviceManagement
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property referencingConfigurationPolicies for deviceManagement
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable), nil
}
// Patch update the navigation property referencingConfigurationPolicies in deviceManagement
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable), nil
}
// Reorder provides operations to call the reorder method.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Reorder()(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderRequestBuilder) {
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Settings()(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsRequestBuilder) {
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) SettingsById(id string)(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting%2Did"] = id
    }
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder provides operations to manage the referencingConfigurationPolicies property of the microsoft.graph.deviceManagementReusablePolicySetting entity.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters configuration policies referencing the current reusable setting. This property is read-only.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters
}
// ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Assign()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Assignments()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignmentsRequestBuilder) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) AssignmentsById(id string)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment%2Did"] = id
    }
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal instantiates a new DeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    m := &ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder instantiates a new DeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateCopy provides operations to call the createCopy method.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) CreateCopy()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyRequestBuilder) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemCreateCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property referencingConfigurationPolicies for deviceManagement
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Get configuration policies referencing the current reusable setting. This property is read-only.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable), nil
}
// Patch update the navigation property referencingConfigurationPolicies in deviceManagement
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable), nil
}
// Reorder provides operations to call the reorder method.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Reorder()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderRequestBuilder) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemReorderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Settings()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsRequestBuilder) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) SettingsById(id string)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting%2Did"] = id
    }
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property referencingConfigurationPolicies for deviceManagement
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation configuration policies referencing the current reusable setting. This property is read-only.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property referencingConfigurationPolicies in deviceManagement
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

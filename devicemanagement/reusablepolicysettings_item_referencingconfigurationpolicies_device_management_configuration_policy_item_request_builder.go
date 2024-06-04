package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder provides operations to manage the referencingConfigurationPolicies property of the microsoft.graph.deviceManagementReusablePolicySetting entity.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters configuration policies referencing the current reusable setting. This property is read-only.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetQueryParameters
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Assign()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Assignments()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ClearEnrollmentTimeDeviceMembershipTarget provides operations to call the clearEnrollmentTimeDeviceMembershipTarget method.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ClearEnrollmentTimeDeviceMembershipTarget()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    m := &ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateCopy provides operations to call the createCopy method.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) CreateCopy()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemCreatecopyCreateCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property referencingConfigurationPolicies for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get configuration policies referencing the current reusable setting. This property is read-only.
// returns a DeviceManagementConfigurationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a DeviceManagementConfigurationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemReorderRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Reorder()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemReorderRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemReorderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveEnrollmentTimeDeviceMembershipTarget provides operations to call the retrieveEnrollmentTimeDeviceMembershipTarget method.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrieveenrollmenttimedevicemembershiptargetRetrieveEnrollmentTimeDeviceMembershipTargetRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) RetrieveEnrollmentTimeDeviceMembershipTarget()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrieveenrollmenttimedevicemembershiptargetRetrieveEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrieveenrollmenttimedevicemembershiptargetRetrieveEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveLatestUpgradeDefaultBaselinePolicy provides operations to call the retrieveLatestUpgradeDefaultBaselinePolicy method.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) RetrieveLatestUpgradeDefaultBaselinePolicy()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Settings()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property referencingConfigurationPolicies for deviceManagement
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation configuration policies referencing the current reusable setting. This property is read-only.
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property referencingConfigurationPolicies in deviceManagement
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesDeviceManagementConfigurationPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

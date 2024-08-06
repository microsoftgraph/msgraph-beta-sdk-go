package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
type ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// Assign provides operations to call the assign method.
// returns a *ConfigurationPoliciesItemAssignRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Assign()(*ConfigurationPoliciesItemAssignRequestBuilder) {
    return NewConfigurationPoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
// returns a *ConfigurationPoliciesItemAssignmentsRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Assignments()(*ConfigurationPoliciesItemAssignmentsRequestBuilder) {
    return NewConfigurationPoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ClearEnrollmentTimeDeviceMembershipTarget provides operations to call the clearEnrollmentTimeDeviceMembershipTarget method.
// returns a *ConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ClearEnrollmentTimeDeviceMembershipTarget()(*ConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    return NewConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal instantiates a new ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    m := &ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder instantiates a new ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder and sets the default values.
func NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateCopy provides operations to call the createCopy method.
// returns a *ConfigurationPoliciesItemCreateCopyRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) CreateCopy()(*ConfigurationPoliciesItemCreateCopyRequestBuilder) {
    return NewConfigurationPoliciesItemCreateCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property configurationPolicies for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of all Configuration policies
// returns a DeviceManagementConfigurationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
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
// Patch update the navigation property configurationPolicies in deviceManagement
// returns a DeviceManagementConfigurationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
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
// returns a *ConfigurationPoliciesItemReorderRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Reorder()(*ConfigurationPoliciesItemReorderRequestBuilder) {
    return NewConfigurationPoliciesItemReorderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveEnrollmentTimeDeviceMembershipTarget provides operations to call the retrieveEnrollmentTimeDeviceMembershipTarget method.
// returns a *ConfigurationPoliciesItemRetrieveEnrollmentTimeDeviceMembershipTargetRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) RetrieveEnrollmentTimeDeviceMembershipTarget()(*ConfigurationPoliciesItemRetrieveEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    return NewConfigurationPoliciesItemRetrieveEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveLatestUpgradeDefaultBaselinePolicy provides operations to call the retrieveLatestUpgradeDefaultBaselinePolicy method.
// returns a *ConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) RetrieveLatestUpgradeDefaultBaselinePolicy()(*ConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) {
    return NewConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetEnrollmentTimeDeviceMembershipTarget provides operations to call the setEnrollmentTimeDeviceMembershipTarget method.
// returns a *ConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) SetEnrollmentTimeDeviceMembershipTarget()(*ConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    return NewConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
// returns a *ConfigurationPoliciesItemSettingsRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) Settings()(*ConfigurationPoliciesItemSettingsRequestBuilder) {
    return NewConfigurationPoliciesItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property configurationPolicies for deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of all Configuration policies
// returns a *RequestInformation when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property configurationPolicies in deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, requestConfiguration *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder when successful
func (m *ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    return NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

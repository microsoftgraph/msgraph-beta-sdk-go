package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder provides operations to call the retrieveLatestUpgradeDefaultBaselinePolicy method.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderInternal instantiates a new RetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) {
    m := &ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}/retrieveLatestUpgradeDefaultBaselinePolicy()", pathParameters),
    }
    return m
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder instantiates a new RetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveLatestUpgradeDefaultBaselinePolicy
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
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
// ToGetRequestInformation invoke function retrieveLatestUpgradeDefaultBaselinePolicy
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) WithUrl(rawUrl string)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

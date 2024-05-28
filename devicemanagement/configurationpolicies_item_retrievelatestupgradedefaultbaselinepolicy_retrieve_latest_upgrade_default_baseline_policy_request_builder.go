package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder provides operations to call the retrieveLatestUpgradeDefaultBaselinePolicy method.
type ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderInternal instantiates a new ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder and sets the default values.
func NewConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) {
    m := &ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy%2Did}/retrieveLatestUpgradeDefaultBaselinePolicy()", pathParameters),
    }
    return m
}
// NewConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder instantiates a new ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder and sets the default values.
func NewConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveLatestUpgradeDefaultBaselinePolicy
// returns a DeviceManagementConfigurationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
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
// ToGetRequestInformation invoke function retrieveLatestUpgradeDefaultBaselinePolicy
// returns a *RequestInformation when successful
func (m *ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder when successful
func (m *ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) WithUrl(rawUrl string)(*ConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder) {
    return NewConfigurationpoliciesItemRetrievelatestupgradedefaultbaselinepolicyRetrieveLatestUpgradeDefaultBaselinePolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder provides operations to call the clearEnrollmentTimeDeviceMembershipTarget method.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal instantiates a new ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    m := &ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}/clearEnrollmentTimeDeviceMembershipTarget", pathParameters),
    }
    return m
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder instantiates a new ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action clearEnrollmentTimeDeviceMembershipTarget
// Deprecated: This method is obsolete. Use PostAsClearEnrollmentTimeDeviceMembershipTargetPostResponse instead.
// returns a ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) Post(ctx context.Context, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration)(ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetResponseable), nil
}
// PostAsClearEnrollmentTimeDeviceMembershipTargetPostResponse invoke action clearEnrollmentTimeDeviceMembershipTarget
// returns a ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) PostAsClearEnrollmentTimeDeviceMembershipTargetPostResponse(ctx context.Context, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration)(ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetPostResponseable), nil
}
// ToPostRequestInformation invoke action clearEnrollmentTimeDeviceMembershipTarget
// returns a *RequestInformation when successful
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder when successful
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) WithUrl(rawUrl string)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemClearEnrollmentTimeDeviceMembershipTargetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

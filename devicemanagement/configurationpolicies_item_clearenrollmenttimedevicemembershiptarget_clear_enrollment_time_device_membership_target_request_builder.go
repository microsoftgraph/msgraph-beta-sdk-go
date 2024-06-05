package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder provides operations to call the clearEnrollmentTimeDeviceMembershipTarget method.
type ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal instantiates a new ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder and sets the default values.
func NewConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    m := &ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy%2Did}/clearEnrollmentTimeDeviceMembershipTarget", pathParameters),
    }
    return m
}
// NewConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder instantiates a new ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder and sets the default values.
func NewConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action clearEnrollmentTimeDeviceMembershipTarget
// Deprecated: This method is obsolete. Use PostAsClearEnrollmentTimeDeviceMembershipTargetPostResponse instead.
// returns a ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) Post(ctx context.Context, requestConfiguration *ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration)(ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetResponseable), nil
}
// PostAsClearEnrollmentTimeDeviceMembershipTargetPostResponse invoke action clearEnrollmentTimeDeviceMembershipTarget
// returns a ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) PostAsClearEnrollmentTimeDeviceMembershipTargetPostResponse(ctx context.Context, requestConfiguration *ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration)(ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetPostResponseable), nil
}
// ToPostRequestInformation invoke action clearEnrollmentTimeDeviceMembershipTarget
// returns a *RequestInformation when successful
func (m *ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder when successful
func (m *ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) WithUrl(rawUrl string)(*ConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder) {
    return NewConfigurationpoliciesItemClearenrollmenttimedevicemembershiptargetClearEnrollmentTimeDeviceMembershipTargetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

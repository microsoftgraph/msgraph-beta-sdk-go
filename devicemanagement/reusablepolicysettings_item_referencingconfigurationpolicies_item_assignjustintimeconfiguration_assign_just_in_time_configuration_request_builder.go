package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder provides operations to call the assignJustInTimeConfiguration method.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilderInternal instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder) {
    m := &ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}/assignJustInTimeConfiguration", pathParameters),
    }
    return m
}
// NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder instantiates a new ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder and sets the default values.
func NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assignJustInTimeConfiguration
// Deprecated: This method is obsolete. Use PostAsAssignJustInTimeConfigurationPostResponse instead.
// returns a ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder) Post(ctx context.Context, body ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationPostRequestBodyable, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration)(ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationResponseable), nil
}
// PostAsAssignJustInTimeConfigurationPostResponse invoke action assignJustInTimeConfiguration
// returns a ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder) PostAsAssignJustInTimeConfigurationPostResponse(ctx context.Context, body ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationPostRequestBodyable, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration)(ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationPostResponseable), nil
}
// ToPostRequestInformation invoke action assignJustInTimeConfiguration
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationPostRequestBodyable, requestConfiguration *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder when successful
func (m *ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder) WithUrl(rawUrl string)(*ReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesItemAssignjustintimeconfigurationAssignJustInTimeConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

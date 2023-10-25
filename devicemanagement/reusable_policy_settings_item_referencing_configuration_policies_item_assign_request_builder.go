package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder provides operations to call the assign method.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderInternal instantiates a new AssignRequestBuilder and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder) {
    m := &ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}/assign", pathParameters),
    }
    return m
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder instantiates a new AssignRequestBuilder and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assign
// Deprecated: This method is obsolete. Use PostAsAssignPostResponse instead.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder) Post(ctx context.Context, body ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBodyable, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderPostRequestConfiguration)(ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignResponseable), nil
}
// PostAsAssignPostResponse invoke action assign
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder) PostAsAssignPostResponse(ctx context.Context, body ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBodyable, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderPostRequestConfiguration)(ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostResponseable), nil
}
// ToPostRequestInformation invoke action assign
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBodyable, requestConfiguration *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder) WithUrl(rawUrl string)(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

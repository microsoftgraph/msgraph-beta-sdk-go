package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder provides operations to call the assignJustInTimeConfiguration method.
type ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilderInternal instantiates a new ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder and sets the default values.
func NewConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder) {
    m := &ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy%2Did}/assignJustInTimeConfiguration", pathParameters),
    }
    return m
}
// NewConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder instantiates a new ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder and sets the default values.
func NewConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assignJustInTimeConfiguration
// Deprecated: This method is obsolete. Use PostAsAssignJustInTimeConfigurationPostResponse instead.
// returns a ConfigurationPoliciesItemAssignJustInTimeConfigurationResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder) Post(ctx context.Context, body ConfigurationPoliciesItemAssignJustInTimeConfigurationPostRequestBodyable, requestConfiguration *ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration)(ConfigurationPoliciesItemAssignJustInTimeConfigurationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConfigurationPoliciesItemAssignJustInTimeConfigurationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConfigurationPoliciesItemAssignJustInTimeConfigurationResponseable), nil
}
// PostAsAssignJustInTimeConfigurationPostResponse invoke action assignJustInTimeConfiguration
// returns a ConfigurationPoliciesItemAssignJustInTimeConfigurationPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder) PostAsAssignJustInTimeConfigurationPostResponse(ctx context.Context, body ConfigurationPoliciesItemAssignJustInTimeConfigurationPostRequestBodyable, requestConfiguration *ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration)(ConfigurationPoliciesItemAssignJustInTimeConfigurationPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConfigurationPoliciesItemAssignJustInTimeConfigurationPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConfigurationPoliciesItemAssignJustInTimeConfigurationPostResponseable), nil
}
// ToPostRequestInformation invoke action assignJustInTimeConfiguration
// returns a *RequestInformation when successful
func (m *ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ConfigurationPoliciesItemAssignJustInTimeConfigurationPostRequestBodyable, requestConfiguration *ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder when successful
func (m *ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder) WithUrl(rawUrl string)(*ConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder) {
    return NewConfigurationPoliciesItemAssignJustInTimeConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

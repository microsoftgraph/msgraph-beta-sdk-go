package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder provides operations to call the createCopy method.
type ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilderInternal instantiates a new ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder and sets the default values.
func NewConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder) {
    m := &ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy%2Did}/createCopy", pathParameters),
    }
    return m
}
// NewConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder instantiates a new ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder and sets the default values.
func NewConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createCopy
// returns a DeviceManagementConfigurationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder) Post(ctx context.Context, body ConfigurationpoliciesItemCreatecopyCreateCopyPostRequestBodyable, requestConfiguration *ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action createCopy
// returns a *RequestInformation when successful
func (m *ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ConfigurationpoliciesItemCreatecopyCreateCopyPostRequestBodyable, requestConfiguration *ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder when successful
func (m *ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder) WithUrl(rawUrl string)(*ConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder) {
    return NewConfigurationpoliciesItemCreatecopyCreateCopyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

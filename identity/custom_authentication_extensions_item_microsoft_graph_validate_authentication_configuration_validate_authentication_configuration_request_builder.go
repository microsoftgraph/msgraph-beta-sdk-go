package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilder provides operations to call the validateAuthenticationConfiguration method.
type CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilderInternal instantiates a new ValidateAuthenticationConfigurationRequestBuilder and sets the default values.
func NewCustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilder) {
    m := &CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/customAuthenticationExtensions/{customAuthenticationExtension%2Did}/microsoft.graph.validateAuthenticationConfiguration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilder instantiates a new ValidateAuthenticationConfigurationRequestBuilder and sets the default values.
func NewCustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action validateAuthenticationConfiguration
func (m *CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilder) Post(ctx context.Context, requestConfiguration *CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConfigurationValidationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationConfigurationValidationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationConfigurationValidationable), nil
}
// ToPostRequestInformation invoke action validateAuthenticationConfiguration
func (m *CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CustomAuthenticationExtensionsItemMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

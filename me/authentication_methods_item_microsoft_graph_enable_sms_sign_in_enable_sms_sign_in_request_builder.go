package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilder provides operations to call the enableSmsSignIn method.
type AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilderInternal instantiates a new EnableSmsSignInRequestBuilder and sets the default values.
func NewAuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilder) {
    m := &AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/methods/{authenticationMethod%2Did}/microsoft.graph.enableSmsSignIn";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilder instantiates a new EnableSmsSignInRequestBuilder and sets the default values.
func NewAuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action enableSmsSignIn
func (m *AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilder) Post(ctx context.Context, requestConfiguration *AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action enableSmsSignIn
func (m *AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodsItemMicrosoftGraphEnableSmsSignInEnableSmsSignInRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

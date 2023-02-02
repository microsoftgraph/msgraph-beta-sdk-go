package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilder provides operations to call the disableSmsSignIn method.
type ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilderInternal instantiates a new DisableSmsSignInRequestBuilder and sets the default values.
func NewItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilder) {
    m := &ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/methods/{authenticationMethod%2Did}/microsoft.graph.disableSmsSignIn";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilder instantiates a new DisableSmsSignInRequestBuilder and sets the default values.
func NewItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action disableSmsSignIn
func (m *ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action disableSmsSignIn
func (m *ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationMethodsItemMicrosoftGraphDisableSmsSignInDisableSmsSignInRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

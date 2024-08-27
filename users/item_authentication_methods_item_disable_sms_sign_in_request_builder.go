package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder provides operations to call the disableSmsSignIn method.
type ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationMethodsItemDisableSmsSignInRequestBuilderInternal instantiates a new ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder and sets the default values.
func NewItemAuthenticationMethodsItemDisableSmsSignInRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder) {
    m := &ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/methods/{authenticationMethod%2Did}/disableSmsSignIn", pathParameters),
    }
    return m
}
// NewItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder instantiates a new ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder and sets the default values.
func NewItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationMethodsItemDisableSmsSignInRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action disableSmsSignIn
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action disableSmsSignIn
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder when successful
func (m *ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder) {
    return NewItemAuthenticationMethodsItemDisableSmsSignInRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

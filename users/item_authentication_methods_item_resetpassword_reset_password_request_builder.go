package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder provides operations to call the resetPassword method.
type ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilderInternal instantiates a new ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder and sets the default values.
func NewItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder) {
    m := &ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/methods/{authenticationMethod%2Did}/resetPassword", pathParameters),
    }
    return m
}
// NewItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder instantiates a new ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder and sets the default values.
func NewItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate a reset for the password associated with a password authentication method object. This can only be done by an administrator with appropriate permissions and can't be performed on a user's own account. This flow writes the new password to Microsoft Entra ID and pushes it to on-premises Active Directory if configured using password writeback. The admin can either provide a new password or have the system generate one. The user is prompted to change their password on their next sign in. This reset is a long-running operation and will return a Location header with a link where the caller can periodically check for the status of the reset operation.
// returns a PasswordResetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationmethod-resetpassword?view=graph-rest-beta
func (m *ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder) Post(ctx context.Context, body ItemAuthenticationMethodsItemResetpasswordResetPasswordPostRequestBodyable, requestConfiguration *ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordResetResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePasswordResetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordResetResponseable), nil
}
// ToPostRequestInformation initiate a reset for the password associated with a password authentication method object. This can only be done by an administrator with appropriate permissions and can't be performed on a user's own account. This flow writes the new password to Microsoft Entra ID and pushes it to on-premises Active Directory if configured using password writeback. The admin can either provide a new password or have the system generate one. The user is prompted to change their password on their next sign in. This reset is a long-running operation and will return a Location header with a link where the caller can periodically check for the status of the reset operation.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAuthenticationMethodsItemResetpasswordResetPasswordPostRequestBodyable, requestConfiguration *ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder when successful
func (m *ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder) {
    return NewItemAuthenticationMethodsItemResetpasswordResetPasswordRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

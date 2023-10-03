package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder provides operations to manage the passwordlessMicrosoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetQueryParameters retrieve a user's single Microsoft Authenticator Passwordless Phone Sign-in method object. This API is supported in the following national cloud deployments.
type ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetQueryParameters
}
// NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal instantiates a new PasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    m := &ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder instantiates a new PasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a user's Microsoft Authenticator Passwordless Phone Sign-in method object. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/passwordlessmicrosoftauthenticatorauthenticationmethod-delete?view=graph-rest-1.0
func (m *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Device provides operations to manage the device property of the microsoft.graph.passwordlessMicrosoftAuthenticatorAuthenticationMethod entity.
func (m *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) Device()(*ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilder) {
    return NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a user's single Microsoft Authenticator Passwordless Phone Sign-in method object. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/passwordlessmicrosoftauthenticatorauthenticationmethod-get?view=graph-rest-1.0
func (m *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordlessMicrosoftAuthenticatorAuthenticationMethodable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePasswordlessMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordlessMicrosoftAuthenticatorAuthenticationMethodable), nil
}
// ToDeleteRequestInformation deletes a user's Microsoft Authenticator Passwordless Phone Sign-in method object. This API is supported in the following national cloud deployments.
func (m *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a user's single Microsoft Authenticator Passwordless Phone Sign-in method object. This API is supported in the following national cloud deployments.
func (m *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    return NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

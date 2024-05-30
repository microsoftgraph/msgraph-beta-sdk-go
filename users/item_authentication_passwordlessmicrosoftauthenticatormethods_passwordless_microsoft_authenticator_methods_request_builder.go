package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder provides operations to manage the passwordlessMicrosoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetQueryParameters represents the Microsoft Authenticator Passwordless Phone Sign-in methods registered to a user for authentication.
type ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetQueryParameters
}
// ByPasswordlessMicrosoftAuthenticatorAuthenticationMethodId provides operations to manage the passwordlessMicrosoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder when successful
func (m *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) ByPasswordlessMicrosoftAuthenticatorAuthenticationMethodId(passwordlessMicrosoftAuthenticatorAuthenticationMethodId string)(*ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if passwordlessMicrosoftAuthenticatorAuthenticationMethodId != "" {
        urlTplParams["passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did"] = passwordlessMicrosoftAuthenticatorAuthenticationMethodId
    }
    return NewItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal instantiates a new ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder and sets the default values.
func NewItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    m := &ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder instantiates a new ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder and sets the default values.
func NewItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsCountRequestBuilder when successful
func (m *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) Count()(*ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsCountRequestBuilder) {
    return NewItemAuthenticationPasswordlessmicrosoftauthenticatormethodsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents the Microsoft Authenticator Passwordless Phone Sign-in methods registered to a user for authentication.
// returns a PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponseable), nil
}
// ToGetRequestInformation represents the Microsoft Authenticator Passwordless Phone Sign-in methods registered to a user for authentication.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder when successful
func (m *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    return NewItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

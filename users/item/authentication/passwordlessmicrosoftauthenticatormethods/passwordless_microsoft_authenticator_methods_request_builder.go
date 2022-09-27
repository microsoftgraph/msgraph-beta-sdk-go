package passwordlessmicrosoftauthenticatormethods

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i64a95688aba3f8daa031a76499d4e9264c8a3125b44b58f8fd2e8b0439cc117f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/count"
)

// PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder provides operations to manage the passwordlessMicrosoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
type PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetQueryParameters retrieve a list of a user's Microsoft Authenticator Passwordless Phone Sign-in method objects and their properties.
type PasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetQueryParameters struct {
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
// PasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetQueryParameters
}
// NewPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal instantiates a new PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder and sets the default values.
func NewPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    m := &PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder instantiates a new PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder and sets the default values.
func NewPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) Count()(*i64a95688aba3f8daa031a76499d4e9264c8a3125b44b58f8fd2e8b0439cc117f.CountRequestBuilder) {
    return i64a95688aba3f8daa031a76499d4e9264c8a3125b44b58f8fd2e8b0439cc117f.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation retrieve a list of a user's Microsoft Authenticator Passwordless Phone Sign-in method objects and their properties.
func (m *PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieve a list of a user's Microsoft Authenticator Passwordless Phone Sign-in method objects and their properties.
func (m *PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get retrieve a list of a user's Microsoft Authenticator Passwordless Phone Sign-in method objects and their properties.
func (m *PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) Get(ctx context.Context, requestConfiguration *PasswordlessMicrosoftAuthenticatorMethodsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordlessMicrosoftAuthenticatorAuthenticationMethodCollectionResponseable), nil
}

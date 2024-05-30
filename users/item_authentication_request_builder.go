package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationRequestBuilder provides operations to manage the authentication property of the microsoft.graph.user entity.
type ItemAuthenticationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationRequestBuilderGetQueryParameters the authentication methods that are supported for the user.
type ItemAuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationRequestBuilderGetQueryParameters
}
// ItemAuthenticationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationRequestBuilderInternal instantiates a new ItemAuthenticationRequestBuilder and sets the default values.
func NewItemAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationRequestBuilder) {
    m := &ItemAuthenticationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAuthenticationRequestBuilder instantiates a new ItemAuthenticationRequestBuilder and sets the default values.
func NewItemAuthenticationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property authentication for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// EmailMethods provides operations to manage the emailMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationEmailmethodsEmailMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) EmailMethods()(*ItemAuthenticationEmailmethodsEmailMethodsRequestBuilder) {
    return NewItemAuthenticationEmailmethodsEmailMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Fido2Methods provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationFido2methodsFido2MethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) Fido2Methods()(*ItemAuthenticationFido2methodsFido2MethodsRequestBuilder) {
    return NewItemAuthenticationFido2methodsFido2MethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the authentication methods that are supported for the user.
// returns a Authenticationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable), nil
}
// Methods provides operations to manage the methods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) Methods()(*ItemAuthenticationMethodsRequestBuilder) {
    return NewItemAuthenticationMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftAuthenticatorMethods provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorMethodsRequestBuilder) {
    return NewItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationOperationsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) Operations()(*ItemAuthenticationOperationsRequestBuilder) {
    return NewItemAuthenticationOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PasswordlessMicrosoftAuthenticatorMethods provides operations to manage the passwordlessMicrosoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethods()(*ItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    return NewItemAuthenticationPasswordlessmicrosoftauthenticatormethodsPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PasswordMethods provides operations to manage the passwordMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationPasswordmethodsPasswordMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) PasswordMethods()(*ItemAuthenticationPasswordmethodsPasswordMethodsRequestBuilder) {
    return NewItemAuthenticationPasswordmethodsPasswordMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property authentication in users
// returns a Authenticationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, requestConfiguration *ItemAuthenticationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable), nil
}
// PhoneMethods provides operations to manage the phoneMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationPhonemethodsPhoneMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) PhoneMethods()(*ItemAuthenticationPhonemethodsPhoneMethodsRequestBuilder) {
    return NewItemAuthenticationPhonemethodsPhoneMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PlatformCredentialMethods provides operations to manage the platformCredentialMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationPlatformcredentialmethodsPlatformCredentialMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) PlatformCredentialMethods()(*ItemAuthenticationPlatformcredentialmethodsPlatformCredentialMethodsRequestBuilder) {
    return NewItemAuthenticationPlatformcredentialmethodsPlatformCredentialMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SignInPreferences the signInPreferences property
// returns a *ItemAuthenticationSigninpreferencesSignInPreferencesRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) SignInPreferences()(*ItemAuthenticationSigninpreferencesSignInPreferencesRequestBuilder) {
    return NewItemAuthenticationSigninpreferencesSignInPreferencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareOathMethods provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationSoftwareoathmethodsSoftwareOathMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) SoftwareOathMethods()(*ItemAuthenticationSoftwareoathmethodsSoftwareOathMethodsRequestBuilder) {
    return NewItemAuthenticationSoftwareoathmethodsSoftwareOathMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TemporaryAccessPassMethods provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) TemporaryAccessPassMethods()(*ItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassMethodsRequestBuilder) {
    return NewItemAuthenticationTemporaryaccesspassmethodsTemporaryAccessPassMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property authentication for users
// returns a *RequestInformation when successful
func (m *ItemAuthenticationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the authentication methods that are supported for the user.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property authentication in users
// returns a *RequestInformation when successful
func (m *ItemAuthenticationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, requestConfiguration *ItemAuthenticationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WindowsHelloForBusinessMethods provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationWindowshelloforbusinessmethodsWindowsHelloForBusinessMethodsRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*ItemAuthenticationWindowshelloforbusinessmethodsWindowsHelloForBusinessMethodsRequestBuilder) {
    return NewItemAuthenticationWindowshelloforbusinessmethodsWindowsHelloForBusinessMethodsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuthenticationRequestBuilder when successful
func (m *ItemAuthenticationRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationRequestBuilder) {
    return NewItemAuthenticationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

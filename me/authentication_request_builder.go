package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationRequestBuilder provides operations to manage the authentication property of the microsoft.graph.user entity.
type AuthenticationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationRequestBuilderGetQueryParameters get authentication from me
type AuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationRequestBuilderGetQueryParameters
}
// AuthenticationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationRequestBuilderInternal instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationRequestBuilder) {
    m := &AuthenticationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationRequestBuilder instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewAuthenticationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get authentication from me
func (m *AuthenticationRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// CreatePatchRequestInformation update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, requestConfiguration *AuthenticationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EmailMethods provides operations to manage the emailMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) EmailMethods()(*AuthenticationEmailMethodsRequestBuilder) {
    return NewAuthenticationEmailMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailMethodsById provides operations to manage the emailMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) EmailMethodsById(id string)(*EmailAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailAuthenticationMethod%2Did"] = id
    }
    return NewEmailAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Fido2Methods provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) Fido2Methods()(*AuthenticationFido2MethodsRequestBuilder) {
    return NewAuthenticationFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fido2MethodsById provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) Fido2MethodsById(id string)(*Fido2AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod%2Did"] = id
    }
    return NewFido2AuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get authentication from me
func (m *AuthenticationRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable), nil
}
// Methods provides operations to manage the methods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) Methods()(*AuthenticationMethodsRequestBuilder) {
    return NewAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MethodsById provides operations to manage the methods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) MethodsById(id string)(*AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod%2Did"] = id
    }
    return NewAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftAuthenticatorMethods provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*AuthenticationMicrosoftAuthenticatorMethodsRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftAuthenticatorMethodsById provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*MicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return NewMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) Operations()(*AuthenticationOperationsRequestBuilder) {
    return NewAuthenticationOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) OperationsById(id string)(*LongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation%2Did"] = id
    }
    return NewLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PasswordlessMicrosoftAuthenticatorMethods provides operations to manage the passwordlessMicrosoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethods()(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordlessMicrosoftAuthenticatorMethodsById provides operations to manage the passwordlessMicrosoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethodsById(id string)(*PasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return NewPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PasswordMethods provides operations to manage the passwordMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) PasswordMethods()(*AuthenticationPasswordMethodsRequestBuilder) {
    return NewAuthenticationPasswordMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordMethodsById provides operations to manage the passwordMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) PasswordMethodsById(id string)(*PasswordAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordAuthenticationMethod%2Did"] = id
    }
    return NewPasswordAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, requestConfiguration *AuthenticationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable), nil
}
// PhoneMethods provides operations to manage the phoneMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) PhoneMethods()(*AuthenticationPhoneMethodsRequestBuilder) {
    return NewAuthenticationPhoneMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhoneMethodsById provides operations to manage the phoneMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) PhoneMethodsById(id string)(*PhoneAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["phoneAuthenticationMethod%2Did"] = id
    }
    return NewPhoneAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SoftwareOathMethods provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) SoftwareOathMethods()(*AuthenticationSoftwareOathMethodsRequestBuilder) {
    return NewAuthenticationSoftwareOathMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftwareOathMethodsById provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) SoftwareOathMethodsById(id string)(*SoftwareOathAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["softwareOathAuthenticationMethod%2Did"] = id
    }
    return NewSoftwareOathAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TemporaryAccessPassMethods provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethods()(*AuthenticationTemporaryAccessPassMethodsRequestBuilder) {
    return NewAuthenticationTemporaryAccessPassMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemporaryAccessPassMethodsById provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethodsById(id string)(*TemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod%2Did"] = id
    }
    return NewTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsHelloForBusinessMethods provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*AuthenticationWindowsHelloForBusinessMethodsRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsHelloForBusinessMethodsById provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod%2Did"] = id
    }
    return NewWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package authentication

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i28848822f3980fdce6726498e9fc09d751b8a59cf5da6b8c082377a9dd61fa94 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/operations"
    i2a0e0508cb5dac8c4a5ff74321688dfcdaa826e9faa19bdc6a70268c36ac402f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordmethods"
    i3c82d3c069d877192717ae253184fedace9ae2e72f03648dc1e909f4a3085727 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/emailmethods"
    i3de92ba094c6412eaa07dac178f38f33e50162805f2531cf1357ebe544a96d18 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods"
    i553aa88228c73bfaaeeef34824d8c3229b6b7db8227c301e41c771b57001dd1f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods"
    i81c2901312c57f1ce22103393de382e2806a19bfa6c163ef15b5a2c5aaefc745 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/softwareoathmethods"
    ia55b6dd6613f097e6d74a8806a220afcb427d9a8f6bffc469027552b9486b552 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/temporaryaccesspassmethods"
    ib8d8d8244d3eb013f4bc094ce5af2285946de30a3abcfc827de92ab9d365368b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/fido2methods"
    ibbe313e164cc65dcee7e84b36906cd6535ea52eb5ee5e2be453de2cd747ca472 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/methods"
    ida0d057327dcf5ae8936c3ca5f5443bcab41b07f01cc7dc59b58c500be7c9f87 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods"
    if78b7f6c376cc85f14bbf69cb183a8aaa85ded820418b3fdcd6a32b685c875c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/phonemethods"
    i0199bdc7d0e31ba40e9079a671271952b6c664dd52228e453e3aa4029592e679 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordmethods/item"
    i11e9821a8e25762609e9e981c9f33988dac6988220959b859b3acf5016e24454 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/softwareoathmethods/item"
    i2fa2bed3b2bf85fb6d420cebcc1e3022721013a245e0c8e9488904d31b8c82cf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/phonemethods/item"
    i467ded3a4537cbf8edac7f6b51c34bc6260059e41501cb558b59f959687f3ead "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/fido2methods/item"
    i5959cb956fb6583f70d31a45fd1df7ec1f6063d4908ec79970275f6476e0c9c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/temporaryaccesspassmethods/item"
    i8b94c7ec519f49a4ce821818ba1007cd615cbf88b35e60111f6243e2381a0a30 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item"
    i8e4613fe9beae47166d3199e35bebd1eacc7b9fb65d987f75b5d6cf3fa8be437 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/emailmethods/item"
    ia0ffdc8f006fc6f17388d1138910ea760bda5327cb7954a36443dafdd3d6fbfa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/methods/item"
    ib16a6b42b7b0fb91aacf9e83541f6f5c6a5266c8492af0c9840118f0f8230fda "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/operations/item"
    ib3cdab32f2c2a8097f7c565b3cb7cb087b04b57f65b7f9c169711dc829cef01c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item"
    ief9d84a1c52ab385faa5763ce3c89c06416ddd2ebf7f8093933c9cb404554238 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item"
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
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationRequestBuilderGetQueryParameters the authentication methods that are supported for the user.
type AuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationRequestBuilderGetQueryParameters
}
// AuthenticationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationRequestBuilderInternal instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationRequestBuilder) {
    m := &AuthenticationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property authentication for users
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property authentication for users
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AuthenticationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the authentication methods that are supported for the user.
func (m *AuthenticationRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the authentication methods that are supported for the user.
func (m *AuthenticationRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AuthenticationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property authentication in users
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property authentication in users
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, requestConfiguration *AuthenticationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property authentication for users
func (m *AuthenticationRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property authentication for users
func (m *AuthenticationRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *AuthenticationRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EmailMethods the emailMethods property
func (m *AuthenticationRequestBuilder) EmailMethods()(*i3c82d3c069d877192717ae253184fedace9ae2e72f03648dc1e909f4a3085727.EmailMethodsRequestBuilder) {
    return i3c82d3c069d877192717ae253184fedace9ae2e72f03648dc1e909f4a3085727.NewEmailMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.emailMethods.item collection
func (m *AuthenticationRequestBuilder) EmailMethodsById(id string)(*i8e4613fe9beae47166d3199e35bebd1eacc7b9fb65d987f75b5d6cf3fa8be437.EmailAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailAuthenticationMethod%2Did"] = id
    }
    return i8e4613fe9beae47166d3199e35bebd1eacc7b9fb65d987f75b5d6cf3fa8be437.NewEmailAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Fido2Methods the fido2Methods property
func (m *AuthenticationRequestBuilder) Fido2Methods()(*ib8d8d8244d3eb013f4bc094ce5af2285946de30a3abcfc827de92ab9d365368b.Fido2MethodsRequestBuilder) {
    return ib8d8d8244d3eb013f4bc094ce5af2285946de30a3abcfc827de92ab9d365368b.NewFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fido2MethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.fido2Methods.item collection
func (m *AuthenticationRequestBuilder) Fido2MethodsById(id string)(*i467ded3a4537cbf8edac7f6b51c34bc6260059e41501cb558b59f959687f3ead.Fido2AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod%2Did"] = id
    }
    return i467ded3a4537cbf8edac7f6b51c34bc6260059e41501cb558b59f959687f3ead.NewFido2AuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the authentication methods that are supported for the user.
func (m *AuthenticationRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the authentication methods that are supported for the user.
func (m *AuthenticationRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AuthenticationRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable), nil
}
// Methods the methods property
func (m *AuthenticationRequestBuilder) Methods()(*ibbe313e164cc65dcee7e84b36906cd6535ea52eb5ee5e2be453de2cd747ca472.MethodsRequestBuilder) {
    return ibbe313e164cc65dcee7e84b36906cd6535ea52eb5ee5e2be453de2cd747ca472.NewMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.methods.item collection
func (m *AuthenticationRequestBuilder) MethodsById(id string)(*ia0ffdc8f006fc6f17388d1138910ea760bda5327cb7954a36443dafdd3d6fbfa.AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod%2Did"] = id
    }
    return ia0ffdc8f006fc6f17388d1138910ea760bda5327cb7954a36443dafdd3d6fbfa.NewAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftAuthenticatorMethods the microsoftAuthenticatorMethods property
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*ida0d057327dcf5ae8936c3ca5f5443bcab41b07f01cc7dc59b58c500be7c9f87.MicrosoftAuthenticatorMethodsRequestBuilder) {
    return ida0d057327dcf5ae8936c3ca5f5443bcab41b07f01cc7dc59b58c500be7c9f87.NewMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftAuthenticatorMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.microsoftAuthenticatorMethods.item collection
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*ib3cdab32f2c2a8097f7c565b3cb7cb087b04b57f65b7f9c169711dc829cef01c.MicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return ib3cdab32f2c2a8097f7c565b3cb7cb087b04b57f65b7f9c169711dc829cef01c.NewMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *AuthenticationRequestBuilder) Operations()(*i28848822f3980fdce6726498e9fc09d751b8a59cf5da6b8c082377a9dd61fa94.OperationsRequestBuilder) {
    return i28848822f3980fdce6726498e9fc09d751b8a59cf5da6b8c082377a9dd61fa94.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.operations.item collection
func (m *AuthenticationRequestBuilder) OperationsById(id string)(*ib16a6b42b7b0fb91aacf9e83541f6f5c6a5266c8492af0c9840118f0f8230fda.LongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation%2Did"] = id
    }
    return ib16a6b42b7b0fb91aacf9e83541f6f5c6a5266c8492af0c9840118f0f8230fda.NewLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PasswordlessMicrosoftAuthenticatorMethods the passwordlessMicrosoftAuthenticatorMethods property
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethods()(*i553aa88228c73bfaaeeef34824d8c3229b6b7db8227c301e41c771b57001dd1f.PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    return i553aa88228c73bfaaeeef34824d8c3229b6b7db8227c301e41c771b57001dd1f.NewPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordlessMicrosoftAuthenticatorMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordlessMicrosoftAuthenticatorMethods.item collection
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethodsById(id string)(*i8b94c7ec519f49a4ce821818ba1007cd615cbf88b35e60111f6243e2381a0a30.PasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return i8b94c7ec519f49a4ce821818ba1007cd615cbf88b35e60111f6243e2381a0a30.NewPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PasswordMethods the passwordMethods property
func (m *AuthenticationRequestBuilder) PasswordMethods()(*i2a0e0508cb5dac8c4a5ff74321688dfcdaa826e9faa19bdc6a70268c36ac402f.PasswordMethodsRequestBuilder) {
    return i2a0e0508cb5dac8c4a5ff74321688dfcdaa826e9faa19bdc6a70268c36ac402f.NewPasswordMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordMethods.item collection
func (m *AuthenticationRequestBuilder) PasswordMethodsById(id string)(*i0199bdc7d0e31ba40e9079a671271952b6c664dd52228e453e3aa4029592e679.PasswordAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordAuthenticationMethod%2Did"] = id
    }
    return i0199bdc7d0e31ba40e9079a671271952b6c664dd52228e453e3aa4029592e679.NewPasswordAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property authentication in users
func (m *AuthenticationRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property authentication in users
func (m *AuthenticationRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, requestConfiguration *AuthenticationRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PhoneMethods the phoneMethods property
func (m *AuthenticationRequestBuilder) PhoneMethods()(*if78b7f6c376cc85f14bbf69cb183a8aaa85ded820418b3fdcd6a32b685c875c6.PhoneMethodsRequestBuilder) {
    return if78b7f6c376cc85f14bbf69cb183a8aaa85ded820418b3fdcd6a32b685c875c6.NewPhoneMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhoneMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.phoneMethods.item collection
func (m *AuthenticationRequestBuilder) PhoneMethodsById(id string)(*i2fa2bed3b2bf85fb6d420cebcc1e3022721013a245e0c8e9488904d31b8c82cf.PhoneAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["phoneAuthenticationMethod%2Did"] = id
    }
    return i2fa2bed3b2bf85fb6d420cebcc1e3022721013a245e0c8e9488904d31b8c82cf.NewPhoneAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SoftwareOathMethods the softwareOathMethods property
func (m *AuthenticationRequestBuilder) SoftwareOathMethods()(*i81c2901312c57f1ce22103393de382e2806a19bfa6c163ef15b5a2c5aaefc745.SoftwareOathMethodsRequestBuilder) {
    return i81c2901312c57f1ce22103393de382e2806a19bfa6c163ef15b5a2c5aaefc745.NewSoftwareOathMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftwareOathMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.softwareOathMethods.item collection
func (m *AuthenticationRequestBuilder) SoftwareOathMethodsById(id string)(*i11e9821a8e25762609e9e981c9f33988dac6988220959b859b3acf5016e24454.SoftwareOathAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["softwareOathAuthenticationMethod%2Did"] = id
    }
    return i11e9821a8e25762609e9e981c9f33988dac6988220959b859b3acf5016e24454.NewSoftwareOathAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TemporaryAccessPassMethods the temporaryAccessPassMethods property
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethods()(*ia55b6dd6613f097e6d74a8806a220afcb427d9a8f6bffc469027552b9486b552.TemporaryAccessPassMethodsRequestBuilder) {
    return ia55b6dd6613f097e6d74a8806a220afcb427d9a8f6bffc469027552b9486b552.NewTemporaryAccessPassMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemporaryAccessPassMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.temporaryAccessPassMethods.item collection
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethodsById(id string)(*i5959cb956fb6583f70d31a45fd1df7ec1f6063d4908ec79970275f6476e0c9c4.TemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod%2Did"] = id
    }
    return i5959cb956fb6583f70d31a45fd1df7ec1f6063d4908ec79970275f6476e0c9c4.NewTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsHelloForBusinessMethods the windowsHelloForBusinessMethods property
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*i3de92ba094c6412eaa07dac178f38f33e50162805f2531cf1357ebe544a96d18.WindowsHelloForBusinessMethodsRequestBuilder) {
    return i3de92ba094c6412eaa07dac178f38f33e50162805f2531cf1357ebe544a96d18.NewWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsHelloForBusinessMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item collection
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*ief9d84a1c52ab385faa5763ce3c89c06416ddd2ebf7f8093933c9cb404554238.WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod%2Did"] = id
    }
    return ief9d84a1c52ab385faa5763ce3c89c06416ddd2ebf7f8093933c9cb404554238.NewWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

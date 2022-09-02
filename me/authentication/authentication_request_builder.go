package authentication

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i031886e282a72eeb426bbd544e45b4457ce2e270fc8c1a19fa2b007b18eeae51 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods"
    i09a669c0839eb56b1267ec6c47053774b28212601399fa8e394fe3659d37cbe8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/operations"
    i7181776de67abe7bec5b6b136451c674def2604be35b184f0aa413473fcc5701 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods"
    i9d172c880291778c29968c982a1fde24417fe56d76705022dd57555d79c61a0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/fido2methods"
    i9fafed747e4fd8dc14a1cd08e98cf200425cadc1d857dc77905604b4d814cebc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods"
    ic0756de5303fb580740539d17e229ce1a3ba8b519f526de32096715f3410b0dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/temporaryaccesspassmethods"
    ic150efac30f986d5b426bfc0c6529ba6d28832fce2cd65152e990426885bceda "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods"
    ic1f96216136db886ff1afedeb281a1126c04793184c3ea4371f0f38018c9112d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/phonemethods"
    id91dbdbdfcc1f5a69fdff3c77ff512c8b4d1c46f0a74fb12686474b14fc391bf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordmethods"
    iefd4f63a7e855b21e1b4c41ac37f80e63563cce5d2ef96a65386065e5ddc1e68 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/emailmethods"
    ieff400c498c615c0e3401955f75e873207fd99c040e4912f844a5ec9aa1ee806 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/softwareoathmethods"
    i0059c2e4e23bd64f21f59deef18c5332cc205989dbcfed5623a379da011bf37c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/fido2methods/item"
    i01fb633a3822873540b4a37a19ffef4b3877b2e3724ff2a9e0ef15556e878183 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item"
    i1d6e1ee1e40c5bb5120f696f065400747d9f8b7c528e9db461c049031f562672 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item"
    i2b7ed9d4e22b5e330477a489ae98a8b1036a0bf39a69812e6c8f86f7902e6637 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/temporaryaccesspassmethods/item"
    i522f450607780460342b3cadbe08d49ca77462560e1b365ab8ce05addc5fb48e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/phonemethods/item"
    i96111e0c5ad54467d200eedd5128b62e6efcead28ff274fda8911bca83e2901f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item"
    id34bbd90b1fa28ce673d316987b63e1a63ca88a556236a01b0fc6f3f0dda42e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/operations/item"
    idb337e103dafe8aa86e240ec88da925669064777080064e8250bd1af4600efc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods/item"
    ide04a99d52e0e569423a7b87e2dd9af319adae8ce3b7091b35326abddb082980 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/emailmethods/item"
    iec467b36e4fdcc4d4a86d09fbaf10732c444a78b7a0400959fa67b6b0ca78e8e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordmethods/item"
    if150209e25231db73554a8080d70c03307ce471bae20e97cc930aebfc021ba64 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/softwareoathmethods/item"
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
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property authentication for me
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
// CreateGetRequestInformation get authentication from me
func (m *AuthenticationRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get authentication from me
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
// CreatePatchRequestInformation update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property authentication in me
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
// Delete delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// EmailMethods the emailMethods property
func (m *AuthenticationRequestBuilder) EmailMethods()(*iefd4f63a7e855b21e1b4c41ac37f80e63563cce5d2ef96a65386065e5ddc1e68.EmailMethodsRequestBuilder) {
    return iefd4f63a7e855b21e1b4c41ac37f80e63563cce5d2ef96a65386065e5ddc1e68.NewEmailMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.emailMethods.item collection
func (m *AuthenticationRequestBuilder) EmailMethodsById(id string)(*ide04a99d52e0e569423a7b87e2dd9af319adae8ce3b7091b35326abddb082980.EmailAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailAuthenticationMethod%2Did"] = id
    }
    return ide04a99d52e0e569423a7b87e2dd9af319adae8ce3b7091b35326abddb082980.NewEmailAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Fido2Methods the fido2Methods property
func (m *AuthenticationRequestBuilder) Fido2Methods()(*i9d172c880291778c29968c982a1fde24417fe56d76705022dd57555d79c61a0e.Fido2MethodsRequestBuilder) {
    return i9d172c880291778c29968c982a1fde24417fe56d76705022dd57555d79c61a0e.NewFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fido2MethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.fido2Methods.item collection
func (m *AuthenticationRequestBuilder) Fido2MethodsById(id string)(*i0059c2e4e23bd64f21f59deef18c5332cc205989dbcfed5623a379da011bf37c.Fido2AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod%2Did"] = id
    }
    return i0059c2e4e23bd64f21f59deef18c5332cc205989dbcfed5623a379da011bf37c.NewFido2AuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get authentication from me
func (m *AuthenticationRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Methods the methods property
func (m *AuthenticationRequestBuilder) Methods()(*i7181776de67abe7bec5b6b136451c674def2604be35b184f0aa413473fcc5701.MethodsRequestBuilder) {
    return i7181776de67abe7bec5b6b136451c674def2604be35b184f0aa413473fcc5701.NewMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.methods.item collection
func (m *AuthenticationRequestBuilder) MethodsById(id string)(*idb337e103dafe8aa86e240ec88da925669064777080064e8250bd1af4600efc7.AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod%2Did"] = id
    }
    return idb337e103dafe8aa86e240ec88da925669064777080064e8250bd1af4600efc7.NewAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftAuthenticatorMethods the microsoftAuthenticatorMethods property
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*i9fafed747e4fd8dc14a1cd08e98cf200425cadc1d857dc77905604b4d814cebc.MicrosoftAuthenticatorMethodsRequestBuilder) {
    return i9fafed747e4fd8dc14a1cd08e98cf200425cadc1d857dc77905604b4d814cebc.NewMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftAuthenticatorMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.microsoftAuthenticatorMethods.item collection
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*i01fb633a3822873540b4a37a19ffef4b3877b2e3724ff2a9e0ef15556e878183.MicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return i01fb633a3822873540b4a37a19ffef4b3877b2e3724ff2a9e0ef15556e878183.NewMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *AuthenticationRequestBuilder) Operations()(*i09a669c0839eb56b1267ec6c47053774b28212601399fa8e394fe3659d37cbe8.OperationsRequestBuilder) {
    return i09a669c0839eb56b1267ec6c47053774b28212601399fa8e394fe3659d37cbe8.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.operations.item collection
func (m *AuthenticationRequestBuilder) OperationsById(id string)(*id34bbd90b1fa28ce673d316987b63e1a63ca88a556236a01b0fc6f3f0dda42e7.LongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation%2Did"] = id
    }
    return id34bbd90b1fa28ce673d316987b63e1a63ca88a556236a01b0fc6f3f0dda42e7.NewLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PasswordlessMicrosoftAuthenticatorMethods the passwordlessMicrosoftAuthenticatorMethods property
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethods()(*i031886e282a72eeb426bbd544e45b4457ce2e270fc8c1a19fa2b007b18eeae51.PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    return i031886e282a72eeb426bbd544e45b4457ce2e270fc8c1a19fa2b007b18eeae51.NewPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordlessMicrosoftAuthenticatorMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.passwordlessMicrosoftAuthenticatorMethods.item collection
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethodsById(id string)(*i1d6e1ee1e40c5bb5120f696f065400747d9f8b7c528e9db461c049031f562672.PasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return i1d6e1ee1e40c5bb5120f696f065400747d9f8b7c528e9db461c049031f562672.NewPasswordlessMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PasswordMethods the passwordMethods property
func (m *AuthenticationRequestBuilder) PasswordMethods()(*id91dbdbdfcc1f5a69fdff3c77ff512c8b4d1c46f0a74fb12686474b14fc391bf.PasswordMethodsRequestBuilder) {
    return id91dbdbdfcc1f5a69fdff3c77ff512c8b4d1c46f0a74fb12686474b14fc391bf.NewPasswordMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.passwordMethods.item collection
func (m *AuthenticationRequestBuilder) PasswordMethodsById(id string)(*iec467b36e4fdcc4d4a86d09fbaf10732c444a78b7a0400959fa67b6b0ca78e8e.PasswordAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordAuthenticationMethod%2Did"] = id
    }
    return iec467b36e4fdcc4d4a86d09fbaf10732c444a78b7a0400959fa67b6b0ca78e8e.NewPasswordAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Authenticationable, requestConfiguration *AuthenticationRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// PhoneMethods the phoneMethods property
func (m *AuthenticationRequestBuilder) PhoneMethods()(*ic1f96216136db886ff1afedeb281a1126c04793184c3ea4371f0f38018c9112d.PhoneMethodsRequestBuilder) {
    return ic1f96216136db886ff1afedeb281a1126c04793184c3ea4371f0f38018c9112d.NewPhoneMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhoneMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.phoneMethods.item collection
func (m *AuthenticationRequestBuilder) PhoneMethodsById(id string)(*i522f450607780460342b3cadbe08d49ca77462560e1b365ab8ce05addc5fb48e.PhoneAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["phoneAuthenticationMethod%2Did"] = id
    }
    return i522f450607780460342b3cadbe08d49ca77462560e1b365ab8ce05addc5fb48e.NewPhoneAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SoftwareOathMethods the softwareOathMethods property
func (m *AuthenticationRequestBuilder) SoftwareOathMethods()(*ieff400c498c615c0e3401955f75e873207fd99c040e4912f844a5ec9aa1ee806.SoftwareOathMethodsRequestBuilder) {
    return ieff400c498c615c0e3401955f75e873207fd99c040e4912f844a5ec9aa1ee806.NewSoftwareOathMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftwareOathMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.softwareOathMethods.item collection
func (m *AuthenticationRequestBuilder) SoftwareOathMethodsById(id string)(*if150209e25231db73554a8080d70c03307ce471bae20e97cc930aebfc021ba64.SoftwareOathAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["softwareOathAuthenticationMethod%2Did"] = id
    }
    return if150209e25231db73554a8080d70c03307ce471bae20e97cc930aebfc021ba64.NewSoftwareOathAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TemporaryAccessPassMethods the temporaryAccessPassMethods property
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethods()(*ic0756de5303fb580740539d17e229ce1a3ba8b519f526de32096715f3410b0dc.TemporaryAccessPassMethodsRequestBuilder) {
    return ic0756de5303fb580740539d17e229ce1a3ba8b519f526de32096715f3410b0dc.NewTemporaryAccessPassMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemporaryAccessPassMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.temporaryAccessPassMethods.item collection
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethodsById(id string)(*i2b7ed9d4e22b5e330477a489ae98a8b1036a0bf39a69812e6c8f86f7902e6637.TemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod%2Did"] = id
    }
    return i2b7ed9d4e22b5e330477a489ae98a8b1036a0bf39a69812e6c8f86f7902e6637.NewTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsHelloForBusinessMethods the windowsHelloForBusinessMethods property
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*ic150efac30f986d5b426bfc0c6529ba6d28832fce2cd65152e990426885bceda.WindowsHelloForBusinessMethodsRequestBuilder) {
    return ic150efac30f986d5b426bfc0c6529ba6d28832fce2cd65152e990426885bceda.NewWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsHelloForBusinessMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.authentication.windowsHelloForBusinessMethods.item collection
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*i96111e0c5ad54467d200eedd5128b62e6efcead28ff274fda8911bca83e2901f.WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod%2Did"] = id
    }
    return i96111e0c5ad54467d200eedd5128b62e6efcead28ff274fda8911bca83e2901f.NewWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

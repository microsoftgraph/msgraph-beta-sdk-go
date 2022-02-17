package authentication

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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

// AuthenticationRequestBuilder builds and executes requests for operations under \users\{user-id}\authentication
type AuthenticationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AuthenticationRequestBuilderDeleteOptions options for Delete
type AuthenticationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuthenticationRequestBuilderGetOptions options for Get
type AuthenticationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AuthenticationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuthenticationRequestBuilderGetQueryParameters get authentication from users
type AuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AuthenticationRequestBuilderPatchOptions options for Patch
type AuthenticationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Authentication;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAuthenticationRequestBuilderInternal instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationRequestBuilder) {
    m := &AuthenticationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/authentication{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationRequestBuilder instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewAuthenticationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property authentication for users
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformation(options *AuthenticationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get authentication from users
func (m *AuthenticationRequestBuilder) CreateGetRequestInformation(options *AuthenticationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property authentication in users
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformation(options *AuthenticationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property authentication for users
func (m *AuthenticationRequestBuilder) Delete(options *AuthenticationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationRequestBuilder) EmailMethods()(*i3c82d3c069d877192717ae253184fedace9ae2e72f03648dc1e909f4a3085727.EmailMethodsRequestBuilder) {
    return i3c82d3c069d877192717ae253184fedace9ae2e72f03648dc1e909f4a3085727.NewEmailMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.emailMethods.item collection
func (m *AuthenticationRequestBuilder) EmailMethodsById(id string)(*i8e4613fe9beae47166d3199e35bebd1eacc7b9fb65d987f75b5d6cf3fa8be437.EmailAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailAuthenticationMethod_id"] = id
    }
    return i8e4613fe9beae47166d3199e35bebd1eacc7b9fb65d987f75b5d6cf3fa8be437.NewEmailAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) Fido2Methods()(*ib8d8d8244d3eb013f4bc094ce5af2285946de30a3abcfc827de92ab9d365368b.Fido2MethodsRequestBuilder) {
    return ib8d8d8244d3eb013f4bc094ce5af2285946de30a3abcfc827de92ab9d365368b.NewFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fido2MethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.fido2Methods.item collection
func (m *AuthenticationRequestBuilder) Fido2MethodsById(id string)(*i467ded3a4537cbf8edac7f6b51c34bc6260059e41501cb558b59f959687f3ead.Fido2AuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod_id"] = id
    }
    return i467ded3a4537cbf8edac7f6b51c34bc6260059e41501cb558b59f959687f3ead.NewFido2AuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get authentication from users
func (m *AuthenticationRequestBuilder) Get(options *AuthenticationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Authentication, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAuthentication() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Authentication), nil
}
func (m *AuthenticationRequestBuilder) Methods()(*ibbe313e164cc65dcee7e84b36906cd6535ea52eb5ee5e2be453de2cd747ca472.MethodsRequestBuilder) {
    return ibbe313e164cc65dcee7e84b36906cd6535ea52eb5ee5e2be453de2cd747ca472.NewMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.methods.item collection
func (m *AuthenticationRequestBuilder) MethodsById(id string)(*ia0ffdc8f006fc6f17388d1138910ea760bda5327cb7954a36443dafdd3d6fbfa.AuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod_id"] = id
    }
    return ia0ffdc8f006fc6f17388d1138910ea760bda5327cb7954a36443dafdd3d6fbfa.NewAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*ida0d057327dcf5ae8936c3ca5f5443bcab41b07f01cc7dc59b58c500be7c9f87.MicrosoftAuthenticatorMethodsRequestBuilder) {
    return ida0d057327dcf5ae8936c3ca5f5443bcab41b07f01cc7dc59b58c500be7c9f87.NewMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftAuthenticatorMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.microsoftAuthenticatorMethods.item collection
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*ib3cdab32f2c2a8097f7c565b3cb7cb087b04b57f65b7f9c169711dc829cef01c.MicrosoftAuthenticatorAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod_id"] = id
    }
    return ib3cdab32f2c2a8097f7c565b3cb7cb087b04b57f65b7f9c169711dc829cef01c.NewMicrosoftAuthenticatorAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) Operations()(*i28848822f3980fdce6726498e9fc09d751b8a59cf5da6b8c082377a9dd61fa94.OperationsRequestBuilder) {
    return i28848822f3980fdce6726498e9fc09d751b8a59cf5da6b8c082377a9dd61fa94.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.operations.item collection
func (m *AuthenticationRequestBuilder) OperationsById(id string)(*ib16a6b42b7b0fb91aacf9e83541f6f5c6a5266c8492af0c9840118f0f8230fda.LongRunningOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation_id"] = id
    }
    return ib16a6b42b7b0fb91aacf9e83541f6f5c6a5266c8492af0c9840118f0f8230fda.NewLongRunningOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethods()(*i553aa88228c73bfaaeeef34824d8c3229b6b7db8227c301e41c771b57001dd1f.PasswordlessMicrosoftAuthenticatorMethodsRequestBuilder) {
    return i553aa88228c73bfaaeeef34824d8c3229b6b7db8227c301e41c771b57001dd1f.NewPasswordlessMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordlessMicrosoftAuthenticatorMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordlessMicrosoftAuthenticatorMethods.item collection
func (m *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethodsById(id string)(*i8b94c7ec519f49a4ce821818ba1007cd615cbf88b35e60111f6243e2381a0a30.PasswordlessMicrosoftAuthenticatorAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordlessMicrosoftAuthenticatorAuthenticationMethod_id"] = id
    }
    return i8b94c7ec519f49a4ce821818ba1007cd615cbf88b35e60111f6243e2381a0a30.NewPasswordlessMicrosoftAuthenticatorAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) PasswordMethods()(*i2a0e0508cb5dac8c4a5ff74321688dfcdaa826e9faa19bdc6a70268c36ac402f.PasswordMethodsRequestBuilder) {
    return i2a0e0508cb5dac8c4a5ff74321688dfcdaa826e9faa19bdc6a70268c36ac402f.NewPasswordMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.passwordMethods.item collection
func (m *AuthenticationRequestBuilder) PasswordMethodsById(id string)(*i0199bdc7d0e31ba40e9079a671271952b6c664dd52228e453e3aa4029592e679.PasswordAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordAuthenticationMethod_id"] = id
    }
    return i0199bdc7d0e31ba40e9079a671271952b6c664dd52228e453e3aa4029592e679.NewPasswordAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property authentication in users
func (m *AuthenticationRequestBuilder) Patch(options *AuthenticationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationRequestBuilder) PhoneMethods()(*if78b7f6c376cc85f14bbf69cb183a8aaa85ded820418b3fdcd6a32b685c875c6.PhoneMethodsRequestBuilder) {
    return if78b7f6c376cc85f14bbf69cb183a8aaa85ded820418b3fdcd6a32b685c875c6.NewPhoneMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhoneMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.phoneMethods.item collection
func (m *AuthenticationRequestBuilder) PhoneMethodsById(id string)(*i2fa2bed3b2bf85fb6d420cebcc1e3022721013a245e0c8e9488904d31b8c82cf.PhoneAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["phoneAuthenticationMethod_id"] = id
    }
    return i2fa2bed3b2bf85fb6d420cebcc1e3022721013a245e0c8e9488904d31b8c82cf.NewPhoneAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) SoftwareOathMethods()(*i81c2901312c57f1ce22103393de382e2806a19bfa6c163ef15b5a2c5aaefc745.SoftwareOathMethodsRequestBuilder) {
    return i81c2901312c57f1ce22103393de382e2806a19bfa6c163ef15b5a2c5aaefc745.NewSoftwareOathMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftwareOathMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.softwareOathMethods.item collection
func (m *AuthenticationRequestBuilder) SoftwareOathMethodsById(id string)(*i11e9821a8e25762609e9e981c9f33988dac6988220959b859b3acf5016e24454.SoftwareOathAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["softwareOathAuthenticationMethod_id"] = id
    }
    return i11e9821a8e25762609e9e981c9f33988dac6988220959b859b3acf5016e24454.NewSoftwareOathAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethods()(*ia55b6dd6613f097e6d74a8806a220afcb427d9a8f6bffc469027552b9486b552.TemporaryAccessPassMethodsRequestBuilder) {
    return ia55b6dd6613f097e6d74a8806a220afcb427d9a8f6bffc469027552b9486b552.NewTemporaryAccessPassMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemporaryAccessPassMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.temporaryAccessPassMethods.item collection
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethodsById(id string)(*i5959cb956fb6583f70d31a45fd1df7ec1f6063d4908ec79970275f6476e0c9c4.TemporaryAccessPassAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod_id"] = id
    }
    return i5959cb956fb6583f70d31a45fd1df7ec1f6063d4908ec79970275f6476e0c9c4.NewTemporaryAccessPassAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*i3de92ba094c6412eaa07dac178f38f33e50162805f2531cf1357ebe544a96d18.WindowsHelloForBusinessMethodsRequestBuilder) {
    return i3de92ba094c6412eaa07dac178f38f33e50162805f2531cf1357ebe544a96d18.NewWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsHelloForBusinessMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item collection
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*ief9d84a1c52ab385faa5763ce3c89c06416ddd2ebf7f8093933c9cb404554238.WindowsHelloForBusinessAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod_id"] = id
    }
    return ief9d84a1c52ab385faa5763ce3c89c06416ddd2ebf7f8093933c9cb404554238.NewWindowsHelloForBusinessAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

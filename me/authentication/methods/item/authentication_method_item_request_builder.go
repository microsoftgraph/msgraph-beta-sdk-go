package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i168734de056992d72cab15e81d68305aaeaeb4ea4a0542fd088e7d19aef1952e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods/item/enablesmssignin"
    ic4978741cb0a6da22ff1661ecb8d06e4df9cd8f4a9a73d4d8821bc9ffd545054 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods/item/disablesmssignin"
    ida6186e59b7d46b210722f00360284059fb02e165f3d3a13c9cbd4943a50a6f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods/item/resetpassword"
)

// AuthenticationMethodItemRequestBuilder builds and executes requests for operations under \me\authentication\methods\{authenticationMethod-id}
type AuthenticationMethodItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AuthenticationMethodItemRequestBuilderDeleteOptions options for Delete
type AuthenticationMethodItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuthenticationMethodItemRequestBuilderGetOptions options for Get
type AuthenticationMethodItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AuthenticationMethodItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuthenticationMethodItemRequestBuilderGetQueryParameters get methods from me
type AuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AuthenticationMethodItemRequestBuilderPatchOptions options for Patch
type AuthenticationMethodItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethod;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAuthenticationMethodItemRequestBuilderInternal instantiates a new AuthenticationMethodItemRequestBuilder and sets the default values.
func NewAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationMethodItemRequestBuilder) {
    m := &AuthenticationMethodItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/methods/{authenticationMethod_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationMethodItemRequestBuilder instantiates a new AuthenticationMethodItemRequestBuilder and sets the default values.
func NewAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property methods for me
func (m *AuthenticationMethodItemRequestBuilder) CreateDeleteRequestInformation(options *AuthenticationMethodItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get methods from me
func (m *AuthenticationMethodItemRequestBuilder) CreateGetRequestInformation(options *AuthenticationMethodItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property methods in me
func (m *AuthenticationMethodItemRequestBuilder) CreatePatchRequestInformation(options *AuthenticationMethodItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property methods for me
func (m *AuthenticationMethodItemRequestBuilder) Delete(options *AuthenticationMethodItemRequestBuilderDeleteOptions)(error) {
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
func (m *AuthenticationMethodItemRequestBuilder) DisableSmsSignIn()(*ic4978741cb0a6da22ff1661ecb8d06e4df9cd8f4a9a73d4d8821bc9ffd545054.DisableSmsSignInRequestBuilder) {
    return ic4978741cb0a6da22ff1661ecb8d06e4df9cd8f4a9a73d4d8821bc9ffd545054.NewDisableSmsSignInRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationMethodItemRequestBuilder) EnableSmsSignIn()(*i168734de056992d72cab15e81d68305aaeaeb4ea4a0542fd088e7d19aef1952e.EnableSmsSignInRequestBuilder) {
    return i168734de056992d72cab15e81d68305aaeaeb4ea4a0542fd088e7d19aef1952e.NewEnableSmsSignInRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get methods from me
func (m *AuthenticationMethodItemRequestBuilder) Get(options *AuthenticationMethodItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethod, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAuthenticationMethod() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethod), nil
}
// Patch update the navigation property methods in me
func (m *AuthenticationMethodItemRequestBuilder) Patch(options *AuthenticationMethodItemRequestBuilderPatchOptions)(error) {
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
func (m *AuthenticationMethodItemRequestBuilder) ResetPassword()(*ida6186e59b7d46b210722f00360284059fb02e165f3d3a13c9cbd4943a50a6f1.ResetPasswordRequestBuilder) {
    return ida6186e59b7d46b210722f00360284059fb02e165f3d3a13c9cbd4943a50a6f1.NewResetPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

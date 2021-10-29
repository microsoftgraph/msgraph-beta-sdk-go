package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i168734de056992d72cab15e81d68305aaeaeb4ea4a0542fd088e7d19aef1952e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods/item/enablesmssignin"
    ic4978741cb0a6da22ff1661ecb8d06e4df9cd8f4a9a73d4d8821bc9ffd545054 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods/item/disablesmssignin"
    ida6186e59b7d46b210722f00360284059fb02e165f3d3a13c9cbd4943a50a6f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/methods/item/resetpassword"
)

// Builds and executes requests for operations under \me\authentication\methods\{authenticationMethod-id}
type AuthenticationMethodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AuthenticationMethodRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AuthenticationMethodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AuthenticationMethodRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get methods from me
type AuthenticationMethodRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AuthenticationMethodRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethod;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new AuthenticationMethodRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAuthenticationMethodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationMethodRequestBuilder) {
    m := &AuthenticationMethodRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/authentication/methods/{authenticationMethod_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AuthenticationMethodRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAuthenticationMethodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationMethodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property methods for me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodRequestBuilder) CreateDeleteRequestInformation(options *AuthenticationMethodRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get methods from me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodRequestBuilder) CreateGetRequestInformation(options *AuthenticationMethodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update the navigation property methods in me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodRequestBuilder) CreatePatchRequestInformation(options *AuthenticationMethodRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete navigation property methods for me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodRequestBuilder) Delete(options *AuthenticationMethodRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationMethodRequestBuilder) DisableSmsSignIn()(*ic4978741cb0a6da22ff1661ecb8d06e4df9cd8f4a9a73d4d8821bc9ffd545054.DisableSmsSignInRequestBuilder) {
    return ic4978741cb0a6da22ff1661ecb8d06e4df9cd8f4a9a73d4d8821bc9ffd545054.NewDisableSmsSignInRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AuthenticationMethodRequestBuilder) EnableSmsSignIn()(*i168734de056992d72cab15e81d68305aaeaeb4ea4a0542fd088e7d19aef1952e.EnableSmsSignInRequestBuilder) {
    return i168734de056992d72cab15e81d68305aaeaeb4ea4a0542fd088e7d19aef1952e.NewEnableSmsSignInRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get methods from me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodRequestBuilder) Get(options *AuthenticationMethodRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethod, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAuthenticationMethod() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethod), nil
}
// Update the navigation property methods in me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodRequestBuilder) Patch(options *AuthenticationMethodRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationMethodRequestBuilder) ResetPassword()(*ida6186e59b7d46b210722f00360284059fb02e165f3d3a13c9cbd4943a50a6f1.ResetPasswordRequestBuilder) {
    return ida6186e59b7d46b210722f00360284059fb02e165f3d3a13c9cbd4943a50a6f1.NewResetPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

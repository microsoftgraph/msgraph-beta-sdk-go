package authenticationmethods

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i61af2017ea563365bd6fa5b54a17b5c9d272c07133dab054c245c0ba802588c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/authenticationmethods/usersregisteredbymethod"
    ib7f16628d84c0be87b68eb41df4b24ca4202644b1a9a828575242e338acc8524 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/authenticationmethods/usersregisteredbymethodwithincludedusertypeswithincludeduserroles"
    iea66ad25029df78159398965a02cd69d8925573317b76bc5330047b2d9b0baf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/authenticationmethods/usersregisteredbyfeaturewithincludedusertypeswithincludeduserroles"
    ieba9935f611e8dffd6c4a595723c566237eca57407b2ce3912312f1d48019de6 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/authenticationmethods/usersregisteredbyfeature"
)

// Builds and executes requests for operations under \reports\authenticationMethods
type AuthenticationMethodsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AuthenticationMethodsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AuthenticationMethodsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AuthenticationMethodsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get authenticationMethods from reports
type AuthenticationMethodsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AuthenticationMethodsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethodsRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new AuthenticationMethodsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAuthenticationMethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationMethodsRequestBuilder) {
    m := &AuthenticationMethodsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/authenticationMethods{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AuthenticationMethodsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAuthenticationMethodsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationMethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property authenticationMethods for reports
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodsRequestBuilder) CreateDeleteRequestInformation(options *AuthenticationMethodsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get authenticationMethods from reports
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodsRequestBuilder) CreateGetRequestInformation(options *AuthenticationMethodsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update the navigation property authenticationMethods in reports
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodsRequestBuilder) CreatePatchRequestInformation(options *AuthenticationMethodsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete navigation property authenticationMethods for reports
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodsRequestBuilder) Delete(options *AuthenticationMethodsRequestBuilderDeleteOptions)(error) {
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
// Get authenticationMethods from reports
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodsRequestBuilder) Get(options *AuthenticationMethodsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethodsRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAuthenticationMethodsRoot() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethodsRoot), nil
}
// Update the navigation property authenticationMethods in reports
// Parameters:
//  - options : Options for the request
func (m *AuthenticationMethodsRequestBuilder) Patch(options *AuthenticationMethodsRequestBuilderPatchOptions)(error) {
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
// Builds and executes requests for operations under \reports\authenticationMethods\microsoft.graph.usersRegisteredByFeature()
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByFeature()(*ieba9935f611e8dffd6c4a595723c566237eca57407b2ce3912312f1d48019de6.UsersRegisteredByFeatureRequestBuilder) {
    return ieba9935f611e8dffd6c4a595723c566237eca57407b2ce3912312f1d48019de6.NewUsersRegisteredByFeatureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \reports\authenticationMethods\microsoft.graph.usersRegisteredByFeature(includedUserTypes={includedUserTypes},includedUserRoles={includedUserRoles})
// Parameters:
//  - includedUserRoles : Usage: includedUserRoles={includedUserRoles}
//  - includedUserTypes : Usage: includedUserTypes={includedUserTypes}
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRoles(includedUserTypes *string, includedUserRoles *string)(*iea66ad25029df78159398965a02cd69d8925573317b76bc5330047b2d9b0baf0.UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    return iea66ad25029df78159398965a02cd69d8925573317b76bc5330047b2d9b0baf0.NewUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter, includedUserTypes, includedUserRoles);
}
// Builds and executes requests for operations under \reports\authenticationMethods\microsoft.graph.usersRegisteredByMethod()
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByMethod()(*i61af2017ea563365bd6fa5b54a17b5c9d272c07133dab054c245c0ba802588c7.UsersRegisteredByMethodRequestBuilder) {
    return i61af2017ea563365bd6fa5b54a17b5c9d272c07133dab054c245c0ba802588c7.NewUsersRegisteredByMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \reports\authenticationMethods\microsoft.graph.usersRegisteredByMethod(includedUserTypes={includedUserTypes},includedUserRoles={includedUserRoles})
// Parameters:
//  - includedUserRoles : Usage: includedUserRoles={includedUserRoles}
//  - includedUserTypes : Usage: includedUserTypes={includedUserTypes}
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRoles(includedUserTypes *string, includedUserRoles *string)(*ib7f16628d84c0be87b68eb41df4b24ca4202644b1a9a828575242e338acc8524.UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    return ib7f16628d84c0be87b68eb41df4b24ca4202644b1a9a828575242e338acc8524.NewUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter, includedUserTypes, includedUserRoles);
}

package authenticationmethods

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0dd7ffa9993603b98e542fb6765c2bc58b69867a61d8a29faf5aa5bd28dc5134 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/usersregisteredbymethod"
    i1bba8e4b27ab9c120aa0ff2d39a209b233e51379ecc5ae53660f9bc4e1dd4439 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/usersregisteredbymethodwithincludedusertypeswithincludeduserroles"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i47f70708cebf48f687242c96efc5507b89cca5b077a72c1d685b5c61a39d46c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/usersregisteredbyfeaturewithincludedusertypeswithincludeduserroles"
    ie54d2504e9f581782cfc818a41a8b995a582cdd5f7976792d36e04e314a0ca31 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/userregistrationdetails"
    ieccbed6a2bbaa40cf04b762541a9ab50dd6ed3158a44574052a51b40a95f8487 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/usersregisteredbyfeature"
    i8465f6f7d0d944a33802781ad4947a28b996037208b7f788a58b7c1946465bf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/userregistrationdetails/item"
)

// AuthenticationMethodsRequestBuilder provides operations to manage the authenticationMethods property of the microsoft.graph.reportRoot entity.
type AuthenticationMethodsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AuthenticationMethodsRequestBuilderDeleteOptions options for Delete
type AuthenticationMethodsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuthenticationMethodsRequestBuilderGetOptions options for Get
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
// AuthenticationMethodsRequestBuilderGetQueryParameters container for navigation properties for Azure AD authentication methods resources.
type AuthenticationMethodsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AuthenticationMethodsRequestBuilderPatchOptions options for Patch
type AuthenticationMethodsRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethodsRootable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAuthenticationMethodsRequestBuilderInternal instantiates a new AuthenticationMethodsRequestBuilder and sets the default values.
func NewAuthenticationMethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationMethodsRequestBuilder) {
    m := &AuthenticationMethodsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports/authenticationMethods{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationMethodsRequestBuilder instantiates a new AuthenticationMethodsRequestBuilder and sets the default values.
func NewAuthenticationMethodsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationMethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property authenticationMethods for print
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
// CreateGetRequestInformation container for navigation properties for Azure AD authentication methods resources.
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
// CreatePatchRequestInformation update the navigation property authenticationMethods in print
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
// Delete delete navigation property authenticationMethods for print
func (m *AuthenticationMethodsRequestBuilder) Delete(options *AuthenticationMethodsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get container for navigation properties for Azure AD authentication methods resources.
func (m *AuthenticationMethodsRequestBuilder) Get(options *AuthenticationMethodsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethodsRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAuthenticationMethodsRootFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthenticationMethodsRootable), nil
}
// Patch update the navigation property authenticationMethods in print
func (m *AuthenticationMethodsRequestBuilder) Patch(options *AuthenticationMethodsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationMethodsRequestBuilder) UserRegistrationDetails()(*ie54d2504e9f581782cfc818a41a8b995a582cdd5f7976792d36e04e314a0ca31.UserRegistrationDetailsRequestBuilder) {
    return ie54d2504e9f581782cfc818a41a8b995a582cdd5f7976792d36e04e314a0ca31.NewUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRegistrationDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.authenticationMethods.userRegistrationDetails.item collection
func (m *AuthenticationMethodsRequestBuilder) UserRegistrationDetailsById(id string)(*i8465f6f7d0d944a33802781ad4947a28b996037208b7f788a58b7c1946465bf0.UserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userRegistrationDetails_id"] = id
    }
    return i8465f6f7d0d944a33802781ad4947a28b996037208b7f788a58b7c1946465bf0.NewUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsersRegisteredByFeature provides operations to call the usersRegisteredByFeature method.
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByFeature()(*ieccbed6a2bbaa40cf04b762541a9ab50dd6ed3158a44574052a51b40a95f8487.UsersRegisteredByFeatureRequestBuilder) {
    return ieccbed6a2bbaa40cf04b762541a9ab50dd6ed3158a44574052a51b40a95f8487.NewUsersRegisteredByFeatureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRoles provides operations to call the usersRegisteredByFeature method.
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRoles(includedUserRoles *string, includedUserTypes *string)(*i47f70708cebf48f687242c96efc5507b89cca5b077a72c1d685b5c61a39d46c0.UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    return i47f70708cebf48f687242c96efc5507b89cca5b077a72c1d685b5c61a39d46c0.NewUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter, includedUserRoles, includedUserTypes);
}
// UsersRegisteredByMethod provides operations to call the usersRegisteredByMethod method.
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByMethod()(*i0dd7ffa9993603b98e542fb6765c2bc58b69867a61d8a29faf5aa5bd28dc5134.UsersRegisteredByMethodRequestBuilder) {
    return i0dd7ffa9993603b98e542fb6765c2bc58b69867a61d8a29faf5aa5bd28dc5134.NewUsersRegisteredByMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRoles provides operations to call the usersRegisteredByMethod method.
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRoles(includedUserRoles *string, includedUserTypes *string)(*i1bba8e4b27ab9c120aa0ff2d39a209b233e51379ecc5ae53660f9bc4e1dd4439.UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    return i1bba8e4b27ab9c120aa0ff2d39a209b233e51379ecc5ae53660f9bc4e1dd4439.NewUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter, includedUserRoles, includedUserTypes);
}

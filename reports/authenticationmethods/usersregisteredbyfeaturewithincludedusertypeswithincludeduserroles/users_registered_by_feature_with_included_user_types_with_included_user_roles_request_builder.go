package usersregisteredbyfeaturewithincludedusertypeswithincludeduserroles

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder builds and executes requests for operations under \reports\authenticationMethods\microsoft.graph.usersRegisteredByFeature(includedUserTypes={includedUserTypes},includedUserRoles={includedUserRoles})
type UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetOptions options for Get
type UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal instantiates a new UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder and sets the default values.
func NewUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, includedUserTypes *string, includedUserRoles *string)(*UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    m := &UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/authenticationMethods/microsoft.graph.usersRegisteredByFeature(includedUserTypes={includedUserTypes},includedUserRoles={includedUserRoles})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if includedUserTypes != nil {
        urlTplParams["includedUserTypes"] = *includedUserTypes
    }
    if includedUserRoles != nil {
        urlTplParams["includedUserRoles"] = *includedUserRoles
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder instantiates a new UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder and sets the default values.
func NewUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function usersRegisteredByFeature
func (m *UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) CreateGetRequestInformation(options *UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function usersRegisteredByFeature
func (m *UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) Get(options *UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserRegistrationFeatureSummary, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserRegistrationFeatureSummary() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserRegistrationFeatureSummary), nil
}

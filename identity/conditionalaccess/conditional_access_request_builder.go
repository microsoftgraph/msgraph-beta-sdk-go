package conditionalaccess

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0b765a6506e9c469fc394944929f3f3f68abc1cd67221843e72d0023cdc7b9a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/authenticationcontextclassreferences"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i74336fcd63e509c2dff28b731f926303088803045222a57ca21181ed72fcf778 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/policies"
    i8b450d93e3a0398b0e364118a8d1d65374bc93acb38b538683e67f6f68a0df16 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/namedlocations"
    i1e1e27fc92484640c1240b9cb5fb9a5b81b2571a142348b22dc8e7cfa597bdea "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/authenticationcontextclassreferences/item"
    i4e66627208bf17ad44b017955b0a38ce27d6db98bb96808a5974cac73445707c "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/policies/item"
    ib6f953303fb5d962a08cf74168e4f0d71eb89d20c00a8164e1326cd252166379 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess/namedlocations/item"
)

// Builds and executes requests for operations under \identity\conditionalAccess
type ConditionalAccessRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ConditionalAccessRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ConditionalAccessRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ConditionalAccessRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// the entry point for the Conditional Access (CA) object model.
type ConditionalAccessRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ConditionalAccessRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConditionalAccessRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ConditionalAccessRequestBuilder) AuthenticationContextClassReferences()(*i0b765a6506e9c469fc394944929f3f3f68abc1cd67221843e72d0023cdc7b9a5.AuthenticationContextClassReferencesRequestBuilder) {
    return i0b765a6506e9c469fc394944929f3f3f68abc1cd67221843e72d0023cdc7b9a5.NewAuthenticationContextClassReferencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.conditionalAccess.authenticationContextClassReferences.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ConditionalAccessRequestBuilder) AuthenticationContextClassReferencesById(id string)(*i1e1e27fc92484640c1240b9cb5fb9a5b81b2571a142348b22dc8e7cfa597bdea.AuthenticationContextClassReferenceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationContextClassReference_id"] = id
    }
    return i1e1e27fc92484640c1240b9cb5fb9a5b81b2571a142348b22dc8e7cfa597bdea.NewAuthenticationContextClassReferenceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ConditionalAccessRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewConditionalAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConditionalAccessRequestBuilder) {
    m := &ConditionalAccessRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ConditionalAccessRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewConditionalAccessRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConditionalAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) CreateDeleteRequestInformation(options *ConditionalAccessRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) CreateGetRequestInformation(options *ConditionalAccessRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) CreatePatchRequestInformation(options *ConditionalAccessRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) Delete(options *ConditionalAccessRequestBuilderDeleteOptions)(error) {
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
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) Get(options *ConditionalAccessRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConditionalAccessRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewConditionalAccessRoot() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConditionalAccessRoot), nil
}
func (m *ConditionalAccessRequestBuilder) NamedLocations()(*i8b450d93e3a0398b0e364118a8d1d65374bc93acb38b538683e67f6f68a0df16.NamedLocationsRequestBuilder) {
    return i8b450d93e3a0398b0e364118a8d1d65374bc93acb38b538683e67f6f68a0df16.NewNamedLocationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.conditionalAccess.namedLocations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ConditionalAccessRequestBuilder) NamedLocationsById(id string)(*ib6f953303fb5d962a08cf74168e4f0d71eb89d20c00a8164e1326cd252166379.NamedLocationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["namedLocation_id"] = id
    }
    return ib6f953303fb5d962a08cf74168e4f0d71eb89d20c00a8164e1326cd252166379.NewNamedLocationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - options : Options for the request
func (m *ConditionalAccessRequestBuilder) Patch(options *ConditionalAccessRequestBuilderPatchOptions)(error) {
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
func (m *ConditionalAccessRequestBuilder) Policies()(*i74336fcd63e509c2dff28b731f926303088803045222a57ca21181ed72fcf778.PoliciesRequestBuilder) {
    return i74336fcd63e509c2dff28b731f926303088803045222a57ca21181ed72fcf778.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.conditionalAccess.policies.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ConditionalAccessRequestBuilder) PoliciesById(id string)(*i4e66627208bf17ad44b017955b0a38ce27d6db98bb96808a5974cac73445707c.ConditionalAccessPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy_id"] = id
    }
    return i4e66627208bf17ad44b017955b0a38ce27d6db98bb96808a5974cac73445707c.NewConditionalAccessPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

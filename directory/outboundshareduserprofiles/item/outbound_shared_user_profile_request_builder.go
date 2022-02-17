package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i49ddd37d26248284ea64a19067478e30c75ede845247b7b5de9cf41e742ed1d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/outboundshareduserprofiles/item/tenants"
    i3a7f0a2cf3e5cbb22ae18d4d09233b740418e3869d44c3de04c038c8644da19a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/outboundshareduserprofiles/item/tenants/item"
)

// OutboundSharedUserProfileRequestBuilder builds and executes requests for operations under \directory\outboundSharedUserProfiles\{outboundSharedUserProfile-userId}
type OutboundSharedUserProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OutboundSharedUserProfileRequestBuilderDeleteOptions options for Delete
type OutboundSharedUserProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutboundSharedUserProfileRequestBuilderGetOptions options for Get
type OutboundSharedUserProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OutboundSharedUserProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutboundSharedUserProfileRequestBuilderGetQueryParameters get outboundSharedUserProfiles from directory
type OutboundSharedUserProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OutboundSharedUserProfileRequestBuilderPatchOptions options for Patch
type OutboundSharedUserProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutboundSharedUserProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOutboundSharedUserProfileRequestBuilderInternal instantiates a new OutboundSharedUserProfileRequestBuilder and sets the default values.
func NewOutboundSharedUserProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutboundSharedUserProfileRequestBuilder) {
    m := &OutboundSharedUserProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/outboundSharedUserProfiles/{outboundSharedUserProfile_userId}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOutboundSharedUserProfileRequestBuilder instantiates a new OutboundSharedUserProfileRequestBuilder and sets the default values.
func NewOutboundSharedUserProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutboundSharedUserProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutboundSharedUserProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property outboundSharedUserProfiles for directory
func (m *OutboundSharedUserProfileRequestBuilder) CreateDeleteRequestInformation(options *OutboundSharedUserProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get outboundSharedUserProfiles from directory
func (m *OutboundSharedUserProfileRequestBuilder) CreateGetRequestInformation(options *OutboundSharedUserProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property outboundSharedUserProfiles in directory
func (m *OutboundSharedUserProfileRequestBuilder) CreatePatchRequestInformation(options *OutboundSharedUserProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property outboundSharedUserProfiles for directory
func (m *OutboundSharedUserProfileRequestBuilder) Delete(options *OutboundSharedUserProfileRequestBuilderDeleteOptions)(error) {
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
// Get get outboundSharedUserProfiles from directory
func (m *OutboundSharedUserProfileRequestBuilder) Get(options *OutboundSharedUserProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutboundSharedUserProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutboundSharedUserProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutboundSharedUserProfile), nil
}
// Patch update the navigation property outboundSharedUserProfiles in directory
func (m *OutboundSharedUserProfileRequestBuilder) Patch(options *OutboundSharedUserProfileRequestBuilderPatchOptions)(error) {
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
func (m *OutboundSharedUserProfileRequestBuilder) Tenants()(*i49ddd37d26248284ea64a19067478e30c75ede845247b7b5de9cf41e742ed1d4.TenantsRequestBuilder) {
    return i49ddd37d26248284ea64a19067478e30c75ede845247b7b5de9cf41e742ed1d4.NewTenantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.outboundSharedUserProfiles.item.tenants.item collection
func (m *OutboundSharedUserProfileRequestBuilder) TenantsById(id string)(*i3a7f0a2cf3e5cbb22ae18d4d09233b740418e3869d44c3de04c038c8644da19a.TenantReferenceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantReference_tenantId"] = id
    }
    return i3a7f0a2cf3e5cbb22ae18d4d09233b740418e3869d44c3de04c038c8644da19a.NewTenantReferenceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

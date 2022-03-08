package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i49ddd37d26248284ea64a19067478e30c75ede845247b7b5de9cf41e742ed1d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/outboundshareduserprofiles/item/tenants"
    i3a7f0a2cf3e5cbb22ae18d4d09233b740418e3869d44c3de04c038c8644da19a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/outboundshareduserprofiles/item/tenants/item"
)

// OutboundSharedUserProfileItemRequestBuilder provides operations to manage the outboundSharedUserProfiles property of the microsoft.graph.directory entity.
type OutboundSharedUserProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OutboundSharedUserProfileItemRequestBuilderDeleteOptions options for Delete
type OutboundSharedUserProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutboundSharedUserProfileItemRequestBuilderGetOptions options for Get
type OutboundSharedUserProfileItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OutboundSharedUserProfileItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutboundSharedUserProfileItemRequestBuilderGetQueryParameters get outboundSharedUserProfiles from directory
type OutboundSharedUserProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OutboundSharedUserProfileItemRequestBuilderPatchOptions options for Patch
type OutboundSharedUserProfileItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutboundSharedUserProfileable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOutboundSharedUserProfileItemRequestBuilderInternal instantiates a new OutboundSharedUserProfileItemRequestBuilder and sets the default values.
func NewOutboundSharedUserProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutboundSharedUserProfileItemRequestBuilder) {
    m := &OutboundSharedUserProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/outboundSharedUserProfiles/{outboundSharedUserProfile_userId}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOutboundSharedUserProfileItemRequestBuilder instantiates a new OutboundSharedUserProfileItemRequestBuilder and sets the default values.
func NewOutboundSharedUserProfileItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutboundSharedUserProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutboundSharedUserProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property outboundSharedUserProfiles for directory
func (m *OutboundSharedUserProfileItemRequestBuilder) CreateDeleteRequestInformation(options *OutboundSharedUserProfileItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OutboundSharedUserProfileItemRequestBuilder) CreateGetRequestInformation(options *OutboundSharedUserProfileItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OutboundSharedUserProfileItemRequestBuilder) CreatePatchRequestInformation(options *OutboundSharedUserProfileItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OutboundSharedUserProfileItemRequestBuilder) Delete(options *OutboundSharedUserProfileItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get outboundSharedUserProfiles from directory
func (m *OutboundSharedUserProfileItemRequestBuilder) Get(options *OutboundSharedUserProfileItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutboundSharedUserProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOutboundSharedUserProfileFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutboundSharedUserProfileable), nil
}
// Patch update the navigation property outboundSharedUserProfiles in directory
func (m *OutboundSharedUserProfileItemRequestBuilder) Patch(options *OutboundSharedUserProfileItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *OutboundSharedUserProfileItemRequestBuilder) Tenants()(*i49ddd37d26248284ea64a19067478e30c75ede845247b7b5de9cf41e742ed1d4.TenantsRequestBuilder) {
    return i49ddd37d26248284ea64a19067478e30c75ede845247b7b5de9cf41e742ed1d4.NewTenantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.outboundSharedUserProfiles.item.tenants.item collection
func (m *OutboundSharedUserProfileItemRequestBuilder) TenantsById(id string)(*i3a7f0a2cf3e5cbb22ae18d4d09233b740418e3869d44c3de04c038c8644da19a.TenantReferenceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantReference_tenantId"] = id
    }
    return i3a7f0a2cf3e5cbb22ae18d4d09233b740418e3869d44c3de04c038c8644da19a.NewTenantReferenceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

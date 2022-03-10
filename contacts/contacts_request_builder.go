package contacts

import (
    i078866701ebd014ef41908c3e576ecccdc54753b23645955da30d52b28cf8773 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/getuserownedobjects"
    i263e3a49923c8b3d5d4677d4bffaef4520d3839dbb9e191c0ead10a25a4595b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/delta"
    i8fa8f6aa0a9b25f3a77255c9f362eca9df147df0488d2dc3b9c1a7a0c9afe5f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/getbyids"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    idc9921cc70b53e0e0915a24e8dbffa0e67b4c8c3233929cbec9702d7a8422ce4 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/count"
    ife4ab247d2377531ea11bb4c8a0ab1e2b8f2156beb597f561bf1f7d9bd436500 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/validateproperties"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// ContactsRequestBuilder provides operations to manage the collection of orgContact entities.
type ContactsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContactsRequestBuilderGetOptions options for Get
type ContactsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContactsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContactsRequestBuilderGetQueryParameters get entities from contacts
type ContactsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// ContactsRequestBuilderPostOptions options for Post
type ContactsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContactable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewContactsRequestBuilderInternal instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactsRequestBuilder) {
    m := &ContactsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContactsRequestBuilder instantiates a new ContactsRequestBuilder and sets the default values.
func NewContactsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContactsRequestBuilder) Count()(*idc9921cc70b53e0e0915a24e8dbffa0e67b4c8c3233929cbec9702d7a8422ce4.CountRequestBuilder) {
    return idc9921cc70b53e0e0915a24e8dbffa0e67b4c8c3233929cbec9702d7a8422ce4.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from contacts
func (m *ContactsRequestBuilder) CreateGetRequestInformation(options *ContactsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to contacts
func (m *ContactsRequestBuilder) CreatePostRequestInformation(options *ContactsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Delta provides operations to call the delta method.
func (m *ContactsRequestBuilder) Delta()(*i263e3a49923c8b3d5d4677d4bffaef4520d3839dbb9e191c0ead10a25a4595b3.DeltaRequestBuilder) {
    return i263e3a49923c8b3d5d4677d4bffaef4520d3839dbb9e191c0ead10a25a4595b3.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from contacts
func (m *ContactsRequestBuilder) Get(options *ContactsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContactCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOrgContactCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContactCollectionResponseable), nil
}
func (m *ContactsRequestBuilder) GetByIds()(*i8fa8f6aa0a9b25f3a77255c9f362eca9df147df0488d2dc3b9c1a7a0c9afe5f4.GetByIdsRequestBuilder) {
    return i8fa8f6aa0a9b25f3a77255c9f362eca9df147df0488d2dc3b9c1a7a0c9afe5f4.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactsRequestBuilder) GetUserOwnedObjects()(*i078866701ebd014ef41908c3e576ecccdc54753b23645955da30d52b28cf8773.GetUserOwnedObjectsRequestBuilder) {
    return i078866701ebd014ef41908c3e576ecccdc54753b23645955da30d52b28cf8773.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to contacts
func (m *ContactsRequestBuilder) Post(options *ContactsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContactable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOrgContactFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContactable), nil
}
func (m *ContactsRequestBuilder) ValidateProperties()(*ife4ab247d2377531ea11bb4c8a0ab1e2b8f2156beb597f561bf1f7d9bd436500.ValidatePropertiesRequestBuilder) {
    return ife4ab247d2377531ea11bb4c8a0ab1e2b8f2156beb597f561bf1f7d9bd436500.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

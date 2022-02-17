package permissiongrants

import (
    i202a750c6a1428068d8dd5672e744e09ad9f8fda997188cbd223f9347890372b "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/getuserownedobjects"
    i4e27de1f5297ce970dc5746ca0bb947c75587bf74148d8b31b6afcb558ceb4c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/getbyids"
    i9fb61394240d7d93c9556b366b36bb2ee252e683760ce0939e65ef81293661ba "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/validateproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// PermissionGrantsRequestBuilder builds and executes requests for operations under \permissionGrants
type PermissionGrantsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PermissionGrantsRequestBuilderGetOptions options for Get
type PermissionGrantsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PermissionGrantsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PermissionGrantsRequestBuilderGetQueryParameters get entities from permissionGrants
type PermissionGrantsRequestBuilderGetQueryParameters struct {
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
}
// PermissionGrantsRequestBuilderPostOptions options for Post
type PermissionGrantsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrant;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewPermissionGrantsRequestBuilderInternal instantiates a new PermissionGrantsRequestBuilder and sets the default values.
func NewPermissionGrantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantsRequestBuilder) {
    m := &PermissionGrantsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/permissionGrants{?search,filter,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPermissionGrantsRequestBuilder instantiates a new PermissionGrantsRequestBuilder and sets the default values.
func NewPermissionGrantsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get entities from permissionGrants
func (m *PermissionGrantsRequestBuilder) CreateGetRequestInformation(options *PermissionGrantsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to permissionGrants
func (m *PermissionGrantsRequestBuilder) CreatePostRequestInformation(options *PermissionGrantsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get entities from permissionGrants
func (m *PermissionGrantsRequestBuilder) Get(options *PermissionGrantsRequestBuilderGetOptions)(*PermissionGrantsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionGrantsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*PermissionGrantsResponse), nil
}
func (m *PermissionGrantsRequestBuilder) GetByIds()(*i4e27de1f5297ce970dc5746ca0bb947c75587bf74148d8b31b6afcb558ceb4c6.GetByIdsRequestBuilder) {
    return i4e27de1f5297ce970dc5746ca0bb947c75587bf74148d8b31b6afcb558ceb4c6.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PermissionGrantsRequestBuilder) GetUserOwnedObjects()(*i202a750c6a1428068d8dd5672e744e09ad9f8fda997188cbd223f9347890372b.GetUserOwnedObjectsRequestBuilder) {
    return i202a750c6a1428068d8dd5672e744e09ad9f8fda997188cbd223f9347890372b.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to permissionGrants
func (m *PermissionGrantsRequestBuilder) Post(options *PermissionGrantsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrant, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewResourceSpecificPermissionGrant() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrant), nil
}
func (m *PermissionGrantsRequestBuilder) ValidateProperties()(*i9fb61394240d7d93c9556b366b36bb2ee252e683760ce0939e65ef81293661ba.ValidatePropertiesRequestBuilder) {
    return i9fb61394240d7d93c9556b366b36bb2ee252e683760ce0939e65ef81293661ba.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

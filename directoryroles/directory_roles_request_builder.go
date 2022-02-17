package directoryroles

import (
    i686e52110d509abde81e2ec02a4b1cacb26be255a28efb139b5fe07005585dcb "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/getuserownedobjects"
    i6c75901f09b564ec245b7b3a868272e273083db1b9879e0b05d7329117fa6e34 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/getbyids"
    i89f3d18e33abe74e12cff2eb112be1d34453a000675f0e7670b07807419b57ff "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/delta"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ifd6f8f935853943530de1cc46edc30d54c90acd20d9483c7525660fa0c70d33f "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/validateproperties"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// DirectoryRolesRequestBuilder builds and executes requests for operations under \directoryRoles
type DirectoryRolesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRolesRequestBuilderGetOptions options for Get
type DirectoryRolesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRolesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRolesRequestBuilderGetQueryParameters get entities from directoryRoles
type DirectoryRolesRequestBuilderGetQueryParameters struct {
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
}
// DirectoryRolesRequestBuilderPostOptions options for Post
type DirectoryRolesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRole;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDirectoryRolesRequestBuilderInternal instantiates a new DirectoryRolesRequestBuilder and sets the default values.
func NewDirectoryRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    m := &DirectoryRolesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles{?skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRolesRequestBuilder instantiates a new DirectoryRolesRequestBuilder and sets the default values.
func NewDirectoryRolesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get entities from directoryRoles
func (m *DirectoryRolesRequestBuilder) CreateGetRequestInformation(options *DirectoryRolesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to directoryRoles
func (m *DirectoryRolesRequestBuilder) CreatePostRequestInformation(options *DirectoryRolesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delta builds and executes requests for operations under \directoryRoles\microsoft.graph.delta()
func (m *DirectoryRolesRequestBuilder) Delta()(*i89f3d18e33abe74e12cff2eb112be1d34453a000675f0e7670b07807419b57ff.DeltaRequestBuilder) {
    return i89f3d18e33abe74e12cff2eb112be1d34453a000675f0e7670b07807419b57ff.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from directoryRoles
func (m *DirectoryRolesRequestBuilder) Get(options *DirectoryRolesRequestBuilderGetOptions)(*DirectoryRolesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryRolesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DirectoryRolesResponse), nil
}
func (m *DirectoryRolesRequestBuilder) GetByIds()(*i6c75901f09b564ec245b7b3a868272e273083db1b9879e0b05d7329117fa6e34.GetByIdsRequestBuilder) {
    return i6c75901f09b564ec245b7b3a868272e273083db1b9879e0b05d7329117fa6e34.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRolesRequestBuilder) GetUserOwnedObjects()(*i686e52110d509abde81e2ec02a4b1cacb26be255a28efb139b5fe07005585dcb.GetUserOwnedObjectsRequestBuilder) {
    return i686e52110d509abde81e2ec02a4b1cacb26be255a28efb139b5fe07005585dcb.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to directoryRoles
func (m *DirectoryRolesRequestBuilder) Post(options *DirectoryRolesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRole, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectoryRole() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRole), nil
}
func (m *DirectoryRolesRequestBuilder) ValidateProperties()(*ifd6f8f935853943530de1cc46edc30d54c90acd20d9483c7525660fa0c70d33f.ValidatePropertiesRequestBuilder) {
    return ifd6f8f935853943530de1cc46edc30d54c90acd20d9483c7525660fa0c70d33f.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

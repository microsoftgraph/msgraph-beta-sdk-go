package importeddeviceidentities

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2eb32d0eb94562e9b6f79567b48fe799dded573649031029fb86f8965772271e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importeddeviceidentities/searchexistingidentities"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i76339fd2d40042286c7a9b4333ff7246712608a8ec9a407e5fc00e4a39473d6a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importeddeviceidentities/importdeviceidentitylist"
)

// ImportedDeviceIdentitiesRequestBuilder builds and executes requests for operations under \deviceManagement\importedDeviceIdentities
type ImportedDeviceIdentitiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ImportedDeviceIdentitiesRequestBuilderGetOptions options for Get
type ImportedDeviceIdentitiesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ImportedDeviceIdentitiesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ImportedDeviceIdentitiesRequestBuilderGetQueryParameters the imported device identities.
type ImportedDeviceIdentitiesRequestBuilderGetQueryParameters struct {
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
// ImportedDeviceIdentitiesRequestBuilderPostOptions options for Post
type ImportedDeviceIdentitiesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentity;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewImportedDeviceIdentitiesRequestBuilderInternal instantiates a new ImportedDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedDeviceIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ImportedDeviceIdentitiesRequestBuilder) {
    m := &ImportedDeviceIdentitiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/importedDeviceIdentities{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportedDeviceIdentitiesRequestBuilder instantiates a new ImportedDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ImportedDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the imported device identities.
func (m *ImportedDeviceIdentitiesRequestBuilder) CreateGetRequestInformation(options *ImportedDeviceIdentitiesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation the imported device identities.
func (m *ImportedDeviceIdentitiesRequestBuilder) CreatePostRequestInformation(options *ImportedDeviceIdentitiesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the imported device identities.
func (m *ImportedDeviceIdentitiesRequestBuilder) Get(options *ImportedDeviceIdentitiesRequestBuilderGetOptions)(*ImportedDeviceIdentitiesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImportedDeviceIdentitiesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ImportedDeviceIdentitiesResponse), nil
}
func (m *ImportedDeviceIdentitiesRequestBuilder) ImportDeviceIdentityList()(*i76339fd2d40042286c7a9b4333ff7246712608a8ec9a407e5fc00e4a39473d6a.ImportDeviceIdentityListRequestBuilder) {
    return i76339fd2d40042286c7a9b4333ff7246712608a8ec9a407e5fc00e4a39473d6a.NewImportDeviceIdentityListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post the imported device identities.
func (m *ImportedDeviceIdentitiesRequestBuilder) Post(options *ImportedDeviceIdentitiesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentity, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewImportedDeviceIdentity() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentity), nil
}
func (m *ImportedDeviceIdentitiesRequestBuilder) SearchExistingIdentities()(*i2eb32d0eb94562e9b6f79567b48fe799dded573649031029fb86f8965772271e.SearchExistingIdentitiesRequestBuilder) {
    return i2eb32d0eb94562e9b6f79567b48fe799dded573649031029fb86f8965772271e.NewSearchExistingIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

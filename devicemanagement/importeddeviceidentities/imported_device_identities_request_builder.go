package importeddeviceidentities

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i2eb32d0eb94562e9b6f79567b48fe799dded573649031029fb86f8965772271e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importeddeviceidentities/searchexistingidentities"
    i480d1fa84691984b46caafa0c9d056d9c38c0452d47b9d4829c099dde2a6f199 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importeddeviceidentities/count"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i76339fd2d40042286c7a9b4333ff7246712608a8ec9a407e5fc00e4a39473d6a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importeddeviceidentities/importdeviceidentitylist"
)

// ImportedDeviceIdentitiesRequestBuilder provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
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
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentityable;
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
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportedDeviceIdentitiesRequestBuilder instantiates a new ImportedDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ImportedDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ImportedDeviceIdentitiesRequestBuilder) Count()(*i480d1fa84691984b46caafa0c9d056d9c38c0452d47b9d4829c099dde2a6f199.CountRequestBuilder) {
    return i480d1fa84691984b46caafa0c9d056d9c38c0452d47b9d4829c099dde2a6f199.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreatePostRequestInformation create new navigation property to importedDeviceIdentities for deviceManagement
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
func (m *ImportedDeviceIdentitiesRequestBuilder) Get(options *ImportedDeviceIdentitiesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateImportedDeviceIdentityCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentityCollectionResponseable), nil
}
func (m *ImportedDeviceIdentitiesRequestBuilder) ImportDeviceIdentityList()(*i76339fd2d40042286c7a9b4333ff7246712608a8ec9a407e5fc00e4a39473d6a.ImportDeviceIdentityListRequestBuilder) {
    return i76339fd2d40042286c7a9b4333ff7246712608a8ec9a407e5fc00e4a39473d6a.NewImportDeviceIdentityListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to importedDeviceIdentities for deviceManagement
func (m *ImportedDeviceIdentitiesRequestBuilder) Post(options *ImportedDeviceIdentitiesRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentityable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateImportedDeviceIdentityFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedDeviceIdentityable), nil
}
func (m *ImportedDeviceIdentitiesRequestBuilder) SearchExistingIdentities()(*i2eb32d0eb94562e9b6f79567b48fe799dded573649031029fb86f8965772271e.SearchExistingIdentitiesRequestBuilder) {
    return i2eb32d0eb94562e9b6f79567b48fe799dded573649031029fb86f8965772271e.NewSearchExistingIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package importedwindowsautopilotdeviceidentities

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i611aa4102282c0f13028b90d38cf1cff4567e84626f417e14f811917b4228636 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importedwindowsautopilotdeviceidentities/count"
    i8a259409215ade9230da33cd772df83ffd3fb56d6e0feedce6800757cfe3df34 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importedwindowsautopilotdeviceidentities/import_escaped"
)

// ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetOptions options for Get
type ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters collection of imported Windows autopilot devices.
type ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters struct {
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
// ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostOptions options for Post
type ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedWindowsAutopilotDeviceIdentityable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal instantiates a new ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    m := &ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/importedWindowsAutopilotDeviceIdentities{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilder instantiates a new ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Count()(*i611aa4102282c0f13028b90d38cf1cff4567e84626f417e14f811917b4228636.CountRequestBuilder) {
    return i611aa4102282c0f13028b90d38cf1cff4567e84626f417e14f811917b4228636.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation collection of imported Windows autopilot devices.
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) CreateGetRequestInformation(options *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to importedWindowsAutopilotDeviceIdentities for deviceManagement
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) CreatePostRequestInformation(options *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get collection of imported Windows autopilot devices.
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Get(options *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedWindowsAutopilotDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateImportedWindowsAutopilotDeviceIdentityCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedWindowsAutopilotDeviceIdentityCollectionResponseable), nil
}
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Import_escaped()(*i8a259409215ade9230da33cd772df83ffd3fb56d6e0feedce6800757cfe3df34.ImportRequestBuilder) {
    return i8a259409215ade9230da33cd772df83ffd3fb56d6e0feedce6800757cfe3df34.NewImportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to importedWindowsAutopilotDeviceIdentities for deviceManagement
func (m *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Post(options *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedWindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ImportedWindowsAutopilotDeviceIdentityable), nil
}

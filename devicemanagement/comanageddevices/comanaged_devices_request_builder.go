package comanageddevices

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0952f14bd3d84883aa9509ce20c436cb3a9becc286ff26416d1a9ab4f9a5188e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/executeaction"
    i21b46d2a36281265e82fedbdddef1c20e278c5a4bbf3cf85fa0d860eccee5b55 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/movedevicestoou"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7fe672ce2ded63745de2480f0a175b485195547bff2ac413663b22ef3172bb4e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/bulkreprovisioncloudpc"
    i97970601371fcbb5ba7ab85926c15aaf372bd11b13e0e69958d51b9047db72f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/bulkrestorecloudpc"
)

// ComanagedDevicesRequestBuilder builds and executes requests for operations under \deviceManagement\comanagedDevices
type ComanagedDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ComanagedDevicesRequestBuilderGetOptions options for Get
type ComanagedDevicesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ComanagedDevicesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ComanagedDevicesRequestBuilderGetQueryParameters the list of co-managed devices report
type ComanagedDevicesRequestBuilderGetQueryParameters struct {
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
// ComanagedDevicesRequestBuilderPostOptions options for Post
type ComanagedDevicesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ComanagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i7fe672ce2ded63745de2480f0a175b485195547bff2ac413663b22ef3172bb4e.BulkReprovisionCloudPcRequestBuilder) {
    return i7fe672ce2ded63745de2480f0a175b485195547bff2ac413663b22ef3172bb4e.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ComanagedDevicesRequestBuilder) BulkRestoreCloudPc()(*i97970601371fcbb5ba7ab85926c15aaf372bd11b13e0e69958d51b9047db72f5.BulkRestoreCloudPcRequestBuilder) {
    return i97970601371fcbb5ba7ab85926c15aaf372bd11b13e0e69958d51b9047db72f5.NewBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewComanagedDevicesRequestBuilderInternal instantiates a new ComanagedDevicesRequestBuilder and sets the default values.
func NewComanagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ComanagedDevicesRequestBuilder) {
    m := &ComanagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComanagedDevicesRequestBuilder instantiates a new ComanagedDevicesRequestBuilder and sets the default values.
func NewComanagedDevicesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ComanagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the list of co-managed devices report
func (m *ComanagedDevicesRequestBuilder) CreateGetRequestInformation(options *ComanagedDevicesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation the list of co-managed devices report
func (m *ComanagedDevicesRequestBuilder) CreatePostRequestInformation(options *ComanagedDevicesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ComanagedDevicesRequestBuilder) ExecuteAction()(*i0952f14bd3d84883aa9509ce20c436cb3a9becc286ff26416d1a9ab4f9a5188e.ExecuteActionRequestBuilder) {
    return i0952f14bd3d84883aa9509ce20c436cb3a9becc286ff26416d1a9ab4f9a5188e.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of co-managed devices report
func (m *ComanagedDevicesRequestBuilder) Get(options *ComanagedDevicesRequestBuilderGetOptions)(*ComanagedDevicesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComanagedDevicesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ComanagedDevicesResponse), nil
}
func (m *ComanagedDevicesRequestBuilder) MoveDevicesToOU()(*i21b46d2a36281265e82fedbdddef1c20e278c5a4bbf3cf85fa0d860eccee5b55.MoveDevicesToOURequestBuilder) {
    return i21b46d2a36281265e82fedbdddef1c20e278c5a4bbf3cf85fa0d860eccee5b55.NewMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post the list of co-managed devices report
func (m *ComanagedDevicesRequestBuilder) Post(options *ComanagedDevicesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedDevice() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice), nil
}

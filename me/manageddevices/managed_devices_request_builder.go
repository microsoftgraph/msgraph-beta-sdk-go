package manageddevices

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i025773964677aae1b280d645edcafd7be2abcbf2d4e52db52a446e3624435c0d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/bulkreprovisioncloudpc"
    i0376a560b32355e38b49876a437f694bca534f547ea086c5232875988b209b18 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/executeaction"
    i2fa0a3981204d622a8e17b16cd6457fbd3f21735cada71cd2ec3e6dc48806633 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/bulkrestorecloudpc"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8d74a2aa0c51df791a2ec12c0b6d34c1f7541b2b112e11876a6a570c4da51784 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/movedevicestoou"
    i97d3e0c4b436b89b39e01602180f507b9e6f9fbb21bff9015709335c6ebe8051 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/count"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// ManagedDevicesRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ManagedDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDevicesRequestBuilderGetOptions options for Get
type ManagedDevicesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedDevicesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDevicesRequestBuilderGetQueryParameters the managed devices associated with the user.
type ManagedDevicesRequestBuilderGetQueryParameters struct {
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
// ManagedDevicesRequestBuilderPostOptions options for Post
type ManagedDevicesRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i025773964677aae1b280d645edcafd7be2abcbf2d4e52db52a446e3624435c0d.BulkReprovisionCloudPcRequestBuilder) {
    return i025773964677aae1b280d645edcafd7be2abcbf2d4e52db52a446e3624435c0d.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDevicesRequestBuilder) BulkRestoreCloudPc()(*i2fa0a3981204d622a8e17b16cd6457fbd3f21735cada71cd2ec3e6dc48806633.BulkRestoreCloudPcRequestBuilder) {
    return i2fa0a3981204d622a8e17b16cd6457fbd3f21735cada71cd2ec3e6dc48806633.NewBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDevicesRequestBuilderInternal instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    m := &ManagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDevicesRequestBuilder instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedDevicesRequestBuilder) Count()(*i97d3e0c4b436b89b39e01602180f507b9e6f9fbb21bff9015709335c6ebe8051.CountRequestBuilder) {
    return i97d3e0c4b436b89b39e01602180f507b9e6f9fbb21bff9015709335c6ebe8051.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) CreateGetRequestInformation(options *ManagedDevicesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to managedDevices for me
func (m *ManagedDevicesRequestBuilder) CreatePostRequestInformation(options *ManagedDevicesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedDevicesRequestBuilder) ExecuteAction()(*i0376a560b32355e38b49876a437f694bca534f547ea086c5232875988b209b18.ExecuteActionRequestBuilder) {
    return i0376a560b32355e38b49876a437f694bca534f547ea086c5232875988b209b18.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) Get(options *ManagedDevicesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateManagedDeviceCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceCollectionResponseable), nil
}
func (m *ManagedDevicesRequestBuilder) MoveDevicesToOU()(*i8d74a2aa0c51df791a2ec12c0b6d34c1f7541b2b112e11876a6a570c4da51784.MoveDevicesToOURequestBuilder) {
    return i8d74a2aa0c51df791a2ec12c0b6d34c1f7541b2b112e11876a6a570c4da51784.NewMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to managedDevices for me
func (m *ManagedDevicesRequestBuilder) Post(options *ManagedDevicesRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateManagedDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceable), nil
}

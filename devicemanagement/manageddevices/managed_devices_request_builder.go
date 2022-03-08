package manageddevices

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i06835d5500346b7bc8d008091fe2b42358e5cd0e649c006b5492948c0d336599 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/movedevicestoou"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5f8ca051de0bcbbe5b881cdaf7d6d9a13cc6fca814cdda55f7a1fc121c17598e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/bulkrestorecloudpc"
    i9f3c0ec10d79caf7b25f90c50c18ffe886ea494018312ae756500b4daa0b1e5a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/bulkreprovisioncloudpc"
    iaf61f975bbfeb9c90d8d01c8ad1c458f6730caee7a72f9bb3801c4a97ad62d62 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/count"
    ib1edad87ec29538295d4b5f56002379bc4d111c6ca85e79b9e3b58f1a12313aa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/executeaction"
)

// ManagedDevicesRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
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
// ManagedDevicesRequestBuilderGetQueryParameters the list of managed devices.
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
func (m *ManagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i9f3c0ec10d79caf7b25f90c50c18ffe886ea494018312ae756500b4daa0b1e5a.BulkReprovisionCloudPcRequestBuilder) {
    return i9f3c0ec10d79caf7b25f90c50c18ffe886ea494018312ae756500b4daa0b1e5a.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDevicesRequestBuilder) BulkRestoreCloudPc()(*i5f8ca051de0bcbbe5b881cdaf7d6d9a13cc6fca814cdda55f7a1fc121c17598e.BulkRestoreCloudPcRequestBuilder) {
    return i5f8ca051de0bcbbe5b881cdaf7d6d9a13cc6fca814cdda55f7a1fc121c17598e.NewBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDevicesRequestBuilderInternal instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    m := &ManagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices{?top,skip,search,filter,count,orderby,select,expand}";
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
func (m *ManagedDevicesRequestBuilder) Count()(*iaf61f975bbfeb9c90d8d01c8ad1c458f6730caee7a72f9bb3801c4a97ad62d62.CountRequestBuilder) {
    return iaf61f975bbfeb9c90d8d01c8ad1c458f6730caee7a72f9bb3801c4a97ad62d62.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of managed devices.
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
// CreatePostRequestInformation create new navigation property to managedDevices for deviceManagement
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
func (m *ManagedDevicesRequestBuilder) ExecuteAction()(*ib1edad87ec29538295d4b5f56002379bc4d111c6ca85e79b9e3b58f1a12313aa.ExecuteActionRequestBuilder) {
    return ib1edad87ec29538295d4b5f56002379bc4d111c6ca85e79b9e3b58f1a12313aa.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of managed devices.
func (m *ManagedDevicesRequestBuilder) Get(options *ManagedDevicesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateManagedDeviceCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceCollectionResponseable), nil
}
func (m *ManagedDevicesRequestBuilder) MoveDevicesToOU()(*i06835d5500346b7bc8d008091fe2b42358e5cd0e649c006b5492948c0d336599.MoveDevicesToOURequestBuilder) {
    return i06835d5500346b7bc8d008091fe2b42358e5cd0e649c006b5492948c0d336599.NewMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to managedDevices for deviceManagement
func (m *ManagedDevicesRequestBuilder) Post(options *ManagedDevicesRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateManagedDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceable), nil
}

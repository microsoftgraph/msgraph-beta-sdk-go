package manageddevices

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1018d4869692703b9b231a9d9796ea086b51d50b4ab465fc18bfa4bfec0e9951 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/detectedapps/item/manageddevices/executeaction"
    i179bbad3399daa7952a7169991dad820c980669224959d8820b27df3705cfea2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/detectedapps/item/manageddevices/movedevicestoou"
    i247cbfb29f4a7ba9c2ae9e5a868d95fbcc0128e2c169700e507e3ba4e042bb27 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/detectedapps/item/manageddevices/bulkreprovisioncloudpc"
    i860b3077a0ea6a4a5feada40f3b0d1f1d9795bd2e63ecbee4ef47f25ada7c0f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/detectedapps/item/manageddevices/bulkrestorecloudpc"
    idc3947230f9ec3c69eb3e851f6ac557afab9214686da7582beed0fef3c1a0fd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/detectedapps/item/manageddevices/ref"
)

// ManagedDevicesRequestBuilder builds and executes requests for operations under \deviceManagement\detectedApps\{detectedApp-id}\managedDevices
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
// ManagedDevicesRequestBuilderGetQueryParameters the devices that have the discovered application installed
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
func (m *ManagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i247cbfb29f4a7ba9c2ae9e5a868d95fbcc0128e2c169700e507e3ba4e042bb27.BulkReprovisionCloudPcRequestBuilder) {
    return i247cbfb29f4a7ba9c2ae9e5a868d95fbcc0128e2c169700e507e3ba4e042bb27.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDevicesRequestBuilder) BulkRestoreCloudPc()(*i860b3077a0ea6a4a5feada40f3b0d1f1d9795bd2e63ecbee4ef47f25ada7c0f6.BulkRestoreCloudPcRequestBuilder) {
    return i860b3077a0ea6a4a5feada40f3b0d1f1d9795bd2e63ecbee4ef47f25ada7c0f6.NewBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDevicesRequestBuilderInternal instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    m := &ManagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/detectedApps/{detectedApp_id}/managedDevices{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDevicesRequestBuilder instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the devices that have the discovered application installed
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
func (m *ManagedDevicesRequestBuilder) ExecuteAction()(*i1018d4869692703b9b231a9d9796ea086b51d50b4ab465fc18bfa4bfec0e9951.ExecuteActionRequestBuilder) {
    return i1018d4869692703b9b231a9d9796ea086b51d50b4ab465fc18bfa4bfec0e9951.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the devices that have the discovered application installed
func (m *ManagedDevicesRequestBuilder) Get(options *ManagedDevicesRequestBuilderGetOptions)(*ManagedDevicesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevicesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagedDevicesResponse), nil
}
func (m *ManagedDevicesRequestBuilder) MoveDevicesToOU()(*i179bbad3399daa7952a7169991dad820c980669224959d8820b27df3705cfea2.MoveDevicesToOURequestBuilder) {
    return i179bbad3399daa7952a7169991dad820c980669224959d8820b27df3705cfea2.NewMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDevicesRequestBuilder) Ref()(*idc3947230f9ec3c69eb3e851f6ac557afab9214686da7582beed0fef3c1a0fd8.RefRequestBuilder) {
    return idc3947230f9ec3c69eb3e851f6ac557afab9214686da7582beed0fef3c1a0fd8.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

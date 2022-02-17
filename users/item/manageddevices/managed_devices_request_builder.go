package manageddevices

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i221b9cde99e476699da097a8d903ef5a79f3532dc9f4f0e5ee20a66cefd5214f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/bulkreprovisioncloudpc"
    i354da0a7c77a3e4dbbf4c2908e40eb9ba102f2173af0e90bc4d10956f0e44dee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/movedevicestoou"
    i8d8a434efe54b7a3fa263217b0c04ff177a4bc1c2230b98a73b80cd1d561c96d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/executeaction"
    i992082cec9bab654e1839fe55796615f8620ef8a4078776f6e65f83f16c347b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/bulkrestorecloudpc"
)

// ManagedDevicesRequestBuilder builds and executes requests for operations under \users\{user-id}\managedDevices
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
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i221b9cde99e476699da097a8d903ef5a79f3532dc9f4f0e5ee20a66cefd5214f.BulkReprovisionCloudPcRequestBuilder) {
    return i221b9cde99e476699da097a8d903ef5a79f3532dc9f4f0e5ee20a66cefd5214f.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDevicesRequestBuilder) BulkRestoreCloudPc()(*i992082cec9bab654e1839fe55796615f8620ef8a4078776f6e65f83f16c347b1.BulkRestoreCloudPcRequestBuilder) {
    return i992082cec9bab654e1839fe55796615f8620ef8a4078776f6e65f83f16c347b1.NewBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDevicesRequestBuilderInternal instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    m := &ManagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/managedDevices{?top,skip,search,filter,count,orderby,select,expand}";
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
// CreatePostRequestInformation the managed devices associated with the user.
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
func (m *ManagedDevicesRequestBuilder) ExecuteAction()(*i8d8a434efe54b7a3fa263217b0c04ff177a4bc1c2230b98a73b80cd1d561c96d.ExecuteActionRequestBuilder) {
    return i8d8a434efe54b7a3fa263217b0c04ff177a4bc1c2230b98a73b80cd1d561c96d.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices associated with the user.
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
func (m *ManagedDevicesRequestBuilder) MoveDevicesToOU()(*i354da0a7c77a3e4dbbf4c2908e40eb9ba102f2173af0e90bc4d10956f0e44dee.MoveDevicesToOURequestBuilder) {
    return i354da0a7c77a3e4dbbf4c2908e40eb9ba102f2173af0e90bc4d10956f0e44dee.NewMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) Post(options *ManagedDevicesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice, error) {
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

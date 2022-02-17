package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    if76c2cb80ed357d07ffd85f409f06a5249ec510099f57edb418f41e7bdb74e03 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/autopilotevents/item/policystatusdetails"
    ib32d34f50255bdea0699208e320acf08a90078315ca8ef82a9c10a80d7799d4b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/autopilotevents/item/policystatusdetails/item"
)

// DeviceManagementAutopilotEventRequestBuilder builds and executes requests for operations under \deviceManagement\autopilotEvents\{deviceManagementAutopilotEvent-id}
type DeviceManagementAutopilotEventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementAutopilotEventRequestBuilderDeleteOptions options for Delete
type DeviceManagementAutopilotEventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementAutopilotEventRequestBuilderGetOptions options for Get
type DeviceManagementAutopilotEventRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementAutopilotEventRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementAutopilotEventRequestBuilderGetQueryParameters the list of autopilot events for the tenant.
type DeviceManagementAutopilotEventRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementAutopilotEventRequestBuilderPatchOptions options for Patch
type DeviceManagementAutopilotEventRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementAutopilotEvent;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceManagementAutopilotEventRequestBuilderInternal instantiates a new DeviceManagementAutopilotEventRequestBuilder and sets the default values.
func NewDeviceManagementAutopilotEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementAutopilotEventRequestBuilder) {
    m := &DeviceManagementAutopilotEventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/autopilotEvents/{deviceManagementAutopilotEvent_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementAutopilotEventRequestBuilder instantiates a new DeviceManagementAutopilotEventRequestBuilder and sets the default values.
func NewDeviceManagementAutopilotEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementAutopilotEventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementAutopilotEventRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of autopilot events for the tenant.
func (m *DeviceManagementAutopilotEventRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementAutopilotEventRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation the list of autopilot events for the tenant.
func (m *DeviceManagementAutopilotEventRequestBuilder) CreateGetRequestInformation(options *DeviceManagementAutopilotEventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of autopilot events for the tenant.
func (m *DeviceManagementAutopilotEventRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementAutopilotEventRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete the list of autopilot events for the tenant.
func (m *DeviceManagementAutopilotEventRequestBuilder) Delete(options *DeviceManagementAutopilotEventRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of autopilot events for the tenant.
func (m *DeviceManagementAutopilotEventRequestBuilder) Get(options *DeviceManagementAutopilotEventRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementAutopilotEvent, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementAutopilotEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementAutopilotEvent), nil
}
// Patch the list of autopilot events for the tenant.
func (m *DeviceManagementAutopilotEventRequestBuilder) Patch(options *DeviceManagementAutopilotEventRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceManagementAutopilotEventRequestBuilder) PolicyStatusDetails()(*if76c2cb80ed357d07ffd85f409f06a5249ec510099f57edb418f41e7bdb74e03.PolicyStatusDetailsRequestBuilder) {
    return if76c2cb80ed357d07ffd85f409f06a5249ec510099f57edb418f41e7bdb74e03.NewPolicyStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PolicyStatusDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.autopilotEvents.item.policyStatusDetails.item collection
func (m *DeviceManagementAutopilotEventRequestBuilder) PolicyStatusDetailsById(id string)(*ib32d34f50255bdea0699208e320acf08a90078315ca8ef82a9c10a80d7799d4b.DeviceManagementAutopilotPolicyStatusDetailRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementAutopilotPolicyStatusDetail_id"] = id
    }
    return ib32d34f50255bdea0699208e320acf08a90078315ca8ef82a9c10a80d7799d4b.NewDeviceManagementAutopilotPolicyStatusDetailRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

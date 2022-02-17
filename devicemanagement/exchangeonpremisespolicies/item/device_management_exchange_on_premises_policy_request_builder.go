package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0ddcec3ebabd376004d9ebf44592b0aee802011a19e8f490cebd899271d38877 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/exchangeonpremisespolicies/item/conditionalaccesssettings"
)

// DeviceManagementExchangeOnPremisesPolicyRequestBuilder builds and executes requests for operations under \deviceManagement\exchangeOnPremisesPolicies\{deviceManagementExchangeOnPremisesPolicy-id}
type DeviceManagementExchangeOnPremisesPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementExchangeOnPremisesPolicyRequestBuilderDeleteOptions options for Delete
type DeviceManagementExchangeOnPremisesPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementExchangeOnPremisesPolicyRequestBuilderGetOptions options for Get
type DeviceManagementExchangeOnPremisesPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementExchangeOnPremisesPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementExchangeOnPremisesPolicyRequestBuilderGetQueryParameters the list of Exchange On Premisis policies configured by the tenant.
type DeviceManagementExchangeOnPremisesPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementExchangeOnPremisesPolicyRequestBuilderPatchOptions options for Patch
type DeviceManagementExchangeOnPremisesPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementExchangeOnPremisesPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementExchangeOnPremisesPolicyRequestBuilder) ConditionalAccessSettings()(*i0ddcec3ebabd376004d9ebf44592b0aee802011a19e8f490cebd899271d38877.ConditionalAccessSettingsRequestBuilder) {
    return i0ddcec3ebabd376004d9ebf44592b0aee802011a19e8f490cebd899271d38877.NewConditionalAccessSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementExchangeOnPremisesPolicyRequestBuilderInternal instantiates a new DeviceManagementExchangeOnPremisesPolicyRequestBuilder and sets the default values.
func NewDeviceManagementExchangeOnPremisesPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementExchangeOnPremisesPolicyRequestBuilder) {
    m := &DeviceManagementExchangeOnPremisesPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/exchangeOnPremisesPolicies/{deviceManagementExchangeOnPremisesPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementExchangeOnPremisesPolicyRequestBuilder instantiates a new DeviceManagementExchangeOnPremisesPolicyRequestBuilder and sets the default values.
func NewDeviceManagementExchangeOnPremisesPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementExchangeOnPremisesPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementExchangeOnPremisesPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagementExchangeOnPremisesPolicyRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementExchangeOnPremisesPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagementExchangeOnPremisesPolicyRequestBuilder) CreateGetRequestInformation(options *DeviceManagementExchangeOnPremisesPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagementExchangeOnPremisesPolicyRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementExchangeOnPremisesPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagementExchangeOnPremisesPolicyRequestBuilder) Delete(options *DeviceManagementExchangeOnPremisesPolicyRequestBuilderDeleteOptions)(error) {
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
// Get the list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagementExchangeOnPremisesPolicyRequestBuilder) Get(options *DeviceManagementExchangeOnPremisesPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementExchangeOnPremisesPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementExchangeOnPremisesPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementExchangeOnPremisesPolicy), nil
}
// Patch the list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagementExchangeOnPremisesPolicyRequestBuilder) Patch(options *DeviceManagementExchangeOnPremisesPolicyRequestBuilderPatchOptions)(error) {
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

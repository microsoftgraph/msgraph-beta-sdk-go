package devicecompliancepolicies

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i58651a263a770247eef91ec5783125ddffea19f0485aeac13f756043009a7093 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/validatecompliancescript"
    i5f664252c457e99b8d5db6b834754f966442a3164d883ca219d4dbf2362b3cdc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/setscheduledretirestate"
    i8c922b0106272c33024960c0cdcce96f4a55b9e9d3d906cb19b2281dbb71993a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/haspayloadlinks"
    ib90740a26262712164b9160ce4822a569344cfc69adafea017f8563b8e220b51 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/refreshdevicecompliancereportsummarization"
    ice8333de4faa45aa6e283d6de5cec5fb2d91ddb4641c608a190a6c8ebda92bc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/getdevicesscheduledtoretire"
)

// DeviceCompliancePoliciesRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCompliancePolicies
type DeviceCompliancePoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceCompliancePoliciesRequestBuilderGetOptions options for Get
type DeviceCompliancePoliciesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceCompliancePoliciesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCompliancePoliciesRequestBuilderGetQueryParameters the device compliance policies.
type DeviceCompliancePoliciesRequestBuilderGetQueryParameters struct {
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
// DeviceCompliancePoliciesRequestBuilderPostOptions options for Post
type DeviceCompliancePoliciesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceCompliancePoliciesRequestBuilderInternal instantiates a new DeviceCompliancePoliciesRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePoliciesRequestBuilder) {
    m := &DeviceCompliancePoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCompliancePoliciesRequestBuilder instantiates a new DeviceCompliancePoliciesRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the device compliance policies.
func (m *DeviceCompliancePoliciesRequestBuilder) CreateGetRequestInformation(options *DeviceCompliancePoliciesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation the device compliance policies.
func (m *DeviceCompliancePoliciesRequestBuilder) CreatePostRequestInformation(options *DeviceCompliancePoliciesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the device compliance policies.
func (m *DeviceCompliancePoliciesRequestBuilder) Get(options *DeviceCompliancePoliciesRequestBuilderGetOptions)(*DeviceCompliancePoliciesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePoliciesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DeviceCompliancePoliciesResponse), nil
}
func (m *DeviceCompliancePoliciesRequestBuilder) GetDevicesScheduledToRetire()(*ice8333de4faa45aa6e283d6de5cec5fb2d91ddb4641c608a190a6c8ebda92bc5.GetDevicesScheduledToRetireRequestBuilder) {
    return ice8333de4faa45aa6e283d6de5cec5fb2d91ddb4641c608a190a6c8ebda92bc5.NewGetDevicesScheduledToRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePoliciesRequestBuilder) HasPayloadLinks()(*i8c922b0106272c33024960c0cdcce96f4a55b9e9d3d906cb19b2281dbb71993a.HasPayloadLinksRequestBuilder) {
    return i8c922b0106272c33024960c0cdcce96f4a55b9e9d3d906cb19b2281dbb71993a.NewHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post the device compliance policies.
func (m *DeviceCompliancePoliciesRequestBuilder) Post(options *DeviceCompliancePoliciesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicy, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceCompliancePolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicy), nil
}
func (m *DeviceCompliancePoliciesRequestBuilder) RefreshDeviceComplianceReportSummarization()(*ib90740a26262712164b9160ce4822a569344cfc69adafea017f8563b8e220b51.RefreshDeviceComplianceReportSummarizationRequestBuilder) {
    return ib90740a26262712164b9160ce4822a569344cfc69adafea017f8563b8e220b51.NewRefreshDeviceComplianceReportSummarizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePoliciesRequestBuilder) SetScheduledRetireState()(*i5f664252c457e99b8d5db6b834754f966442a3164d883ca219d4dbf2362b3cdc.SetScheduledRetireStateRequestBuilder) {
    return i5f664252c457e99b8d5db6b834754f966442a3164d883ca219d4dbf2362b3cdc.NewSetScheduledRetireStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePoliciesRequestBuilder) ValidateComplianceScript()(*i58651a263a770247eef91ec5783125ddffea19f0485aeac13f756043009a7093.ValidateComplianceScriptRequestBuilder) {
    return i58651a263a770247eef91ec5783125ddffea19f0485aeac13f756043009a7093.NewValidateComplianceScriptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

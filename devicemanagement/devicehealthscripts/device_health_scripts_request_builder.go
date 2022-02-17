package devicehealthscripts

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/getremediationsummary"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6a873728e2a8769c250276e0c4dc8fa971581156a07a4c225120f5c5242df0a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/areglobalscriptsavailable"
    i87802417fac8dafea49fd0889f538e357ba4f700228dd931789b2c3a9194e714 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/enableglobalscripts"
)

// DeviceHealthScriptsRequestBuilder builds and executes requests for operations under \deviceManagement\deviceHealthScripts
type DeviceHealthScriptsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceHealthScriptsRequestBuilderGetOptions options for Get
type DeviceHealthScriptsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceHealthScriptsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceHealthScriptsRequestBuilderGetQueryParameters the list of device health scripts associated with the tenant.
type DeviceHealthScriptsRequestBuilderGetQueryParameters struct {
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
// DeviceHealthScriptsRequestBuilderPostOptions options for Post
type DeviceHealthScriptsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScript;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AreGlobalScriptsAvailable builds and executes requests for operations under \deviceManagement\deviceHealthScripts\microsoft.graph.areGlobalScriptsAvailable()
func (m *DeviceHealthScriptsRequestBuilder) AreGlobalScriptsAvailable()(*i6a873728e2a8769c250276e0c4dc8fa971581156a07a4c225120f5c5242df0a3.AreGlobalScriptsAvailableRequestBuilder) {
    return i6a873728e2a8769c250276e0c4dc8fa971581156a07a4c225120f5c5242df0a3.NewAreGlobalScriptsAvailableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceHealthScriptsRequestBuilderInternal instantiates a new DeviceHealthScriptsRequestBuilder and sets the default values.
func NewDeviceHealthScriptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceHealthScriptsRequestBuilder) {
    m := &DeviceHealthScriptsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceHealthScriptsRequestBuilder instantiates a new DeviceHealthScriptsRequestBuilder and sets the default values.
func NewDeviceHealthScriptsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceHealthScriptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) CreateGetRequestInformation(options *DeviceHealthScriptsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) CreatePostRequestInformation(options *DeviceHealthScriptsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceHealthScriptsRequestBuilder) EnableGlobalScripts()(*i87802417fac8dafea49fd0889f538e357ba4f700228dd931789b2c3a9194e714.EnableGlobalScriptsRequestBuilder) {
    return i87802417fac8dafea49fd0889f538e357ba4f700228dd931789b2c3a9194e714.NewEnableGlobalScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) Get(options *DeviceHealthScriptsRequestBuilderGetOptions)(*DeviceHealthScriptsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DeviceHealthScriptsResponse), nil
}
// GetRemediationSummary builds and executes requests for operations under \deviceManagement\deviceHealthScripts\microsoft.graph.getRemediationSummary()
func (m *DeviceHealthScriptsRequestBuilder) GetRemediationSummary()(*i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547.GetRemediationSummaryRequestBuilder) {
    return i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547.NewGetRemediationSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) Post(options *DeviceHealthScriptsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScript, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceHealthScript() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScript), nil
}

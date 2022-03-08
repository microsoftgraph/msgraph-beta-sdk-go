package devicehealthscripts

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/getremediationsummary"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6a873728e2a8769c250276e0c4dc8fa971581156a07a4c225120f5c5242df0a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/areglobalscriptsavailable"
    i87802417fac8dafea49fd0889f538e357ba4f700228dd931789b2c3a9194e714 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/enableglobalscripts"
    ic4de1ae6615c5df9eec3108346be9478c79d3a7aa8645c40d6f5091d453580f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/count"
)

// DeviceHealthScriptsRequestBuilder provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
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
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AreGlobalScriptsAvailable provides operations to call the areGlobalScriptsAvailable method.
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
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceHealthScriptsRequestBuilder instantiates a new DeviceHealthScriptsRequestBuilder and sets the default values.
func NewDeviceHealthScriptsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceHealthScriptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceHealthScriptsRequestBuilder) Count()(*ic4de1ae6615c5df9eec3108346be9478c79d3a7aa8645c40d6f5091d453580f1.CountRequestBuilder) {
    return ic4de1ae6615c5df9eec3108346be9478c79d3a7aa8645c40d6f5091d453580f1.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreatePostRequestInformation create new navigation property to deviceHealthScripts for deviceManagement
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
func (m *DeviceHealthScriptsRequestBuilder) Get(options *DeviceHealthScriptsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceHealthScriptCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptCollectionResponseable), nil
}
// GetRemediationSummary provides operations to call the getRemediationSummary method.
func (m *DeviceHealthScriptsRequestBuilder) GetRemediationSummary()(*i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547.GetRemediationSummaryRequestBuilder) {
    return i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547.NewGetRemediationSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptsRequestBuilder) Post(options *DeviceHealthScriptsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceHealthScriptFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptable), nil
}

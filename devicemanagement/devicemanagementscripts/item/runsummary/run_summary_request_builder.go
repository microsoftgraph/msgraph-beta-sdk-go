package runsummary

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i88cff6105125945cfaafb039ac893486ca373c7334c06a63e3a1d067b1d5ad96 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/runsummary/ref"
)

// RunSummaryRequestBuilder builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\runSummary
type RunSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RunSummaryRequestBuilderGetOptions options for Get
type RunSummaryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RunSummaryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RunSummaryRequestBuilderGetQueryParameters run summary for device management script.
type RunSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewRunSummaryRequestBuilderInternal instantiates a new RunSummaryRequestBuilder and sets the default values.
func NewRunSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RunSummaryRequestBuilder) {
    m := &RunSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript_id}/runSummary{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRunSummaryRequestBuilder instantiates a new RunSummaryRequestBuilder and sets the default values.
func NewRunSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RunSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation run summary for device management script.
func (m *RunSummaryRequestBuilder) CreateGetRequestInformation(options *RunSummaryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get run summary for device management script.
func (m *RunSummaryRequestBuilder) Get(options *RunSummaryRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementScriptRunSummary, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementScriptRunSummary() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementScriptRunSummary), nil
}
func (m *RunSummaryRequestBuilder) Ref()(*i88cff6105125945cfaafb039ac893486ca373c7334c06a63e3a1d067b1d5ad96.RefRequestBuilder) {
    return i88cff6105125945cfaafb039ac893486ca373c7334c06a63e3a1d067b1d5ad96.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

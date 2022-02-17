package getwindowsqualityupdatealertsperpolicyperdevicereport

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder builds and executes requests for operations under \deviceManagement\reports\microsoft.graph.getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport
type GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions options for Post
type GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions struct {
    // 
    Body *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal instantiates a new GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    m := &GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder instantiates a new GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) CreatePostRequestInformation(options *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) Post(options *GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "byte", nil, nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}

package getwindowsupdatealertsperpolicyperdevicereport

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder provides operations to call the getWindowsUpdateAlertsPerPolicyPerDeviceReport method.
type GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions options for Post
type GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions struct {
    // 
    Body GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal instantiates a new GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    m := &GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getWindowsUpdateAlertsPerPolicyPerDeviceReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder instantiates a new GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder and sets the default values.
func NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getWindowsUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) CreatePostRequestInformation(options *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getWindowsUpdateAlertsPerPolicyPerDeviceReport
func (m *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) Post(options *GetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderPostOptions)(GetWindowsUpdateAlertsPerPolicyPerDeviceReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetWindowsUpdateAlertsPerPolicyPerDeviceReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetWindowsUpdateAlertsPerPolicyPerDeviceReportResponseable), nil
}

package getwindowsupdatealertsummaryreport

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetWindowsUpdateAlertSummaryReportRequestBuilder builds and executes requests for operations under \deviceManagement\reports\microsoft.graph.getWindowsUpdateAlertSummaryReport
type GetWindowsUpdateAlertSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetWindowsUpdateAlertSummaryReportRequestBuilderPostOptions options for Post
type GetWindowsUpdateAlertSummaryReportRequestBuilderPostOptions struct {
    // 
    Body *GetWindowsUpdateAlertSummaryReportRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetWindowsUpdateAlertSummaryReportRequestBuilderInternal instantiates a new GetWindowsUpdateAlertSummaryReportRequestBuilder and sets the default values.
func NewGetWindowsUpdateAlertSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetWindowsUpdateAlertSummaryReportRequestBuilder) {
    m := &GetWindowsUpdateAlertSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getWindowsUpdateAlertSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetWindowsUpdateAlertSummaryReportRequestBuilder instantiates a new GetWindowsUpdateAlertSummaryReportRequestBuilder and sets the default values.
func NewGetWindowsUpdateAlertSummaryReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetWindowsUpdateAlertSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetWindowsUpdateAlertSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getWindowsUpdateAlertSummaryReport
func (m *GetWindowsUpdateAlertSummaryReportRequestBuilder) CreatePostRequestInformation(options *GetWindowsUpdateAlertSummaryReportRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getWindowsUpdateAlertSummaryReport
func (m *GetWindowsUpdateAlertSummaryReportRequestBuilder) Post(options *GetWindowsUpdateAlertSummaryReportRequestBuilderPostOptions)([]byte, error) {
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

package getfailedmobileappssummaryreport

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetFailedMobileAppsSummaryReportRequestBuilder provides operations to call the getFailedMobileAppsSummaryReport method.
type GetFailedMobileAppsSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetFailedMobileAppsSummaryReportRequestBuilderPostOptions options for Post
type GetFailedMobileAppsSummaryReportRequestBuilderPostOptions struct {
    // 
    Body GetFailedMobileAppsSummaryReportRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetFailedMobileAppsSummaryReportRequestBuilderInternal instantiates a new GetFailedMobileAppsSummaryReportRequestBuilder and sets the default values.
func NewGetFailedMobileAppsSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetFailedMobileAppsSummaryReportRequestBuilder) {
    m := &GetFailedMobileAppsSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getFailedMobileAppsSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetFailedMobileAppsSummaryReportRequestBuilder instantiates a new GetFailedMobileAppsSummaryReportRequestBuilder and sets the default values.
func NewGetFailedMobileAppsSummaryReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetFailedMobileAppsSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetFailedMobileAppsSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getFailedMobileAppsSummaryReport
func (m *GetFailedMobileAppsSummaryReportRequestBuilder) CreatePostRequestInformation(options *GetFailedMobileAppsSummaryReportRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getFailedMobileAppsSummaryReport
func (m *GetFailedMobileAppsSummaryReportRequestBuilder) Post(options *GetFailedMobileAppsSummaryReportRequestBuilderPostOptions)(GetFailedMobileAppsSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetFailedMobileAppsSummaryReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetFailedMobileAppsSummaryReportResponseable), nil
}

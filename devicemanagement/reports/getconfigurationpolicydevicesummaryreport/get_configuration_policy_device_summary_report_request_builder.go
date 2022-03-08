package getconfigurationpolicydevicesummaryreport

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetConfigurationPolicyDeviceSummaryReportRequestBuilder provides operations to call the getConfigurationPolicyDeviceSummaryReport method.
type GetConfigurationPolicyDeviceSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetConfigurationPolicyDeviceSummaryReportRequestBuilderPostOptions options for Post
type GetConfigurationPolicyDeviceSummaryReportRequestBuilderPostOptions struct {
    // 
    Body GetConfigurationPolicyDeviceSummaryReportRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal instantiates a new GetConfigurationPolicyDeviceSummaryReportRequestBuilder and sets the default values.
func NewGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    m := &GetConfigurationPolicyDeviceSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getConfigurationPolicyDeviceSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetConfigurationPolicyDeviceSummaryReportRequestBuilder instantiates a new GetConfigurationPolicyDeviceSummaryReportRequestBuilder and sets the default values.
func NewGetConfigurationPolicyDeviceSummaryReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getConfigurationPolicyDeviceSummaryReport
func (m *GetConfigurationPolicyDeviceSummaryReportRequestBuilder) CreatePostRequestInformation(options *GetConfigurationPolicyDeviceSummaryReportRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getConfigurationPolicyDeviceSummaryReport
func (m *GetConfigurationPolicyDeviceSummaryReportRequestBuilder) Post(options *GetConfigurationPolicyDeviceSummaryReportRequestBuilderPostOptions)(GetConfigurationPolicyDeviceSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetConfigurationPolicyDeviceSummaryReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetConfigurationPolicyDeviceSummaryReportResponseable), nil
}

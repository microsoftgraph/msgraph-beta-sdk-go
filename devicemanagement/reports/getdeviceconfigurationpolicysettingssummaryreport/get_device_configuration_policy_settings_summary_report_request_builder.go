package getdeviceconfigurationpolicysettingssummaryreport

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder provides operations to call the getDeviceConfigurationPolicySettingsSummaryReport method.
type GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderPostOptions options for Post
type GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderPostOptions struct {
    // 
    Body GetDeviceConfigurationPolicySettingsSummaryReportRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal instantiates a new GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder and sets the default values.
func NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) {
    m := &GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getDeviceConfigurationPolicySettingsSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder instantiates a new GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder and sets the default values.
func NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getDeviceConfigurationPolicySettingsSummaryReport
func (m *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) CreatePostRequestInformation(options *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getDeviceConfigurationPolicySettingsSummaryReport
func (m *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) Post(options *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderPostOptions)(GetDeviceConfigurationPolicySettingsSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetDeviceConfigurationPolicySettingsSummaryReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetDeviceConfigurationPolicySettingsSummaryReportResponseable), nil
}

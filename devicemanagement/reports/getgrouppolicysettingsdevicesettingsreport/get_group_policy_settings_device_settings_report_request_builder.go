package getgrouppolicysettingsdevicesettingsreport

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetGroupPolicySettingsDeviceSettingsReportRequestBuilder provides operations to call the getGroupPolicySettingsDeviceSettingsReport method.
type GetGroupPolicySettingsDeviceSettingsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetGroupPolicySettingsDeviceSettingsReportRequestBuilderPostOptions options for Post
type GetGroupPolicySettingsDeviceSettingsReportRequestBuilderPostOptions struct {
    // 
    Body GetGroupPolicySettingsDeviceSettingsReportRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal instantiates a new GetGroupPolicySettingsDeviceSettingsReportRequestBuilder and sets the default values.
func NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) {
    m := &GetGroupPolicySettingsDeviceSettingsReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getGroupPolicySettingsDeviceSettingsReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilder instantiates a new GetGroupPolicySettingsDeviceSettingsReportRequestBuilder and sets the default values.
func NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getGroupPolicySettingsDeviceSettingsReport
func (m *GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) CreatePostRequestInformation(options *GetGroupPolicySettingsDeviceSettingsReportRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getGroupPolicySettingsDeviceSettingsReport
func (m *GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) Post(options *GetGroupPolicySettingsDeviceSettingsReportRequestBuilderPostOptions)(GetGroupPolicySettingsDeviceSettingsReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetGroupPolicySettingsDeviceSettingsReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetGroupPolicySettingsDeviceSettingsReportResponseable), nil
}

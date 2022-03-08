package getdeviceconfigurationpolicystatussummary

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetDeviceConfigurationPolicyStatusSummaryRequestBuilder provides operations to call the getDeviceConfigurationPolicyStatusSummary method.
type GetDeviceConfigurationPolicyStatusSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetDeviceConfigurationPolicyStatusSummaryRequestBuilderPostOptions options for Post
type GetDeviceConfigurationPolicyStatusSummaryRequestBuilderPostOptions struct {
    // 
    Body GetDeviceConfigurationPolicyStatusSummaryRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal instantiates a new GetDeviceConfigurationPolicyStatusSummaryRequestBuilder and sets the default values.
func NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) {
    m := &GetDeviceConfigurationPolicyStatusSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getDeviceConfigurationPolicyStatusSummary";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilder instantiates a new GetDeviceConfigurationPolicyStatusSummaryRequestBuilder and sets the default values.
func NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getDeviceConfigurationPolicyStatusSummary
func (m *GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) CreatePostRequestInformation(options *GetDeviceConfigurationPolicyStatusSummaryRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getDeviceConfigurationPolicyStatusSummary
func (m *GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) Post(options *GetDeviceConfigurationPolicyStatusSummaryRequestBuilderPostOptions)(GetDeviceConfigurationPolicyStatusSummaryResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetDeviceConfigurationPolicyStatusSummaryResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetDeviceConfigurationPolicyStatusSummaryResponseable), nil
}

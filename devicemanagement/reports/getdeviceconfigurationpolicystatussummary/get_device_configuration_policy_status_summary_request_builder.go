package getdeviceconfigurationpolicystatussummary

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetDeviceConfigurationPolicyStatusSummaryRequestBuilder provides operations to call the getDeviceConfigurationPolicyStatusSummary method.
type GetDeviceConfigurationPolicyStatusSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetDeviceConfigurationPolicyStatusSummaryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetDeviceConfigurationPolicyStatusSummaryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal instantiates a new GetDeviceConfigurationPolicyStatusSummaryRequestBuilder and sets the default values.
func NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) {
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
func NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getDeviceConfigurationPolicyStatusSummary
func (m *GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) CreatePostRequestInformation(body GetDeviceConfigurationPolicyStatusSummaryRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getDeviceConfigurationPolicyStatusSummary
func (m *GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetDeviceConfigurationPolicyStatusSummaryRequestBodyable, requestConfiguration *GetDeviceConfigurationPolicyStatusSummaryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getDeviceConfigurationPolicyStatusSummary
func (m *GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) Post(body GetDeviceConfigurationPolicyStatusSummaryRequestBodyable)(GetDeviceConfigurationPolicyStatusSummaryResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getDeviceConfigurationPolicyStatusSummary
func (m *GetDeviceConfigurationPolicyStatusSummaryRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetDeviceConfigurationPolicyStatusSummaryRequestBodyable, requestConfiguration *GetDeviceConfigurationPolicyStatusSummaryRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetDeviceConfigurationPolicyStatusSummaryResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetDeviceConfigurationPolicyStatusSummaryResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetDeviceConfigurationPolicyStatusSummaryResponseable), nil
}

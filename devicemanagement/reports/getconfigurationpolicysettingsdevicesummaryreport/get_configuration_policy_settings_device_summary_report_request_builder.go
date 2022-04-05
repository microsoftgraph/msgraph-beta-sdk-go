package getconfigurationpolicysettingsdevicesummaryreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder provides operations to call the getConfigurationPolicySettingsDeviceSummaryReport method.
type GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderPostOptions options for Post
type GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderPostOptions struct {
    // 
    Body GetConfigurationPolicySettingsDeviceSummaryReportRequestBodyable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderInternal instantiates a new GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder and sets the default values.
func NewGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) {
    m := &GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getConfigurationPolicySettingsDeviceSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder instantiates a new GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder and sets the default values.
func NewGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getConfigurationPolicySettingsDeviceSummaryReport
func (m *GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) CreatePostRequestInformation(options *GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action getConfigurationPolicySettingsDeviceSummaryReport
func (m *GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) Post(options *GetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderPostOptions)(GetConfigurationPolicySettingsDeviceSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetConfigurationPolicySettingsDeviceSummaryReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetConfigurationPolicySettingsDeviceSummaryReportResponseable), nil
}

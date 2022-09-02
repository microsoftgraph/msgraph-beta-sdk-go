package getgrouppolicysettingsdevicesettingsreport

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetGroupPolicySettingsDeviceSettingsReportRequestBuilder provides operations to call the getGroupPolicySettingsDeviceSettingsReport method.
type GetGroupPolicySettingsDeviceSettingsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetGroupPolicySettingsDeviceSettingsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetGroupPolicySettingsDeviceSettingsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal instantiates a new GetGroupPolicySettingsDeviceSettingsReportRequestBuilder and sets the default values.
func NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) {
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
func NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getGroupPolicySettingsDeviceSettingsReport
func (m *GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) CreatePostRequestInformation(body GetGroupPolicySettingsDeviceSettingsReportPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getGroupPolicySettingsDeviceSettingsReport
func (m *GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetGroupPolicySettingsDeviceSettingsReportPostRequestBodyable, requestConfiguration *GetGroupPolicySettingsDeviceSettingsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getGroupPolicySettingsDeviceSettingsReport
func (m *GetGroupPolicySettingsDeviceSettingsReportRequestBuilder) Post(ctx context.Context, body GetGroupPolicySettingsDeviceSettingsReportPostRequestBodyable, requestConfiguration *GetGroupPolicySettingsDeviceSettingsReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}

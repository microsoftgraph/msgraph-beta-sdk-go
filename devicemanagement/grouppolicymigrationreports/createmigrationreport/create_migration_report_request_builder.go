package createmigrationreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CreateMigrationReportRequestBuilder provides operations to call the createMigrationReport method.
type CreateMigrationReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CreateMigrationReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CreateMigrationReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCreateMigrationReportRequestBuilderInternal instantiates a new CreateMigrationReportRequestBuilder and sets the default values.
func NewCreateMigrationReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateMigrationReportRequestBuilder) {
    m := &CreateMigrationReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyMigrationReports/microsoft.graph.createMigrationReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreateMigrationReportRequestBuilder instantiates a new CreateMigrationReportRequestBuilder and sets the default values.
func NewCreateMigrationReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateMigrationReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateMigrationReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action createMigrationReport
func (m *CreateMigrationReportRequestBuilder) CreatePostRequestInformation(body CreateMigrationReportPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action createMigrationReport
func (m *CreateMigrationReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body CreateMigrationReportPostRequestBodyable, requestConfiguration *CreateMigrationReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action createMigrationReport
func (m *CreateMigrationReportRequestBuilder) Post(body CreateMigrationReportPostRequestBodyable)(CreateMigrationReportResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action createMigrationReport
func (m *CreateMigrationReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body CreateMigrationReportPostRequestBodyable, requestConfiguration *CreateMigrationReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(CreateMigrationReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateCreateMigrationReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(CreateMigrationReportResponseable), nil
}

package getzebrafotadeploymentreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetZebraFotaDeploymentReportRequestBuilder provides operations to call the getZebraFotaDeploymentReport method.
type GetZebraFotaDeploymentReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetZebraFotaDeploymentReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetZebraFotaDeploymentReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetZebraFotaDeploymentReportRequestBuilderInternal instantiates a new GetZebraFotaDeploymentReportRequestBuilder and sets the default values.
func NewGetZebraFotaDeploymentReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetZebraFotaDeploymentReportRequestBuilder) {
    m := &GetZebraFotaDeploymentReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getZebraFotaDeploymentReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetZebraFotaDeploymentReportRequestBuilder instantiates a new GetZebraFotaDeploymentReportRequestBuilder and sets the default values.
func NewGetZebraFotaDeploymentReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetZebraFotaDeploymentReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetZebraFotaDeploymentReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getZebraFotaDeploymentReport
func (m *GetZebraFotaDeploymentReportRequestBuilder) CreatePostRequestInformation(body GetZebraFotaDeploymentReportPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getZebraFotaDeploymentReport
func (m *GetZebraFotaDeploymentReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetZebraFotaDeploymentReportPostRequestBodyable, requestConfiguration *GetZebraFotaDeploymentReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getZebraFotaDeploymentReport
func (m *GetZebraFotaDeploymentReportRequestBuilder) Post(body GetZebraFotaDeploymentReportPostRequestBodyable)([]byte, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getZebraFotaDeploymentReport
func (m *GetZebraFotaDeploymentReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetZebraFotaDeploymentReportPostRequestBodyable, requestConfiguration *GetZebraFotaDeploymentReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(requestInfo, "byte", responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}

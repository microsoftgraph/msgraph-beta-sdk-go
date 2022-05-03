package evaluateassignmentfilter

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EvaluateAssignmentFilterRequestBuilder provides operations to call the evaluateAssignmentFilter method.
type EvaluateAssignmentFilterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EvaluateAssignmentFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EvaluateAssignmentFilterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEvaluateAssignmentFilterRequestBuilderInternal instantiates a new EvaluateAssignmentFilterRequestBuilder and sets the default values.
func NewEvaluateAssignmentFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EvaluateAssignmentFilterRequestBuilder) {
    m := &EvaluateAssignmentFilterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.evaluateAssignmentFilter";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEvaluateAssignmentFilterRequestBuilder instantiates a new EvaluateAssignmentFilterRequestBuilder and sets the default values.
func NewEvaluateAssignmentFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EvaluateAssignmentFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEvaluateAssignmentFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action evaluateAssignmentFilter
func (m *EvaluateAssignmentFilterRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body EvaluateAssignmentFilterRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action evaluateAssignmentFilter
func (m *EvaluateAssignmentFilterRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body EvaluateAssignmentFilterRequestBodyable, requestConfiguration *EvaluateAssignmentFilterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action evaluateAssignmentFilter
func (m *EvaluateAssignmentFilterRequestBuilder) PostWithResponseHandler(body EvaluateAssignmentFilterRequestBodyable, requestConfiguration *EvaluateAssignmentFilterRequestBuilderPostRequestConfiguration)(EvaluateAssignmentFilterResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action evaluateAssignmentFilter
func (m *EvaluateAssignmentFilterRequestBuilder) PostWithResponseHandler(body EvaluateAssignmentFilterRequestBodyable, requestConfiguration *EvaluateAssignmentFilterRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(EvaluateAssignmentFilterResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateEvaluateAssignmentFilterResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(EvaluateAssignmentFilterResponseable), nil
}

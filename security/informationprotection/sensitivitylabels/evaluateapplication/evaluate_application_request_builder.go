package evaluateapplication

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type EvaluateApplicationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEvaluateApplicationRequestBuilderInternal instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EvaluateApplicationRequestBuilder) {
    m := &EvaluateApplicationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateApplication";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEvaluateApplicationRequestBuilder instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action evaluateApplication
func (m *EvaluateApplicationRequestBuilder) CreatePostRequestInformation(body EvaluateApplicationPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action evaluateApplication
func (m *EvaluateApplicationRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body EvaluateApplicationPostRequestBodyable, requestConfiguration *EvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action evaluateApplication
func (m *EvaluateApplicationRequestBuilder) Post(body EvaluateApplicationPostRequestBodyable)(EvaluateApplicationResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action evaluateApplication
func (m *EvaluateApplicationRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body EvaluateApplicationPostRequestBodyable, requestConfiguration *EvaluateApplicationRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(EvaluateApplicationResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateEvaluateApplicationResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(EvaluateApplicationResponseable), nil
}

package setscheduledretirestate

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SetScheduledRetireStateRequestBuilder provides operations to call the setScheduledRetireState method.
type SetScheduledRetireStateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SetScheduledRetireStateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetScheduledRetireStateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSetScheduledRetireStateRequestBuilderInternal instantiates a new SetScheduledRetireStateRequestBuilder and sets the default values.
func NewSetScheduledRetireStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetScheduledRetireStateRequestBuilder) {
    m := &SetScheduledRetireStateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/microsoft.graph.setScheduledRetireState";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetScheduledRetireStateRequestBuilder instantiates a new SetScheduledRetireStateRequestBuilder and sets the default values.
func NewSetScheduledRetireStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetScheduledRetireStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetScheduledRetireStateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action setScheduledRetireState
func (m *SetScheduledRetireStateRequestBuilder) CreatePostRequestInformation(body SetScheduledRetireStatePostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action setScheduledRetireState
func (m *SetScheduledRetireStateRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SetScheduledRetireStatePostRequestBodyable, requestConfiguration *SetScheduledRetireStateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action setScheduledRetireState
func (m *SetScheduledRetireStateRequestBuilder) Post(body SetScheduledRetireStatePostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action setScheduledRetireState
func (m *SetScheduledRetireStateRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body SetScheduledRetireStatePostRequestBodyable, requestConfiguration *SetScheduledRetireStateRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}

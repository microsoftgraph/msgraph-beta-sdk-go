package updatepriorities

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UpdatePrioritiesRequestBuilder provides operations to call the updatePriorities method.
type UpdatePrioritiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UpdatePrioritiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UpdatePrioritiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUpdatePrioritiesRequestBuilderInternal instantiates a new UpdatePrioritiesRequestBuilder and sets the default values.
func NewUpdatePrioritiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdatePrioritiesRequestBuilder) {
    m := &UpdatePrioritiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/officeConfiguration/clientConfigurations/microsoft.graph.updatePriorities";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdatePrioritiesRequestBuilder instantiates a new UpdatePrioritiesRequestBuilder and sets the default values.
func NewUpdatePrioritiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdatePrioritiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatePrioritiesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updatePriorities
func (m *UpdatePrioritiesRequestBuilder) CreatePostRequestInformation(body UpdatePrioritiesPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action updatePriorities
func (m *UpdatePrioritiesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body UpdatePrioritiesPostRequestBodyable, requestConfiguration *UpdatePrioritiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action updatePriorities
func (m *UpdatePrioritiesRequestBuilder) Post(body UpdatePrioritiesPostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action updatePriorities
func (m *UpdatePrioritiesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body UpdatePrioritiesPostRequestBodyable, requestConfiguration *UpdatePrioritiesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

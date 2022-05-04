package querybyplatformtype

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// QueryByPlatformTypeRequestBuilder provides operations to call the queryByPlatformType method.
type QueryByPlatformTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// QueryByPlatformTypeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type QueryByPlatformTypeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewQueryByPlatformTypeRequestBuilderInternal instantiates a new QueryByPlatformTypeRequestBuilder and sets the default values.
func NewQueryByPlatformTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*QueryByPlatformTypeRequestBuilder) {
    m := &QueryByPlatformTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/resourceAccessProfiles/microsoft.graph.queryByPlatformType";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewQueryByPlatformTypeRequestBuilder instantiates a new QueryByPlatformTypeRequestBuilder and sets the default values.
func NewQueryByPlatformTypeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*QueryByPlatformTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewQueryByPlatformTypeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action queryByPlatformType
func (m *QueryByPlatformTypeRequestBuilder) CreatePostRequestInformation(body QueryByPlatformTypeRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action queryByPlatformType
func (m *QueryByPlatformTypeRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body QueryByPlatformTypeRequestBodyable, requestConfiguration *QueryByPlatformTypeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action queryByPlatformType
func (m *QueryByPlatformTypeRequestBuilder) Post(body QueryByPlatformTypeRequestBodyable)(QueryByPlatformTypeResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action queryByPlatformType
func (m *QueryByPlatformTypeRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body QueryByPlatformTypeRequestBodyable, requestConfiguration *QueryByPlatformTypeRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(QueryByPlatformTypeResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateQueryByPlatformTypeResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(QueryByPlatformTypeResponseable), nil
}

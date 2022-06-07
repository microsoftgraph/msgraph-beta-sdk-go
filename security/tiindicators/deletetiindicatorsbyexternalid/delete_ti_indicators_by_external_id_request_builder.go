package deletetiindicatorsbyexternalid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeleteTiIndicatorsByExternalIdRequestBuilder provides operations to call the deleteTiIndicatorsByExternalId method.
type DeleteTiIndicatorsByExternalIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeleteTiIndicatorsByExternalIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeleteTiIndicatorsByExternalIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeleteTiIndicatorsByExternalIdRequestBuilderInternal instantiates a new DeleteTiIndicatorsByExternalIdRequestBuilder and sets the default values.
func NewDeleteTiIndicatorsByExternalIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteTiIndicatorsByExternalIdRequestBuilder) {
    m := &DeleteTiIndicatorsByExternalIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/tiIndicators/microsoft.graph.deleteTiIndicatorsByExternalId";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeleteTiIndicatorsByExternalIdRequestBuilder instantiates a new DeleteTiIndicatorsByExternalIdRequestBuilder and sets the default values.
func NewDeleteTiIndicatorsByExternalIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteTiIndicatorsByExternalIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeleteTiIndicatorsByExternalIdRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action deleteTiIndicatorsByExternalId
func (m *DeleteTiIndicatorsByExternalIdRequestBuilder) CreatePostRequestInformation(body DeleteTiIndicatorsByExternalIdPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action deleteTiIndicatorsByExternalId
func (m *DeleteTiIndicatorsByExternalIdRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body DeleteTiIndicatorsByExternalIdPostRequestBodyable, requestConfiguration *DeleteTiIndicatorsByExternalIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action deleteTiIndicatorsByExternalId
func (m *DeleteTiIndicatorsByExternalIdRequestBuilder) Post(body DeleteTiIndicatorsByExternalIdPostRequestBodyable)(DeleteTiIndicatorsByExternalIdResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action deleteTiIndicatorsByExternalId
func (m *DeleteTiIndicatorsByExternalIdRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body DeleteTiIndicatorsByExternalIdPostRequestBodyable, requestConfiguration *DeleteTiIndicatorsByExternalIdRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(DeleteTiIndicatorsByExternalIdResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateDeleteTiIndicatorsByExternalIdResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(DeleteTiIndicatorsByExternalIdResponseable), nil
}

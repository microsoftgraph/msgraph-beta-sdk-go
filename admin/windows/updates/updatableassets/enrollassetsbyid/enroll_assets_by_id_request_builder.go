package enrollassetsbyid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EnrollAssetsByIdRequestBuilder provides operations to call the enrollAssetsById method.
type EnrollAssetsByIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EnrollAssetsByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnrollAssetsByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEnrollAssetsByIdRequestBuilderInternal instantiates a new EnrollAssetsByIdRequestBuilder and sets the default values.
func NewEnrollAssetsByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnrollAssetsByIdRequestBuilder) {
    m := &EnrollAssetsByIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/updatableAssets/microsoft.graph.windowsUpdates.enrollAssetsById";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEnrollAssetsByIdRequestBuilder instantiates a new EnrollAssetsByIdRequestBuilder and sets the default values.
func NewEnrollAssetsByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnrollAssetsByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnrollAssetsByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action enrollAssetsById
func (m *EnrollAssetsByIdRequestBuilder) CreatePostRequestInformation(body EnrollAssetsByIdPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action enrollAssetsById
func (m *EnrollAssetsByIdRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body EnrollAssetsByIdPostRequestBodyable, requestConfiguration *EnrollAssetsByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action enrollAssetsById
func (m *EnrollAssetsByIdRequestBuilder) Post(body EnrollAssetsByIdPostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action enrollAssetsById
func (m *EnrollAssetsByIdRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body EnrollAssetsByIdPostRequestBodyable, requestConfiguration *EnrollAssetsByIdRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

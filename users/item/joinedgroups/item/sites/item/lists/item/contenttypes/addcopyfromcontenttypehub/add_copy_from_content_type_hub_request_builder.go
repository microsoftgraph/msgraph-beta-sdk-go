package addcopyfromcontenttypehub

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AddCopyFromContentTypeHubRequestBuilder provides operations to call the addCopyFromContentTypeHub method.
type AddCopyFromContentTypeHubRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAddCopyFromContentTypeHubRequestBuilderInternal instantiates a new AddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewAddCopyFromContentTypeHubRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AddCopyFromContentTypeHubRequestBuilder) {
    m := &AddCopyFromContentTypeHubRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/microsoft.graph.addCopyFromContentTypeHub";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAddCopyFromContentTypeHubRequestBuilder instantiates a new AddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewAddCopyFromContentTypeHubRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AddCopyFromContentTypeHubRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAddCopyFromContentTypeHubRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action addCopyFromContentTypeHub
func (m *AddCopyFromContentTypeHubRequestBuilder) CreatePostRequestInformation(body AddCopyFromContentTypeHubRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action addCopyFromContentTypeHub
func (m *AddCopyFromContentTypeHubRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body AddCopyFromContentTypeHubRequestBodyable, requestConfiguration *AddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action addCopyFromContentTypeHub
func (m *AddCopyFromContentTypeHubRequestBuilder) Post(body AddCopyFromContentTypeHubRequestBodyable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action addCopyFromContentTypeHub
func (m *AddCopyFromContentTypeHubRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body AddCopyFromContentTypeHubRequestBodyable, requestConfiguration *AddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}

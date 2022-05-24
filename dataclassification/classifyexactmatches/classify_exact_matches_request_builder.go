package classifyexactmatches

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ClassifyExactMatchesRequestBuilder provides operations to call the classifyExactMatches method.
type ClassifyExactMatchesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ClassifyExactMatchesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassifyExactMatchesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewClassifyExactMatchesRequestBuilderInternal instantiates a new ClassifyExactMatchesRequestBuilder and sets the default values.
func NewClassifyExactMatchesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassifyExactMatchesRequestBuilder) {
    m := &ClassifyExactMatchesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/dataClassification/microsoft.graph.classifyExactMatches";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewClassifyExactMatchesRequestBuilder instantiates a new ClassifyExactMatchesRequestBuilder and sets the default values.
func NewClassifyExactMatchesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassifyExactMatchesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassifyExactMatchesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action classifyExactMatches
func (m *ClassifyExactMatchesRequestBuilder) CreatePostRequestInformation(body ClassifyExactMatchesPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action classifyExactMatches
func (m *ClassifyExactMatchesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ClassifyExactMatchesPostRequestBodyable, requestConfiguration *ClassifyExactMatchesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action classifyExactMatches
func (m *ClassifyExactMatchesRequestBuilder) Post(body ClassifyExactMatchesPostRequestBodyable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchClassificationResultable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action classifyExactMatches
func (m *ClassifyExactMatchesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ClassifyExactMatchesPostRequestBodyable, requestConfiguration *ClassifyExactMatchesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchClassificationResultable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExactMatchClassificationResultFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExactMatchClassificationResultable), nil
}

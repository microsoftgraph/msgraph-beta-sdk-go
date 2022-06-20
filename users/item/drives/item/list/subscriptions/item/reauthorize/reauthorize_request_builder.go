package reauthorize

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReauthorizeRequestBuilder provides operations to call the reauthorize method.
type ReauthorizeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReauthorizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReauthorizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReauthorizeRequestBuilderInternal instantiates a new ReauthorizeRequestBuilder and sets the default values.
func NewReauthorizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReauthorizeRequestBuilder) {
    m := &ReauthorizeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/list/subscriptions/{subscription%2Did}/microsoft.graph.reauthorize";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReauthorizeRequestBuilder instantiates a new ReauthorizeRequestBuilder and sets the default values.
func NewReauthorizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReauthorizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReauthorizeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action reauthorize
func (m *ReauthorizeRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action reauthorize
func (m *ReauthorizeRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *ReauthorizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action reauthorize
func (m *ReauthorizeRequestBuilder) Post()(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action reauthorize
func (m *ReauthorizeRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *ReauthorizeRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}

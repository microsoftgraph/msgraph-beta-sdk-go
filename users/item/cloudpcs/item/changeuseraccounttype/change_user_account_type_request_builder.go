package changeuseraccounttype

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChangeUserAccountTypeRequestBuilder provides operations to call the changeUserAccountType method.
type ChangeUserAccountTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChangeUserAccountTypeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChangeUserAccountTypeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChangeUserAccountTypeRequestBuilderInternal instantiates a new ChangeUserAccountTypeRequestBuilder and sets the default values.
func NewChangeUserAccountTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChangeUserAccountTypeRequestBuilder) {
    m := &ChangeUserAccountTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/microsoft.graph.changeUserAccountType";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChangeUserAccountTypeRequestBuilder instantiates a new ChangeUserAccountTypeRequestBuilder and sets the default values.
func NewChangeUserAccountTypeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChangeUserAccountTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChangeUserAccountTypeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action changeUserAccountType
func (m *ChangeUserAccountTypeRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ChangeUserAccountTypeRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action changeUserAccountType
func (m *ChangeUserAccountTypeRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ChangeUserAccountTypeRequestBodyable, requestConfiguration *ChangeUserAccountTypeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action changeUserAccountType
func (m *ChangeUserAccountTypeRequestBuilder) PostWithResponseHandler(body ChangeUserAccountTypeRequestBodyable, requestConfiguration *ChangeUserAccountTypeRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action changeUserAccountType
func (m *ChangeUserAccountTypeRequestBuilder) PostWithResponseHandler(body ChangeUserAccountTypeRequestBodyable, requestConfiguration *ChangeUserAccountTypeRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

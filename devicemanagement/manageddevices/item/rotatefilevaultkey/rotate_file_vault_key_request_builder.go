package rotatefilevaultkey

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RotateFileVaultKeyRequestBuilder provides operations to call the rotateFileVaultKey method.
type RotateFileVaultKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RotateFileVaultKeyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RotateFileVaultKeyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRotateFileVaultKeyRequestBuilderInternal instantiates a new RotateFileVaultKeyRequestBuilder and sets the default values.
func NewRotateFileVaultKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RotateFileVaultKeyRequestBuilder) {
    m := &RotateFileVaultKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/microsoft.graph.rotateFileVaultKey";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRotateFileVaultKeyRequestBuilder instantiates a new RotateFileVaultKeyRequestBuilder and sets the default values.
func NewRotateFileVaultKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RotateFileVaultKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRotateFileVaultKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action rotateFileVaultKey
func (m *RotateFileVaultKeyRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action rotateFileVaultKey
func (m *RotateFileVaultKeyRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *RotateFileVaultKeyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action rotateFileVaultKey
func (m *RotateFileVaultKeyRequestBuilder) PostWithResponseHandler(requestConfiguration *RotateFileVaultKeyRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler invoke action rotateFileVaultKey
func (m *RotateFileVaultKeyRequestBuilder) PostWithResponseHandler(requestConfiguration *RotateFileVaultKeyRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}

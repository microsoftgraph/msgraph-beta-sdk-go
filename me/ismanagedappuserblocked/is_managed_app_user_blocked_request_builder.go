package ismanagedappuserblocked

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IsManagedAppUserBlockedRequestBuilder provides operations to call the isManagedAppUserBlocked method.
type IsManagedAppUserBlockedRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IsManagedAppUserBlockedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IsManagedAppUserBlockedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIsManagedAppUserBlockedRequestBuilderInternal instantiates a new IsManagedAppUserBlockedRequestBuilder and sets the default values.
func NewIsManagedAppUserBlockedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IsManagedAppUserBlockedRequestBuilder) {
    m := &IsManagedAppUserBlockedRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.isManagedAppUserBlocked()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIsManagedAppUserBlockedRequestBuilder instantiates a new IsManagedAppUserBlockedRequestBuilder and sets the default values.
func NewIsManagedAppUserBlockedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IsManagedAppUserBlockedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIsManagedAppUserBlockedRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformationWithRequestConfiguration gets the blocked state of a managed app user.
func (m *IsManagedAppUserBlockedRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration gets the blocked state of a managed app user.
func (m *IsManagedAppUserBlockedRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *IsManagedAppUserBlockedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler gets the blocked state of a managed app user.
func (m *IsManagedAppUserBlockedRequestBuilder) GetWithResponseHandler(requestConfiguration *IsManagedAppUserBlockedRequestBuilderGetRequestConfiguration)(IsManagedAppUserBlockedResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler gets the blocked state of a managed app user.
func (m *IsManagedAppUserBlockedRequestBuilder) GetWithResponseHandler(requestConfiguration *IsManagedAppUserBlockedRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(IsManagedAppUserBlockedResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateIsManagedAppUserBlockedResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(IsManagedAppUserBlockedResponseable), nil
}

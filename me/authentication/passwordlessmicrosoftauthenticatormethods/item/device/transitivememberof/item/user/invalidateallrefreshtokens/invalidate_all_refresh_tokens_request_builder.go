package invalidateallrefreshtokens

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InvalidateAllRefreshTokensRequestBuilder provides operations to call the invalidateAllRefreshTokens method.
type InvalidateAllRefreshTokensRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInvalidateAllRefreshTokensRequestBuilderInternal instantiates a new InvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewInvalidateAllRefreshTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InvalidateAllRefreshTokensRequestBuilder) {
    m := &InvalidateAllRefreshTokensRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user/microsoft.graph.invalidateAllRefreshTokens";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInvalidateAllRefreshTokensRequestBuilder instantiates a new InvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewInvalidateAllRefreshTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InvalidateAllRefreshTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInvalidateAllRefreshTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action invalidateAllRefreshTokens
func (m *InvalidateAllRefreshTokensRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action invalidateAllRefreshTokens
func (m *InvalidateAllRefreshTokensRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *InvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action invalidateAllRefreshTokens
func (m *InvalidateAllRefreshTokensRequestBuilder) Post()(InvalidateAllRefreshTokensResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action invalidateAllRefreshTokens
func (m *InvalidateAllRefreshTokensRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *InvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(InvalidateAllRefreshTokensResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateInvalidateAllRefreshTokensResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(InvalidateAllRefreshTokensResponseable), nil
}

package wipeandblockmanagedapps

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WipeAndBlockManagedAppsRequestBuilder provides operations to call the wipeAndBlockManagedApps method.
type WipeAndBlockManagedAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WipeAndBlockManagedAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WipeAndBlockManagedAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWipeAndBlockManagedAppsRequestBuilderInternal instantiates a new WipeAndBlockManagedAppsRequestBuilder and sets the default values.
func NewWipeAndBlockManagedAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WipeAndBlockManagedAppsRequestBuilder) {
    m := &WipeAndBlockManagedAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user/microsoft.graph.wipeAndBlockManagedApps";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWipeAndBlockManagedAppsRequestBuilder instantiates a new WipeAndBlockManagedAppsRequestBuilder and sets the default values.
func NewWipeAndBlockManagedAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WipeAndBlockManagedAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWipeAndBlockManagedAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation blocks the managed app user from app check-in.
func (m *WipeAndBlockManagedAppsRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration blocks the managed app user from app check-in.
func (m *WipeAndBlockManagedAppsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *WipeAndBlockManagedAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post blocks the managed app user from app check-in.
func (m *WipeAndBlockManagedAppsRequestBuilder) Post()(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler blocks the managed app user from app check-in.
func (m *WipeAndBlockManagedAppsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *WipeAndBlockManagedAppsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

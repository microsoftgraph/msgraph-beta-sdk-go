package extendqualityupdatespause

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExtendQualityUpdatesPauseRequestBuilder provides operations to call the extendQualityUpdatesPause method.
type ExtendQualityUpdatesPauseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExtendQualityUpdatesPauseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExtendQualityUpdatesPauseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExtendQualityUpdatesPauseRequestBuilderInternal instantiates a new ExtendQualityUpdatesPauseRequestBuilder and sets the default values.
func NewExtendQualityUpdatesPauseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExtendQualityUpdatesPauseRequestBuilder) {
    m := &ExtendQualityUpdatesPauseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/microsoft.graph.windowsUpdateForBusinessConfiguration/microsoft.graph.extendQualityUpdatesPause";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExtendQualityUpdatesPauseRequestBuilder instantiates a new ExtendQualityUpdatesPauseRequestBuilder and sets the default values.
func NewExtendQualityUpdatesPauseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExtendQualityUpdatesPauseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExtendQualityUpdatesPauseRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration extend Quality Updates Pause for a Windows Update for Business ring.
func (m *ExtendQualityUpdatesPauseRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration extend Quality Updates Pause for a Windows Update for Business ring.
func (m *ExtendQualityUpdatesPauseRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *ExtendQualityUpdatesPauseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler extend Quality Updates Pause for a Windows Update for Business ring.
func (m *ExtendQualityUpdatesPauseRequestBuilder) PostWithResponseHandler(requestConfiguration *ExtendQualityUpdatesPauseRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler extend Quality Updates Pause for a Windows Update for Business ring.
func (m *ExtendQualityUpdatesPauseRequestBuilder) PostWithResponseHandler(requestConfiguration *ExtendQualityUpdatesPauseRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

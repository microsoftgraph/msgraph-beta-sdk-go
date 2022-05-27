package triggerdevicescopeaction

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TriggerDeviceScopeActionRequestBuilder provides operations to call the triggerDeviceScopeAction method.
type TriggerDeviceScopeActionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TriggerDeviceScopeActionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggerDeviceScopeActionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTriggerDeviceScopeActionRequestBuilderInternal instantiates a new TriggerDeviceScopeActionRequestBuilder and sets the default values.
func NewTriggerDeviceScopeActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggerDeviceScopeActionRequestBuilder) {
    m := &TriggerDeviceScopeActionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceScopes/{userExperienceAnalyticsDeviceScope%2Did}/microsoft.graph.triggerDeviceScopeAction";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTriggerDeviceScopeActionRequestBuilder instantiates a new TriggerDeviceScopeActionRequestBuilder and sets the default values.
func NewTriggerDeviceScopeActionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggerDeviceScopeActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTriggerDeviceScopeActionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action triggerDeviceScopeAction
func (m *TriggerDeviceScopeActionRequestBuilder) CreatePostRequestInformation(body TriggerDeviceScopeActionPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action triggerDeviceScopeAction
func (m *TriggerDeviceScopeActionRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body TriggerDeviceScopeActionPostRequestBodyable, requestConfiguration *TriggerDeviceScopeActionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action triggerDeviceScopeAction
func (m *TriggerDeviceScopeActionRequestBuilder) Post(body TriggerDeviceScopeActionPostRequestBodyable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceScopeActionResultable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action triggerDeviceScopeAction
func (m *TriggerDeviceScopeActionRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body TriggerDeviceScopeActionPostRequestBodyable, requestConfiguration *TriggerDeviceScopeActionRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceScopeActionResultable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceScopeActionResultFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceScopeActionResultable), nil
}

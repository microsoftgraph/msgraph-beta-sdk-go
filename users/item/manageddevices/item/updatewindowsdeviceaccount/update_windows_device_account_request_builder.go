package updatewindowsdeviceaccount

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UpdateWindowsDeviceAccountRequestBuilder provides operations to call the updateWindowsDeviceAccount method.
type UpdateWindowsDeviceAccountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UpdateWindowsDeviceAccountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UpdateWindowsDeviceAccountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUpdateWindowsDeviceAccountRequestBuilderInternal instantiates a new UpdateWindowsDeviceAccountRequestBuilder and sets the default values.
func NewUpdateWindowsDeviceAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdateWindowsDeviceAccountRequestBuilder) {
    m := &UpdateWindowsDeviceAccountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/microsoft.graph.updateWindowsDeviceAccount";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdateWindowsDeviceAccountRequestBuilder instantiates a new UpdateWindowsDeviceAccountRequestBuilder and sets the default values.
func NewUpdateWindowsDeviceAccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UpdateWindowsDeviceAccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdateWindowsDeviceAccountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updateWindowsDeviceAccount
func (m *UpdateWindowsDeviceAccountRequestBuilder) CreatePostRequestInformation(body UpdateWindowsDeviceAccountPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action updateWindowsDeviceAccount
func (m *UpdateWindowsDeviceAccountRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body UpdateWindowsDeviceAccountPostRequestBodyable, requestConfiguration *UpdateWindowsDeviceAccountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action updateWindowsDeviceAccount
func (m *UpdateWindowsDeviceAccountRequestBuilder) Post(body UpdateWindowsDeviceAccountPostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action updateWindowsDeviceAccount
func (m *UpdateWindowsDeviceAccountRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body UpdateWindowsDeviceAccountPostRequestBodyable, requestConfiguration *UpdateWindowsDeviceAccountRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

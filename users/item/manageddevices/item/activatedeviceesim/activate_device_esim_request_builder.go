package activatedeviceesim

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ActivateDeviceEsimRequestBuilder provides operations to call the activateDeviceEsim method.
type ActivateDeviceEsimRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ActivateDeviceEsimRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActivateDeviceEsimRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewActivateDeviceEsimRequestBuilderInternal instantiates a new ActivateDeviceEsimRequestBuilder and sets the default values.
func NewActivateDeviceEsimRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActivateDeviceEsimRequestBuilder) {
    m := &ActivateDeviceEsimRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/microsoft.graph.activateDeviceEsim";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewActivateDeviceEsimRequestBuilder instantiates a new ActivateDeviceEsimRequestBuilder and sets the default values.
func NewActivateDeviceEsimRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActivateDeviceEsimRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActivateDeviceEsimRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation activate eSIM on the device.
func (m *ActivateDeviceEsimRequestBuilder) CreatePostRequestInformation(body ActivateDeviceEsimRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration activate eSIM on the device.
func (m *ActivateDeviceEsimRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ActivateDeviceEsimRequestBodyable, requestConfiguration *ActivateDeviceEsimRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post activate eSIM on the device.
func (m *ActivateDeviceEsimRequestBuilder) Post(body ActivateDeviceEsimRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler activate eSIM on the device.
func (m *ActivateDeviceEsimRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ActivateDeviceEsimRequestBodyable, requestConfiguration *ActivateDeviceEsimRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

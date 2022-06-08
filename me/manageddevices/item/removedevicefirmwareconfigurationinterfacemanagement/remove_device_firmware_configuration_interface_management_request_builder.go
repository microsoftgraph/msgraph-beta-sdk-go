package removedevicefirmwareconfigurationinterfacemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
type RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal instantiates a new RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    m := &RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}/microsoft.graph.removeDeviceFirmwareConfigurationInterfaceManagement";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder instantiates a new RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action removeDeviceFirmwareConfigurationInterfaceManagement
func (m *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action removeDeviceFirmwareConfigurationInterfaceManagement
func (m *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action removeDeviceFirmwareConfigurationInterfaceManagement
func (m *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) Post()(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action removeDeviceFirmwareConfigurationInterfaceManagement
func (m *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *RemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

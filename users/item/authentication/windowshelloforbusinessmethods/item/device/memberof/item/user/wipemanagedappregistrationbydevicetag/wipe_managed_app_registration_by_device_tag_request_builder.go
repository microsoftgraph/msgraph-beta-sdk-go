package wipemanagedappregistrationbydevicetag

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WipeManagedAppRegistrationByDeviceTagRequestBuilder provides operations to call the wipeManagedAppRegistrationByDeviceTag method.
type WipeManagedAppRegistrationByDeviceTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal instantiates a new WipeManagedAppRegistrationByDeviceTagRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    m := &WipeManagedAppRegistrationByDeviceTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user/microsoft.graph.wipeManagedAppRegistrationByDeviceTag";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWipeManagedAppRegistrationByDeviceTagRequestBuilder instantiates a new WipeManagedAppRegistrationByDeviceTagRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationByDeviceTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation issues a wipe operation on an app registration with specified device tag.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBuilder) CreatePostRequestInformation(body WipeManagedAppRegistrationByDeviceTagPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration issues a wipe operation on an app registration with specified device tag.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body WipeManagedAppRegistrationByDeviceTagPostRequestBodyable, requestConfiguration *WipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post issues a wipe operation on an app registration with specified device tag.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBuilder) Post(body WipeManagedAppRegistrationByDeviceTagPostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler issues a wipe operation on an app registration with specified device tag.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body WipeManagedAppRegistrationByDeviceTagPostRequestBodyable, requestConfiguration *WipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

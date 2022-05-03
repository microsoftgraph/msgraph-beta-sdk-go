package windowsprivacyaccesscontrols

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WindowsPrivacyAccessControlsRequestBuilder provides operations to call the windowsPrivacyAccessControls method.
type WindowsPrivacyAccessControlsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsPrivacyAccessControlsRequestBuilderInternal instantiates a new WindowsPrivacyAccessControlsRequestBuilder and sets the default values.
func NewWindowsPrivacyAccessControlsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsPrivacyAccessControlsRequestBuilder) {
    m := &WindowsPrivacyAccessControlsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/microsoft.graph.windowsPrivacyAccessControls";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsPrivacyAccessControlsRequestBuilder instantiates a new WindowsPrivacyAccessControlsRequestBuilder and sets the default values.
func NewWindowsPrivacyAccessControlsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsPrivacyAccessControlsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsPrivacyAccessControlsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action windowsPrivacyAccessControls
func (m *WindowsPrivacyAccessControlsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body WindowsPrivacyAccessControlsRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action windowsPrivacyAccessControls
func (m *WindowsPrivacyAccessControlsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body WindowsPrivacyAccessControlsRequestBodyable, requestConfiguration *WindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action windowsPrivacyAccessControls
func (m *WindowsPrivacyAccessControlsRequestBuilder) PostWithResponseHandler(body WindowsPrivacyAccessControlsRequestBodyable, requestConfiguration *WindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action windowsPrivacyAccessControls
func (m *WindowsPrivacyAccessControlsRequestBuilder) PostWithResponseHandler(body WindowsPrivacyAccessControlsRequestBodyable, requestConfiguration *WindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

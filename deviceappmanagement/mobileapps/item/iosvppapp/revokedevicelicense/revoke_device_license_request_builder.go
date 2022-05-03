package revokedevicelicense

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RevokeDeviceLicenseRequestBuilder provides operations to call the revokeDeviceLicense method.
type RevokeDeviceLicenseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RevokeDeviceLicenseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RevokeDeviceLicenseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRevokeDeviceLicenseRequestBuilderInternal instantiates a new RevokeDeviceLicenseRequestBuilder and sets the default values.
func NewRevokeDeviceLicenseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RevokeDeviceLicenseRequestBuilder) {
    m := &RevokeDeviceLicenseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/microsoft.graph.iosVppApp/microsoft.graph.revokeDeviceLicense";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRevokeDeviceLicenseRequestBuilder instantiates a new RevokeDeviceLicenseRequestBuilder and sets the default values.
func NewRevokeDeviceLicenseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RevokeDeviceLicenseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRevokeDeviceLicenseRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation revoke assigned iOS VPP device license for given app.
func (m *RevokeDeviceLicenseRequestBuilder) CreatePostRequestInformation(body RevokeDeviceLicenseRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration revoke assigned iOS VPP device license for given app.
func (m *RevokeDeviceLicenseRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body RevokeDeviceLicenseRequestBodyable, requestConfiguration *RevokeDeviceLicenseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post revoke assigned iOS VPP device license for given app.
func (m *RevokeDeviceLicenseRequestBuilder) Post(body RevokeDeviceLicenseRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler revoke assigned iOS VPP device license for given app.
func (m *RevokeDeviceLicenseRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body RevokeDeviceLicenseRequestBodyable, requestConfiguration *RevokeDeviceLicenseRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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

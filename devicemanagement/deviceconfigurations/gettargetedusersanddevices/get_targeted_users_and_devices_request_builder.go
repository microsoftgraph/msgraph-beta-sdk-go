package gettargetedusersanddevices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetTargetedUsersAndDevicesRequestBuilder provides operations to call the getTargetedUsersAndDevices method.
type GetTargetedUsersAndDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetTargetedUsersAndDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetTargetedUsersAndDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetTargetedUsersAndDevicesRequestBuilderInternal instantiates a new GetTargetedUsersAndDevicesRequestBuilder and sets the default values.
func NewGetTargetedUsersAndDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTargetedUsersAndDevicesRequestBuilder) {
    m := &GetTargetedUsersAndDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/microsoft.graph.getTargetedUsersAndDevices";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetTargetedUsersAndDevicesRequestBuilder instantiates a new GetTargetedUsersAndDevicesRequestBuilder and sets the default values.
func NewGetTargetedUsersAndDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTargetedUsersAndDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTargetedUsersAndDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getTargetedUsersAndDevices
func (m *GetTargetedUsersAndDevicesRequestBuilder) CreatePostRequestInformation(body GetTargetedUsersAndDevicesRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getTargetedUsersAndDevices
func (m *GetTargetedUsersAndDevicesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetTargetedUsersAndDevicesRequestBodyable, requestConfiguration *GetTargetedUsersAndDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getTargetedUsersAndDevices
func (m *GetTargetedUsersAndDevicesRequestBuilder) Post(body GetTargetedUsersAndDevicesRequestBodyable)(GetTargetedUsersAndDevicesResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getTargetedUsersAndDevices
func (m *GetTargetedUsersAndDevicesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetTargetedUsersAndDevicesRequestBodyable, requestConfiguration *GetTargetedUsersAndDevicesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetTargetedUsersAndDevicesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetTargetedUsersAndDevicesResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetTargetedUsersAndDevicesResponseable), nil
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder provides operations to call the retrieveRemoteHelpSession method.
type DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal instantiates a new RetrieveRemoteHelpSessionWithSessionKeyRequestBuilder and sets the default values.
func NewDeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, sessionKey *string)(*DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    m := &DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/microsoft.graph.retrieveRemoteHelpSession(sessionKey='{sessionKey}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if sessionKey != nil {
        urlTplParams["sessionKey"] = *sessionKey
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder instantiates a new RetrieveRemoteHelpSessionWithSessionKeyRequestBuilder and sets the default values.
func NewDeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function retrieveRemoteHelpSession
func (m *DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function retrieveRemoteHelpSession
func (m *DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RetrieveRemoteHelpSessionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRetrieveRemoteHelpSessionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RetrieveRemoteHelpSessionResponseable), nil
}

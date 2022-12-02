package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder provides operations to call the getRealTimeRemoteConnectionStatus method.
type DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal instantiates a new GetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder and sets the default values.
func NewDeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, cloudPcId *string)(*DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    m := &DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/reports/microsoft.graph.getRealTimeRemoteConnectionStatus(cloudPcId='{cloudPcId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if cloudPcId != nil {
        urlTplParams["cloudPcId"] = *cloudPcId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder instantiates a new GetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder and sets the default values.
func NewDeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getRealTimeRemoteConnectionStatus
func (m *DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getRealTimeRemoteConnectionStatus
func (m *DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}

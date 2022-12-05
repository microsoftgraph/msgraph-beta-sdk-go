package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder provides operations to call the sendCustomNotificationToCompanyPortal method.
type DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal instantiates a new SendCustomNotificationToCompanyPortalRequestBuilder and sets the default values.
func NewDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    m := &DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/microsoft.graph.sendCustomNotificationToCompanyPortal";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder instantiates a new SendCustomNotificationToCompanyPortalRequestBuilder and sets the default values.
func NewDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action sendCustomNotificationToCompanyPortal
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBodyable, requestConfiguration *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action sendCustomNotificationToCompanyPortal
func (m *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBodyable, requestConfiguration *DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}

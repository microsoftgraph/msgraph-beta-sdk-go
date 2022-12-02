package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder provides operations to call the sendCustomNotificationToCompanyPortal method.
type MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal instantiates a new SendCustomNotificationToCompanyPortalRequestBuilder and sets the default values.
func NewMeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    m := &MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}/microsoft.graph.sendCustomNotificationToCompanyPortal";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder instantiates a new SendCustomNotificationToCompanyPortalRequestBuilder and sets the default values.
func NewMeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action sendCustomNotificationToCompanyPortal
func (m *MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeManagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBodyable, requestConfiguration *MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeManagedDevicesItemSendCustomNotificationToCompanyPortalPostRequestBodyable, requestConfiguration *MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderPostRequestConfiguration)(error) {
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

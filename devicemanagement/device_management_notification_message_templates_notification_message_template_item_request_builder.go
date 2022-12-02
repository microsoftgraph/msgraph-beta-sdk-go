package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
type DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderGetQueryParameters the Notification Message Templates.
type DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderGetQueryParameters
}
// DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderInternal instantiates a new NotificationMessageTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) {
    m := &DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/notificationMessageTemplates/{notificationMessageTemplate%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder instantiates a new NotificationMessageTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property notificationMessageTemplates for deviceManagement
func (m *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the Notification Message Templates.
func (m *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property notificationMessageTemplates in deviceManagement
func (m *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NotificationMessageTemplateable, requestConfiguration *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property notificationMessageTemplates for deviceManagement
func (m *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the Notification Message Templates.
func (m *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NotificationMessageTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateNotificationMessageTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NotificationMessageTemplateable), nil
}
// LocalizedNotificationMessages provides operations to manage the localizedNotificationMessages property of the microsoft.graph.notificationMessageTemplate entity.
func (m *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) LocalizedNotificationMessages()(*DeviceManagementNotificationMessageTemplatesItemLocalizedNotificationMessagesRequestBuilder) {
    return NewDeviceManagementNotificationMessageTemplatesItemLocalizedNotificationMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizedNotificationMessagesById provides operations to manage the localizedNotificationMessages property of the microsoft.graph.notificationMessageTemplate entity.
func (m *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) LocalizedNotificationMessagesById(id string)(*DeviceManagementNotificationMessageTemplatesItemLocalizedNotificationMessagesLocalizedNotificationMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["localizedNotificationMessage%2Did"] = id
    }
    return NewDeviceManagementNotificationMessageTemplatesItemLocalizedNotificationMessagesLocalizedNotificationMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property notificationMessageTemplates in deviceManagement
func (m *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NotificationMessageTemplateable, requestConfiguration *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NotificationMessageTemplateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateNotificationMessageTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NotificationMessageTemplateable), nil
}
// SendTestMessage provides operations to call the sendTestMessage method.
func (m *DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) SendTestMessage()(*DeviceManagementNotificationMessageTemplatesItemSendTestMessageRequestBuilder) {
    return NewDeviceManagementNotificationMessageTemplatesItemSendTestMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

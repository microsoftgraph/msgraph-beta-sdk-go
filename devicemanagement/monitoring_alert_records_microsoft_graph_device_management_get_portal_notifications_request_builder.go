package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder provides operations to call the getPortalNotifications method.
type MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetQueryParameters invoke function getPortalNotifications
type MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetQueryParameters
}
// NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderInternal instantiates a new MicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder and sets the default values.
func NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) {
    m := &MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/monitoring/alertRecords/microsoft.graph.deviceManagement.getPortalNotifications(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder instantiates a new MicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder and sets the default values.
func NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getPortalNotifications
// Deprecated: This method is obsolete. Use GetAsGetPortalNotificationsGetResponse instead.
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) Get(ctx context.Context, requestConfiguration *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration)(MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponseable), nil
}
// GetAsGetPortalNotificationsGetResponse invoke function getPortalNotifications
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) GetAsGetPortalNotificationsGetResponse(ctx context.Context, requestConfiguration *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration)(MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsGetResponseable), nil
}
// ToGetRequestInformation invoke function getPortalNotifications
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) WithUrl(rawUrl string)(*MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) {
    return NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

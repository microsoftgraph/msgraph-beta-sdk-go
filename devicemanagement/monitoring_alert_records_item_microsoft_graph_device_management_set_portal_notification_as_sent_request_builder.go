package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder provides operations to call the setPortalNotificationAsSent method.
type MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderInternal instantiates a new MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder and sets the default values.
func NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) {
    m := &MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/monitoring/alertRecords/{alertRecord%2Did}/microsoft.graph.deviceManagement.setPortalNotificationAsSent", pathParameters),
    }
    return m
}
// NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder instantiates a new MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder and sets the default values.
func NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the status of the notification associated with the specified alertRecord on the Microsoft EndPoint Manager admin center as sent, by setting the isPortalNotificationSent property of the portal notification to true.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/devicemanagement-alertrecord-setportalnotificationassent?view=graph-rest-beta
func (m *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) Post(ctx context.Context, requestConfiguration *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation set the status of the notification associated with the specified alertRecord on the Microsoft EndPoint Manager admin center as sent, by setting the isPortalNotificationSent property of the portal notification to true.
// returns a *RequestInformation when successful
func (m *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder when successful
func (m *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) WithUrl(rawUrl string)(*MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) {
    return NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

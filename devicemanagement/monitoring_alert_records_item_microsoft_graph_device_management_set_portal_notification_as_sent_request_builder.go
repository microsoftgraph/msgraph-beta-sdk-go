package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder provides operations to call the setPortalNotificationAsSent method.
type MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderInternal instantiates a new MicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder and sets the default values.
func NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) {
    m := &MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/monitoring/alertRecords/{alertRecord%2Did}/microsoft.graph.deviceManagement.setPortalNotificationAsSent";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder instantiates a new MicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder and sets the default values.
func NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the status of the notification associated with the specified alertRecord on the Microsoft EndPoint Manager admin center as sent, by setting the **isPortalNotificationSent** property of the portal notification to `true`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/devicemanagement-alertrecord-setportalnotificationassent?view=graph-rest-1.0
func (m *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) Post(ctx context.Context, requestConfiguration *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation set the status of the notification associated with the specified alertRecord on the Microsoft EndPoint Manager admin center as sent, by setting the **isPortalNotificationSent** property of the portal notification to `true`.
func (m *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MonitoringAlertRecordsItemMicrosoftGraphDeviceManagementSetPortalNotificationAsSentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

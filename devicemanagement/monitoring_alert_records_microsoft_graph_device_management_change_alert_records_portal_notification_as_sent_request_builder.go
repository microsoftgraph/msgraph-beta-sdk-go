package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder provides operations to call the changeAlertRecordsPortalNotificationAsSent method.
type MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilderInternal instantiates a new MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder and sets the default values.
func NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder) {
    m := &MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/monitoring/alertRecords/microsoft.graph.deviceManagement.changeAlertRecordsPortalNotificationAsSent", pathParameters),
    }
    return m
}
// NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder instantiates a new MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder and sets the default values.
func NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the isPortalNotificationSent property of all portal notification resources associated with the specified alertRecord to true, marking them as sent. A maximum of 100 alertRecord IDs can be received at one time, and a maximum of 100 portal notification resources can be changed in the isPortalNotificationSent property status.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder) Post(ctx context.Context, body MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentChangeAlertRecordsPortalNotificationAsSentPostRequestBodyable, requestConfiguration *MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation set the isPortalNotificationSent property of all portal notification resources associated with the specified alertRecord to true, marking them as sent. A maximum of 100 alertRecord IDs can be received at one time, and a maximum of 100 portal notification resources can be changed in the isPortalNotificationSent property status.
// returns a *RequestInformation when successful
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder) ToPostRequestInformation(ctx context.Context, body MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentChangeAlertRecordsPortalNotificationAsSentPostRequestBodyable, requestConfiguration *MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder when successful
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder) WithUrl(rawUrl string)(*MonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder) {
    return NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

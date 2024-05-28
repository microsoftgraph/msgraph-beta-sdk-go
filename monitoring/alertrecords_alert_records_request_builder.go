package monitoring

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf "github.com/microsoftgraph/msgraph-beta-sdk-go/models/devicemanagement"
)

// AlertrecordsAlertRecordsRequestBuilder provides operations to manage the alertRecords property of the microsoft.graph.deviceManagement.monitoring entity.
type AlertrecordsAlertRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AlertrecordsAlertRecordsRequestBuilderGetQueryParameters the collection of records of alert events.
type AlertrecordsAlertRecordsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AlertrecordsAlertRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AlertrecordsAlertRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AlertrecordsAlertRecordsRequestBuilderGetQueryParameters
}
// AlertrecordsAlertRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AlertrecordsAlertRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAlertRecordId provides operations to manage the alertRecords property of the microsoft.graph.deviceManagement.monitoring entity.
// returns a *AlertrecordsAlertRecordItemRequestBuilder when successful
func (m *AlertrecordsAlertRecordsRequestBuilder) ByAlertRecordId(alertRecordId string)(*AlertrecordsAlertRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if alertRecordId != "" {
        urlTplParams["alertRecord%2Did"] = alertRecordId
    }
    return NewAlertrecordsAlertRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAlertrecordsAlertRecordsRequestBuilderInternal instantiates a new AlertrecordsAlertRecordsRequestBuilder and sets the default values.
func NewAlertrecordsAlertRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertrecordsAlertRecordsRequestBuilder) {
    m := &AlertrecordsAlertRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/monitoring/alertRecords{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAlertrecordsAlertRecordsRequestBuilder instantiates a new AlertrecordsAlertRecordsRequestBuilder and sets the default values.
func NewAlertrecordsAlertRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertrecordsAlertRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAlertrecordsAlertRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AlertrecordsCountRequestBuilder when successful
func (m *AlertrecordsAlertRecordsRequestBuilder) Count()(*AlertrecordsCountRequestBuilder) {
    return NewAlertrecordsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of records of alert events.
// returns a AlertRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AlertrecordsAlertRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *AlertrecordsAlertRecordsRequestBuilderGetRequestConfiguration)(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.CreateAlertRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordCollectionResponseable), nil
}
// MicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSent provides operations to call the changeAlertRecordsPortalNotificationAsSent method.
// returns a *AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder when successful
func (m *AlertrecordsAlertRecordsRequestBuilder) MicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSent()(*AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilder) {
    return NewAlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentMicrosoftGraphDeviceManagementChangeAlertRecordsPortalNotificationAsSentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphDeviceManagementGetPortalNotifications provides operations to call the getPortalNotifications method.
// returns a *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder when successful
func (m *AlertrecordsAlertRecordsRequestBuilder) MicrosoftGraphDeviceManagementGetPortalNotifications()(*AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) {
    return NewAlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to alertRecords for monitoring
// returns a AlertRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AlertrecordsAlertRecordsRequestBuilder) Post(ctx context.Context, body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordable, requestConfiguration *AlertrecordsAlertRecordsRequestBuilderPostRequestConfiguration)(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.CreateAlertRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordable), nil
}
// ToGetRequestInformation the collection of records of alert events.
// returns a *RequestInformation when successful
func (m *AlertrecordsAlertRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AlertrecordsAlertRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to alertRecords for monitoring
// returns a *RequestInformation when successful
func (m *AlertrecordsAlertRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordable, requestConfiguration *AlertrecordsAlertRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AlertrecordsAlertRecordsRequestBuilder when successful
func (m *AlertrecordsAlertRecordsRequestBuilder) WithUrl(rawUrl string)(*AlertrecordsAlertRecordsRequestBuilder) {
    return NewAlertrecordsAlertRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

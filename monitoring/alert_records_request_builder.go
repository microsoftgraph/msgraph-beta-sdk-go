package monitoring

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf "github.com/microsoftgraph/msgraph-beta-sdk-go/models/devicemanagement"
)

// AlertRecordsRequestBuilder provides operations to manage the alertRecords property of the microsoft.graph.deviceManagement.monitoring entity.
type AlertRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AlertRecordsRequestBuilderGetQueryParameters get a list of the alertRecord objects and their properties.
type AlertRecordsRequestBuilderGetQueryParameters struct {
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
// AlertRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AlertRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AlertRecordsRequestBuilderGetQueryParameters
}
// AlertRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AlertRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAlertRecordId provides operations to manage the alertRecords property of the microsoft.graph.deviceManagement.monitoring entity.
// returns a *AlertRecordsAlertRecordItemRequestBuilder when successful
func (m *AlertRecordsRequestBuilder) ByAlertRecordId(alertRecordId string)(*AlertRecordsAlertRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if alertRecordId != "" {
        urlTplParams["alertRecord%2Did"] = alertRecordId
    }
    return NewAlertRecordsAlertRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAlertRecordsRequestBuilderInternal instantiates a new AlertRecordsRequestBuilder and sets the default values.
func NewAlertRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertRecordsRequestBuilder) {
    m := &AlertRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/monitoring/alertRecords{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAlertRecordsRequestBuilder instantiates a new AlertRecordsRequestBuilder and sets the default values.
func NewAlertRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAlertRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AlertRecordsCountRequestBuilder when successful
func (m *AlertRecordsRequestBuilder) Count()(*AlertRecordsCountRequestBuilder) {
    return NewAlertRecordsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the alertRecord objects and their properties.
// returns a AlertRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/devicemanagement-alertrecord-list?view=graph-rest-1.0
func (m *AlertRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *AlertRecordsRequestBuilderGetRequestConfiguration)(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordCollectionResponseable, error) {
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
// MicrosoftGraphDeviceManagementGetPortalNotifications provides operations to call the getPortalNotifications method.
// returns a *AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder when successful
func (m *AlertRecordsRequestBuilder) MicrosoftGraphDeviceManagementGetPortalNotifications()(*AlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) {
    return NewAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to alertRecords for monitoring
// returns a AlertRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AlertRecordsRequestBuilder) Post(ctx context.Context, body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordable, requestConfiguration *AlertRecordsRequestBuilderPostRequestConfiguration)(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordable, error) {
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
// ToGetRequestInformation get a list of the alertRecord objects and their properties.
// returns a *RequestInformation when successful
func (m *AlertRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AlertRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AlertRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.AlertRecordable, requestConfiguration *AlertRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/monitoring/alertRecords", m.BaseRequestBuilder.PathParameters)
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
// returns a *AlertRecordsRequestBuilder when successful
func (m *AlertRecordsRequestBuilder) WithUrl(rawUrl string)(*AlertRecordsRequestBuilder) {
    return NewAlertRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

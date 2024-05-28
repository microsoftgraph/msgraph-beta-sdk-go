package monitoring

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder provides operations to call the getPortalNotifications method.
type AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetQueryParameters invoke function getPortalNotifications
type AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetQueryParameters struct {
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
// AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetQueryParameters
}
// NewAlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderInternal instantiates a new AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder and sets the default values.
func NewAlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) {
    m := &AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/monitoring/alertRecords/microsoft.graph.deviceManagement.getPortalNotifications(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder instantiates a new AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder and sets the default values.
func NewAlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getPortalNotifications
// Deprecated: This method is obsolete. Use GetAsGetPortalNotificationsGetResponse instead.
// returns a AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsGetPortalNotificationsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) Get(ctx context.Context, requestConfiguration *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration)(AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsGetPortalNotificationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsGetPortalNotificationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsGetPortalNotificationsResponseable), nil
}
// GetAsGetPortalNotificationsGetResponse invoke function getPortalNotifications
// returns a AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsGetPortalNotificationsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) GetAsGetPortalNotificationsGetResponse(ctx context.Context, requestConfiguration *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration)(AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsGetPortalNotificationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsGetPortalNotificationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsGetPortalNotificationsGetResponseable), nil
}
// ToGetRequestInformation invoke function getPortalNotifications
// returns a *RequestInformation when successful
func (m *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder when successful
func (m *AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) WithUrl(rawUrl string)(*AlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder) {
    return NewAlertrecordsMicrosoftgraphdevicemanagementgetportalnotificationsMicrosoftGraphDeviceManagementGetPortalNotificationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

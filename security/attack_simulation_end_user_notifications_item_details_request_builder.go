package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AttackSimulationEndUserNotificationsItemDetailsRequestBuilder provides operations to manage the details property of the microsoft.graph.endUserNotification entity.
type AttackSimulationEndUserNotificationsItemDetailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttackSimulationEndUserNotificationsItemDetailsRequestBuilderGetQueryParameters get details from security
type AttackSimulationEndUserNotificationsItemDetailsRequestBuilderGetQueryParameters struct {
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
// AttackSimulationEndUserNotificationsItemDetailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationEndUserNotificationsItemDetailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttackSimulationEndUserNotificationsItemDetailsRequestBuilderGetQueryParameters
}
// AttackSimulationEndUserNotificationsItemDetailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationEndUserNotificationsItemDetailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEndUserNotificationDetailId provides operations to manage the details property of the microsoft.graph.endUserNotification entity.
// returns a *AttackSimulationEndUserNotificationsItemDetailsEndUserNotificationDetailItemRequestBuilder when successful
func (m *AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) ByEndUserNotificationDetailId(endUserNotificationDetailId string)(*AttackSimulationEndUserNotificationsItemDetailsEndUserNotificationDetailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if endUserNotificationDetailId != "" {
        urlTplParams["endUserNotificationDetail%2Did"] = endUserNotificationDetailId
    }
    return NewAttackSimulationEndUserNotificationsItemDetailsEndUserNotificationDetailItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAttackSimulationEndUserNotificationsItemDetailsRequestBuilderInternal instantiates a new AttackSimulationEndUserNotificationsItemDetailsRequestBuilder and sets the default values.
func NewAttackSimulationEndUserNotificationsItemDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) {
    m := &AttackSimulationEndUserNotificationsItemDetailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/endUserNotifications/{endUserNotification%2Did}/details{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAttackSimulationEndUserNotificationsItemDetailsRequestBuilder instantiates a new AttackSimulationEndUserNotificationsItemDetailsRequestBuilder and sets the default values.
func NewAttackSimulationEndUserNotificationsItemDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttackSimulationEndUserNotificationsItemDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AttackSimulationEndUserNotificationsItemDetailsCountRequestBuilder when successful
func (m *AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) Count()(*AttackSimulationEndUserNotificationsItemDetailsCountRequestBuilder) {
    return NewAttackSimulationEndUserNotificationsItemDetailsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get details from security
// returns a EndUserNotificationDetailCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) Get(ctx context.Context, requestConfiguration *AttackSimulationEndUserNotificationsItemDetailsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationDetailCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEndUserNotificationDetailCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationDetailCollectionResponseable), nil
}
// Post create new navigation property to details for security
// returns a EndUserNotificationDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationDetailable, requestConfiguration *AttackSimulationEndUserNotificationsItemDetailsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationDetailable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEndUserNotificationDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationDetailable), nil
}
// ToGetRequestInformation get details from security
// returns a *RequestInformation when successful
func (m *AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttackSimulationEndUserNotificationsItemDetailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to details for security
// returns a *RequestInformation when successful
func (m *AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EndUserNotificationDetailable, requestConfiguration *AttackSimulationEndUserNotificationsItemDetailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AttackSimulationEndUserNotificationsItemDetailsRequestBuilder when successful
func (m *AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) WithUrl(rawUrl string)(*AttackSimulationEndUserNotificationsItemDetailsRequestBuilder) {
    return NewAttackSimulationEndUserNotificationsItemDetailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

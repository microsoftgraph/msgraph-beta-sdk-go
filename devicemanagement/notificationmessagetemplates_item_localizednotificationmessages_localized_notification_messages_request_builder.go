package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder provides operations to manage the localizedNotificationMessages property of the microsoft.graph.notificationMessageTemplate entity.
type NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderGetQueryParameters the list of localized messages for this Notification Message Template.
type NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderGetQueryParameters struct {
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
// NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderGetQueryParameters
}
// NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByLocalizedNotificationMessageId provides operations to manage the localizedNotificationMessages property of the microsoft.graph.notificationMessageTemplate entity.
// returns a *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessageItemRequestBuilder when successful
func (m *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) ByLocalizedNotificationMessageId(localizedNotificationMessageId string)(*NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if localizedNotificationMessageId != "" {
        urlTplParams["localizedNotificationMessage%2Did"] = localizedNotificationMessageId
    }
    return NewNotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewNotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderInternal instantiates a new NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder and sets the default values.
func NewNotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) {
    m := &NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/notificationMessageTemplates/{notificationMessageTemplate%2Did}/localizedNotificationMessages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewNotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder instantiates a new NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder and sets the default values.
func NewNotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *NotificationmessagetemplatesItemLocalizednotificationmessagesCountRequestBuilder when successful
func (m *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) Count()(*NotificationmessagetemplatesItemLocalizednotificationmessagesCountRequestBuilder) {
    return NewNotificationmessagetemplatesItemLocalizednotificationmessagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of localized messages for this Notification Message Template.
// returns a LocalizedNotificationMessageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocalizedNotificationMessageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLocalizedNotificationMessageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocalizedNotificationMessageCollectionResponseable), nil
}
// Post create new navigation property to localizedNotificationMessages for deviceManagement
// returns a LocalizedNotificationMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocalizedNotificationMessageable, requestConfiguration *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocalizedNotificationMessageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLocalizedNotificationMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocalizedNotificationMessageable), nil
}
// ToGetRequestInformation the list of localized messages for this Notification Message Template.
// returns a *RequestInformation when successful
func (m *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to localizedNotificationMessages for deviceManagement
// returns a *RequestInformation when successful
func (m *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocalizedNotificationMessageable, requestConfiguration *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder when successful
func (m *NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) WithUrl(rawUrl string)(*NotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder) {
    return NewNotificationmessagetemplatesItemLocalizednotificationmessagesLocalizedNotificationMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

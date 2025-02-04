package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder provides operations to manage the events property of the microsoft.graph.messageRecipient entity.
type ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderGetQueryParameters get events from admin
type ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderGetQueryParameters struct {
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
// ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderGetQueryParameters
}
// ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMessageEventId provides operations to manage the events property of the microsoft.graph.messageRecipient entity.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *ExchangeMessageTracesItemRecipientsItemEventsMessageEventItemRequestBuilder when successful
func (m *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) ByMessageEventId(messageEventId string)(*ExchangeMessageTracesItemRecipientsItemEventsMessageEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if messageEventId != "" {
        urlTplParams["messageEvent%2Did"] = messageEventId
    }
    return NewExchangeMessageTracesItemRecipientsItemEventsMessageEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewExchangeMessageTracesItemRecipientsItemEventsRequestBuilderInternal instantiates a new ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder and sets the default values.
func NewExchangeMessageTracesItemRecipientsItemEventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) {
    m := &ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/messageTraces/{messageTrace%2Did}/recipients/{messageRecipient%2Did}/events{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewExchangeMessageTracesItemRecipientsItemEventsRequestBuilder instantiates a new ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder and sets the default values.
func NewExchangeMessageTracesItemRecipientsItemEventsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMessageTracesItemRecipientsItemEventsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ExchangeMessageTracesItemRecipientsItemEventsCountRequestBuilder when successful
func (m *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) Count()(*ExchangeMessageTracesItemRecipientsItemEventsCountRequestBuilder) {
    return NewExchangeMessageTracesItemRecipientsItemEventsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get events from admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a MessageEventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageEventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageEventCollectionResponseable), nil
}
// Post create new navigation property to events for admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a MessageEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageEventable, requestConfiguration *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageEventable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageEventable), nil
}
// ToGetRequestInformation get events from admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to events for admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageEventable, requestConfiguration *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder when successful
func (m *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) WithUrl(rawUrl string)(*ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) {
    return NewExchangeMessageTracesItemRecipientsItemEventsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

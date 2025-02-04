package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder provides operations to manage the recipients property of the microsoft.graph.messageTrace entity.
type ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderGetQueryParameters get recipients from admin
type ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderGetQueryParameters
}
// ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderInternal instantiates a new ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder and sets the default values.
func NewExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) {
    m := &ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/messageTraces/{messageTrace%2Did}/recipients/{messageRecipient%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder instantiates a new ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder and sets the default values.
func NewExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property recipients for admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Events provides operations to manage the events property of the microsoft.graph.messageRecipient entity.
// returns a *ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder when successful
func (m *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) Events()(*ExchangeMessageTracesItemRecipientsItemEventsRequestBuilder) {
    return NewExchangeMessageTracesItemRecipientsItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get recipients from admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a MessageRecipientable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageRecipientable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageRecipientFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageRecipientable), nil
}
// Patch update the navigation property recipients in admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a MessageRecipientable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageRecipientable, requestConfiguration *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageRecipientable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageRecipientFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageRecipientable), nil
}
// ToDeleteRequestInformation delete navigation property recipients for admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get recipients from admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property recipients in admin
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MessageRecipientable, requestConfiguration *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder when successful
func (m *ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) WithUrl(rawUrl string)(*ExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder) {
    return NewExchangeMessageTracesItemRecipientsMessageRecipientItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

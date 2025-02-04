package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMailboxesItemExportItemsRequestBuilder provides operations to call the exportItems method.
type ExchangeMailboxesItemExportItemsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMailboxesItemExportItemsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMailboxesItemExportItemsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExchangeMailboxesItemExportItemsRequestBuilderInternal instantiates a new ExchangeMailboxesItemExportItemsRequestBuilder and sets the default values.
func NewExchangeMailboxesItemExportItemsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemExportItemsRequestBuilder) {
    m := &ExchangeMailboxesItemExportItemsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/mailboxes/{mailbox%2Did}/exportItems", pathParameters),
    }
    return m
}
// NewExchangeMailboxesItemExportItemsRequestBuilder instantiates a new ExchangeMailboxesItemExportItemsRequestBuilder and sets the default values.
func NewExchangeMailboxesItemExportItemsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemExportItemsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMailboxesItemExportItemsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post export Exchange mailboxItem objects in full-fidelity FastTransfer stream format for backup purposes. This item format can be restored to the same mailbox or a different one. You can export up to 20 items in a single export request.
// Deprecated: This method is obsolete. Use PostAsExportItemsPostResponse instead.
// returns a ExchangeMailboxesItemExportItemsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailbox-exportitems?view=graph-rest-beta
func (m *ExchangeMailboxesItemExportItemsRequestBuilder) Post(ctx context.Context, body ExchangeMailboxesItemExportItemsPostRequestBodyable, requestConfiguration *ExchangeMailboxesItemExportItemsRequestBuilderPostRequestConfiguration)(ExchangeMailboxesItemExportItemsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExchangeMailboxesItemExportItemsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExchangeMailboxesItemExportItemsResponseable), nil
}
// PostAsExportItemsPostResponse export Exchange mailboxItem objects in full-fidelity FastTransfer stream format for backup purposes. This item format can be restored to the same mailbox or a different one. You can export up to 20 items in a single export request.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a ExchangeMailboxesItemExportItemsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailbox-exportitems?view=graph-rest-beta
func (m *ExchangeMailboxesItemExportItemsRequestBuilder) PostAsExportItemsPostResponse(ctx context.Context, body ExchangeMailboxesItemExportItemsPostRequestBodyable, requestConfiguration *ExchangeMailboxesItemExportItemsRequestBuilderPostRequestConfiguration)(ExchangeMailboxesItemExportItemsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExchangeMailboxesItemExportItemsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExchangeMailboxesItemExportItemsPostResponseable), nil
}
// ToPostRequestInformation export Exchange mailboxItem objects in full-fidelity FastTransfer stream format for backup purposes. This item format can be restored to the same mailbox or a different one. You can export up to 20 items in a single export request.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMailboxesItemExportItemsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ExchangeMailboxesItemExportItemsPostRequestBodyable, requestConfiguration *ExchangeMailboxesItemExportItemsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeMailboxesItemExportItemsRequestBuilder when successful
func (m *ExchangeMailboxesItemExportItemsRequestBuilder) WithUrl(rawUrl string)(*ExchangeMailboxesItemExportItemsRequestBuilder) {
    return NewExchangeMailboxesItemExportItemsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

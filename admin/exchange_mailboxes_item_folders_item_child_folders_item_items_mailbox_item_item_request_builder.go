package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.mailboxFolder entity.
type ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderGetQueryParameters the collection of items in this folder.
type ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderGetQueryParameters
}
// NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderInternal instantiates a new ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder) {
    m := &ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/mailboxes/{mailbox%2Did}/folders/{mailboxFolder%2Did}/childFolders/{mailboxFolder%2Did1}/items/{mailboxItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder instantiates a new ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the collection of items in this folder.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport on 2021-08-19 and will be removed 2021-11-15
// returns a MailboxItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxItemable), nil
}
// ToGetRequestInformation the collection of items in this folder.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport on 2021-08-19 and will be removed 2021-11-15
// returns a *RequestInformation when successful
func (m *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport on 2021-08-19 and will be removed 2021-11-15
// returns a *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder when successful
func (m *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder) WithUrl(rawUrl string)(*ExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder) {
    return NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsMailboxItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

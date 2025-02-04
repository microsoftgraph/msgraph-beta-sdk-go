package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.mailboxFolder entity.
type ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderGetQueryParameters the collection of child folders in this folder.
type ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderGetQueryParameters
}
// NewExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderInternal instantiates a new ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder) {
    m := &ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/mailboxes/{mailbox%2Did}/folders/{mailboxFolder%2Did}/childFolders/{mailboxFolder%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder instantiates a new ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the collection of child folders in this folder.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a MailboxFolderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxFolderable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxFolderable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.mailboxFolder entity.
// returns a *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsRequestBuilder when successful
func (m *ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder) Items()(*ExchangeMailboxesItemFoldersItemChildFoldersItemItemsRequestBuilder) {
    return NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the collection of child folders in this folder.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder when successful
func (m *ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder) WithUrl(rawUrl string)(*ExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder) {
    return NewExchangeMailboxesItemFoldersItemChildFoldersMailboxFolderItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

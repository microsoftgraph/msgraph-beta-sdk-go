package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder provides operations to manage the folders property of the microsoft.graph.mailbox entity.
type ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderGetQueryParameters read the properties and relationships of a mailboxFolder object.
type ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderGetQueryParameters
}
// ChildFolders provides operations to manage the childFolders property of the microsoft.graph.mailboxFolder entity.
// returns a *ExchangeMailboxesItemFoldersItemChildFoldersRequestBuilder when successful
func (m *ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder) ChildFolders()(*ExchangeMailboxesItemFoldersItemChildFoldersRequestBuilder) {
    return NewExchangeMailboxesItemFoldersItemChildFoldersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderInternal instantiates a new ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder) {
    m := &ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/mailboxes/{mailbox%2Did}/folders/{mailboxFolder%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder instantiates a new ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of a mailboxFolder object.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a MailboxFolderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailboxfolder-get?view=graph-rest-beta
func (m *ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxFolderable, error) {
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
// returns a *ExchangeMailboxesItemFoldersItemItemsRequestBuilder when successful
func (m *ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder) Items()(*ExchangeMailboxesItemFoldersItemItemsRequestBuilder) {
    return NewExchangeMailboxesItemFoldersItemItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read the properties and relationships of a mailboxFolder object.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder when successful
func (m *ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder) WithUrl(rawUrl string)(*ExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder) {
    return NewExchangeMailboxesItemFoldersMailboxFolderItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

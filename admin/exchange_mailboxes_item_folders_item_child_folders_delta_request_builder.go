package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder provides operations to call the delta method.
type ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderGetQueryParameters get a set of mailboxFolder objects that have been added, deleted, or removed from the user's mailbox. A delta function call for folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the folders. This approach allows you to maintain and synchronize a local store of a user's mail folders without having to fetch all the folders of that mailbox from the server every time.
type ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderGetQueryParameters struct {
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
// ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderGetQueryParameters
}
// NewExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderInternal instantiates a new ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder) {
    m := &ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/mailboxes/{mailbox%2Did}/folders/{mailboxFolder%2Did}/childFolders/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder instantiates a new ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of mailboxFolder objects that have been added, deleted, or removed from the user's mailbox. A delta function call for folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the folders. This approach allows you to maintain and synchronize a local store of a user's mail folders without having to fetch all the folders of that mailbox from the server every time.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ExchangeMailboxesItemFoldersItemChildFoldersDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailboxfolder-delta?view=graph-rest-beta
func (m *ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderGetRequestConfiguration)(ExchangeMailboxesItemFoldersItemChildFoldersDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExchangeMailboxesItemFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExchangeMailboxesItemFoldersItemChildFoldersDeltaResponseable), nil
}
// GetAsDeltaGetResponse get a set of mailboxFolder objects that have been added, deleted, or removed from the user's mailbox. A delta function call for folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the folders. This approach allows you to maintain and synchronize a local store of a user's mail folders without having to fetch all the folders of that mailbox from the server every time.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a ExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mailboxfolder-delta?view=graph-rest-beta
func (m *ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderGetRequestConfiguration)(ExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExchangeMailboxesItemFoldersItemChildFoldersDeltaGetResponseable), nil
}
// ToGetRequestInformation get a set of mailboxFolder objects that have been added, deleted, or removed from the user's mailbox. A delta function call for folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the folders. This approach allows you to maintain and synchronize a local store of a user's mail folders without having to fetch all the folders of that mailbox from the server every time.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder when successful
func (m *ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder) WithUrl(rawUrl string)(*ExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder) {
    return NewExchangeMailboxesItemFoldersItemChildFoldersDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder provides operations to call the delta method.
type ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderGetQueryParameters get a set of mailboxItem objects that have been added, deleted, or updated in a specified mailboxFolder. A delta function call for items in a folder is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the items in that folder. This approach allows you to maintain and synchronize a local store of a user's mailbox items without having to fetch the entire set of items from the server every time.
type ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderGetQueryParameters struct {
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
// ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderGetQueryParameters
}
// NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderInternal instantiates a new ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder) {
    m := &ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/exchange/mailboxes/{mailbox%2Did}/folders/{mailboxFolder%2Did}/childFolders/{mailboxFolder%2Did1}/items/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder instantiates a new ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder and sets the default values.
func NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of mailboxItem objects that have been added, deleted, or updated in a specified mailboxFolder. A delta function call for items in a folder is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the items in that folder. This approach allows you to maintain and synchronize a local store of a user's mailbox items without having to fetch the entire set of items from the server every time.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderGetRequestConfiguration)(ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaResponseable), nil
}
// GetAsDeltaGetResponse get a set of mailboxItem objects that have been added, deleted, or updated in a specified mailboxFolder. A delta function call for items in a folder is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the items in that folder. This approach allows you to maintain and synchronize a local store of a user's mailbox items without having to fetch the entire set of items from the server every time.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport on 2021-08-19 and will be removed 2021-11-15
// returns a ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderGetRequestConfiguration)(ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaGetResponseable), nil
}
// ToGetRequestInformation get a set of mailboxItem objects that have been added, deleted, or updated in a specified mailboxFolder. A delta function call for items in a folder is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the items in that folder. This approach allows you to maintain and synchronize a local store of a user's mailbox items without having to fetch the entire set of items from the server every time.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport on 2021-08-19 and will be removed 2021-11-15
// returns a *RequestInformation when successful
func (m *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder when successful
func (m *ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder) WithUrl(rawUrl string)(*ExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder) {
    return NewExchangeMailboxesItemFoldersItemChildFoldersItemItemsDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

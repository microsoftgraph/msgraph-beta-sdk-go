package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder provides operations to manage the listItem property of the microsoft.graph.itemActivityOLD entity.
type ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderGetQueryParameters get listItem from groups
type ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderGetQueryParameters
}
// NewItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderInternal instantiates a new ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder) {
    m := &ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/activities/{itemActivityOLD%2Did}/listItem{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder instantiates a new ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get listItem from groups
// returns a ListItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// ToGetRequestInformation get listItem from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder when successful
func (m *ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemActivitiesItemListItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

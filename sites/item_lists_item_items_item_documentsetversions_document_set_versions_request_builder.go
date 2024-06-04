package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
type ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetQueryParameters get a list of the versions of a document set item in a list.
type ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetQueryParameters struct {
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
// ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetQueryParameters
}
// ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDocumentSetVersionId provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
// returns a *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder when successful
func (m *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) ByDocumentSetVersionId(documentSetVersionId string)(*ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if documentSetVersionId != "" {
        urlTplParams["documentSetVersion%2Did"] = documentSetVersionId
    }
    return NewItemListsItemItemsItemDocumentsetversionsDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderInternal instantiates a new ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder and sets the default values.
func NewItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) {
    m := &ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/documentSetVersions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder instantiates a new ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder and sets the default values.
func NewItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemListsItemItemsItemDocumentsetversionsCountRequestBuilder when successful
func (m *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) Count()(*ItemListsItemItemsItemDocumentsetversionsCountRequestBuilder) {
    return NewItemListsItemItemsItemDocumentsetversionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the versions of a document set item in a list.
// returns a DocumentSetVersionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/listitem-list-documentsetversions?view=graph-rest-beta
func (m *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDocumentSetVersionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionCollectionResponseable), nil
}
// Post create a new version of a document set item in a list.
// returns a DocumentSetVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/listitem-post-documentsetversions?view=graph-rest-beta
func (m *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionable, requestConfiguration *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDocumentSetVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionable), nil
}
// ToGetRequestInformation get a list of the versions of a document set item in a list.
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new version of a document set item in a list.
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionable, requestConfiguration *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder when successful
func (m *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) {
    return NewItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

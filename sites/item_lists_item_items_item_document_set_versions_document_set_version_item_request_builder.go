package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
type ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderGetQueryParameters version information for a document set version created by a user.
type ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderGetQueryParameters
}
// ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderInternal instantiates a new ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder and sets the default values.
func NewItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) {
    m := &ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/documentSetVersions/{documentSetVersion%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder instantiates a new ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder and sets the default values.
func NewItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property documentSetVersions for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Fields provides operations to manage the fields property of the microsoft.graph.listItemVersion entity.
// returns a *ItemListsItemItemsItemDocumentSetVersionsItemFieldsRequestBuilder when successful
func (m *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) Fields()(*ItemListsItemItemsItemDocumentSetVersionsItemFieldsRequestBuilder) {
    return NewItemListsItemItemsItemDocumentSetVersionsItemFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get version information for a document set version created by a user.
// returns a DocumentSetVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property documentSetVersions in sites
// returns a DocumentSetVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionable, requestConfiguration *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// Restore provides operations to call the restore method.
// returns a *ItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder when successful
func (m *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) Restore()(*ItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder) {
    return NewItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property documentSetVersions for sites
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation version information for a document set version created by a user.
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property documentSetVersions in sites
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DocumentSetVersionable, requestConfiguration *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder when successful
func (m *ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) {
    return NewItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

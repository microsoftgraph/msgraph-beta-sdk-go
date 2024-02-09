package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.site entity.
type ItemSitesItemOnenoteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemOnenoteRequestBuilderGetQueryParameters get onenote from groups
type ItemSitesItemOnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteRequestBuilderGetQueryParameters
}
// ItemSitesItemOnenoteRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemOnenoteRequestBuilderInternal instantiates a new ItemSitesItemOnenoteRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteRequestBuilder) {
    m := &ItemSitesItemOnenoteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteRequestBuilder instantiates a new ItemSitesItemOnenoteRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property onenote for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get onenote from groups
// returns a Onenoteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable), nil
}
// Notebooks provides operations to manage the notebooks property of the microsoft.graph.onenote entity.
// returns a *ItemSitesItemOnenoteNotebooksRequestBuilder when successful
func (m *ItemSitesItemOnenoteRequestBuilder) Notebooks()(*ItemSitesItemOnenoteNotebooksRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.onenote entity.
// returns a *ItemSitesItemOnenoteOperationsRequestBuilder when successful
func (m *ItemSitesItemOnenoteRequestBuilder) Operations()(*ItemSitesItemOnenoteOperationsRequestBuilder) {
    return NewItemSitesItemOnenoteOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pages provides operations to manage the pages property of the microsoft.graph.onenote entity.
// returns a *ItemSitesItemOnenotePagesRequestBuilder when successful
func (m *ItemSitesItemOnenoteRequestBuilder) Pages()(*ItemSitesItemOnenotePagesRequestBuilder) {
    return NewItemSitesItemOnenotePagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property onenote in groups
// returns a Onenoteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemOnenoteRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, requestConfiguration *ItemSitesItemOnenoteRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable), nil
}
// Resources provides operations to manage the resources property of the microsoft.graph.onenote entity.
// returns a *ItemSitesItemOnenoteResourcesRequestBuilder when successful
func (m *ItemSitesItemOnenoteRequestBuilder) Resources()(*ItemSitesItemOnenoteResourcesRequestBuilder) {
    return NewItemSitesItemOnenoteResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SectionGroups provides operations to manage the sectionGroups property of the microsoft.graph.onenote entity.
// returns a *ItemSitesItemOnenoteSectionGroupsRequestBuilder when successful
func (m *ItemSitesItemOnenoteRequestBuilder) SectionGroups()(*ItemSitesItemOnenoteSectionGroupsRequestBuilder) {
    return NewItemSitesItemOnenoteSectionGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sections provides operations to manage the sections property of the microsoft.graph.onenote entity.
// returns a *ItemSitesItemOnenoteSectionsRequestBuilder when successful
func (m *ItemSitesItemOnenoteRequestBuilder) Sections()(*ItemSitesItemOnenoteSectionsRequestBuilder) {
    return NewItemSitesItemOnenoteSectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property onenote for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get onenote from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property onenote in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemOnenoteRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, requestConfiguration *ItemSitesItemOnenoteRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemSitesItemOnenoteRequestBuilder when successful
func (m *ItemSitesItemOnenoteRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteRequestBuilder) {
    return NewItemSitesItemOnenoteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

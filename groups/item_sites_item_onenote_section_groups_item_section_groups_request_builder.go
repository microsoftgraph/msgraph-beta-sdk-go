package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
type ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderGetQueryParameters retrieve a list of section groups from the specified section group. This API is available in the following national cloud deployments.
type ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderGetQueryParameters
}
// BySectionGroupId1 provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder) BySectionGroupId1(sectionGroupId1 string)(*ItemSitesItemOnenoteSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sectionGroupId1 != "" {
        urlTplParams["sectionGroup%2Did1"] = sectionGroupId1
    }
    return NewItemSitesItemOnenoteSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderInternal instantiates a new SectionGroupsRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder) {
    m := &ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sectionGroups{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder instantiates a new SectionGroupsRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder) Count()(*ItemSitesItemOnenoteSectionGroupsItemSectionGroupsCountRequestBuilder) {
    return NewItemSitesItemOnenoteSectionGroupsItemSectionGroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of section groups from the specified section group. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/sectiongroup-list-sectiongroups?view=graph-rest-1.0
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSectionGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupCollectionResponseable), nil
}
// ToGetRequestInformation retrieve a list of section groups from the specified section group. This API is available in the following national cloud deployments.
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder) {
    return NewItemSitesItemOnenoteSectionGroupsItemSectionGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

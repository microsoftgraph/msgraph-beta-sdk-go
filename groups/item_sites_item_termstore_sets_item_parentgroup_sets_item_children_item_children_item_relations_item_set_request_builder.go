package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder provides operations to manage the set property of the microsoft.graph.termStore.relation entity.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderGetQueryParameters the [set] in which the relation is relevant.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderGetQueryParameters
}
// NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderInternal instantiates a new ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder) {
    m := &ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}/children/{term%2Did1}/relations/{relation%2Did}/set{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder instantiates a new ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder and sets the default values.
func NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the [set] in which the relation is relevant.
// returns a Setable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable), nil
}
// ToGetRequestInformation the [set] in which the relation is relevant.
// returns a *RequestInformation when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder when successful
func (m *ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder) {
    return NewItemSitesItemTermstoreSetsItemParentgroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

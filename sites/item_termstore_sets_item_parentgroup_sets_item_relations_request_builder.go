package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.set entity.
type ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetQueryParameters indicates which terms have been pinned or reused directly under the set.
type ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetQueryParameters struct {
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
// ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetQueryParameters
}
// ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRelationId provides operations to manage the relations property of the microsoft.graph.termStore.set entity.
// returns a *ItemTermstoreSetsItemParentgroupSetsItemRelationsRelationItemRequestBuilder when successful
func (m *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) ByRelationId(relationId string)(*ItemTermstoreSetsItemParentgroupSetsItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if relationId != "" {
        urlTplParams["relation%2Did"] = relationId
    }
    return NewItemTermstoreSetsItemParentgroupSetsItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderInternal instantiates a new ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder and sets the default values.
func NewItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) {
    m := &ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/relations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder instantiates a new ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder and sets the default values.
func NewItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTermstoreSetsItemParentgroupSetsItemRelationsCountRequestBuilder when successful
func (m *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) Count()(*ItemTermstoreSetsItemParentgroupSetsItemRelationsCountRequestBuilder) {
    return NewItemTermstoreSetsItemParentgroupSetsItemRelationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get indicates which terms have been pinned or reused directly under the set.
// returns a RelationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.RelationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateRelationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.RelationCollectionResponseable), nil
}
// Post create new navigation property to relations for sites
// returns a Relationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) Post(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Relationable, requestConfiguration *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderPostRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Relationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateRelationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Relationable), nil
}
// ToGetRequestInformation indicates which terms have been pinned or reused directly under the set.
// returns a *RequestInformation when successful
func (m *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to relations for sites
// returns a *RequestInformation when successful
func (m *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Relationable, requestConfiguration *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder when successful
func (m *ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) WithUrl(rawUrl string)(*ItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder) {
    return NewItemTermstoreSetsItemParentgroupSetsItemRelationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

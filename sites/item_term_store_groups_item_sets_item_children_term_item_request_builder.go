package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.set entity.
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetQueryParameters children terms of set in term [store].
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetQueryParameters
}
// ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.term entity.
// returns a *ItemTermStoreGroupsItemSetsItemChildrenItemChildrenRequestBuilder when successful
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Children()(*ItemTermStoreGroupsItemSetsItemChildrenItemChildrenRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemChildrenItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderInternal instantiates a new ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) {
    m := &ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStore/groups/{group%2Did}/sets/{set%2Did}/children/{term%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder instantiates a new ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property children for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get children terms of set in term [store].
// returns a Termable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable), nil
}
// Patch update the navigation property children in sites
// returns a Termable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Patch(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable), nil
}
// Relations provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
// returns a *ItemTermStoreGroupsItemSetsItemChildrenItemRelationsRequestBuilder when successful
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Relations()(*ItemTermStoreGroupsItemSetsItemChildrenItemRelationsRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemChildrenItemRelationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.term entity.
// returns a *ItemTermStoreGroupsItemSetsItemChildrenItemSetRequestBuilder when successful
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) Set()(*ItemTermStoreGroupsItemSetsItemChildrenItemSetRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemChildrenItemSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property children for sites
// returns a *RequestInformation when successful
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation children terms of set in term [store].
// returns a *RequestInformation when successful
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property children in sites
// returns a *RequestInformation when successful
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder when successful
func (m *ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) WithUrl(rawUrl string)(*ItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemChildrenTermItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

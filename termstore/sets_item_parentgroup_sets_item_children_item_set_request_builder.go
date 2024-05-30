package termstore

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder provides operations to manage the set property of the microsoft.graph.termStore.term entity.
type SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SetsItemParentgroupSetsItemChildrenItemSetRequestBuilderGetQueryParameters the [set] in which the term is created.
type SetsItemParentgroupSetsItemChildrenItemSetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SetsItemParentgroupSetsItemChildrenItemSetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetsItemParentgroupSetsItemChildrenItemSetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SetsItemParentgroupSetsItemChildrenItemSetRequestBuilderGetQueryParameters
}
// NewSetsItemParentgroupSetsItemChildrenItemSetRequestBuilderInternal instantiates a new SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder and sets the default values.
func NewSetsItemParentgroupSetsItemChildrenItemSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder) {
    m := &SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}/set{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSetsItemParentgroupSetsItemChildrenItemSetRequestBuilder instantiates a new SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder and sets the default values.
func NewSetsItemParentgroupSetsItemChildrenItemSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetsItemParentgroupSetsItemChildrenItemSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the [set] in which the term is created.
// returns a Setable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder) Get(ctx context.Context, requestConfiguration *SetsItemParentgroupSetsItemChildrenItemSetRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, error) {
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
// ToGetRequestInformation the [set] in which the term is created.
// returns a *RequestInformation when successful
func (m *SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SetsItemParentgroupSetsItemChildrenItemSetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder when successful
func (m *SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder) WithUrl(rawUrl string)(*SetsItemParentgroupSetsItemChildrenItemSetRequestBuilder) {
    return NewSetsItemParentgroupSetsItemChildrenItemSetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

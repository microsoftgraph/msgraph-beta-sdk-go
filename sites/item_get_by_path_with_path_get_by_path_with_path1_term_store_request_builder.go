package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder provides operations to manage the termStore property of the microsoft.graph.site entity.
type ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderGetQueryParameters read the properties and relationships of a store object.
type ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderGetQueryParameters
}
// ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderInternal instantiates a new ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder and sets the default values.
func NewItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) {
    m := &ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/termStore{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder instantiates a new ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder and sets the default values.
func NewItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property termStore for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a store object.
// returns a Storeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/termstore-store-get?view=graph-rest-1.0
func (m *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Storeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateStoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Storeable), nil
}
// Patch update the navigation property termStore in sites
// returns a Storeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) Patch(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Storeable, requestConfiguration *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderPatchRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Storeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateStoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Storeable), nil
}
// ToDeleteRequestInformation delete navigation property termStore for sites
// returns a *RequestInformation when successful
func (m *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/termStore", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a store object.
// returns a *RequestInformation when successful
func (m *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property termStore in sites
// returns a *RequestInformation when successful
func (m *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Storeable, requestConfiguration *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/termStore", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder when successful
func (m *ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) WithUrl(rawUrl string)(*ItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) {
    return NewItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

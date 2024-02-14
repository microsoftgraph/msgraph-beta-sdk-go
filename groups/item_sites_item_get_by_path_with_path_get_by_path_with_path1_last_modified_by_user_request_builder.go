package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
type ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderGetQueryParameters get lastModifiedByUser from groups
type ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderGetQueryParameters
}
// NewItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderInternal instantiates a new ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder and sets the default values.
func NewItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder) {
    m := &ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/lastModifiedByUser{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder instantiates a new ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder and sets the default values.
func NewItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get lastModifiedByUser from groups
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// ToGetRequestInformation get lastModifiedByUser from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder when successful
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

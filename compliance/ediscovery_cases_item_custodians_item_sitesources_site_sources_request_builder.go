package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder provides operations to manage the siteSources property of the microsoft.graph.ediscovery.custodian entity.
type EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetQueryParameters get a list of siteSource objects and their properties.
type EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetQueryParameters struct {
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
// EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetQueryParameters
}
// EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySiteSourceId provides operations to manage the siteSources property of the microsoft.graph.ediscovery.custodian entity.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourceItemRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) BySiteSourceId(siteSourceId string)(*EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if siteSourceId != "" {
        urlTplParams["siteSource%2Did"] = siteSourceId
    }
    return NewEdiscoveryCasesItemCustodiansItemSitesourcesSiteSourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderInternal instantiates a new EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) {
    m := &EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians/{custodian%2Did}/siteSources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder instantiates a new EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EdiscoveryCasesItemCustodiansItemSitesourcesCountRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) Count()(*EdiscoveryCasesItemCustodiansItemSitesourcesCountRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemSitesourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of siteSource objects and their properties.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a SiteSourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-custodian-list-sitesources?view=graph-rest-beta
func (m *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SiteSourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSiteSourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SiteSourceCollectionResponseable), nil
}
// Post create a new custodian siteSource object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a SiteSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-custodian-post-sitesources?view=graph-rest-beta
func (m *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) Post(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SiteSourceable, requestConfiguration *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderPostRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SiteSourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSiteSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SiteSourceable), nil
}
// ToGetRequestInformation get a list of siteSource objects and their properties.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new custodian siteSource object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SiteSourceable, requestConfiguration *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder provides operations to manage the siteSources property of the microsoft.graph.security.ediscoveryCustodian entity.
type CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetQueryParameters get a list of the siteSource objects associated with an ediscoveryCustodian.
type CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySiteSourceId provides operations to manage the siteSources property of the microsoft.graph.security.ediscoveryCustodian entity.
// returns a *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourceItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) BySiteSourceId(siteSourceId string)(*CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if siteSourceId != "" {
        urlTplParams["siteSource%2Did"] = siteSourceId
    }
    return NewCasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) {
    m := &CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/siteSources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder instantiates a new CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CasesEdiscoverycasesItemCustodiansItemSitesourcesCountRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) Count()(*CasesEdiscoverycasesItemCustodiansItemSitesourcesCountRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemSitesourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the siteSource objects associated with an ediscoveryCustodian.
// returns a SiteSourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-list-sitesources?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SiteSourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSiteSourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SiteSourceCollectionResponseable), nil
}
// Post create a new siteSource object associated with an eDiscovery custodian.
// returns a SiteSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-post-sitesources?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SiteSourceable, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SiteSourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSiteSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SiteSourceable), nil
}
// ToGetRequestInformation get a list of the siteSource objects associated with an ediscoveryCustodian.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new siteSource object associated with an eDiscovery custodian.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SiteSourceable, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

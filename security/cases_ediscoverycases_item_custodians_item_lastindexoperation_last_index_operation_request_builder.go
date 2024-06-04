package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder provides operations to manage the lastIndexOperation property of the microsoft.graph.security.ediscoveryCustodian entity.
type CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderGetQueryParameters get a list of the ediscoveryIndexOperation associated with an ediscoveryCustodian.
type CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderGetQueryParameters
}
// NewCasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder) {
    m := &CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/lastIndexOperation{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder instantiates a new CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of the ediscoveryIndexOperation associated with an ediscoveryCustodian.
// returns a EdiscoveryIndexOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-list-lastindexoperation?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryIndexOperationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryIndexOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryIndexOperationable), nil
}
// ToGetRequestInformation get a list of the ediscoveryIndexOperation associated with an ediscoveryCustodian.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

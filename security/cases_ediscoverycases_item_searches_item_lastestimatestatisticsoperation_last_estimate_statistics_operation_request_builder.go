package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder provides operations to manage the lastEstimateStatisticsOperation property of the microsoft.graph.security.ediscoverySearch entity.
type CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderGetQueryParameters get the last ediscoveryEstimateOperation objects and their properties.
type CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderGetQueryParameters
}
// NewCasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder) {
    m := &CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/lastEstimateStatisticsOperation{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder instantiates a new CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the last ediscoveryEstimateOperation objects and their properties.
// returns a EdiscoveryEstimateOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverysearch-list-lastestimatestatisticsoperation?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryEstimateOperationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryEstimateOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryEstimateOperationable), nil
}
// ToGetRequestInformation get the last ediscoveryEstimateOperation objects and their properties.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

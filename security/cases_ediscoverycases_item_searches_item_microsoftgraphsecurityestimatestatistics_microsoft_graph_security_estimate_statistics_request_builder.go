package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder provides operations to call the estimateStatistics method.
type CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) {
    m := &CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/microsoft.graph.security.estimateStatistics", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder instantiates a new CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run an estimate of the number of emails and documents in the eDiscovery search. To learn more about searches in eDiscovery, see Collect data for a case in eDiscovery (Premium).
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverysearch-estimatestatistics?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) Post(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation run an estimate of the number of emails and documents in the eDiscovery search. To learn more about searches in eDiscovery, see Collect data for a case in eDiscovery (Premium).
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

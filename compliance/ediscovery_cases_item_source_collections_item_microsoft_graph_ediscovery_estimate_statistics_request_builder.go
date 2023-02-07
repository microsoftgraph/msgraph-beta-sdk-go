package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder provides operations to call the estimateStatistics method.
type EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilderInternal instantiates a new MicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder) {
    m := &EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}/microsoft.graph.ediscovery.estimateStatistics";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder instantiates a new MicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run an estimate of the number of emails and documents in the source collection. To learn more about source collections (also known as searches in eDiscovery), see Collect data for a case in Advanced eDiscovery.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/ediscovery-sourcecollection-estimatestatistics?view=graph-rest-1.0
func (m *EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation run an estimate of the number of emails and documents in the source collection. To learn more about source collections (also known as searches in eDiscovery), see Collect data for a case in Advanced eDiscovery.
func (m *EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

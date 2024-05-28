package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
type EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderGetQueryParameters returns a list of sourceCollection objects associated with this case.
type EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderGetQueryParameters
}
// EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdditionalSources provides operations to manage the additionalSources property of the microsoft.graph.ediscovery.sourceCollection entity.
// returns a *EdiscoveryCasesItemSourcecollectionsItemAdditionalsourcesAdditionalSourcesRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) AdditionalSources()(*EdiscoveryCasesItemSourcecollectionsItemAdditionalsourcesAdditionalSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemAdditionalsourcesAdditionalSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddToReviewSetOperation provides operations to manage the addToReviewSetOperation property of the microsoft.graph.ediscovery.sourceCollection entity.
// returns a *EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) AddToReviewSetOperation()(*EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderInternal instantiates a new EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) {
    m := &EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder instantiates a new EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustodianSources provides operations to manage the custodianSources property of the microsoft.graph.ediscovery.sourceCollection entity.
// returns a *EdiscoveryCasesItemSourcecollectionsItemCustodiansourcesCustodianSourcesRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) CustodianSources()(*EdiscoveryCasesItemSourcecollectionsItemCustodiansourcesCustodianSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemCustodiansourcesCustodianSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a sourceCollection object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-sourcecollection-delete?view=graph-rest-beta
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of sourceCollection objects associated with this case.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a SourceCollectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSourceCollectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable), nil
}
// LastEstimateStatisticsOperation provides operations to manage the lastEstimateStatisticsOperation property of the microsoft.graph.ediscovery.sourceCollection entity.
// returns a *EdiscoveryCasesItemSourcecollectionsItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) LastEstimateStatisticsOperation()(*EdiscoveryCasesItemSourcecollectionsItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryEstimateStatistics provides operations to call the estimateStatistics method.
// returns a *EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoveryestimatestatisticsMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) MicrosoftGraphEdiscoveryEstimateStatistics()(*EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoveryestimatestatisticsMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoveryestimatestatisticsMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryPurgeData provides operations to call the purgeData method.
// returns a *EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) MicrosoftGraphEdiscoveryPurgeData()(*EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NoncustodialSources provides operations to manage the noncustodialSources property of the microsoft.graph.ediscovery.sourceCollection entity.
// returns a *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) NoncustodialSources()(*EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a sourceCollection object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a SourceCollectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-sourcecollection-update?view=graph-rest-beta
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSourceCollectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable), nil
}
// ToDeleteRequestInformation delete a sourceCollection object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns a list of sourceCollection objects associated with this case.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a sourceCollection object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsSourceCollectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

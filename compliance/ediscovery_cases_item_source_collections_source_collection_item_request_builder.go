package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
type EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetQueryParameters returns a list of sourceCollection objects associated with this case.
type EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetQueryParameters
}
// EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdditionalSources provides operations to manage the additionalSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) AdditionalSources()(*EdiscoveryCasesItemSourceCollectionsItemAdditionalSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemAdditionalSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddToReviewSetOperation provides operations to manage the addToReviewSetOperation property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) AddToReviewSetOperation()(*EdiscoveryCasesItemSourceCollectionsItemAddToReviewSetOperationRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemAddToReviewSetOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) {
    m := &EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustodianSources provides operations to manage the custodianSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) CustodianSources()(*EdiscoveryCasesItemSourceCollectionsItemCustodianSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemCustodianSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a sourceCollection object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-sourcecollection-delete?view=graph-rest-1.0
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get returns a list of sourceCollection objects associated with this case.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) LastEstimateStatisticsOperation()(*EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryEstimateStatistics provides operations to call the estimateStatistics method.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) MicrosoftGraphEdiscoveryEstimateStatistics()(*EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryPurgeData provides operations to call the purgeData method.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) MicrosoftGraphEdiscoveryPurgeData()(*EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NoncustodialSources provides operations to manage the noncustodialSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) NoncustodialSources()(*EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a sourceCollection object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-sourcecollection-update?view=graph-rest-1.0
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation returns a list of sourceCollection objects associated with this case.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of a sourceCollection object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

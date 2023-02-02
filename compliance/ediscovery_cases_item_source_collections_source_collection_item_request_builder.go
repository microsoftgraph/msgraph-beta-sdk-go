package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
type EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
    return NewEdiscoveryCasesItemSourceCollectionsItemAdditionalSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AdditionalSourcesById provides operations to manage the additionalSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) AdditionalSourcesById(id string)(*EdiscoveryCasesItemSourceCollectionsItemAdditionalSourcesDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemSourceCollectionsItemAdditionalSourcesDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AddToReviewSetOperation provides operations to manage the addToReviewSetOperation property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) AddToReviewSetOperation()(*EdiscoveryCasesItemSourceCollectionsItemAddToReviewSetOperationRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemAddToReviewSetOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, sourceCollectionId *string)(*EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) {
    m := &EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if sourceCollectionId != nil {
        urlTplParams["sourceCollection%2Did"] = *sourceCollectionId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CustodianSources provides operations to manage the custodianSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) CustodianSources()(*EdiscoveryCasesItemSourceCollectionsItemCustodianSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemCustodianSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CustodianSourcesById provides operations to manage the custodianSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) CustodianSourcesById(id string)(*EdiscoveryCasesItemSourceCollectionsItemCustodianSourcesDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemSourceCollectionsItemCustodianSourcesDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Delete delete navigation property sourceCollections for compliance
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get returns a list of sourceCollection objects associated with this case.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSourceCollectionFromDiscriminatorValue, errorMapping)
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
    return NewEdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEdiscoveryEstimateStatistics provides operations to call the estimateStatistics method.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) MicrosoftGraphEdiscoveryEstimateStatistics()(*EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsEstimateStatisticsRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryEstimateStatisticsEstimateStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEdiscoveryPurgeData provides operations to call the purgeData method.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) MicrosoftGraphEdiscoveryPurgeData()(*EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataPurgeDataRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataPurgeDataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NoncustodialSources provides operations to manage the noncustodialSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) NoncustodialSources()(*EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NoncustodialSourcesById provides operations to manage the noncustodialSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) NoncustodialSourcesById(id string)(*EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Patch update the navigation property sourceCollections in compliance
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSourceCollectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable), nil
}
// ToDeleteRequestInformation delete navigation property sourceCollections for compliance
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation returns a list of sourceCollection objects associated with this case.
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update the navigation property sourceCollections in compliance
func (m *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

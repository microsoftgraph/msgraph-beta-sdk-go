package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
type ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetQueryParameters returns a list of sourceCollection objects associated with this case.
type ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetQueryParameters
}
// ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdditionalSources provides operations to manage the additionalSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) AdditionalSources()(*ComplianceEdiscoveryCasesItemSourceCollectionsItemAdditionalSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemAdditionalSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdditionalSourcesById provides operations to manage the additionalSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) AdditionalSourcesById(id string)(*ComplianceEdiscoveryCasesItemSourceCollectionsItemAdditionalSourcesDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemAdditionalSourcesDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AddToReviewSetOperation provides operations to manage the addToReviewSetOperation property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) AddToReviewSetOperation()(*ComplianceEdiscoveryCasesItemSourceCollectionsItemAddToReviewSetOperationRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemAddToReviewSetOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) {
    m := &ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sourceCollections for compliance
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation returns a list of sourceCollection objects associated with this case.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property sourceCollections in compliance
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CustodianSources provides operations to manage the custodianSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) CustodianSources()(*ComplianceEdiscoveryCasesItemSourceCollectionsItemCustodianSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemCustodianSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustodianSourcesById provides operations to manage the custodianSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) CustodianSourcesById(id string)(*ComplianceEdiscoveryCasesItemSourceCollectionsItemCustodianSourcesDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemCustodianSourcesDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property sourceCollections for compliance
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EstimateStatistics provides operations to call the estimateStatistics method.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) EstimateStatistics()(*ComplianceEdiscoveryCasesItemSourceCollectionsItemEstimateStatisticsRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemEstimateStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns a list of sourceCollection objects associated with this case.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSourceCollectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable), nil
}
// LastEstimateStatisticsOperation provides operations to manage the lastEstimateStatisticsOperation property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) LastEstimateStatisticsOperation()(*ComplianceEdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialSources provides operations to manage the noncustodialSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) NoncustodialSources()(*ComplianceEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialSourcesById provides operations to manage the noncustodialSources property of the microsoft.graph.ediscovery.sourceCollection entity.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) NoncustodialSourcesById(id string)(*ComplianceEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["noncustodialDataSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property sourceCollections in compliance
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSourceCollectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable), nil
}
// PurgeData provides operations to call the purgeData method.
func (m *ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) PurgeData()(*ComplianceEdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

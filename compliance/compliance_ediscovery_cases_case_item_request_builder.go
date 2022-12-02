package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// ComplianceEdiscoveryCasesCaseItemRequestBuilder provides operations to manage the cases property of the microsoft.graph.ediscovery.ediscoveryroot entity.
type ComplianceEdiscoveryCasesCaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComplianceEdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComplianceEdiscoveryCasesCaseItemRequestBuilderGetQueryParameters get cases from compliance
type ComplianceEdiscoveryCasesCaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComplianceEdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComplianceEdiscoveryCasesCaseItemRequestBuilderGetQueryParameters
}
// ComplianceEdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Close provides operations to call the close method.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) Close()(*ComplianceEdiscoveryCasesItemCloseRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCloseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewComplianceEdiscoveryCasesCaseItemRequestBuilderInternal instantiates a new CaseItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesCaseItemRequestBuilder) {
    m := &ComplianceEdiscoveryCasesCaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComplianceEdiscoveryCasesCaseItemRequestBuilder instantiates a new CaseItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesCaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesCaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComplianceEdiscoveryCasesCaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cases for compliance
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get cases from compliance
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property cases in compliance
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable, requestConfiguration *ComplianceEdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Custodians provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) Custodians()(*ComplianceEdiscoveryCasesItemCustodiansRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCustodiansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustodiansById provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) CustodiansById(id string)(*ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["custodian%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property cases for compliance
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get cases from compliance
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCase_escapedFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable), nil
}
// LegalHolds provides operations to manage the legalHolds property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) LegalHolds()(*ComplianceEdiscoveryCasesItemLegalHoldsRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemLegalHoldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LegalHoldsById provides operations to manage the legalHolds property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) LegalHoldsById(id string)(*ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["legalHold%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NoncustodialDataSources provides operations to manage the noncustodialDataSources property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) NoncustodialDataSources()(*ComplianceEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialDataSourcesById provides operations to manage the noncustodialDataSources property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) NoncustodialDataSourcesById(id string)(*ComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["noncustodialDataSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) Operations()(*ComplianceEdiscoveryCasesItemOperationsRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) OperationsById(id string)(*ComplianceEdiscoveryCasesItemOperationsCaseOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["caseOperation%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemOperationsCaseOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property cases in compliance
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable, requestConfiguration *ComplianceEdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCase_escapedFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable), nil
}
// Reopen provides operations to call the reopen method.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) Reopen()(*ComplianceEdiscoveryCasesItemReopenRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemReopenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewSets provides operations to manage the reviewSets property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) ReviewSets()(*ComplianceEdiscoveryCasesItemReviewSetsRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemReviewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewSetsById provides operations to manage the reviewSets property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) ReviewSetsById(id string)(*ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["reviewSet%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) Settings()(*ComplianceEdiscoveryCasesItemSettingsRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SourceCollections provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) SourceCollections()(*ComplianceEdiscoveryCasesItemSourceCollectionsRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemSourceCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SourceCollectionsById provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) SourceCollectionsById(id string)(*ComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sourceCollection%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tags provides operations to manage the tags property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) Tags()(*ComplianceEdiscoveryCasesItemTagsRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById provides operations to manage the tags property of the microsoft.graph.ediscovery.case entity.
func (m *ComplianceEdiscoveryCasesCaseItemRequestBuilder) TagsById(id string)(*ComplianceEdiscoveryCasesItemTagsTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tag%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemTagsTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

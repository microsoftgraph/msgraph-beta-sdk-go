package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesCaseItemRequestBuilder provides operations to manage the cases property of the microsoft.graph.ediscovery.ediscoveryroot entity.
type EdiscoveryCasesCaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCasesCaseItemRequestBuilderGetQueryParameters get cases from compliance
type EdiscoveryCasesCaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesCaseItemRequestBuilderGetQueryParameters
}
// EdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesCaseItemRequestBuilderInternal instantiates a new CaseItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, caseId *string)(*EdiscoveryCasesCaseItemRequestBuilder) {
    m := &EdiscoveryCasesCaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if caseId != nil {
        urlTplParams["case%2Did"] = *caseId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEdiscoveryCasesCaseItemRequestBuilder instantiates a new CaseItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesCaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesCaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesCaseItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Custodians provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) Custodians()(*EdiscoveryCasesItemCustodiansRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CustodiansById provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) CustodiansById(id string)(*EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Delete delete navigation property cases for compliance
func (m *EdiscoveryCasesCaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get cases from compliance
func (m *EdiscoveryCasesCaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCase_escapedFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable), nil
}
// LegalHolds provides operations to manage the legalHolds property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) LegalHolds()(*EdiscoveryCasesItemLegalHoldsRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LegalHoldsById provides operations to manage the legalHolds property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) LegalHoldsById(id string)(*EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MicrosoftGraphEdiscoveryClose provides operations to call the close method.
func (m *EdiscoveryCasesCaseItemRequestBuilder) MicrosoftGraphEdiscoveryClose()(*EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseCloseRequestBuilder) {
    return NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseCloseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEdiscoveryReopen provides operations to call the reopen method.
func (m *EdiscoveryCasesCaseItemRequestBuilder) MicrosoftGraphEdiscoveryReopen()(*EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenReopenRequestBuilder) {
    return NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenReopenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NoncustodialDataSources provides operations to manage the noncustodialDataSources property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) NoncustodialDataSources()(*EdiscoveryCasesItemNoncustodialDataSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NoncustodialDataSourcesById provides operations to manage the noncustodialDataSources property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) NoncustodialDataSourcesById(id string)(*EdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemNoncustodialDataSourcesNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Operations provides operations to manage the operations property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) Operations()(*EdiscoveryCasesItemOperationsRequestBuilder) {
    return NewEdiscoveryCasesItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) OperationsById(id string)(*EdiscoveryCasesItemOperationsCaseOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemOperationsCaseOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Patch update the navigation property cases in compliance
func (m *EdiscoveryCasesCaseItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCase_escapedFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable), nil
}
// ReviewSets provides operations to manage the reviewSets property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) ReviewSets()(*EdiscoveryCasesItemReviewSetsRequestBuilder) {
    return NewEdiscoveryCasesItemReviewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReviewSetsById provides operations to manage the reviewSets property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) ReviewSetsById(id string)(*EdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Settings provides operations to manage the settings property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) Settings()(*EdiscoveryCasesItemSettingsRequestBuilder) {
    return NewEdiscoveryCasesItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SourceCollections provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) SourceCollections()(*EdiscoveryCasesItemSourceCollectionsRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SourceCollectionsById provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) SourceCollectionsById(id string)(*EdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemSourceCollectionsSourceCollectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Tags provides operations to manage the tags property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) Tags()(*EdiscoveryCasesItemTagsRequestBuilder) {
    return NewEdiscoveryCasesItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TagsById provides operations to manage the tags property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) TagsById(id string)(*EdiscoveryCasesItemTagsTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemTagsTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ToDeleteRequestInformation delete navigation property cases for compliance
func (m *EdiscoveryCasesCaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get cases from compliance
func (m *EdiscoveryCasesCaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property cases in compliance
func (m *EdiscoveryCasesCaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Case_escapedable, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

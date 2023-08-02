package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesCaseItemRequestBuilder provides operations to manage the cases property of the microsoft.graph.ediscovery.ediscoveryroot entity.
type EdiscoveryCasesCaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCasesCaseItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a case object.
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
func NewEdiscoveryCasesCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesCaseItemRequestBuilder) {
    m := &EdiscoveryCasesCaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesCaseItemRequestBuilder instantiates a new CaseItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesCaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesCaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesCaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Custodians provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) Custodians()(*EdiscoveryCasesItemCustodiansRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a case object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-case-delete?view=graph-rest-1.0
func (m *EdiscoveryCasesCaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of a case object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-case-get?view=graph-rest-1.0
func (m *EdiscoveryCasesCaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseEscapedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCaseEscapedFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseEscapedable), nil
}
// LegalHolds provides operations to manage the legalHolds property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) LegalHolds()(*EdiscoveryCasesItemLegalHoldsRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryClose provides operations to call the close method.
func (m *EdiscoveryCasesCaseItemRequestBuilder) MicrosoftGraphEdiscoveryClose()(*EdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilder) {
    return NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryCloseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryReopen provides operations to call the reopen method.
func (m *EdiscoveryCasesCaseItemRequestBuilder) MicrosoftGraphEdiscoveryReopen()(*EdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilder) {
    return NewEdiscoveryCasesItemMicrosoftGraphEdiscoveryReopenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NoncustodialDataSources provides operations to manage the noncustodialDataSources property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) NoncustodialDataSources()(*EdiscoveryCasesItemNoncustodialDataSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) Operations()(*EdiscoveryCasesItemOperationsRequestBuilder) {
    return NewEdiscoveryCasesItemOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a case object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-case-update?view=graph-rest-1.0
func (m *EdiscoveryCasesCaseItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseEscapedable, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseEscapedable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCaseEscapedFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseEscapedable), nil
}
// ReviewSets provides operations to manage the reviewSets property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) ReviewSets()(*EdiscoveryCasesItemReviewSetsRequestBuilder) {
    return NewEdiscoveryCasesItemReviewSetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) Settings()(*EdiscoveryCasesItemSettingsRequestBuilder) {
    return NewEdiscoveryCasesItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SourceCollections provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) SourceCollections()(*EdiscoveryCasesItemSourceCollectionsRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.ediscovery.case entity.
func (m *EdiscoveryCasesCaseItemRequestBuilder) Tags()(*EdiscoveryCasesItemTagsRequestBuilder) {
    return NewEdiscoveryCasesItemTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a case object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesCaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve the properties and relationships of a case object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesCaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a case object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesCaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseEscapedable, requestConfiguration *EdiscoveryCasesCaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

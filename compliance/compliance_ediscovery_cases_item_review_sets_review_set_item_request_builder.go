package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder provides operations to manage the reviewSets property of the microsoft.graph.ediscovery.case entity.
type ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderGetQueryParameters returns a list of reviewSet objects in the case. Read-only. Nullable.
type ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderGetQueryParameters
}
// ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddToReviewSet provides operations to call the addToReviewSet method.
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) AddToReviewSet()(*ComplianceEdiscoveryCasesItemReviewSetsItemAddToReviewSetRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemReviewSetsItemAddToReviewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderInternal instantiates a new ReviewSetItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) {
    m := &ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/reviewSets/{reviewSet%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder instantiates a new ReviewSetItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reviewSets for compliance
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property reviewSets in compliance
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable, requestConfiguration *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property reviewSets for compliance
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Export provides operations to call the export method.
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) Export()(*ComplianceEdiscoveryCasesItemReviewSetsItemExportRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemReviewSetsItemExportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateReviewSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable), nil
}
// Patch update the navigation property reviewSets in compliance
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable, requestConfiguration *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateReviewSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable), nil
}
// Queries provides operations to manage the queries property of the microsoft.graph.ediscovery.reviewSet entity.
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) Queries()(*ComplianceEdiscoveryCasesItemReviewSetsItemQueriesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemReviewSetsItemQueriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// QueriesById provides operations to manage the queries property of the microsoft.graph.ediscovery.reviewSet entity.
func (m *ComplianceEdiscoveryCasesItemReviewSetsReviewSetItemRequestBuilder) QueriesById(id string)(*ComplianceEdiscoveryCasesItemReviewSetsItemQueriesReviewSetQueryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["reviewSetQuery%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemReviewSetsItemQueriesReviewSetQueryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

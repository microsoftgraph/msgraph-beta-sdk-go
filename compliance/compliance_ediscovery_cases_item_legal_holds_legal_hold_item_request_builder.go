package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder provides operations to manage the legalHolds property of the microsoft.graph.ediscovery.case entity.
type ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetQueryParameters returns a list of case legalHold objects for this case.  Nullable.
type ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetQueryParameters
}
// ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderInternal instantiates a new LegalHoldItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) {
    m := &ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/legalHolds/{legalHold%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder instantiates a new LegalHoldItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property legalHolds for compliance
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of case legalHold objects for this case.  Nullable.
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property legalHolds in compliance
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable, requestConfiguration *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property legalHolds for compliance
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case legalHold objects for this case.  Nullable.
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateLegalHoldFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable), nil
}
// Patch update the navigation property legalHolds in compliance
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable, requestConfiguration *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateLegalHoldFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable), nil
}
// SiteSources provides operations to manage the siteSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) SiteSources()(*ComplianceEdiscoveryCasesItemLegalHoldsItemSiteSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemLegalHoldsItemSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSourcesById provides operations to manage the siteSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) SiteSourcesById(id string)(*ComplianceEdiscoveryCasesItemLegalHoldsItemSiteSourcesSiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemLegalHoldsItemSiteSourcesSiteSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnifiedGroupSources provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) UnifiedGroupSources()(*ComplianceEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnifiedGroupSourcesById provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) UnifiedGroupSourcesById(id string)(*ComplianceEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserSources provides operations to manage the userSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) UserSources()(*ComplianceEdiscoveryCasesItemLegalHoldsItemUserSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemLegalHoldsItemUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSourcesById provides operations to manage the userSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *ComplianceEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) UserSourcesById(id string)(*ComplianceEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

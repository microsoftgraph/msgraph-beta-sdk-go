package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
type ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetQueryParameters returns a list of case custodian objects for this case.  Nullable.
type ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetQueryParameters
}
// ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate provides operations to call the activate method.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) Activate()(*ComplianceEdiscoveryCasesItemCustodiansItemActivateRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCustodiansItemActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyHold provides operations to call the applyHold method.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) ApplyHold()(*ComplianceEdiscoveryCasesItemCustodiansItemApplyHoldRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCustodiansItemApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderInternal instantiates a new CustodianItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) {
    m := &ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians/{custodian%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder instantiates a new CustodianItemRequestBuilder and sets the default values.
func NewComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property custodians for compliance
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of case custodian objects for this case.  Nullable.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property custodians in compliance
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property custodians for compliance
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case custodian objects for this case.  Nullable.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable), nil
}
// Patch update the navigation property custodians in compliance
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable), nil
}
// Release provides operations to call the release method.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) Release()(*ComplianceEdiscoveryCasesItemCustodiansItemReleaseRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCustodiansItemReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveHold provides operations to call the removeHold method.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) RemoveHold()(*ComplianceEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCustodiansItemRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSources provides operations to manage the siteSources property of the microsoft.graph.ediscovery.custodian entity.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) SiteSources()(*ComplianceEdiscoveryCasesItemCustodiansItemSiteSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCustodiansItemSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSourcesById provides operations to manage the siteSources property of the microsoft.graph.ediscovery.custodian entity.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) SiteSourcesById(id string)(*ComplianceEdiscoveryCasesItemCustodiansItemSiteSourcesSiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemCustodiansItemSiteSourcesSiteSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnifiedGroupSources provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.custodian entity.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) UnifiedGroupSources()(*ComplianceEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnifiedGroupSourcesById provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.custodian entity.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) UnifiedGroupSourcesById(id string)(*ComplianceEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UpdateIndex provides operations to call the updateIndex method.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) UpdateIndex()(*ComplianceEdiscoveryCasesItemCustodiansItemUpdateIndexRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCustodiansItemUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSources provides operations to manage the userSources property of the microsoft.graph.ediscovery.custodian entity.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) UserSources()(*ComplianceEdiscoveryCasesItemCustodiansItemUserSourcesRequestBuilder) {
    return NewComplianceEdiscoveryCasesItemCustodiansItemUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSourcesById provides operations to manage the userSources property of the microsoft.graph.ediscovery.custodian entity.
func (m *ComplianceEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) UserSourcesById(id string)(*ComplianceEdiscoveryCasesItemCustodiansItemUserSourcesUserSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource%2Did"] = id
    }
    return NewComplianceEdiscoveryCasesItemCustodiansItemUserSourcesUserSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

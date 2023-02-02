package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder provides operations to manage the legalHolds property of the microsoft.graph.ediscovery.case entity.
type EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetQueryParameters returns a list of case legalHold objects for this case.  Nullable.
type EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetQueryParameters
}
// EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderInternal instantiates a new LegalHoldItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, legalHoldId *string)(*EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) {
    m := &EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/legalHolds/{legalHold%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if legalHoldId != nil {
        urlTplParams["legalHold%2Did"] = *legalHoldId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder instantiates a new LegalHoldItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property legalHolds for compliance
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case legalHold objects for this case.  Nullable.
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateLegalHoldFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable), nil
}
// Patch update the navigation property legalHolds in compliance
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable, requestConfiguration *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateLegalHoldFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable), nil
}
// SiteSources provides operations to manage the siteSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) SiteSources()(*EdiscoveryCasesItemLegalHoldsItemSiteSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsItemSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SiteSourcesById provides operations to manage the siteSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) SiteSourcesById(id string)(*EdiscoveryCasesItemLegalHoldsItemSiteSourcesSiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemLegalHoldsItemSiteSourcesSiteSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ToDeleteRequestInformation delete navigation property legalHolds for compliance
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation returns a list of case legalHold objects for this case.  Nullable.
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property legalHolds in compliance
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.LegalHoldable, requestConfiguration *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UnifiedGroupSources provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) UnifiedGroupSources()(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UnifiedGroupSourcesById provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) UnifiedGroupSourcesById(id string)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserSources provides operations to manage the userSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) UserSources()(*EdiscoveryCasesItemLegalHoldsItemUserSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsItemUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserSourcesById provides operations to manage the userSources property of the microsoft.graph.ediscovery.legalHold entity.
func (m *EdiscoveryCasesItemLegalHoldsLegalHoldItemRequestBuilder) UserSourcesById(id string)(*EdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEdiscoveryCasesItemLegalHoldsItemUserSourcesUserSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}

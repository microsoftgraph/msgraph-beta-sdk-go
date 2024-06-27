package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.legalHold entity.
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetQueryParameters get unifiedGroupSources from compliance
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetQueryParameters
}
// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal instantiates a new EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) {
    m := &EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/legalHolds/{legalHold%2Did}/unifiedGroupSources/{unifiedGroupSource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder instantiates a new EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property unifiedGroupSources for compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get unifiedGroupSources from compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a UnifiedGroupSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateUnifiedGroupSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.ediscovery.unifiedGroupSource entity.
// returns a *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder when successful
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) Group()(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property unifiedGroupSources in compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a UnifiedGroupSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceable, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateUnifiedGroupSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceable), nil
}
// ToDeleteRequestInformation delete navigation property unifiedGroupSources for compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get unifiedGroupSources from compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property unifiedGroupSources in compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceable, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder when successful
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

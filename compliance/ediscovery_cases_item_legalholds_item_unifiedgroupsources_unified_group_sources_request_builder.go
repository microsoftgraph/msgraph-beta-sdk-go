package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.legalHold entity.
type EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderGetQueryParameters get unifiedGroupSources from compliance
type EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderGetQueryParameters
}
// EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedGroupSourceId provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.legalHold entity.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder when successful
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) ByUnifiedGroupSourceId(unifiedGroupSourceId string)(*EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedGroupSourceId != "" {
        urlTplParams["unifiedGroupSource%2Did"] = unifiedGroupSourceId
    }
    return NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderInternal instantiates a new EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) {
    m := &EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/legalHolds/{legalHold%2Did}/unifiedGroupSources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder instantiates a new EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesCountRequestBuilder when successful
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) Count()(*EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesCountRequestBuilder) {
    return NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get unifiedGroupSources from compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a UnifiedGroupSourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateUnifiedGroupSourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceCollectionResponseable), nil
}
// Post create new navigation property to unifiedGroupSources for compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a UnifiedGroupSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) Post(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceable, requestConfiguration *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderPostRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get unifiedGroupSources from compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to unifiedGroupSources for compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.UnifiedGroupSourceable, requestConfiguration *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder when successful
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

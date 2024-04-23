package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
type EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetQueryParameters returns a list of case custodian objects for this case.  Nullable.
type EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetQueryParameters
}
// EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderInternal instantiates a new EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) {
    m := &EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians/{custodian%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder instantiates a new EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemCustodiansCustodianItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property custodians for compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case custodian objects for this case.  Nullable.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a Custodianable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable), nil
}
// LastIndexOperation provides operations to manage the lastIndexOperation property of the microsoft.graph.ediscovery.dataSourceContainer entity.
// returns a *EdiscoveryCasesItemCustodiansItemLastIndexOperationRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) LastIndexOperation()(*EdiscoveryCasesItemCustodiansItemLastIndexOperationRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemLastIndexOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryActivate provides operations to call the activate method.
// returns a *EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryActivateRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) MicrosoftGraphEdiscoveryActivate()(*EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryActivateRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryApplyHold provides operations to call the applyHold method.
// returns a *EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) MicrosoftGraphEdiscoveryApplyHold()(*EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryApplyHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryRelease provides operations to call the release method.
// returns a *EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryReleaseRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) MicrosoftGraphEdiscoveryRelease()(*EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryReleaseRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryReleaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryRemoveHold provides operations to call the removeHold method.
// returns a *EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) MicrosoftGraphEdiscoveryRemoveHold()(*EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryUpdateIndex provides operations to call the updateIndex method.
// returns a *EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) MicrosoftGraphEdiscoveryUpdateIndex()(*EdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property custodians in compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a Custodianable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable), nil
}
// SiteSources provides operations to manage the siteSources property of the microsoft.graph.ediscovery.custodian entity.
// returns a *EdiscoveryCasesItemCustodiansItemSiteSourcesRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) SiteSources()(*EdiscoveryCasesItemCustodiansItemSiteSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemSiteSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property custodians for compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns a list of case custodian objects for this case.  Nullable.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property custodians in compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UnifiedGroupSources provides operations to manage the unifiedGroupSources property of the microsoft.graph.ediscovery.custodian entity.
// returns a *EdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) UnifiedGroupSources()(*EdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserSources provides operations to manage the userSources property of the microsoft.graph.ediscovery.custodian entity.
// returns a *EdiscoveryCasesItemCustodiansItemUserSourcesRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) UserSources()(*EdiscoveryCasesItemCustodiansItemUserSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansItemUserSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemCustodiansCustodianItemRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansCustodianItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

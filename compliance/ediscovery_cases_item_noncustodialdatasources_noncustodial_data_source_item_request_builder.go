package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder provides operations to manage the noncustodialDataSources property of the microsoft.graph.ediscovery.case entity.
type EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderGetQueryParameters read the properties and relationships of a noncustodialDataSource object.
type EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderGetQueryParameters
}
// EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderInternal instantiates a new EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) {
    m := &EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/noncustodialDataSources/{noncustodialDataSource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder instantiates a new EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataSource provides operations to manage the dataSource property of the microsoft.graph.ediscovery.noncustodialDataSource entity.
// returns a *EdiscoveryCasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) DataSource()(*EdiscoveryCasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property noncustodialDataSources for compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a noncustodialDataSource object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a NoncustodialDataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-noncustodialdatasource-get?view=graph-rest-beta
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable), nil
}
// LastIndexOperation provides operations to manage the lastIndexOperation property of the microsoft.graph.ediscovery.dataSourceContainer entity.
// returns a *EdiscoveryCasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) LastIndexOperation()(*EdiscoveryCasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryApplyHold provides operations to call the applyHold method.
// returns a *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphEdiscoveryApplyHold()(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryapplyholdMicrosoftGraphEdiscoveryApplyHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryRelease provides operations to call the release method.
// returns a *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryreleaseMicrosoftGraphEdiscoveryReleaseRequestBuilder when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphEdiscoveryRelease()(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryreleaseMicrosoftGraphEdiscoveryReleaseRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryreleaseMicrosoftGraphEdiscoveryReleaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryRemoveHold provides operations to call the removeHold method.
// returns a *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphEdiscoveryRemoveHold()(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphEdiscoveryUpdateIndex provides operations to call the updateIndex method.
// returns a *EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphEdiscoveryUpdateIndex()(*EdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialdatasourcesItemMicrosoftgraphediscoveryupdateindexMicrosoftGraphEdiscoveryUpdateIndexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property noncustodialDataSources in compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a NoncustodialDataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable), nil
}
// ToDeleteRequestInformation delete navigation property noncustodialDataSources for compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a noncustodialDataSource object.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property noncustodialDataSources in compliance
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, requestConfiguration *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder when successful
func (m *EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialdatasourcesNoncustodialDataSourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

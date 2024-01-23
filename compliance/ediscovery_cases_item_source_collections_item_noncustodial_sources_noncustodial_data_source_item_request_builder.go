package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder provides operations to manage the noncustodialSources property of the microsoft.graph.ediscovery.sourceCollection entity.
type EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderGetQueryParameters noncustodialDataSource sources that are included in the sourceCollection
type EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderInternal instantiates a new NoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder) {
    m := &EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}/noncustodialSources/{noncustodialDataSource%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder instantiates a new NoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get noncustodialDataSource sources that are included in the sourceCollection
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
func (m *EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation noncustodialDataSource sources that are included in the sourceCollection
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
func (m *EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
func (m *EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemNoncustodialSourcesNoncustodialDataSourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

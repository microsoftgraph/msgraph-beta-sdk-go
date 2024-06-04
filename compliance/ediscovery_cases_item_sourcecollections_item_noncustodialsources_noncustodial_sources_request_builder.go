package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder provides operations to manage the noncustodialSources property of the microsoft.graph.ediscovery.sourceCollection entity.
type EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetQueryParameters get a list of noncustodialDataSource associated with a sourceCollection.
type EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetQueryParameters struct {
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
// EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetQueryParameters
}
// ByNoncustodialDataSourceId provides operations to manage the noncustodialSources property of the microsoft.graph.ediscovery.sourceCollection entity.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialDataSourceItemRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) ByNoncustodialDataSourceId(noncustodialDataSourceId string)(*EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if noncustodialDataSourceId != "" {
        urlTplParams["noncustodialDataSource%2Did"] = noncustodialDataSourceId
    }
    return NewEdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderInternal instantiates a new EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) {
    m := &EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}/noncustodialSources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder instantiates a new EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesCountRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) Count()(*EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesCountRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of noncustodialDataSource associated with a sourceCollection.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a NoncustodialDataSourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-sourcecollection-list-noncustodialsources?view=graph-rest-beta
func (m *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateNoncustodialDataSourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NoncustodialDataSourceCollectionResponseable), nil
}
// ToGetRequestInformation get a list of noncustodialDataSource associated with a sourceCollection.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemNoncustodialsourcesNoncustodialSourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

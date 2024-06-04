package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder provides operations to manage the addToReviewSetOperation property of the microsoft.graph.ediscovery.sourceCollection entity.
type EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderGetQueryParameters get the last addToReviewSetOperation object associated with a source collection. 
type EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderInternal instantiates a new EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder) {
    m := &EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}/addToReviewSetOperation{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder instantiates a new EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the last addToReviewSetOperation object associated with a source collection. 
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a AddToReviewSetOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-sourcecollection-list-addtoreviewsetoperation?view=graph-rest-beta
func (m *EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AddToReviewSetOperationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateAddToReviewSetOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.AddToReviewSetOperationable), nil
}
// ToGetRequestInformation get the last addToReviewSetOperation object associated with a source collection. 
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

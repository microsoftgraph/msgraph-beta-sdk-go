package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder provides operations to manage the reviewSet property of the microsoft.graph.ediscovery.caseExportOperation entity.
type EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderGetQueryParameters the review set the content is being exported from.
type EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderInternal instantiates a new ReviewSetRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder) {
    m := &EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/operations/{caseOperation%2Did}/microsoft.graph.ediscovery.caseExportOperation/reviewSet{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder instantiates a new ReviewSetRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the review set the content is being exported from.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateReviewSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable), nil
}
// ToGetRequestInformation the review set the content is being exported from.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder) {
    return NewEdiscoveryCasesItemOperationsItemMicrosoftGraphEdiscoveryCaseExportOperationReviewSetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

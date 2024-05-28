package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder casts the previous resource to caseExportOperation.
type EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetQueryParameters get the item of type microsoft.graph.ediscovery.caseOperation as microsoft.graph.ediscovery.caseExportOperation
type EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderInternal instantiates a new EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) {
    m := &EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/operations/{caseOperation%2Did}/microsoft.graph.ediscovery.caseExportOperation{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder instantiates a new EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.ediscovery.caseOperation as microsoft.graph.ediscovery.caseExportOperation
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a CaseExportOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseExportOperationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCaseExportOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseExportOperationable), nil
}
// ReviewSet provides operations to manage the reviewSet property of the microsoft.graph.ediscovery.caseExportOperation entity.
// returns a *EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationReviewsetReviewSetRequestBuilder when successful
func (m *EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) ReviewSet()(*EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationReviewsetReviewSetRequestBuilder) {
    return NewEdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationReviewsetReviewSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.ediscovery.caseOperation as microsoft.graph.ediscovery.caseExportOperation
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder when successful
func (m *EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) {
    return NewEdiscoveryCasesItemOperationsItemMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder casts the previous resource to caseExportOperation.
type EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetQueryParameters get the items of type microsoft.graph.ediscovery.caseExportOperation in the microsoft.graph.ediscovery.caseOperation collection
type EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetQueryParameters struct {
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
// EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderInternal instantiates a new EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) {
    m := &EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/operations/microsoft.graph.ediscovery.caseExportOperation{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder instantiates a new EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationCountRequestBuilder when successful
func (m *EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) Count()(*EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationCountRequestBuilder) {
    return NewEdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.ediscovery.caseExportOperation in the microsoft.graph.ediscovery.caseOperation collection
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a CaseExportOperationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseExportOperationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCaseExportOperationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CaseExportOperationCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.ediscovery.caseExportOperation in the microsoft.graph.ediscovery.caseOperation collection
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder when successful
func (m *EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder) {
    return NewEdiscoveryCasesItemOperationsMicrosoftgraphediscoverycaseexportoperationMicrosoftGraphEdiscoveryCaseExportOperationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

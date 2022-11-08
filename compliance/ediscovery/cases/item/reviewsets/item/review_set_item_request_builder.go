package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
    i08b6f0b5a320f941f6b9e04f436d2deaf397e1da914e9bfb4309c34077c320f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/addtoreviewset"
    i331d9fcb9fb3723ce0d6251992ccb1b62085e60afae20de233674ebeb7e41db8 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/export"
    i95d988bf764f142af2cbc9b71e126590844a8f1556f57902f7965dddff422bb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/queries"
    iddd703417744884f23660dc02635324f3f6311317b830c6b4be8a93262f68548 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/queries/item"
)

// ReviewSetItemRequestBuilder provides operations to manage the reviewSets property of the microsoft.graph.ediscovery.case entity.
type ReviewSetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReviewSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReviewSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReviewSetItemRequestBuilderGetQueryParameters returns a list of reviewSet objects in the case. Read-only. Nullable.
type ReviewSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReviewSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReviewSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReviewSetItemRequestBuilderGetQueryParameters
}
// ReviewSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReviewSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddToReviewSet provides operations to call the addToReviewSet method.
func (m *ReviewSetItemRequestBuilder) AddToReviewSet()(*i08b6f0b5a320f941f6b9e04f436d2deaf397e1da914e9bfb4309c34077c320f6.AddToReviewSetRequestBuilder) {
    return i08b6f0b5a320f941f6b9e04f436d2deaf397e1da914e9bfb4309c34077c320f6.NewAddToReviewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewReviewSetItemRequestBuilderInternal instantiates a new ReviewSetItemRequestBuilder and sets the default values.
func NewReviewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReviewSetItemRequestBuilder) {
    m := &ReviewSetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/reviewSets/{reviewSet%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReviewSetItemRequestBuilder instantiates a new ReviewSetItemRequestBuilder and sets the default values.
func NewReviewSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReviewSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReviewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reviewSets for compliance
func (m *ReviewSetItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ReviewSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ReviewSetItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ReviewSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property reviewSets in compliance
func (m *ReviewSetItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable, requestConfiguration *ReviewSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property reviewSets for compliance
func (m *ReviewSetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReviewSetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Export provides operations to call the export method.
func (m *ReviewSetItemRequestBuilder) Export()(*i331d9fcb9fb3723ce0d6251992ccb1b62085e60afae20de233674ebeb7e41db8.ExportRequestBuilder) {
    return i331d9fcb9fb3723ce0d6251992ccb1b62085e60afae20de233674ebeb7e41db8.NewExportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ReviewSetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ReviewSetItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateReviewSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable), nil
}
// Patch update the navigation property reviewSets in compliance
func (m *ReviewSetItemRequestBuilder) Patch(ctx context.Context, body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable, requestConfiguration *ReviewSetItemRequestBuilderPatchRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateReviewSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ReviewSetable), nil
}
// Queries provides operations to manage the queries property of the microsoft.graph.ediscovery.reviewSet entity.
func (m *ReviewSetItemRequestBuilder) Queries()(*i95d988bf764f142af2cbc9b71e126590844a8f1556f57902f7965dddff422bb5.QueriesRequestBuilder) {
    return i95d988bf764f142af2cbc9b71e126590844a8f1556f57902f7965dddff422bb5.NewQueriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// QueriesById provides operations to manage the queries property of the microsoft.graph.ediscovery.reviewSet entity.
func (m *ReviewSetItemRequestBuilder) QueriesById(id string)(*iddd703417744884f23660dc02635324f3f6311317b830c6b4be8a93262f68548.ReviewSetQueryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["reviewSetQuery%2Did"] = id
    }
    return iddd703417744884f23660dc02635324f3f6311317b830c6b4be8a93262f68548.NewReviewSetQueryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder provides operations to manage the queries property of the microsoft.graph.security.ediscoveryReviewSet entity.
type CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetQueryParameters read the properties and relationships of an ediscoveryReviewSetQuery object.
type CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder instantiates a new CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an ediscoveryReviewSetQuery object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewset-delete-queries?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an ediscoveryReviewSetQuery object.
// returns a EdiscoveryReviewSetQueryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewsetquery-get?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryReviewSetQueryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable), nil
}
// MicrosoftGraphSecurityApplyTags provides operations to call the applyTags method.
// returns a *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) MicrosoftGraphSecurityApplyTags()(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityapplytagsMicrosoftGraphSecurityApplyTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityExport provides operations to call the export method.
// returns a *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) MicrosoftGraphSecurityExport()(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityexportMicrosoftGraphSecurityExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityRun provides operations to call the run method.
// returns a *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) MicrosoftGraphSecurityRun()(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an ediscoveryReviewSetQuery object.
// returns a EdiscoveryReviewSetQueryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewsetquery-update?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryReviewSetQueryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable), nil
}
// ToDeleteRequestInformation delete an ediscoveryReviewSetQuery object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an ediscoveryReviewSetQuery object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an ediscoveryReviewSetQuery object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesEdiscoveryReviewSetQueryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

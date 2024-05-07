package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder provides operations to manage the ediscoveryCases property of the microsoft.graph.security.casesRoot entity.
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetQueryParameters read the properties and relationships of an ediscoveryCase object.
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetQueryParameters
}
// CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderInternal instantiates a new CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) {
    m := &CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder instantiates a new CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Custodians provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoveryCasesItemCustodiansRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Custodians()(*CasesEdiscoveryCasesItemCustodiansRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete an ediscoveryCase object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-casesroot-delete-ediscoverycases?view=graph-rest-beta
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an ediscoveryCase object.
// returns a EdiscoveryCaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-get?view=graph-rest-beta
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable), nil
}
// LegalHolds provides operations to manage the legalHolds property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoveryCasesItemLegalHoldsRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) LegalHolds()(*CasesEdiscoveryCasesItemLegalHoldsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemLegalHoldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityClose provides operations to call the close method.
// returns a *CasesEdiscoveryCasesItemMicrosoftGraphSecurityCloseRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) MicrosoftGraphSecurityClose()(*CasesEdiscoveryCasesItemMicrosoftGraphSecurityCloseRequestBuilder) {
    return NewCasesEdiscoveryCasesItemMicrosoftGraphSecurityCloseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityReopen provides operations to call the reopen method.
// returns a *CasesEdiscoveryCasesItemMicrosoftGraphSecurityReopenRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) MicrosoftGraphSecurityReopen()(*CasesEdiscoveryCasesItemMicrosoftGraphSecurityReopenRequestBuilder) {
    return NewCasesEdiscoveryCasesItemMicrosoftGraphSecurityReopenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NoncustodialDataSources provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) NoncustodialDataSources()(*CasesEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoveryCasesItemOperationsRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Operations()(*CasesEdiscoveryCasesItemOperationsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an ediscoveryCase object.
// returns a EdiscoveryCaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-update?view=graph-rest-beta
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable), nil
}
// ReviewSets provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoveryCasesItemReviewSetsRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ReviewSets()(*CasesEdiscoveryCasesItemReviewSetsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemReviewSetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Searches provides operations to manage the searches property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoveryCasesItemSearchesRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Searches()(*CasesEdiscoveryCasesItemSearchesRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoveryCasesItemSettingsRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Settings()(*CasesEdiscoveryCasesItemSettingsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoveryCasesItemTagsRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Tags()(*CasesEdiscoveryCasesItemTagsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete an ediscoveryCase object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an ediscoveryCase object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an ediscoveryCase object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable, requestConfiguration *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder when successful
func (m *CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) {
    return NewCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

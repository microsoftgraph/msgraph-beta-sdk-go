package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder provides operations to manage the legalHolds property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderGetQueryParameters returns a list of case eDiscoveryHoldPolicy objects for this case.
type CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/legalHolds/{ediscoveryHoldPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder instantiates a new CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property legalHolds for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case eDiscoveryHoldPolicy objects for this case.
// returns a EdiscoveryHoldPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryHoldPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryHoldPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryHoldPolicyable), nil
}
// Patch update the navigation property legalHolds in security
// returns a EdiscoveryHoldPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryHoldPolicyable, requestConfiguration *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryHoldPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryHoldPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryHoldPolicyable), nil
}
// SiteSources provides operations to manage the siteSources property of the microsoft.graph.security.ediscoveryHoldPolicy entity.
// returns a *CasesEdiscoverycasesItemLegalholdsItemSitesourcesSiteSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) SiteSources()(*CasesEdiscoverycasesItemLegalholdsItemSitesourcesSiteSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemLegalholdsItemSitesourcesSiteSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property legalHolds for security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns a list of case eDiscoveryHoldPolicy objects for this case.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property legalHolds in security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryHoldPolicyable, requestConfiguration *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserSources provides operations to manage the userSources property of the microsoft.graph.security.ediscoveryHoldPolicy entity.
// returns a *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) UserSources()(*CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemLegalholdsEdiscoveryHoldPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

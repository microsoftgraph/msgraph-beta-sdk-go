package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemCustodiansRequestBuilder provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoveryCasesItemCustodiansRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemCustodiansRequestBuilderGetQueryParameters get a list of the custodian objects and their properties.
type CasesEdiscoveryCasesItemCustodiansRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesItemCustodiansRequestBuilderGetQueryParameters
}
// CasesEdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEdiscoveryCustodianId provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) ByEdiscoveryCustodianId(ediscoveryCustodianId string)(*CasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ediscoveryCustodianId != "" {
        urlTplParams["ediscoveryCustodian%2Did"] = ediscoveryCustodianId
    }
    return NewCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCasesEdiscoveryCasesItemCustodiansRequestBuilderInternal instantiates a new CasesEdiscoveryCasesItemCustodiansRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansRequestBuilder) {
    m := &CasesEdiscoveryCasesItemCustodiansRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemCustodiansRequestBuilder instantiates a new CasesEdiscoveryCasesItemCustodiansRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemCustodiansRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CasesEdiscoveryCasesItemCustodiansCountRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) Count()(*CasesEdiscoveryCasesItemCustodiansCountRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the custodian objects and their properties.
// returns a EdiscoveryCustodianCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-list-custodians?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryCustodianCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianCollectionResponseable), nil
}
// MicrosoftGraphSecurityApplyHold provides operations to call the applyHold method.
// returns a *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) MicrosoftGraphSecurityApplyHold()(*CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityRemoveHold provides operations to call the removeHold method.
// returns a *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityRemoveHoldRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) MicrosoftGraphSecurityRemoveHold()(*CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityUpdateIndex provides operations to call the updateIndex method.
// returns a *CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityUpdateIndexRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) MicrosoftGraphSecurityUpdateIndex()(*CasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create a new ediscoveryCustodian object.After the custodian object is created, you will need to create the custodian's userSource to reference their mailbox and OneDrive for Business site.
// returns a EdiscoveryCustodianable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-post-custodians?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianable, requestConfiguration *CasesEdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianable), nil
}
// ToGetRequestInformation get a list of the custodian objects and their properties.
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new ediscoveryCustodian object.After the custodian object is created, you will need to create the custodian's userSource to reference their mailbox and OneDrive for Business site.
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianable, requestConfiguration *CasesEdiscoveryCasesItemCustodiansRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians", m.BaseRequestBuilder.PathParameters)
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
// returns a *CasesEdiscoveryCasesItemCustodiansRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemCustodiansRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemCustodiansRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

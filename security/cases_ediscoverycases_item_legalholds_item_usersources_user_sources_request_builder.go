package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder provides operations to manage the userSources property of the microsoft.graph.security.ediscoveryHoldPolicy entity.
type CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderGetQueryParameters data sources that represent Exchange mailboxes.
type CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserSourceId provides operations to manage the userSources property of the microsoft.graph.security.ediscoveryHoldPolicy entity.
// returns a *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourceItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) ByUserSourceId(userSourceId string)(*CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userSourceId != "" {
        urlTplParams["userSource%2Did"] = userSourceId
    }
    return NewCasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) {
    m := &CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/legalHolds/{ediscoveryHoldPolicy%2Did}/userSources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder instantiates a new CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CasesEdiscoverycasesItemLegalholdsItemUsersourcesCountRequestBuilder when successful
func (m *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) Count()(*CasesEdiscoverycasesItemLegalholdsItemUsersourcesCountRequestBuilder) {
    return NewCasesEdiscoverycasesItemLegalholdsItemUsersourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get data sources that represent Exchange mailboxes.
// returns a UserSourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateUserSourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceCollectionResponseable), nil
}
// Post create a new userSource object.
// returns a UserSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryholdpolicy-post-usersources?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable, requestConfiguration *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateUserSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable), nil
}
// ToGetRequestInformation data sources that represent Exchange mailboxes.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new userSource object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.UserSourceable, requestConfiguration *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemLegalholdsItemUsersourcesUserSourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

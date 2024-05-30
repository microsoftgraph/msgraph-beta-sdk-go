package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder provides operations to call the run method.
type CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderGetQueryParameters run reviewset query to get the list of files.
type CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderGetQueryParameters
}
// NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder) {
    m := &CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}/microsoft.graph.security.run(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder instantiates a new CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderInternal(urlParams, requestAdapter)
}
// Get run reviewset query to get the list of files.
// Deprecated: This method is obsolete. Use GetAsRunGetResponse instead.
// returns a CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewsetquery-run?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration)(CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunResponseable), nil
}
// GetAsRunGetResponse run reviewset query to get the list of files.
// returns a CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewsetquery-run?view=graph-rest-beta
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder) GetAsRunGetResponse(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration)(CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunRunGetResponseable), nil
}
// ToGetRequestInformation run reviewset query to get the list of files.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsItemQueriesItemMicrosoftgraphsecurityrunMicrosoftGraphSecurityRunRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

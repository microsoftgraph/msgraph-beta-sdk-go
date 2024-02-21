package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder provides operations to call the run method.
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderGetQueryParameters invoke function run
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderGetQueryParameters
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderInternal instantiates a new CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}/microsoft.graph.security.run(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder instantiates a new CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function run
// Deprecated: This method is obsolete. Use GetAsRunGetResponse instead.
// returns a CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRunResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration)(CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRunResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRunResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRunResponseable), nil
}
// GetAsRunGetResponse invoke function run
// returns a CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRunGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder) GetAsRunGetResponse(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration)(CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRunGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRunGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRunGetResponseable), nil
}
// ToGetRequestInformation invoke function run
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder) {
    return NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityRunRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

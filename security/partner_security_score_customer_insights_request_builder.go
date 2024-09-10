package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partner/security"
)

// PartnerSecurityScoreCustomerInsightsRequestBuilder provides operations to manage the customerInsights property of the microsoft.graph.partner.security.partnerSecurityScore entity.
type PartnerSecurityScoreCustomerInsightsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnerSecurityScoreCustomerInsightsRequestBuilderGetQueryParameters get a list of the customerInsight objects and their properties.
type PartnerSecurityScoreCustomerInsightsRequestBuilderGetQueryParameters struct {
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
// PartnerSecurityScoreCustomerInsightsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreCustomerInsightsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PartnerSecurityScoreCustomerInsightsRequestBuilderGetQueryParameters
}
// PartnerSecurityScoreCustomerInsightsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreCustomerInsightsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCustomerInsightTenantId provides operations to manage the customerInsights property of the microsoft.graph.partner.security.partnerSecurityScore entity.
// returns a *PartnerSecurityScoreCustomerInsightsCustomerInsightTenantItemRequestBuilder when successful
func (m *PartnerSecurityScoreCustomerInsightsRequestBuilder) ByCustomerInsightTenantId(customerInsightTenantId string)(*PartnerSecurityScoreCustomerInsightsCustomerInsightTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if customerInsightTenantId != "" {
        urlTplParams["customerInsight%2DtenantId"] = customerInsightTenantId
    }
    return NewPartnerSecurityScoreCustomerInsightsCustomerInsightTenantItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPartnerSecurityScoreCustomerInsightsRequestBuilderInternal instantiates a new PartnerSecurityScoreCustomerInsightsRequestBuilder and sets the default values.
func NewPartnerSecurityScoreCustomerInsightsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreCustomerInsightsRequestBuilder) {
    m := &PartnerSecurityScoreCustomerInsightsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/partner/securityScore/customerInsights{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPartnerSecurityScoreCustomerInsightsRequestBuilder instantiates a new PartnerSecurityScoreCustomerInsightsRequestBuilder and sets the default values.
func NewPartnerSecurityScoreCustomerInsightsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreCustomerInsightsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnerSecurityScoreCustomerInsightsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PartnerSecurityScoreCustomerInsightsCountRequestBuilder when successful
func (m *PartnerSecurityScoreCustomerInsightsRequestBuilder) Count()(*PartnerSecurityScoreCustomerInsightsCountRequestBuilder) {
    return NewPartnerSecurityScoreCustomerInsightsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the customerInsight objects and their properties.
// returns a CustomerInsightCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/partner-security-partnersecurityscore-list-customerinsights?view=graph-rest-beta
func (m *PartnerSecurityScoreCustomerInsightsRequestBuilder) Get(ctx context.Context, requestConfiguration *PartnerSecurityScoreCustomerInsightsRequestBuilderGetRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CustomerInsightCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreateCustomerInsightCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CustomerInsightCollectionResponseable), nil
}
// Post create new navigation property to customerInsights for security
// returns a CustomerInsightable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerSecurityScoreCustomerInsightsRequestBuilder) Post(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CustomerInsightable, requestConfiguration *PartnerSecurityScoreCustomerInsightsRequestBuilderPostRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CustomerInsightable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreateCustomerInsightFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CustomerInsightable), nil
}
// ToGetRequestInformation get a list of the customerInsight objects and their properties.
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreCustomerInsightsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PartnerSecurityScoreCustomerInsightsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to customerInsights for security
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreCustomerInsightsRequestBuilder) ToPostRequestInformation(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CustomerInsightable, requestConfiguration *PartnerSecurityScoreCustomerInsightsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnerSecurityScoreCustomerInsightsRequestBuilder when successful
func (m *PartnerSecurityScoreCustomerInsightsRequestBuilder) WithUrl(rawUrl string)(*PartnerSecurityScoreCustomerInsightsRequestBuilder) {
    return NewPartnerSecurityScoreCustomerInsightsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

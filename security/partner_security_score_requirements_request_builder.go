package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partner/security"
)

// PartnerSecurityScoreRequirementsRequestBuilder provides operations to manage the requirements property of the microsoft.graph.partner.security.partnerSecurityScore entity.
type PartnerSecurityScoreRequirementsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnerSecurityScoreRequirementsRequestBuilderGetQueryParameters get a list of securityRequirement objects and their properties.
type PartnerSecurityScoreRequirementsRequestBuilderGetQueryParameters struct {
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
// PartnerSecurityScoreRequirementsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreRequirementsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PartnerSecurityScoreRequirementsRequestBuilderGetQueryParameters
}
// PartnerSecurityScoreRequirementsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreRequirementsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySecurityRequirementId provides operations to manage the requirements property of the microsoft.graph.partner.security.partnerSecurityScore entity.
// returns a *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder when successful
func (m *PartnerSecurityScoreRequirementsRequestBuilder) BySecurityRequirementId(securityRequirementId string)(*PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if securityRequirementId != "" {
        urlTplParams["securityRequirement%2Did"] = securityRequirementId
    }
    return NewPartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPartnerSecurityScoreRequirementsRequestBuilderInternal instantiates a new PartnerSecurityScoreRequirementsRequestBuilder and sets the default values.
func NewPartnerSecurityScoreRequirementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreRequirementsRequestBuilder) {
    m := &PartnerSecurityScoreRequirementsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/partner/securityScore/requirements{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPartnerSecurityScoreRequirementsRequestBuilder instantiates a new PartnerSecurityScoreRequirementsRequestBuilder and sets the default values.
func NewPartnerSecurityScoreRequirementsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreRequirementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnerSecurityScoreRequirementsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PartnerSecurityScoreRequirementsCountRequestBuilder when successful
func (m *PartnerSecurityScoreRequirementsRequestBuilder) Count()(*PartnerSecurityScoreRequirementsCountRequestBuilder) {
    return NewPartnerSecurityScoreRequirementsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of securityRequirement objects and their properties.
// returns a SecurityRequirementCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/partner-security-partnersecurityscore-list-requirements?view=graph-rest-beta
func (m *PartnerSecurityScoreRequirementsRequestBuilder) Get(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequirementsRequestBuilderGetRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreateSecurityRequirementCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementCollectionResponseable), nil
}
// Post create new navigation property to requirements for security
// returns a SecurityRequirementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerSecurityScoreRequirementsRequestBuilder) Post(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementable, requestConfiguration *PartnerSecurityScoreRequirementsRequestBuilderPostRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreateSecurityRequirementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementable), nil
}
// ToGetRequestInformation get a list of securityRequirement objects and their properties.
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreRequirementsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequirementsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to requirements for security
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreRequirementsRequestBuilder) ToPostRequestInformation(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementable, requestConfiguration *PartnerSecurityScoreRequirementsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnerSecurityScoreRequirementsRequestBuilder when successful
func (m *PartnerSecurityScoreRequirementsRequestBuilder) WithUrl(rawUrl string)(*PartnerSecurityScoreRequirementsRequestBuilder) {
    return NewPartnerSecurityScoreRequirementsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

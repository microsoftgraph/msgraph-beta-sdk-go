package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partner/security"
)

// PartnerSecurityScoreRequestBuilder provides operations to manage the securityScore property of the microsoft.graph.partner.security.partnerSecurity entity.
type PartnerSecurityScoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnerSecurityScoreRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PartnerSecurityScoreRequestBuilderGetQueryParameters read the properties and relationships of a partnerSecurityScore object.
type PartnerSecurityScoreRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PartnerSecurityScoreRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PartnerSecurityScoreRequestBuilderGetQueryParameters
}
// PartnerSecurityScoreRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnerSecurityScoreRequestBuilderInternal instantiates a new PartnerSecurityScoreRequestBuilder and sets the default values.
func NewPartnerSecurityScoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreRequestBuilder) {
    m := &PartnerSecurityScoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/partner/securityScore{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPartnerSecurityScoreRequestBuilder instantiates a new PartnerSecurityScoreRequestBuilder and sets the default values.
func NewPartnerSecurityScoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnerSecurityScoreRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomerInsights provides operations to manage the customerInsights property of the microsoft.graph.partner.security.partnerSecurityScore entity.
// returns a *PartnerSecurityScoreCustomerInsightsRequestBuilder when successful
func (m *PartnerSecurityScoreRequestBuilder) CustomerInsights()(*PartnerSecurityScoreCustomerInsightsRequestBuilder) {
    return NewPartnerSecurityScoreCustomerInsightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property securityScore for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerSecurityScoreRequestBuilder) Delete(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a partnerSecurityScore object.
// returns a PartnerSecurityScoreable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/partner-security-partnersecurityscore-get?view=graph-rest-beta
func (m *PartnerSecurityScoreRequestBuilder) Get(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequestBuilderGetRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityScoreable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreatePartnerSecurityScoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityScoreable), nil
}
// History provides operations to manage the history property of the microsoft.graph.partner.security.partnerSecurityScore entity.
// returns a *PartnerSecurityScoreHistoryRequestBuilder when successful
func (m *PartnerSecurityScoreRequestBuilder) History()(*PartnerSecurityScoreHistoryRequestBuilder) {
    return NewPartnerSecurityScoreHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property securityScore in security
// returns a PartnerSecurityScoreable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerSecurityScoreRequestBuilder) Patch(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityScoreable, requestConfiguration *PartnerSecurityScoreRequestBuilderPatchRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityScoreable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreatePartnerSecurityScoreFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityScoreable), nil
}
// Requirements provides operations to manage the requirements property of the microsoft.graph.partner.security.partnerSecurityScore entity.
// returns a *PartnerSecurityScoreRequirementsRequestBuilder when successful
func (m *PartnerSecurityScoreRequestBuilder) Requirements()(*PartnerSecurityScoreRequirementsRequestBuilder) {
    return NewPartnerSecurityScoreRequirementsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property securityScore for security
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a partnerSecurityScore object.
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property securityScore in security
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreRequestBuilder) ToPatchRequestInformation(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityScoreable, requestConfiguration *PartnerSecurityScoreRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnerSecurityScoreRequestBuilder when successful
func (m *PartnerSecurityScoreRequestBuilder) WithUrl(rawUrl string)(*PartnerSecurityScoreRequestBuilder) {
    return NewPartnerSecurityScoreRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partner/security"
)

// PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder provides operations to manage the history property of the microsoft.graph.partner.security.partnerSecurityScore entity.
type PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderGetQueryParameters read the properties and relationships of a securityScoreHistory object.
type PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderGetQueryParameters
}
// PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderInternal instantiates a new PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder and sets the default values.
func NewPartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) {
    m := &PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/partner/securityScore/history/{securityScoreHistory%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder instantiates a new PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder and sets the default values.
func NewPartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property history for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a securityScoreHistory object.
// returns a SecurityScoreHistoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/partner-security-securityscorehistory-get?view=graph-rest-beta
func (m *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderGetRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityScoreHistoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreateSecurityScoreHistoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityScoreHistoryable), nil
}
// Patch update the navigation property history in security
// returns a SecurityScoreHistoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) Patch(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityScoreHistoryable, requestConfiguration *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderPatchRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityScoreHistoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreateSecurityScoreHistoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityScoreHistoryable), nil
}
// ToDeleteRequestInformation delete navigation property history for security
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a securityScoreHistory object.
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property history in security
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityScoreHistoryable, requestConfiguration *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder when successful
func (m *PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) WithUrl(rawUrl string)(*PartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder) {
    return NewPartnerSecurityScoreHistorySecurityScoreHistoryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

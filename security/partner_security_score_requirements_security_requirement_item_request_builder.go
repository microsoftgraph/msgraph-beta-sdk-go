package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partner/security"
)

// PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder provides operations to manage the requirements property of the microsoft.graph.partner.security.partnerSecurityScore entity.
type PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderGetQueryParameters read the properties and relationships of a securityRequirement object.
type PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderGetQueryParameters
}
// PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderInternal instantiates a new PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder and sets the default values.
func NewPartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) {
    m := &PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/partner/securityScore/requirements/{securityRequirement%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder instantiates a new PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder and sets the default values.
func NewPartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property requirements for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a securityRequirement object.
// returns a SecurityRequirementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/partner-security-securityrequirement-get?view=graph-rest-beta
func (m *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderGetRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property requirements in security
// returns a SecurityRequirementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) Patch(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementable, requestConfiguration *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderPatchRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property requirements for security
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a securityRequirement object.
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property requirements in security
// returns a *RequestInformation when successful
func (m *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.SecurityRequirementable, requestConfiguration *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder when successful
func (m *PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) WithUrl(rawUrl string)(*PartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder) {
    return NewPartnerSecurityScoreRequirementsSecurityRequirementItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

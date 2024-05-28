package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/partner/security"
)

// PartnerRequestBuilder provides operations to manage the partner property of the microsoft.graph.security entity.
type PartnerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PartnerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PartnerRequestBuilderGetQueryParameters a container that safeguards the Microsoft Azure resources of Microsoft Cloud Solution Provider (CSP) partners’ customers, including alerts, scores, and all aspects of security.
type PartnerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PartnerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PartnerRequestBuilderGetQueryParameters
}
// PartnerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PartnerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPartnerRequestBuilderInternal instantiates a new PartnerRequestBuilder and sets the default values.
func NewPartnerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerRequestBuilder) {
    m := &PartnerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/partner{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPartnerRequestBuilder instantiates a new PartnerRequestBuilder and sets the default values.
func NewPartnerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PartnerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPartnerRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property partner for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerRequestBuilder) Delete(ctx context.Context, requestConfiguration *PartnerRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a container that safeguards the Microsoft Azure resources of Microsoft Cloud Solution Provider (CSP) partners’ customers, including alerts, scores, and all aspects of security.
// returns a PartnerSecurityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerRequestBuilder) Get(ctx context.Context, requestConfiguration *PartnerRequestBuilderGetRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreatePartnerSecurityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityable), nil
}
// Patch update the navigation property partner in security
// returns a PartnerSecurityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PartnerRequestBuilder) Patch(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityable, requestConfiguration *PartnerRequestBuilderPatchRequestConfiguration)(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.CreatePartnerSecurityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityable), nil
}
// SecurityAlerts provides operations to manage the securityAlerts property of the microsoft.graph.partner.security.partnerSecurity entity.
// returns a *PartnerSecurityalertsSecurityAlertsRequestBuilder when successful
func (m *PartnerRequestBuilder) SecurityAlerts()(*PartnerSecurityalertsSecurityAlertsRequestBuilder) {
    return NewPartnerSecurityalertsSecurityAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property partner for security
// returns a *RequestInformation when successful
func (m *PartnerRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PartnerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a container that safeguards the Microsoft Azure resources of Microsoft Cloud Solution Provider (CSP) partners’ customers, including alerts, scores, and all aspects of security.
// returns a *RequestInformation when successful
func (m *PartnerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PartnerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property partner in security
// returns a *RequestInformation when successful
func (m *PartnerRequestBuilder) ToPatchRequestInformation(ctx context.Context, body if063350b43a6e623648b394d8f6b8221a26a9ae46bc715ec386eaf438133b186.PartnerSecurityable, requestConfiguration *PartnerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PartnerRequestBuilder when successful
func (m *PartnerRequestBuilder) WithUrl(rawUrl string)(*PartnerRequestBuilder) {
    return NewPartnerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

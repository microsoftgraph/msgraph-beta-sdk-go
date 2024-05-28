package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder provides operations to manage the identitySecurityDefaultsEnforcementPolicy property of the microsoft.graph.policyRoot entity.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetQueryParameters retrieve the properties of an identitySecurityDefaultsEnforcementPolicy object.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetQueryParameters
}
// IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal instantiates a new IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder and sets the default values.
func NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    m := &IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/identitySecurityDefaultsEnforcementPolicy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder instantiates a new IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder and sets the default values.
func NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property identitySecurityDefaultsEnforcementPolicy for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties of an identitySecurityDefaultsEnforcementPolicy object.
// returns a IdentitySecurityDefaultsEnforcementPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitysecuritydefaultsenforcementpolicy-get?view=graph-rest-beta
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySecurityDefaultsEnforcementPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySecurityDefaultsEnforcementPolicyable), nil
}
// Patch update the properties of an identitySecurityDefaultsEnforcementPolicy object.
// returns a IdentitySecurityDefaultsEnforcementPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitysecuritydefaultsenforcementpolicy-update?view=graph-rest-beta
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySecurityDefaultsEnforcementPolicyable, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySecurityDefaultsEnforcementPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySecurityDefaultsEnforcementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySecurityDefaultsEnforcementPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property identitySecurityDefaultsEnforcementPolicy for policies
// returns a *RequestInformation when successful
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties of an identitySecurityDefaultsEnforcementPolicy object.
// returns a *RequestInformation when successful
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an identitySecurityDefaultsEnforcementPolicy object.
// returns a *RequestInformation when successful
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySecurityDefaultsEnforcementPolicyable, requestConfiguration *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder when successful
func (m *IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) WithUrl(rawUrl string)(*IdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return NewIdentitysecuritydefaultsenforcementpolicyIdentitySecurityDefaultsEnforcementPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

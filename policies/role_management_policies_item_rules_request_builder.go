package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementPoliciesItemRulesRequestBuilder provides operations to manage the rules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
type RoleManagementPoliciesItemRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleManagementPoliciesItemRulesRequestBuilderGetQueryParameters get the rules defined for a role management policy. The rules are a collection of following types that are derived from the unifiedRoleManagementPolicyRule object:+ unifiedRoleManagementPolicyApprovalRule+ unifiedRoleManagementPolicyAuthenticationContextRule+ unifiedRoleManagementPolicyEnablementRule+ unifiedRoleManagementPolicyExpirationRule+ unifiedRoleManagementPolicyNotificationRule To retrieve rules for a policy that applies to Azure RBAC, use the Azure REST PIM API for role management policies. This API is available in the following national cloud deployments.
type RoleManagementPoliciesItemRulesRequestBuilderGetQueryParameters struct {
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
// RoleManagementPoliciesItemRulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementPoliciesItemRulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementPoliciesItemRulesRequestBuilderGetQueryParameters
}
// RoleManagementPoliciesItemRulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementPoliciesItemRulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleManagementPolicyRuleId provides operations to manage the rules property of the microsoft.graph.unifiedRoleManagementPolicy entity.
func (m *RoleManagementPoliciesItemRulesRequestBuilder) ByUnifiedRoleManagementPolicyRuleId(unifiedRoleManagementPolicyRuleId string)(*RoleManagementPoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleManagementPolicyRuleId != "" {
        urlTplParams["unifiedRoleManagementPolicyRule%2Did"] = unifiedRoleManagementPolicyRuleId
    }
    return NewRoleManagementPoliciesItemRulesUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRoleManagementPoliciesItemRulesRequestBuilderInternal instantiates a new RulesRequestBuilder and sets the default values.
func NewRoleManagementPoliciesItemRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementPoliciesItemRulesRequestBuilder) {
    m := &RoleManagementPoliciesItemRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy%2Did}/rules{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewRoleManagementPoliciesItemRulesRequestBuilder instantiates a new RulesRequestBuilder and sets the default values.
func NewRoleManagementPoliciesItemRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementPoliciesItemRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementPoliciesItemRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *RoleManagementPoliciesItemRulesRequestBuilder) Count()(*RoleManagementPoliciesItemRulesCountRequestBuilder) {
    return NewRoleManagementPoliciesItemRulesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the rules defined for a role management policy. The rules are a collection of following types that are derived from the unifiedRoleManagementPolicyRule object:+ unifiedRoleManagementPolicyApprovalRule+ unifiedRoleManagementPolicyAuthenticationContextRule+ unifiedRoleManagementPolicyEnablementRule+ unifiedRoleManagementPolicyExpirationRule+ unifiedRoleManagementPolicyNotificationRule To retrieve rules for a policy that applies to Azure RBAC, use the Azure REST PIM API for role management policies. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedrolemanagementpolicy-list-rules?view=graph-rest-1.0
func (m *RoleManagementPoliciesItemRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementPoliciesItemRulesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyRuleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementPolicyRuleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyRuleCollectionResponseable), nil
}
// Post create new navigation property to rules for policies
func (m *RoleManagementPoliciesItemRulesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyRuleable, requestConfiguration *RoleManagementPoliciesItemRulesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementPolicyRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyRuleable), nil
}
// ToGetRequestInformation get the rules defined for a role management policy. The rules are a collection of following types that are derived from the unifiedRoleManagementPolicyRule object:+ unifiedRoleManagementPolicyApprovalRule+ unifiedRoleManagementPolicyAuthenticationContextRule+ unifiedRoleManagementPolicyEnablementRule+ unifiedRoleManagementPolicyExpirationRule+ unifiedRoleManagementPolicyNotificationRule To retrieve rules for a policy that applies to Azure RBAC, use the Azure REST PIM API for role management policies. This API is available in the following national cloud deployments.
func (m *RoleManagementPoliciesItemRulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementPoliciesItemRulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to rules for policies
func (m *RoleManagementPoliciesItemRulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyRuleable, requestConfiguration *RoleManagementPoliciesItemRulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *RoleManagementPoliciesItemRulesRequestBuilder) WithUrl(rawUrl string)(*RoleManagementPoliciesItemRulesRequestBuilder) {
    return NewRoleManagementPoliciesItemRulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

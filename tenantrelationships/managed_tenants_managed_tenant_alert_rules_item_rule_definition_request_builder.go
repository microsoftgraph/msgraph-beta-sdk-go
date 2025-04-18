// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder provides operations to manage the ruleDefinition property of the microsoft.graph.managedTenants.managedTenantAlertRule entity.
type ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderGetQueryParameters get ruleDefinition from tenantRelationships
type ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderGetQueryParameters
}
// NewManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderInternal instantiates a new ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder) {
    m := &ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlertRules/{managedTenantAlertRule%2Did}/ruleDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder instantiates a new ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get ruleDefinition from tenantRelationships
// returns a ManagedTenantAlertRuleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertRuleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable), nil
}
// ToGetRequestInformation get ruleDefinition from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder when successful
func (m *ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

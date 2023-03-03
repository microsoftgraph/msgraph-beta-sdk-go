package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder provides operations to manage the managedTenantAlertRules property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderGetQueryParameters get managedTenantAlertRules from tenantRelationships
type ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderGetQueryParameters
}
// ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Alerts provides operations to manage the alerts property of the microsoft.graph.managedTenants.managedTenantAlertRule entity.
func (m *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) Alerts()(*ManagedTenantsManagedTenantAlertRulesItemAlertsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertRulesItemAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AlertsById provides operations to manage the alerts property of the microsoft.graph.managedTenants.managedTenantAlertRule entity.
func (m *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) AlertsById(id string)(*ManagedTenantsManagedTenantAlertRulesItemAlertsManagedTenantAlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlert%2Did"] = id
    }
    return NewManagedTenantsManagedTenantAlertRulesItemAlertsManagedTenantAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderInternal instantiates a new ManagedTenantAlertRuleItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) {
    m := &ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlertRules/{managedTenantAlertRule%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder instantiates a new ManagedTenantAlertRuleItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedTenantAlertRules for tenantRelationships
func (m *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get managedTenantAlertRules from tenantRelationships
func (m *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleable), nil
}
// Patch update the navigation property managedTenantAlertRules in tenantRelationships
func (m *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleable, requestConfiguration *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleable), nil
}
// RuleDefinition provides operations to manage the ruleDefinition property of the microsoft.graph.managedTenants.managedTenantAlertRule entity.
func (m *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) RuleDefinition()(*ManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertRulesItemRuleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedTenantAlertRules for tenantRelationships
func (m *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get managedTenantAlertRules from tenantRelationships
func (m *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property managedTenantAlertRules in tenantRelationships
func (m *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleable, requestConfiguration *ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

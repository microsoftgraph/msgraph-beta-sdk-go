package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetQueryParameters get managedTenantAlerts from tenantRelationships
type ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetQueryParameters
}
// ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertLogs provides operations to manage the alertLogs property of the microsoft.graph.managedTenants.managedTenantAlert entity.
// returns a *ManagedTenantsManagedTenantAlertsItemAlertLogsRequestBuilder when successful
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) AlertLogs()(*ManagedTenantsManagedTenantAlertsItemAlertLogsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemAlertLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AlertRule provides operations to manage the alertRule property of the microsoft.graph.managedTenants.managedTenantAlert entity.
// returns a *ManagedTenantsManagedTenantAlertsItemAlertRuleRequestBuilder when successful
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) AlertRule()(*ManagedTenantsManagedTenantAlertsItemAlertRuleRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemAlertRuleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApiNotifications provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
// returns a *ManagedTenantsManagedTenantAlertsItemApiNotificationsRequestBuilder when successful
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ApiNotifications()(*ManagedTenantsManagedTenantAlertsItemApiNotificationsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemApiNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal instantiates a new ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) {
    m := &ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder instantiates a new ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedTenantAlerts for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// EmailNotifications provides operations to manage the emailNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
// returns a *ManagedTenantsManagedTenantAlertsItemEmailNotificationsRequestBuilder when successful
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) EmailNotifications()(*ManagedTenantsManagedTenantAlertsItemEmailNotificationsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemEmailNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get managedTenantAlerts from tenantRelationships
// returns a ManagedTenantAlertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable), nil
}
// MicrosoftGraphManagedTenantsAddUserInputLog provides operations to call the addUserInputLog method.
// returns a *ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder when successful
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) MicrosoftGraphManagedTenantsAddUserInputLog()(*ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedTenantAlerts in tenantRelationships
// returns a ManagedTenantAlertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable), nil
}
// ToDeleteRequestInformation delete navigation property managedTenantAlerts for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get managedTenantAlerts from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedTenantAlerts in tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder when successful
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) AlertLogs()(*ManagedTenantsManagedTenantAlertsItemAlertLogsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemAlertLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AlertLogsById provides operations to manage the alertLogs property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) AlertLogsById(id string)(*ManagedTenantsManagedTenantAlertsItemAlertLogsManagedTenantAlertLogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertLog%2Did"] = id
    }
    return NewManagedTenantsManagedTenantAlertsItemAlertLogsManagedTenantAlertLogItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AlertRule provides operations to manage the alertRule property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) AlertRule()(*ManagedTenantsManagedTenantAlertsItemAlertRuleRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemAlertRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApiNotifications provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ApiNotifications()(*ManagedTenantsManagedTenantAlertsItemApiNotificationsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemApiNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApiNotificationsById provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ApiNotificationsById(id string)(*ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantApiNotification%2Did"] = id
    }
    return NewManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal instantiates a new ManagedTenantAlertItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) {
    m := &ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder instantiates a new ManagedTenantAlertItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedTenantAlerts for tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// EmailNotifications provides operations to manage the emailNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) EmailNotifications()(*ManagedTenantsManagedTenantAlertsItemEmailNotificationsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemEmailNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EmailNotificationsById provides operations to manage the emailNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) EmailNotificationsById(id string)(*ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantEmailNotification%2Did"] = id
    }
    return NewManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get managedTenantAlerts from tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable), nil
}
// MicrosoftGraphManagedTenantsAddUserInputLog provides operations to call the addUserInputLog method.
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) MicrosoftGraphManagedTenantsAddUserInputLog()(*ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property managedTenantAlerts in tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable), nil
}
// ToDeleteRequestInformation delete navigation property managedTenantAlerts for tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get managedTenantAlerts from tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedTenantAlerts in tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, requestConfiguration *ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

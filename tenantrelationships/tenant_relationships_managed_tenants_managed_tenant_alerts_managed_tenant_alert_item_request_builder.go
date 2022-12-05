package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetQueryParameters get managedTenantAlerts from tenantRelationships
type TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddUserInputLog provides operations to call the addUserInputLog method.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) AddUserInputLog()(*TenantRelationshipsManagedTenantsManagedTenantAlertsItemAddUserInputLogRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsItemAddUserInputLogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertLogs provides operations to manage the alertLogs property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) AlertLogs()(*TenantRelationshipsManagedTenantsManagedTenantAlertsItemAlertLogsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsItemAlertLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertLogsById provides operations to manage the alertLogs property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) AlertLogsById(id string)(*TenantRelationshipsManagedTenantsManagedTenantAlertsItemAlertLogsManagedTenantAlertLogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertLog%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsItemAlertLogsManagedTenantAlertLogItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AlertRule provides operations to manage the alertRule property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) AlertRule()(*TenantRelationshipsManagedTenantsManagedTenantAlertsItemAlertRuleRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsItemAlertRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApiNotifications provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ApiNotifications()(*TenantRelationshipsManagedTenantsManagedTenantAlertsItemApiNotificationsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsItemApiNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApiNotificationsById provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) ApiNotificationsById(id string)(*TenantRelationshipsManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantApiNotification%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal instantiates a new ManagedTenantAlertItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder instantiates a new ManagedTenantAlertItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedTenantAlerts for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get managedTenantAlerts from tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property managedTenantAlerts in tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property managedTenantAlerts for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EmailNotifications provides operations to manage the emailNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) EmailNotifications()(*TenantRelationshipsManagedTenantsManagedTenantAlertsItemEmailNotificationsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsItemEmailNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailNotificationsById provides operations to manage the emailNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) EmailNotificationsById(id string)(*TenantRelationshipsManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantEmailNotification%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get managedTenantAlerts from tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable), nil
}
// Patch update the navigation property managedTenantAlerts in tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable), nil
}

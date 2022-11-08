package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
    i171a62c77cd8ebd0854e4fee5c43c3cca55e79bef9ab9f674eeb4dfafae99604 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts/item/adduserinputlog"
    i571d9b61ab93d9226d395e33aecf0e06edac249ad78d4fb79d84b85c8f90c266 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts/item/alertrule"
    i60e373ca5c04e149e8654932f54488d564794d69b9651981fde072c0212412be "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts/item/apinotifications"
    ice5d0a0c9f28f0658e650420c1a7d47baabe48282bf796f55df19e6b657622c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts/item/alertlogs"
    if129da66a0e17df3a9b75cd67958633a9890cb813547ef2ef3fdd7ab713b7458 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts/item/emailnotifications"
    i8f92ba0e5ef437ca43460fb8e680765aab203d4827686fdba9c37d3c0b84636c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts/item/apinotifications/item"
    ia0f01951766d8ccc820cd97f356272fa1a2df144cf5e32890e27bc99d35d0167 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts/item/emailnotifications/item"
    ic2e4e4f0c5637f5037a0ef064cabd45dfebe1dfe1b7dcef00bdabe0ec09a33f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts/item/alertlogs/item"
)

// ManagedTenantAlertItemRequestBuilder provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantAlertItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantAlertItemRequestBuilderGetQueryParameters get managedTenantAlerts from tenantRelationships
type ManagedTenantAlertItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantAlertItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantAlertItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantAlertItemRequestBuilderGetQueryParameters
}
// ManagedTenantAlertItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantAlertItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddUserInputLog provides operations to call the addUserInputLog method.
func (m *ManagedTenantAlertItemRequestBuilder) AddUserInputLog()(*i171a62c77cd8ebd0854e4fee5c43c3cca55e79bef9ab9f674eeb4dfafae99604.AddUserInputLogRequestBuilder) {
    return i171a62c77cd8ebd0854e4fee5c43c3cca55e79bef9ab9f674eeb4dfafae99604.NewAddUserInputLogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertLogs provides operations to manage the alertLogs property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantAlertItemRequestBuilder) AlertLogs()(*ice5d0a0c9f28f0658e650420c1a7d47baabe48282bf796f55df19e6b657622c1.AlertLogsRequestBuilder) {
    return ice5d0a0c9f28f0658e650420c1a7d47baabe48282bf796f55df19e6b657622c1.NewAlertLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertLogsById provides operations to manage the alertLogs property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantAlertItemRequestBuilder) AlertLogsById(id string)(*ic2e4e4f0c5637f5037a0ef064cabd45dfebe1dfe1b7dcef00bdabe0ec09a33f0.ManagedTenantAlertLogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertLog%2Did"] = id
    }
    return ic2e4e4f0c5637f5037a0ef064cabd45dfebe1dfe1b7dcef00bdabe0ec09a33f0.NewManagedTenantAlertLogItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AlertRule provides operations to manage the alertRule property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantAlertItemRequestBuilder) AlertRule()(*i571d9b61ab93d9226d395e33aecf0e06edac249ad78d4fb79d84b85c8f90c266.AlertRuleRequestBuilder) {
    return i571d9b61ab93d9226d395e33aecf0e06edac249ad78d4fb79d84b85c8f90c266.NewAlertRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApiNotifications provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantAlertItemRequestBuilder) ApiNotifications()(*i60e373ca5c04e149e8654932f54488d564794d69b9651981fde072c0212412be.ApiNotificationsRequestBuilder) {
    return i60e373ca5c04e149e8654932f54488d564794d69b9651981fde072c0212412be.NewApiNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApiNotificationsById provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantAlertItemRequestBuilder) ApiNotificationsById(id string)(*i8f92ba0e5ef437ca43460fb8e680765aab203d4827686fdba9c37d3c0b84636c.ManagedTenantApiNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantApiNotification%2Did"] = id
    }
    return i8f92ba0e5ef437ca43460fb8e680765aab203d4827686fdba9c37d3c0b84636c.NewManagedTenantApiNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedTenantAlertItemRequestBuilderInternal instantiates a new ManagedTenantAlertItemRequestBuilder and sets the default values.
func NewManagedTenantAlertItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantAlertItemRequestBuilder) {
    m := &ManagedTenantAlertItemRequestBuilder{
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
// NewManagedTenantAlertItemRequestBuilder instantiates a new ManagedTenantAlertItemRequestBuilder and sets the default values.
func NewManagedTenantAlertItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantAlertItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantAlertItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedTenantAlerts for tenantRelationships
func (m *ManagedTenantAlertItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantAlertItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantAlertItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantAlertItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, requestConfiguration *ManagedTenantAlertItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantAlertItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantAlertItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ManagedTenantAlertItemRequestBuilder) EmailNotifications()(*if129da66a0e17df3a9b75cd67958633a9890cb813547ef2ef3fdd7ab713b7458.EmailNotificationsRequestBuilder) {
    return if129da66a0e17df3a9b75cd67958633a9890cb813547ef2ef3fdd7ab713b7458.NewEmailNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailNotificationsById provides operations to manage the emailNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
func (m *ManagedTenantAlertItemRequestBuilder) EmailNotificationsById(id string)(*ia0f01951766d8ccc820cd97f356272fa1a2df144cf5e32890e27bc99d35d0167.ManagedTenantEmailNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantEmailNotification%2Did"] = id
    }
    return ia0f01951766d8ccc820cd97f356272fa1a2df144cf5e32890e27bc99d35d0167.NewManagedTenantEmailNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get managedTenantAlerts from tenantRelationships
func (m *ManagedTenantAlertItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantAlertItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
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
func (m *ManagedTenantAlertItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, requestConfiguration *ManagedTenantAlertItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
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

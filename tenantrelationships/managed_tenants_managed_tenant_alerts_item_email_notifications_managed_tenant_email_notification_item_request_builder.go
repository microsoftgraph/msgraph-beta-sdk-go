package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder provides operations to manage the emailNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
type ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetQueryParameters get emailNotifications from tenantRelationships
type ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetQueryParameters
}
// NewManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal instantiates a new ManagedTenantEmailNotificationItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    m := &ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}/emailNotifications/{managedTenantEmailNotification%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder instantiates a new ManagedTenantEmailNotificationItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get emailNotifications from tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantEmailNotificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable), nil
}
// ToGetRequestInformation get emailNotifications from tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

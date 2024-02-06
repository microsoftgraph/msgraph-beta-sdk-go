package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
type ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderGetQueryParameters get apiNotifications from tenantRelationships
type ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderGetQueryParameters
}
// NewManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderInternal instantiates a new ManagedTenantApiNotificationItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder) {
    m := &ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}/apiNotifications/{managedTenantApiNotification%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder instantiates a new ManagedTenantApiNotificationItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get apiNotifications from tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantApiNotificationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantApiNotificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantApiNotificationable), nil
}
// ToGetRequestInformation get apiNotifications from tenantRelationships
func (m *ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemApiNotificationsManagedTenantApiNotificationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

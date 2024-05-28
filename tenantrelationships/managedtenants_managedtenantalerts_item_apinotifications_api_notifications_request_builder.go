package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
type ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderGetQueryParameters get apiNotifications from tenantRelationships
type ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderGetQueryParameters
}
// ByManagedTenantApiNotificationId provides operations to manage the apiNotifications property of the microsoft.graph.managedTenants.managedTenantAlert entity.
// returns a *ManagedtenantsManagedtenantalertsItemApinotificationsManagedTenantApiNotificationItemRequestBuilder when successful
func (m *ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder) ByManagedTenantApiNotificationId(managedTenantApiNotificationId string)(*ManagedtenantsManagedtenantalertsItemApinotificationsManagedTenantApiNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedTenantApiNotificationId != "" {
        urlTplParams["managedTenantApiNotification%2Did"] = managedTenantApiNotificationId
    }
    return NewManagedtenantsManagedtenantalertsItemApinotificationsManagedTenantApiNotificationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderInternal instantiates a new ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder and sets the default values.
func NewManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder) {
    m := &ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}/apiNotifications{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder instantiates a new ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder and sets the default values.
func NewManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsManagedtenantalertsItemApinotificationsCountRequestBuilder when successful
func (m *ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder) Count()(*ManagedtenantsManagedtenantalertsItemApinotificationsCountRequestBuilder) {
    return NewManagedtenantsManagedtenantalertsItemApinotificationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get apiNotifications from tenantRelationships
// returns a ManagedTenantApiNotificationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantApiNotificationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantApiNotificationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantApiNotificationCollectionResponseable), nil
}
// ToGetRequestInformation get apiNotifications from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder when successful
func (m *ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder) {
    return NewManagedtenantsManagedtenantalertsItemApinotificationsApiNotificationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

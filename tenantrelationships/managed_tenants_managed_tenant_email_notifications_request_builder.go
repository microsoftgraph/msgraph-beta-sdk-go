package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagedTenantEmailNotificationsRequestBuilder provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsManagedTenantEmailNotificationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagedTenantEmailNotificationsRequestBuilderGetQueryParameters get managedTenantEmailNotifications from tenantRelationships
type ManagedTenantsManagedTenantEmailNotificationsRequestBuilderGetQueryParameters struct {
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
// ManagedTenantsManagedTenantEmailNotificationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantEmailNotificationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagedTenantEmailNotificationsRequestBuilderGetQueryParameters
}
// ManagedTenantsManagedTenantEmailNotificationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantEmailNotificationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedTenantEmailNotificationId provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder when successful
func (m *ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) ByManagedTenantEmailNotificationId(managedTenantEmailNotificationId string)(*ManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedTenantEmailNotificationId != "" {
        urlTplParams["managedTenantEmailNotification%2Did"] = managedTenantEmailNotificationId
    }
    return NewManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedTenantsManagedTenantEmailNotificationsRequestBuilderInternal instantiates a new ManagedTenantsManagedTenantEmailNotificationsRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantEmailNotificationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) {
    m := &ManagedTenantsManagedTenantEmailNotificationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantEmailNotifications{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagedTenantEmailNotificationsRequestBuilder instantiates a new ManagedTenantsManagedTenantEmailNotificationsRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantEmailNotificationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagedTenantEmailNotificationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedTenantsManagedTenantEmailNotificationsCountRequestBuilder when successful
func (m *ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) Count()(*ManagedTenantsManagedTenantEmailNotificationsCountRequestBuilder) {
    return NewManagedTenantsManagedTenantEmailNotificationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get managedTenantEmailNotifications from tenantRelationships
// returns a ManagedTenantEmailNotificationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantEmailNotificationsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantEmailNotificationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationCollectionResponseable), nil
}
// Post create new navigation property to managedTenantEmailNotifications for tenantRelationships
// returns a ManagedTenantEmailNotificationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable, requestConfiguration *ManagedTenantsManagedTenantEmailNotificationsRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation get managedTenantEmailNotifications from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantEmailNotificationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to managedTenantEmailNotifications for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable, requestConfiguration *ManagedTenantsManagedTenantEmailNotificationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/tenantRelationships/managedTenants/managedTenantEmailNotifications", m.BaseRequestBuilder.PathParameters)
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
// returns a *ManagedTenantsManagedTenantEmailNotificationsRequestBuilder when successful
func (m *ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) {
    return NewManagedTenantsManagedTenantEmailNotificationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

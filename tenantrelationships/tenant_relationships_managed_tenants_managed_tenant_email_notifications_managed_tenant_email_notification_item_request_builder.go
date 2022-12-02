package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetQueryParameters get managedTenantEmailNotifications from tenantRelationships
type TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Alert provides operations to manage the alert property of the microsoft.graph.managedTenants.managedTenantEmailNotification entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) Alert()(*TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsItemAlertRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantEmailNotificationsItemAlertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal instantiates a new ManagedTenantEmailNotificationItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managedTenantEmailNotifications/{managedTenantEmailNotification%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder instantiates a new ManagedTenantEmailNotificationItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedTenantEmailNotifications for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get managedTenantEmailNotifications from tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedTenantEmailNotifications in tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property managedTenantEmailNotifications for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get managedTenantEmailNotifications from tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantEmailNotificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable), nil
}
// Patch update the navigation property managedTenantEmailNotifications in tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantEmailNotificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantEmailNotificationable), nil
}

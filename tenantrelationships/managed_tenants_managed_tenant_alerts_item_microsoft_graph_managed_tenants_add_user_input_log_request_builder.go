package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder provides operations to call the addUserInputLog method.
type ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderInternal instantiates a new MicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) {
    m := &ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}/microsoft.graph.managedTenants.addUserInputLog", pathParameters),
    }
    return m
}
// NewManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder instantiates a new MicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addUserInputLog
func (m *ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) Post(ctx context.Context, body ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogAddUserInputLogPostRequestBodyable, requestConfiguration *ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToPostRequestInformation invoke action addUserInputLog
func (m *ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogAddUserInputLogPostRequestBodyable, requestConfiguration *ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsItemMicrosoftGraphManagedTenantsAddUserInputLogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

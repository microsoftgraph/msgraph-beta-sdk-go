package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder provides operations to manage the alertDefinition property of the microsoft.graph.unifiedRoleManagementAlert entity.
type RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderGetQueryParameters contains the description, impact, and measures to mitigate or prevent the security alert from being triggered in your tenant. Supports $expand.
type RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderGetQueryParameters
}
// NewRoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderInternal instantiates a new AlertDefinitionRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder) {
    m := &RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/alertDefinition{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewRoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder instantiates a new AlertDefinitionRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get contains the description, impact, and measures to mitigate or prevent the security alert from being triggered in your tenant. Supports $expand.
func (m *RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertDefinitionable), nil
}
// ToGetRequestInformation contains the description, impact, and measures to mitigate or prevent the security alert from being triggered in your tenant. Supports $expand.
func (m *RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder) WithUrl(rawUrl string)(*RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder) {
    return NewRoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

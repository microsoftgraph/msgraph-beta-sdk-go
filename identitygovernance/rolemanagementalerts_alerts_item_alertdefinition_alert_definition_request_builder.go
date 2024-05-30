package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder provides operations to manage the alertDefinition property of the microsoft.graph.unifiedRoleManagementAlert entity.
type RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderGetQueryParameters contains the description, impact, and measures to mitigate or prevent the security alert from being triggered in your tenant. Supports $expand.
type RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderGetQueryParameters
}
// NewRolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderInternal instantiates a new RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder) {
    m := &RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/alertDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder instantiates a new RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get contains the description, impact, and measures to mitigate or prevent the security alert from being triggered in your tenant. Supports $expand.
// returns a UnifiedRoleManagementAlertDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder when successful
func (m *RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder) WithUrl(rawUrl string)(*RolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder) {
    return NewRolemanagementalertsAlertsItemAlertdefinitionAlertDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

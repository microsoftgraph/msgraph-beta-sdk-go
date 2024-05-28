package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolemanagementalertsRoleManagementAlertsRequestBuilder provides operations to manage the roleManagementAlerts property of the microsoft.graph.identityGovernance entity.
type RolemanagementalertsRoleManagementAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementalertsRoleManagementAlertsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsRoleManagementAlertsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RolemanagementalertsRoleManagementAlertsRequestBuilderGetQueryParameters get roleManagementAlerts from identityGovernance
type RolemanagementalertsRoleManagementAlertsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementalertsRoleManagementAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsRoleManagementAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementalertsRoleManagementAlertsRequestBuilderGetQueryParameters
}
// RolemanagementalertsRoleManagementAlertsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsRoleManagementAlertsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertConfigurations provides operations to manage the alertConfigurations property of the microsoft.graph.roleManagementAlert entity.
// returns a *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder when successful
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) AlertConfigurations()(*RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) {
    return NewRolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AlertDefinitions provides operations to manage the alertDefinitions property of the microsoft.graph.roleManagementAlert entity.
// returns a *RolemanagementalertsAlertdefinitionsAlertDefinitionsRequestBuilder when successful
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) AlertDefinitions()(*RolemanagementalertsAlertdefinitionsAlertDefinitionsRequestBuilder) {
    return NewRolemanagementalertsAlertdefinitionsAlertDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Alerts provides operations to manage the alerts property of the microsoft.graph.roleManagementAlert entity.
// returns a *RolemanagementalertsAlertsRequestBuilder when successful
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) Alerts()(*RolemanagementalertsAlertsRequestBuilder) {
    return NewRolemanagementalertsAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRolemanagementalertsRoleManagementAlertsRequestBuilderInternal instantiates a new RolemanagementalertsRoleManagementAlertsRequestBuilder and sets the default values.
func NewRolemanagementalertsRoleManagementAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsRoleManagementAlertsRequestBuilder) {
    m := &RolemanagementalertsRoleManagementAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementalertsRoleManagementAlertsRequestBuilder instantiates a new RolemanagementalertsRoleManagementAlertsRequestBuilder and sets the default values.
func NewRolemanagementalertsRoleManagementAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsRoleManagementAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementalertsRoleManagementAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleManagementAlerts for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) Delete(ctx context.Context, requestConfiguration *RolemanagementalertsRoleManagementAlertsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get roleManagementAlerts from identityGovernance
// returns a RoleManagementAlertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementalertsRoleManagementAlertsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleManagementAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoleManagementAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleManagementAlertable), nil
}
// Operations provides operations to manage the operations property of the microsoft.graph.roleManagementAlert entity.
// returns a *RolemanagementalertsOperationsRequestBuilder when successful
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) Operations()(*RolemanagementalertsOperationsRequestBuilder) {
    return NewRolemanagementalertsOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property roleManagementAlerts in identityGovernance
// returns a RoleManagementAlertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleManagementAlertable, requestConfiguration *RolemanagementalertsRoleManagementAlertsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleManagementAlertable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoleManagementAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleManagementAlertable), nil
}
// ToDeleteRequestInformation delete navigation property roleManagementAlerts for identityGovernance
// returns a *RequestInformation when successful
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsRoleManagementAlertsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get roleManagementAlerts from identityGovernance
// returns a *RequestInformation when successful
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsRoleManagementAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleManagementAlerts in identityGovernance
// returns a *RequestInformation when successful
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleManagementAlertable, requestConfiguration *RolemanagementalertsRoleManagementAlertsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *RolemanagementalertsRoleManagementAlertsRequestBuilder when successful
func (m *RolemanagementalertsRoleManagementAlertsRequestBuilder) WithUrl(rawUrl string)(*RolemanagementalertsRoleManagementAlertsRequestBuilder) {
    return NewRolemanagementalertsRoleManagementAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

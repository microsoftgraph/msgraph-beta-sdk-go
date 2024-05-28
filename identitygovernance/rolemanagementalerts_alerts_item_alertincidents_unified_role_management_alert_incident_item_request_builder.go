package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder provides operations to manage the alertIncidents property of the microsoft.graph.unifiedRoleManagementAlert entity.
type RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetQueryParameters represents the incidents of this type of alert that have been triggered in Privileged Identity Management (PIM) for Microsoft Entra roles in the tenant. Supports $expand.
type RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetQueryParameters
}
// RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderInternal instantiates a new RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) {
    m := &RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/alertIncidents/{unifiedRoleManagementAlertIncident%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder instantiates a new RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property alertIncidents for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the incidents of this type of alert that have been triggered in Privileged Identity Management (PIM) for Microsoft Entra roles in the tenant. Supports $expand.
// returns a UnifiedRoleManagementAlertIncidentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertIncidentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable), nil
}
// Patch update the navigation property alertIncidents in identityGovernance
// returns a UnifiedRoleManagementAlertIncidentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertIncidentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable), nil
}
// Remediate provides operations to call the remediate method.
// returns a *RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) Remediate()(*RolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilder) {
    return NewRolemanagementalertsAlertsItemAlertincidentsItemRemediateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property alertIncidents for identityGovernance
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the incidents of this type of alert that have been triggered in Privileged Identity Management (PIM) for Microsoft Entra roles in the tenant. Supports $expand.
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property alertIncidents in identityGovernance
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) WithUrl(rawUrl string)(*RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) {
    return NewRolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

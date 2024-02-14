package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder provides operations to manage the alertIncidents property of the microsoft.graph.unifiedRoleManagementAlert entity.
type RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetQueryParameters read the properties and relationships of an alert incident. The alert incident can be one of the following types that are derived from the unifiedRoleManagementAlertIncident object:
type RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetQueryParameters
}
// RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderInternal instantiates a new RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) {
    m := &RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/alertIncidents/{unifiedRoleManagementAlertIncident%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder instantiates a new RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property alertIncidents for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an alert incident. The alert incident can be one of the following types that are derived from the unifiedRoleManagementAlertIncident object:
// returns a UnifiedRoleManagementAlertIncidentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedrolemanagementalertincident-get?view=graph-rest-1.0
func (m *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, error) {
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
func (m *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, requestConfiguration *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, error) {
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
// returns a *RoleManagementAlertsAlertsItemAlertIncidentsItemRemediateRequestBuilder when successful
func (m *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) Remediate()(*RoleManagementAlertsAlertsItemAlertIncidentsItemRemediateRequestBuilder) {
    return NewRoleManagementAlertsAlertsItemAlertIncidentsItemRemediateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property alertIncidents for identityGovernance
// returns a *RequestInformation when successful
func (m *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/alertIncidents/{unifiedRoleManagementAlertIncident%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an alert incident. The alert incident can be one of the following types that are derived from the unifiedRoleManagementAlertIncident object:
// returns a *RequestInformation when successful
func (m *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, requestConfiguration *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/alertIncidents/{unifiedRoleManagementAlertIncident%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder when successful
func (m *RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) WithUrl(rawUrl string)(*RoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) {
    return NewRoleManagementAlertsAlertsItemAlertIncidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

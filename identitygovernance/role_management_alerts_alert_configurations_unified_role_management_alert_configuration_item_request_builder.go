package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder provides operations to manage the alertConfigurations property of the microsoft.graph.roleManagementAlert entity.
type RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetQueryParameters get alertConfigurations from identityGovernance
type RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetQueryParameters
}
// RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertDefinition provides operations to manage the alertDefinition property of the microsoft.graph.unifiedRoleManagementAlertConfiguration entity.
func (m *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) AlertDefinition()(*RoleManagementAlertsAlertConfigurationsItemAlertDefinitionRequestBuilder) {
    return NewRoleManagementAlertsAlertConfigurationsItemAlertDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderInternal instantiates a new UnifiedRoleManagementAlertConfigurationItemRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) {
    m := &RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alertConfigurations/{unifiedRoleManagementAlertConfiguration%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewRoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder instantiates a new UnifiedRoleManagementAlertConfigurationItemRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property alertConfigurations for identityGovernance
func (m *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Get get alertConfigurations from identityGovernance
func (m *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable), nil
}
// Patch update the navigation property alertConfigurations in identityGovernance
func (m *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, requestConfiguration *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property alertConfigurations for identityGovernance
func (m *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get alertConfigurations from identityGovernance
func (m *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property alertConfigurations in identityGovernance
func (m *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, requestConfiguration *RoleManagementAlertsAlertConfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

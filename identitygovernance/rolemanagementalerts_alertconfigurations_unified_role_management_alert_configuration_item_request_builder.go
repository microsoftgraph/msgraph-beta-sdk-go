package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder provides operations to manage the alertConfigurations property of the microsoft.graph.roleManagementAlert entity.
type RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetQueryParameters the various configurations of an alert for Microsoft Entra roles. The configurations are predefined and can't be created or deleted, but some of the configurations can be modified.
type RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetQueryParameters
}
// RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertDefinition provides operations to manage the alertDefinition property of the microsoft.graph.unifiedRoleManagementAlertConfiguration entity.
// returns a *RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder when successful
func (m *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) AlertDefinition()(*RolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilder) {
    return NewRolemanagementalertsAlertconfigurationsItemAlertdefinitionAlertDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderInternal instantiates a new RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) {
    m := &RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alertConfigurations/{unifiedRoleManagementAlertConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder instantiates a new RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property alertConfigurations for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the various configurations of an alert for Microsoft Entra roles. The configurations are predefined and can't be created or deleted, but some of the configurations can be modified.
// returns a UnifiedRoleManagementAlertConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a UnifiedRoleManagementAlertConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, requestConfiguration *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the various configurations of an alert for Microsoft Entra roles. The configurations are predefined and can't be created or deleted, but some of the configurations can be modified.
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property alertConfigurations in identityGovernance
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, requestConfiguration *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder when successful
func (m *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) {
    return NewRolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

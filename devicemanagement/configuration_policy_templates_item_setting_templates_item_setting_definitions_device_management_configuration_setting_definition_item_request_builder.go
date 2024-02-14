package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSettingTemplate entity.
type ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetQueryParameters list of related Setting Definitions
type ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetQueryParameters
}
// ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal instantiates a new ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder and sets the default values.
func NewConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    m := &ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate%2Did}/settingTemplates/{deviceManagementConfigurationSettingTemplate%2Did}/settingDefinitions/{deviceManagementConfigurationSettingDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder instantiates a new ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder and sets the default values.
func NewConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property settingDefinitions for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of related Setting Definitions
// returns a DeviceManagementConfigurationSettingDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable), nil
}
// Patch update the navigation property settingDefinitions in deviceManagement
// returns a DeviceManagementConfigurationSettingDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable, requestConfiguration *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable), nil
}
// ToDeleteRequestInformation delete navigation property settingDefinitions for deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate%2Did}/settingTemplates/{deviceManagementConfigurationSettingTemplate%2Did}/settingDefinitions/{deviceManagementConfigurationSettingDefinition%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of related Setting Definitions
// returns a *RequestInformation when successful
func (m *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property settingDefinitions in deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable, requestConfiguration *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate%2Did}/settingTemplates/{deviceManagementConfigurationSettingTemplate%2Did}/settingDefinitions/{deviceManagementConfigurationSettingDefinition%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder when successful
func (m *ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*ConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    return NewConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

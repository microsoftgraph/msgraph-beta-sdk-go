package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSetting entity.
type ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetQueryParameters list of related Setting Definitions. This property is read-only.
type ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetQueryParameters
}
// ByDeviceManagementConfigurationSettingDefinitionId provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSetting entity.
func (m *ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) ByDeviceManagementConfigurationSettingDefinitionId(deviceManagementConfigurationSettingDefinitionId string)(*ConfigurationPoliciesItemSettingsItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementConfigurationSettingDefinitionId != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = deviceManagementConfigurationSettingDefinitionId
    }
    return NewConfigurationPoliciesItemSettingsItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderInternal instantiates a new SettingDefinitionsRequestBuilder and sets the default values.
func NewConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) {
    m := &ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy%2Did}/settings/{deviceManagementConfigurationSetting%2Did}/settingDefinitions{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder instantiates a new SettingDefinitionsRequestBuilder and sets the default values.
func NewConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) Count()(*ConfigurationPoliciesItemSettingsItemSettingDefinitionsCountRequestBuilder) {
    return NewConfigurationPoliciesItemSettingsItemSettingDefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of related Setting Definitions. This property is read-only.
func (m *ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionCollectionResponseable), nil
}
// ToGetRequestInformation list of related Setting Definitions. This property is read-only.
func (m *ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) WithUrl(rawUrl string)(*ConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) {
    return NewConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

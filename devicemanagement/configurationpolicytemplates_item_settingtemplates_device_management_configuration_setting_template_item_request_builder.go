package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder provides operations to manage the settingTemplates property of the microsoft.graph.deviceManagementConfigurationPolicyTemplate entity.
type ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderGetQueryParameters setting templates
type ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderGetQueryParameters
}
// ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderInternal instantiates a new ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder and sets the default values.
func NewConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) {
    m := &ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate%2Did}/settingTemplates/{deviceManagementConfigurationSettingTemplate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder instantiates a new ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder and sets the default values.
func NewConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property settingTemplates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get setting templates
// returns a DeviceManagementConfigurationSettingTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateable), nil
}
// Patch update the navigation property settingTemplates in deviceManagement
// returns a DeviceManagementConfigurationSettingTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateable, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateable), nil
}
// SettingDefinitions provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSettingTemplate entity.
// returns a *ConfigurationpolicytemplatesItemSettingtemplatesItemSettingdefinitionsSettingDefinitionsRequestBuilder when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) SettingDefinitions()(*ConfigurationpolicytemplatesItemSettingtemplatesItemSettingdefinitionsSettingDefinitionsRequestBuilder) {
    return NewConfigurationpolicytemplatesItemSettingtemplatesItemSettingdefinitionsSettingDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property settingTemplates for deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation setting templates
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property settingTemplates in deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateable, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) WithUrl(rawUrl string)(*ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) {
    return NewConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

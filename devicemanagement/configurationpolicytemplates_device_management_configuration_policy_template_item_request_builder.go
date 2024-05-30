package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
type ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetQueryParameters list of all templates
type ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetQueryParameters
}
// ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal instantiates a new ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder and sets the default values.
func NewConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    m := &ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder instantiates a new ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder and sets the default values.
func NewConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property configurationPolicyTemplates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of all templates
// returns a DeviceManagementConfigurationPolicyTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable), nil
}
// Patch update the navigation property configurationPolicyTemplates in deviceManagement
// returns a DeviceManagementConfigurationPolicyTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, requestConfiguration *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable), nil
}
// SettingTemplates provides operations to manage the settingTemplates property of the microsoft.graph.deviceManagementConfigurationPolicyTemplate entity.
// returns a *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder when successful
func (m *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) SettingTemplates()(*ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) {
    return NewConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property configurationPolicyTemplates for deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of all templates
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property configurationPolicyTemplates in deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, requestConfiguration *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder when successful
func (m *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) WithUrl(rawUrl string)(*ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    return NewConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

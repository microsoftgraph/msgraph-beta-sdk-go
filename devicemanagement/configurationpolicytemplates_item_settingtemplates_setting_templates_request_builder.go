package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder provides operations to manage the settingTemplates property of the microsoft.graph.deviceManagementConfigurationPolicyTemplate entity.
type ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderGetQueryParameters setting templates
type ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderGetQueryParameters struct {
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
// ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderGetQueryParameters
}
// ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementConfigurationSettingTemplateId provides operations to manage the settingTemplates property of the microsoft.graph.deviceManagementConfigurationPolicyTemplate entity.
// returns a *ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) ByDeviceManagementConfigurationSettingTemplateId(deviceManagementConfigurationSettingTemplateId string)(*ConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementConfigurationSettingTemplateId != "" {
        urlTplParams["deviceManagementConfigurationSettingTemplate%2Did"] = deviceManagementConfigurationSettingTemplateId
    }
    return NewConfigurationpolicytemplatesItemSettingtemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderInternal instantiates a new ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder and sets the default values.
func NewConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) {
    m := &ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate%2Did}/settingTemplates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder instantiates a new ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder and sets the default values.
func NewConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConfigurationpolicytemplatesItemSettingtemplatesCountRequestBuilder when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) Count()(*ConfigurationpolicytemplatesItemSettingtemplatesCountRequestBuilder) {
    return NewConfigurationpolicytemplatesItemSettingtemplatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get setting templates
// returns a DeviceManagementConfigurationSettingTemplateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingTemplateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateCollectionResponseable), nil
}
// Post create new navigation property to settingTemplates for deviceManagement
// returns a DeviceManagementConfigurationSettingTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateable, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation setting templates
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to settingTemplates for deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingTemplateable, requestConfiguration *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder when successful
func (m *ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) WithUrl(rawUrl string)(*ConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder) {
    return NewConfigurationpolicytemplatesItemSettingtemplatesSettingTemplatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

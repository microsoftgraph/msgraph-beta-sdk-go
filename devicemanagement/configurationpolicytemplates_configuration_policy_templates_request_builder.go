package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
type ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderGetQueryParameters list of all templates
type ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderGetQueryParameters struct {
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
// ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderGetQueryParameters
}
// ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementConfigurationPolicyTemplateId provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder when successful
func (m *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) ByDeviceManagementConfigurationPolicyTemplateId(deviceManagementConfigurationPolicyTemplateId string)(*ConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementConfigurationPolicyTemplateId != "" {
        urlTplParams["deviceManagementConfigurationPolicyTemplate%2Did"] = deviceManagementConfigurationPolicyTemplateId
    }
    return NewConfigurationpolicytemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderInternal instantiates a new ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder and sets the default values.
func NewConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) {
    m := &ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configurationPolicyTemplates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder instantiates a new ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder and sets the default values.
func NewConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConfigurationpolicytemplatesCountRequestBuilder when successful
func (m *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) Count()(*ConfigurationpolicytemplatesCountRequestBuilder) {
    return NewConfigurationpolicytemplatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of all templates
// returns a DeviceManagementConfigurationPolicyTemplateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyTemplateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateCollectionResponseable), nil
}
// Post create new navigation property to configurationPolicyTemplates for deviceManagement
// returns a DeviceManagementConfigurationPolicyTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, requestConfiguration *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation list of all templates
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to configurationPolicyTemplates for deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, requestConfiguration *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder when successful
func (m *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) WithUrl(rawUrl string)(*ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) {
    return NewConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

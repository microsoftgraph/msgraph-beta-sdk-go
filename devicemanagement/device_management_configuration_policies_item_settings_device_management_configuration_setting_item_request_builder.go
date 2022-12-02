package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
type DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetQueryParameters policy settings
type DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetQueryParameters
}
// DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal instantiates a new DeviceManagementConfigurationSettingItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) {
    m := &DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configurationPolicies/{deviceManagementConfigurationPolicy%2Did}/settings/{deviceManagementConfigurationSetting%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder instantiates a new DeviceManagementConfigurationSettingItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property settings for deviceManagement
func (m *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation policy settings
func (m *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property settings in deviceManagement
func (m *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, requestConfiguration *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property settings for deviceManagement
func (m *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get policy settings
func (m *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable), nil
}
// Patch update the navigation property settings in deviceManagement
func (m *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, requestConfiguration *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable), nil
}
// SettingDefinitions provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSetting entity.
func (m *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) SettingDefinitions()(*DeviceManagementConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) {
    return NewDeviceManagementConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingDefinitionsById provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSetting entity.
func (m *DeviceManagementConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) SettingDefinitionsById(id string)(*DeviceManagementConfigurationPoliciesItemSettingsItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementConfigurationPoliciesItemSettingsItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

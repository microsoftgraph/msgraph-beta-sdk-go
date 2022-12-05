package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder provides operations to manage the settings property of the microsoft.graph.deviceManagementConfigurationPolicy entity.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetQueryParameters policy settings
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetQueryParameters
}
// DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal instantiates a new DeviceManagementConfigurationSettingItemRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) {
    m := &DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy%2Did}/settings/{deviceManagementConfigurationSetting%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder instantiates a new DeviceManagementConfigurationSettingItemRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property settings for deviceManagement
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, error) {
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
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, requestConfiguration *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingable, error) {
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
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) SettingDefinitions()(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilder) {
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsItemSettingDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingDefinitionsById provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSetting entity.
func (m *DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) SettingDefinitionsById(id string)(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesItemSettingsItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

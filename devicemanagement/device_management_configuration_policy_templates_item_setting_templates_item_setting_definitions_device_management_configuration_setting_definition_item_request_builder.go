package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSettingTemplate entity.
type DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetQueryParameters list of related Setting Definitions
type DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetQueryParameters
}
// DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal instantiates a new DeviceManagementConfigurationSettingDefinitionItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    m := &DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate%2Did}/settingTemplates/{deviceManagementConfigurationSettingTemplate%2Did}/settingDefinitions/{deviceManagementConfigurationSettingDefinition%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder instantiates a new DeviceManagementConfigurationSettingDefinitionItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property settingDefinitions for deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of related Setting Definitions
func (m *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property settingDefinitions in deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property settingDefinitions for deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of related Setting Definitions
func (m *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable), nil
}
// Patch update the navigation property settingDefinitions in deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionable), nil
}

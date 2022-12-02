package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
type DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetQueryParameters list of all templates
type DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetQueryParameters
}
// DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal instantiates a new DeviceManagementConfigurationPolicyTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    m := &DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configurationPolicyTemplates/{deviceManagementConfigurationPolicyTemplate%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder instantiates a new DeviceManagementConfigurationPolicyTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property configurationPolicyTemplates for deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of all templates
func (m *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property configurationPolicyTemplates in deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property configurationPolicyTemplates for deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of all templates
func (m *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable), nil
}
// Patch update the navigation property configurationPolicyTemplates in deviceManagement
func (m *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, requestConfiguration *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyTemplateable), nil
}
// SettingTemplates provides operations to manage the settingTemplates property of the microsoft.graph.deviceManagementConfigurationPolicyTemplate entity.
func (m *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) SettingTemplates()(*DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesRequestBuilder) {
    return NewDeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingTemplatesById provides operations to manage the settingTemplates property of the microsoft.graph.deviceManagementConfigurationPolicyTemplate entity.
func (m *DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) SettingTemplatesById(id string)(*DeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingTemplate%2Did"] = id
    }
    return NewDeviceManagementConfigurationPolicyTemplatesItemSettingTemplatesDeviceManagementConfigurationSettingTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

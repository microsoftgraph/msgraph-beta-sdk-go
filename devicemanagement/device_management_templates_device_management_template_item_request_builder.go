package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
type DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderGetQueryParameters the available templates
type DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderGetQueryParameters
}
// DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagementTemplate entity.
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) Categories()(*DeviceManagementTemplatesItemCategoriesRequestBuilder) {
    return NewDeviceManagementTemplatesItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.deviceManagementTemplate entity.
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) CategoriesById(id string)(*DeviceManagementTemplatesItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplateSettingCategory%2Did"] = id
    }
    return NewDeviceManagementTemplatesItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompareWithTemplateId provides operations to call the compare method.
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) CompareWithTemplateId(templateId *string)(*DeviceManagementTemplatesItemCompareWithTemplateIdRequestBuilder) {
    return NewDeviceManagementTemplatesItemCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
// NewDeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderInternal instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) {
    m := &DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property templates for deviceManagement
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the available templates
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateInstance provides operations to call the createInstance method.
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) CreateInstance()(*DeviceManagementTemplatesItemCreateInstanceRequestBuilder) {
    return NewDeviceManagementTemplatesItemCreateInstanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property templates in deviceManagement
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, requestConfiguration *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property templates for deviceManagement
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the available templates
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable), nil
}
// MigratableTo provides operations to manage the migratableTo property of the microsoft.graph.deviceManagementTemplate entity.
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) MigratableTo()(*DeviceManagementTemplatesItemMigratableToRequestBuilder) {
    return NewDeviceManagementTemplatesItemMigratableToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MigratableToById provides operations to manage the migratableTo property of the microsoft.graph.deviceManagementTemplate entity.
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) MigratableToById(id string)(*DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplate%2Did1"] = id
    }
    return NewDeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property templates in deviceManagement
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, requestConfiguration *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementTemplate entity.
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) Settings()(*DeviceManagementTemplatesItemSettingsRequestBuilder) {
    return NewDeviceManagementTemplatesItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.deviceManagementTemplate entity.
func (m *DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) SettingsById(id string)(*DeviceManagementTemplatesItemSettingsDeviceManagementSettingInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance%2Did"] = id
    }
    return NewDeviceManagementTemplatesItemSettingsDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder provides operations to manage the migratableTo property of the microsoft.graph.deviceManagementTemplate entity.
type DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetQueryParameters collection of templates this template can migrate to
type DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetQueryParameters
}
// DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagementTemplate entity.
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Categories()(*DeviceManagementTemplatesItemMigratableToItemCategoriesRequestBuilder) {
    return NewDeviceManagementTemplatesItemMigratableToItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.deviceManagementTemplate entity.
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) CategoriesById(id string)(*DeviceManagementTemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplateSettingCategory%2Did"] = id
    }
    return NewDeviceManagementTemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompareWithTemplateId provides operations to call the compare method.
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) CompareWithTemplateId(templateId *string)(*DeviceManagementTemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder) {
    return NewDeviceManagementTemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
// NewDeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderInternal instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) {
    m := &DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/{deviceManagementTemplate%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property migratableTo for deviceManagement
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation collection of templates this template can migrate to
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) CreateInstance()(*DeviceManagementTemplatesItemMigratableToItemCreateInstanceRequestBuilder) {
    return NewDeviceManagementTemplatesItemMigratableToItemCreateInstanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property migratableTo in deviceManagement
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, requestConfiguration *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property migratableTo for deviceManagement
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of templates this template can migrate to
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, error) {
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
// Patch update the navigation property migratableTo in deviceManagement
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, requestConfiguration *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, error) {
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
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Settings()(*DeviceManagementTemplatesItemMigratableToItemSettingsRequestBuilder) {
    return NewDeviceManagementTemplatesItemMigratableToItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.deviceManagementTemplate entity.
func (m *DeviceManagementTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) SettingsById(id string)(*DeviceManagementTemplatesItemMigratableToItemSettingsDeviceManagementSettingInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance%2Did"] = id
    }
    return NewDeviceManagementTemplatesItemMigratableToItemSettingsDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

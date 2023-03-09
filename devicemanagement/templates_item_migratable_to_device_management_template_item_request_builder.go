package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder provides operations to manage the migratableTo property of the microsoft.graph.deviceManagementTemplate entity.
type TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetQueryParameters collection of templates this template can migrate to
type TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetQueryParameters
}
// TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagementTemplate entity.
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Categories()(*TemplatesItemMigratableToItemCategoriesRequestBuilder) {
    return NewTemplatesItemMigratableToItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.deviceManagementTemplate entity.
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) CategoriesById(id string)(*TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplateSettingCategory%2Did"] = id
    }
    return NewTemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CompareWithTemplateId provides operations to call the compare method.
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) CompareWithTemplateId(templateId *string)(*TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder) {
    return NewTemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId)
}
// NewTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderInternal instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) {
    m := &TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/{deviceManagementTemplate%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateInstance provides operations to call the createInstance method.
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) CreateInstance()(*TemplatesItemMigratableToItemCreateInstanceRequestBuilder) {
    return NewTemplatesItemMigratableToItemCreateInstanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property migratableTo for deviceManagement
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of templates this template can migrate to
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable), nil
}
// Patch update the navigation property migratableTo in deviceManagement
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, requestConfiguration *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementTemplate entity.
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) Settings()(*TemplatesItemMigratableToItemSettingsRequestBuilder) {
    return NewTemplatesItemMigratableToItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.deviceManagementTemplate entity.
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) SettingsById(id string)(*TemplatesItemMigratableToItemSettingsDeviceManagementSettingInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance%2Did"] = id
    }
    return NewTemplatesItemMigratableToItemSettingsDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property migratableTo for deviceManagement
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation collection of templates this template can migrate to
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property migratableTo in deviceManagement
func (m *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, requestConfiguration *TemplatesItemMigratableToDeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder provides operations to manage the categories property of the microsoft.graph.deviceManagementTemplate entity.
type TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetQueryParameters collection of setting categories within the template
type TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetQueryParameters
}
// TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal instantiates a new TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder and sets the default values.
func NewTemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    m := &TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/{deviceManagementTemplate%2Did1}/categories/{deviceManagementTemplateSettingCategory%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder instantiates a new TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder and sets the default values.
func NewTemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property categories for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of setting categories within the template
// returns a DeviceManagementTemplateSettingCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementTemplateSettingCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable), nil
}
// Patch update the navigation property categories in deviceManagement
// returns a DeviceManagementTemplateSettingCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable, requestConfiguration *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementTemplateSettingCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable), nil
}
// RecommendedSettings provides operations to manage the recommendedSettings property of the microsoft.graph.deviceManagementTemplateSettingCategory entity.
// returns a *TemplatesItemMigratabletoItemCategoriesItemRecommendedsettingsRecommendedSettingsRequestBuilder when successful
func (m *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) RecommendedSettings()(*TemplatesItemMigratabletoItemCategoriesItemRecommendedsettingsRecommendedSettingsRequestBuilder) {
    return NewTemplatesItemMigratabletoItemCategoriesItemRecommendedsettingsRecommendedSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SettingDefinitions provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementSettingCategory entity.
// returns a *TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder when successful
func (m *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) SettingDefinitions()(*TemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilder) {
    return NewTemplatesItemMigratabletoItemCategoriesItemSettingdefinitionsSettingDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property categories for deviceManagement
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of setting categories within the template
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property categories in deviceManagement
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable, requestConfiguration *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder when successful
func (m *TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) WithUrl(rawUrl string)(*TemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    return NewTemplatesItemMigratabletoItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

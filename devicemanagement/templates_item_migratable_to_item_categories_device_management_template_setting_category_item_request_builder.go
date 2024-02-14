package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder provides operations to manage the categories property of the microsoft.graph.deviceManagementTemplate entity.
type TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetQueryParameters collection of setting categories within the template
type TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetQueryParameters
}
// TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal instantiates a new TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder and sets the default values.
func NewTemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    m := &TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/{deviceManagementTemplate%2Did1}/categories/{deviceManagementTemplateSettingCategory%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder instantiates a new TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder and sets the default values.
func NewTemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property categories for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable, error) {
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
func (m *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable, requestConfiguration *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable, error) {
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
// returns a *TemplatesItemMigratableToItemCategoriesItemRecommendedSettingsRequestBuilder when successful
func (m *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) RecommendedSettings()(*TemplatesItemMigratableToItemCategoriesItemRecommendedSettingsRequestBuilder) {
    return NewTemplatesItemMigratableToItemCategoriesItemRecommendedSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SettingDefinitions provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementSettingCategory entity.
// returns a *TemplatesItemMigratableToItemCategoriesItemSettingDefinitionsRequestBuilder when successful
func (m *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) SettingDefinitions()(*TemplatesItemMigratableToItemCategoriesItemSettingDefinitionsRequestBuilder) {
    return NewTemplatesItemMigratableToItemCategoriesItemSettingDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property categories for deviceManagement
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/{deviceManagementTemplate%2Did1}/categories/{deviceManagementTemplateSettingCategory%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of setting categories within the template
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateSettingCategoryable, requestConfiguration *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/{deviceManagementTemplate%2Did1}/categories/{deviceManagementTemplateSettingCategory%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder when successful
func (m *TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) WithUrl(rawUrl string)(*TemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    return NewTemplatesItemMigratableToItemCategoriesDeviceManagementTemplateSettingCategoryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

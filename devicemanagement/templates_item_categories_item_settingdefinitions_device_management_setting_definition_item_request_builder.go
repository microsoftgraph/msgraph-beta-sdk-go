package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementSettingCategory entity.
type TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetQueryParameters the setting definitions this category contains
type TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetQueryParameters
}
// TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderInternal instantiates a new TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder and sets the default values.
func NewTemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) {
    m := &TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/categories/{deviceManagementTemplateSettingCategory%2Did}/settingDefinitions/{deviceManagementSettingDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder instantiates a new TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder and sets the default values.
func NewTemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property settingDefinitions for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the setting definitions this category contains
// returns a DeviceManagementSettingDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementSettingDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable), nil
}
// Patch update the navigation property settingDefinitions in deviceManagement
// returns a DeviceManagementSettingDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, requestConfiguration *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementSettingDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable), nil
}
// ToDeleteRequestInformation delete navigation property settingDefinitions for deviceManagement
// returns a *RequestInformation when successful
func (m *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the setting definitions this category contains
// returns a *RequestInformation when successful
func (m *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property settingDefinitions in deviceManagement
// returns a *RequestInformation when successful
func (m *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, requestConfiguration *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder when successful
func (m *TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*TemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) {
    return NewTemplatesItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementSettingCategory entity.
type IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetQueryParameters the setting definitions this category contains
type IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetQueryParameters
}
// IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderInternal instantiates a new IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder and sets the default values.
func NewIntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) {
    m := &IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/categories/{deviceManagementIntentSettingCategory%2Did}/settingDefinitions/{deviceManagementSettingDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder instantiates a new IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder and sets the default values.
func NewIntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property settingDefinitions for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, error) {
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
func (m *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, requestConfiguration *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, error) {
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
func (m *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingDefinitionable, requestConfiguration *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder when successful
func (m *IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*IntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) {
    return NewIntentsItemCategoriesItemSettingdefinitionsDeviceManagementSettingDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

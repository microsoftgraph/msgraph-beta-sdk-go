package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/categories"
    i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/migratableto"
    i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/createinstance"
    iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/settings"
    ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/comparewithtemplateid"
    i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/settings/item"
    i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/categories/item"
    ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/migratableto/item"
)

// DeviceManagementTemplateItemRequestBuilder provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
type DeviceManagementTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementTemplateItemRequestBuilderGetQueryParameters the available templates
type DeviceManagementTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementTemplateItemRequestBuilderGetQueryParameters
}
// DeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories the categories property
func (m *DeviceManagementTemplateItemRequestBuilder) Categories()(*i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae.CategoriesRequestBuilder) {
    return i52be5ec8508b6fa8d8f8557172cc0171c708d106947178250d59c1749dfacaae.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.categories.item collection
func (m *DeviceManagementTemplateItemRequestBuilder) CategoriesById(id string)(*i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3.DeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplateSettingCategory%2Did"] = id
    }
    return i17905e8ecb64c4ebfd26149dfc5a13403bbe9a01c3c4e539d54e3c0f880637b3.NewDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompareWithTemplateId provides operations to call the compare method.
func (m *DeviceManagementTemplateItemRequestBuilder) CompareWithTemplateId(templateId *string)(*ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037.CompareWithTemplateIdRequestBuilder) {
    return ife7ac7603ed6fc45bfdb4af24d907bc1d7f9ba26f799ed8fce9d24ce5ca62037.NewCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
// NewDeviceManagementTemplateItemRequestBuilderInternal instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTemplateItemRequestBuilder) {
    m := &DeviceManagementTemplateItemRequestBuilder{
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
// NewDeviceManagementTemplateItemRequestBuilder instantiates a new DeviceManagementTemplateItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property templates for deviceManagement
func (m *DeviceManagementTemplateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementTemplateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateInstance the createInstance property
func (m *DeviceManagementTemplateItemRequestBuilder) CreateInstance()(*i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d.CreateInstanceRequestBuilder) {
    return i5a59489a0b9d17cfe0d7e3e92257f9c9b9345bb2389513e892ab06a210c7842d.NewCreateInstanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property templates in deviceManagement
func (m *DeviceManagementTemplateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, requestConfiguration *DeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceManagementTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, error) {
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
// MigratableTo the migratableTo property
func (m *DeviceManagementTemplateItemRequestBuilder) MigratableTo()(*i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b.MigratableToRequestBuilder) {
    return i5a180acaf78e7d5e370fefd58c8ff553c81e7bd7ac1752ed93ec3e207aa02b1b.NewMigratableToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MigratableToById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.migratableTo.item collection
func (m *DeviceManagementTemplateItemRequestBuilder) MigratableToById(id string)(*ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2.DeviceManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplate%2Did1"] = id
    }
    return ibf9b83d9fefef8461c50c50942484eaf240cb4a55d3a43a374f84fcc1e8525c2.NewDeviceManagementTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property templates in deviceManagement
func (m *DeviceManagementTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, requestConfiguration *DeviceManagementTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateable, error) {
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
// Settings the settings property
func (m *DeviceManagementTemplateItemRequestBuilder) Settings()(*iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c.SettingsRequestBuilder) {
    return iee353ad56f6c1b8862fce131f47548769bc3fed165f19876db52c9d9388a859c.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.settings.item collection
func (m *DeviceManagementTemplateItemRequestBuilder) SettingsById(id string)(*i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef.DeviceManagementSettingInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance%2Did"] = id
    }
    return i073fa9ff1c690c5b660c4e3b78f5b6a09f71cdb4010b8ccb982b8206fe8485ef.NewDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

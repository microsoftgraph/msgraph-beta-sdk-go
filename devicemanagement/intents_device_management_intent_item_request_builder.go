package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsDeviceManagementIntentItemRequestBuilder provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
type IntentsDeviceManagementIntentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IntentsDeviceManagementIntentItemRequestBuilderGetQueryParameters the device management intents
type IntentsDeviceManagementIntentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IntentsDeviceManagementIntentItemRequestBuilderGetQueryParameters
}
// IntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Assignments()(*IntentsItemAssignmentsRequestBuilder) {
    return NewIntentsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) AssignmentsById(id string)(*IntentsItemAssignmentsDeviceManagementIntentAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIntentsItemAssignmentsDeviceManagementIntentAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Categories()(*IntentsItemCategoriesRequestBuilder) {
    return NewIntentsItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) CategoriesById(id string)(*IntentsItemCategoriesDeviceManagementIntentSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIntentsItemCategoriesDeviceManagementIntentSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// NewIntentsDeviceManagementIntentItemRequestBuilderInternal instantiates a new DeviceManagementIntentItemRequestBuilder and sets the default values.
func NewIntentsDeviceManagementIntentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceManagementIntentId *string)(*IntentsDeviceManagementIntentItemRequestBuilder) {
    m := &IntentsDeviceManagementIntentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementIntentId != nil {
        urlTplParams["deviceManagementIntent%2Did"] = *deviceManagementIntentId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewIntentsDeviceManagementIntentItemRequestBuilder instantiates a new DeviceManagementIntentItemRequestBuilder and sets the default values.
func NewIntentsDeviceManagementIntentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsDeviceManagementIntentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsDeviceManagementIntentItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property intents for deviceManagement
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) DeviceSettingStateSummaries()(*IntentsItemDeviceSettingStateSummariesRequestBuilder) {
    return NewIntentsItemDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceSettingStateSummariesById provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) DeviceStates()(*IntentsItemDeviceStatesRequestBuilder) {
    return NewIntentsItemDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceStatesById provides operations to manage the deviceStates property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) DeviceStatesById(id string)(*IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceStateSummary provides operations to manage the deviceStateSummary property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) DeviceStateSummary()(*IntentsItemDeviceStateSummaryRequestBuilder) {
    return NewIntentsItemDeviceStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the device management intents
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable), nil
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) MicrosoftGraphAssign()(*IntentsItemMicrosoftGraphAssignAssignRequestBuilder) {
    return NewIntentsItemMicrosoftGraphAssignAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCompareWithTemplateId provides operations to call the compare method.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) MicrosoftGraphCompareWithTemplateId(templateId *string)(*IntentsItemMicrosoftGraphCompareWithTemplateIdCompareWithTemplateIdRequestBuilder) {
    return NewIntentsItemMicrosoftGraphCompareWithTemplateIdCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId)
}
// MicrosoftGraphCreateCopy provides operations to call the createCopy method.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) MicrosoftGraphCreateCopy()(*IntentsItemMicrosoftGraphCreateCopyCreateCopyRequestBuilder) {
    return NewIntentsItemMicrosoftGraphCreateCopyCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMigrateToTemplate provides operations to call the migrateToTemplate method.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) MicrosoftGraphMigrateToTemplate()(*IntentsItemMicrosoftGraphMigrateToTemplateMigrateToTemplateRequestBuilder) {
    return NewIntentsItemMicrosoftGraphMigrateToTemplateMigrateToTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateSettings provides operations to call the updateSettings method.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) MicrosoftGraphUpdateSettings()(*IntentsItemMicrosoftGraphUpdateSettingsUpdateSettingsRequestBuilder) {
    return NewIntentsItemMicrosoftGraphUpdateSettingsUpdateSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property intents in deviceManagement
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Settings()(*IntentsItemSettingsRequestBuilder) {
    return NewIntentsItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) SettingsById(id string)(*IntentsItemSettingsDeviceManagementSettingInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIntentsItemSettingsDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ToDeleteRequestInformation delete navigation property intents for deviceManagement
func (m *IntentsDeviceManagementIntentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the device management intents
func (m *IntentsDeviceManagementIntentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property intents in deviceManagement
func (m *IntentsDeviceManagementIntentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserStates provides operations to manage the userStates property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) UserStates()(*IntentsItemUserStatesRequestBuilder) {
    return NewIntentsItemUserStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserStatesById provides operations to manage the userStates property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) UserStatesById(id string)(*IntentsItemUserStatesDeviceManagementIntentUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIntentsItemUserStatesDeviceManagementIntentUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserStateSummary provides operations to manage the userStateSummary property of the microsoft.graph.deviceManagementIntent entity.
func (m *IntentsDeviceManagementIntentItemRequestBuilder) UserStateSummary()(*IntentsItemUserStateSummaryRequestBuilder) {
    return NewIntentsItemUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsDeviceManagementIntentItemRequestBuilder provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
type IntentsDeviceManagementIntentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// Assign provides operations to call the assign method.
// returns a *IntentsItemAssignRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Assign()(*IntentsItemAssignRequestBuilder) {
    return NewIntentsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementIntent entity.
// returns a *IntentsItemAssignmentsRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Assignments()(*IntentsItemAssignmentsRequestBuilder) {
    return NewIntentsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagementIntent entity.
// returns a *IntentsItemCategoriesRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Categories()(*IntentsItemCategoriesRequestBuilder) {
    return NewIntentsItemCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompareWithTemplateId provides operations to call the compare method.
// returns a *IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) CompareWithTemplateId(templateId *string)(*IntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilder) {
    return NewIntentsItemComparewithtemplateidCompareWithTemplateIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, templateId)
}
// NewIntentsDeviceManagementIntentItemRequestBuilderInternal instantiates a new IntentsDeviceManagementIntentItemRequestBuilder and sets the default values.
func NewIntentsDeviceManagementIntentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsDeviceManagementIntentItemRequestBuilder) {
    m := &IntentsDeviceManagementIntentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIntentsDeviceManagementIntentItemRequestBuilder instantiates a new IntentsDeviceManagementIntentItemRequestBuilder and sets the default values.
func NewIntentsDeviceManagementIntentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsDeviceManagementIntentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsDeviceManagementIntentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateCopy provides operations to call the createCopy method.
// returns a *IntentsItemCreatecopyCreateCopyRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) CreateCopy()(*IntentsItemCreatecopyCreateCopyRequestBuilder) {
    return NewIntentsItemCreatecopyCreateCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property intents for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceManagementIntent entity.
// returns a *IntentsItemDevicesettingstatesummariesDeviceSettingStateSummariesRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) DeviceSettingStateSummaries()(*IntentsItemDevicesettingstatesummariesDeviceSettingStateSummariesRequestBuilder) {
    return NewIntentsItemDevicesettingstatesummariesDeviceSettingStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.deviceManagementIntent entity.
// returns a *IntentsItemDevicestatesDeviceStatesRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) DeviceStates()(*IntentsItemDevicestatesDeviceStatesRequestBuilder) {
    return NewIntentsItemDevicestatesDeviceStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStateSummary provides operations to manage the deviceStateSummary property of the microsoft.graph.deviceManagementIntent entity.
// returns a *IntentsItemDevicestatesummaryDeviceStateSummaryRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) DeviceStateSummary()(*IntentsItemDevicestatesummaryDeviceStateSummaryRequestBuilder) {
    return NewIntentsItemDevicestatesummaryDeviceStateSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the device management intents
// returns a DeviceManagementIntentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable), nil
}
// GetCustomizedSettings provides operations to call the getCustomizedSettings method.
// returns a *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) GetCustomizedSettings()(*IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder) {
    return NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MigrateToTemplate provides operations to call the migrateToTemplate method.
// returns a *IntentsItemMigratetotemplateMigrateToTemplateRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) MigrateToTemplate()(*IntentsItemMigratetotemplateMigrateToTemplateRequestBuilder) {
    return NewIntentsItemMigratetotemplateMigrateToTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property intents in deviceManagement
// returns a DeviceManagementIntentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementIntent entity.
// returns a *IntentsItemSettingsRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) Settings()(*IntentsItemSettingsRequestBuilder) {
    return NewIntentsItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property intents for deviceManagement
// returns a *RequestInformation when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the device management intents
// returns a *RequestInformation when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property intents in deviceManagement
// returns a *RequestInformation when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, requestConfiguration *IntentsDeviceManagementIntentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateSettings provides operations to call the updateSettings method.
// returns a *IntentsItemUpdatesettingsUpdateSettingsRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) UpdateSettings()(*IntentsItemUpdatesettingsUpdateSettingsRequestBuilder) {
    return NewIntentsItemUpdatesettingsUpdateSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserStates provides operations to manage the userStates property of the microsoft.graph.deviceManagementIntent entity.
// returns a *IntentsItemUserstatesUserStatesRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) UserStates()(*IntentsItemUserstatesUserStatesRequestBuilder) {
    return NewIntentsItemUserstatesUserStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserStateSummary provides operations to manage the userStateSummary property of the microsoft.graph.deviceManagementIntent entity.
// returns a *IntentsItemUserstatesummaryUserStateSummaryRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) UserStateSummary()(*IntentsItemUserstatesummaryUserStateSummaryRequestBuilder) {
    return NewIntentsItemUserstatesummaryUserStateSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IntentsDeviceManagementIntentItemRequestBuilder when successful
func (m *IntentsDeviceManagementIntentItemRequestBuilder) WithUrl(rawUrl string)(*IntentsDeviceManagementIntentItemRequestBuilder) {
    return NewIntentsDeviceManagementIntentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

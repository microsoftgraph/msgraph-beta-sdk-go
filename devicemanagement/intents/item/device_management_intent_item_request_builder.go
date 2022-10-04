package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2195331110855a9498295e2fc8c02d9e419cef49a131e88a43c884e5f36bf2df "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/userstates"
    i23e71622e02878eed00b3312aa65af0ceb8cee6cb6853cdeedd5ba33f3bd24d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/devicestates"
    i4f1a4c082319bee8e877418cf5db676aa138d2c05f5586a2c19d6a0251406722 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/comparewithtemplateid"
    i5e948c2b0dc0310c910b94d64fbeab1a9145ed93f22976d0c6a733f47d53ba47 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/devicesettingstatesummaries"
    i77dceab723bbe034f83895d00554f2801769e77222afcf5f03698bf9f20065f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/assign"
    i8744b36b6ccd58cfb5566d8bcfaaaa5a91c2ca5a7d99218a0d837bd74414e860 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/devicestatesummary"
    i8baeef7b87df0daa79bc3d6db8884f08954d45824e962bdd11e8d114daa152f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/migratetotemplate"
    ia8fb096dcfa8bad706be070aca9e18df3a577a485e053215faf223148f7c812f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/categories"
    ib10ff44e65c3ab5e68ee1890f6442340272e8295671c255a8cfe146e7d186f6c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/settings"
    id4e165bb07af7f89bf1f315b6258bc4ffa52e23f05e7b625bca9be4b284ad5b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/assignments"
    ie85946f822147a5126c80dde1abb6503c5a243e2db7c41afd5c15aa0236e4778 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/userstatesummary"
    ie911e70a120660c48ad60cfc3d6cc4ec217fdcaa92a6d7f3f220127d107b81db "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/updatesettings"
    ifad589ef3ce91c5f33b39c8260cfbe764d5d9bb0a50352392359f13ea1148770 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/createcopy"
    i050af1e40f4b04a3234cd6eaef05c2dba88a1a4304262e163378332314ce51be "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/categories/item"
    i292492bc5eb43f5aa2fe6644e7c97af199a20c37d91e9b0991d9fbca2bb024de "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/userstates/item"
    i52e44e28ad3f36bea30cd68ebb25b8657782d8b7b8c755f6add0e96288a7800a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/settings/item"
    i844844c782791d7d5f694de937d8ad00ac682c52a360e46aa3f1340df417d61c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/devicesettingstatesummaries/item"
    ib490048065b539bfe5b6cfa49328e603b61cfae55d4624781d0df2de10463072 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/devicestates/item"
    ie20d3efb9c88b7a4dececad1e9328ce5526d4cdbd23f6257aa72d18d9126ed81 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intents/item/assignments/item"
)

// DeviceManagementIntentItemRequestBuilder provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
type DeviceManagementIntentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementIntentItemRequestBuilderGetQueryParameters the device management intents
type DeviceManagementIntentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementIntentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementIntentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementIntentItemRequestBuilderGetQueryParameters
}
// DeviceManagementIntentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementIntentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *DeviceManagementIntentItemRequestBuilder) Assign()(*i77dceab723bbe034f83895d00554f2801769e77222afcf5f03698bf9f20065f0.AssignRequestBuilder) {
    return i77dceab723bbe034f83895d00554f2801769e77222afcf5f03698bf9f20065f0.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceManagementIntentItemRequestBuilder) Assignments()(*id4e165bb07af7f89bf1f315b6258bc4ffa52e23f05e7b625bca9be4b284ad5b0.AssignmentsRequestBuilder) {
    return id4e165bb07af7f89bf1f315b6258bc4ffa52e23f05e7b625bca9be4b284ad5b0.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.assignments.item collection
func (m *DeviceManagementIntentItemRequestBuilder) AssignmentsById(id string)(*ie20d3efb9c88b7a4dececad1e9328ce5526d4cdbd23f6257aa72d18d9126ed81.DeviceManagementIntentAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentAssignment%2Did"] = id
    }
    return ie20d3efb9c88b7a4dececad1e9328ce5526d4cdbd23f6257aa72d18d9126ed81.NewDeviceManagementIntentAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories the categories property
func (m *DeviceManagementIntentItemRequestBuilder) Categories()(*ia8fb096dcfa8bad706be070aca9e18df3a577a485e053215faf223148f7c812f.CategoriesRequestBuilder) {
    return ia8fb096dcfa8bad706be070aca9e18df3a577a485e053215faf223148f7c812f.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.categories.item collection
func (m *DeviceManagementIntentItemRequestBuilder) CategoriesById(id string)(*i050af1e40f4b04a3234cd6eaef05c2dba88a1a4304262e163378332314ce51be.DeviceManagementIntentSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentSettingCategory%2Did"] = id
    }
    return i050af1e40f4b04a3234cd6eaef05c2dba88a1a4304262e163378332314ce51be.NewDeviceManagementIntentSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompareWithTemplateId provides operations to call the compare method.
func (m *DeviceManagementIntentItemRequestBuilder) CompareWithTemplateId(templateId *string)(*i4f1a4c082319bee8e877418cf5db676aa138d2c05f5586a2c19d6a0251406722.CompareWithTemplateIdRequestBuilder) {
    return i4f1a4c082319bee8e877418cf5db676aa138d2c05f5586a2c19d6a0251406722.NewCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
// NewDeviceManagementIntentItemRequestBuilderInternal instantiates a new DeviceManagementIntentItemRequestBuilder and sets the default values.
func NewDeviceManagementIntentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementIntentItemRequestBuilder) {
    m := &DeviceManagementIntentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementIntentItemRequestBuilder instantiates a new DeviceManagementIntentItemRequestBuilder and sets the default values.
func NewDeviceManagementIntentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementIntentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementIntentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateCopy the createCopy property
func (m *DeviceManagementIntentItemRequestBuilder) CreateCopy()(*ifad589ef3ce91c5f33b39c8260cfbe764d5d9bb0a50352392359f13ea1148770.CreateCopyRequestBuilder) {
    return ifad589ef3ce91c5f33b39c8260cfbe764d5d9bb0a50352392359f13ea1148770.NewCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property intents for deviceManagement
func (m *DeviceManagementIntentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the device management intents
func (m *DeviceManagementIntentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementIntentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property intents in deviceManagement
func (m *DeviceManagementIntentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, requestConfiguration *DeviceManagementIntentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property intents for deviceManagement
func (m *DeviceManagementIntentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementIntentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceSettingStateSummaries the deviceSettingStateSummaries property
func (m *DeviceManagementIntentItemRequestBuilder) DeviceSettingStateSummaries()(*i5e948c2b0dc0310c910b94d64fbeab1a9145ed93f22976d0c6a733f47d53ba47.DeviceSettingStateSummariesRequestBuilder) {
    return i5e948c2b0dc0310c910b94d64fbeab1a9145ed93f22976d0c6a733f47d53ba47.NewDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.deviceSettingStateSummaries.item collection
func (m *DeviceManagementIntentItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*i844844c782791d7d5f694de937d8ad00ac682c52a360e46aa3f1340df417d61c.DeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentDeviceSettingStateSummary%2Did"] = id
    }
    return i844844c782791d7d5f694de937d8ad00ac682c52a360e46aa3f1340df417d61c.NewDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStates the deviceStates property
func (m *DeviceManagementIntentItemRequestBuilder) DeviceStates()(*i23e71622e02878eed00b3312aa65af0ceb8cee6cb6853cdeedd5ba33f3bd24d7.DeviceStatesRequestBuilder) {
    return i23e71622e02878eed00b3312aa65af0ceb8cee6cb6853cdeedd5ba33f3bd24d7.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.deviceStates.item collection
func (m *DeviceManagementIntentItemRequestBuilder) DeviceStatesById(id string)(*ib490048065b539bfe5b6cfa49328e603b61cfae55d4624781d0df2de10463072.DeviceManagementIntentDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentDeviceState%2Did"] = id
    }
    return ib490048065b539bfe5b6cfa49328e603b61cfae55d4624781d0df2de10463072.NewDeviceManagementIntentDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStateSummary the deviceStateSummary property
func (m *DeviceManagementIntentItemRequestBuilder) DeviceStateSummary()(*i8744b36b6ccd58cfb5566d8bcfaaaa5a91c2ca5a7d99218a0d837bd74414e860.DeviceStateSummaryRequestBuilder) {
    return i8744b36b6ccd58cfb5566d8bcfaaaa5a91c2ca5a7d99218a0d837bd74414e860.NewDeviceStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device management intents
func (m *DeviceManagementIntentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementIntentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable), nil
}
// MigrateToTemplate the migrateToTemplate property
func (m *DeviceManagementIntentItemRequestBuilder) MigrateToTemplate()(*i8baeef7b87df0daa79bc3d6db8884f08954d45824e962bdd11e8d114daa152f2.MigrateToTemplateRequestBuilder) {
    return i8baeef7b87df0daa79bc3d6db8884f08954d45824e962bdd11e8d114daa152f2.NewMigrateToTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property intents in deviceManagement
func (m *DeviceManagementIntentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, requestConfiguration *DeviceManagementIntentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentable), nil
}
// Settings the settings property
func (m *DeviceManagementIntentItemRequestBuilder) Settings()(*ib10ff44e65c3ab5e68ee1890f6442340272e8295671c255a8cfe146e7d186f6c.SettingsRequestBuilder) {
    return ib10ff44e65c3ab5e68ee1890f6442340272e8295671c255a8cfe146e7d186f6c.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.settings.item collection
func (m *DeviceManagementIntentItemRequestBuilder) SettingsById(id string)(*i52e44e28ad3f36bea30cd68ebb25b8657782d8b7b8c755f6add0e96288a7800a.DeviceManagementSettingInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance%2Did"] = id
    }
    return i52e44e28ad3f36bea30cd68ebb25b8657782d8b7b8c755f6add0e96288a7800a.NewDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UpdateSettings the updateSettings property
func (m *DeviceManagementIntentItemRequestBuilder) UpdateSettings()(*ie911e70a120660c48ad60cfc3d6cc4ec217fdcaa92a6d7f3f220127d107b81db.UpdateSettingsRequestBuilder) {
    return ie911e70a120660c48ad60cfc3d6cc4ec217fdcaa92a6d7f3f220127d107b81db.NewUpdateSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStates the userStates property
func (m *DeviceManagementIntentItemRequestBuilder) UserStates()(*i2195331110855a9498295e2fc8c02d9e419cef49a131e88a43c884e5f36bf2df.UserStatesRequestBuilder) {
    return i2195331110855a9498295e2fc8c02d9e419cef49a131e88a43c884e5f36bf2df.NewUserStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.userStates.item collection
func (m *DeviceManagementIntentItemRequestBuilder) UserStatesById(id string)(*i292492bc5eb43f5aa2fe6644e7c97af199a20c37d91e9b0991d9fbca2bb024de.DeviceManagementIntentUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentUserState%2Did"] = id
    }
    return i292492bc5eb43f5aa2fe6644e7c97af199a20c37d91e9b0991d9fbca2bb024de.NewDeviceManagementIntentUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStateSummary the userStateSummary property
func (m *DeviceManagementIntentItemRequestBuilder) UserStateSummary()(*ie85946f822147a5126c80dde1abb6503c5a243e2db7c41afd5c15aa0236e4778.UserStateSummaryRequestBuilder) {
    return ie85946f822147a5126c80dde1abb6503c5a243e2db7c41afd5c15aa0236e4778.NewUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementIntentItemRequestBuilderDeleteOptions options for Delete
type DeviceManagementIntentItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementIntentItemRequestBuilderGetOptions options for Get
type DeviceManagementIntentItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementIntentItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementIntentItemRequestBuilderGetQueryParameters the device management intents
type DeviceManagementIntentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementIntentItemRequestBuilderPatchOptions options for Patch
type DeviceManagementIntentItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntentable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementIntentItemRequestBuilder) Assign()(*i77dceab723bbe034f83895d00554f2801769e77222afcf5f03698bf9f20065f0.AssignRequestBuilder) {
    return i77dceab723bbe034f83895d00554f2801769e77222afcf5f03698bf9f20065f0.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["deviceManagementIntentAssignment_id"] = id
    }
    return ie20d3efb9c88b7a4dececad1e9328ce5526d4cdbd23f6257aa72d18d9126ed81.NewDeviceManagementIntentAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["deviceManagementIntentSettingCategory_id"] = id
    }
    return i050af1e40f4b04a3234cd6eaef05c2dba88a1a4304262e163378332314ce51be.NewDeviceManagementIntentSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompareWithTemplateId provides operations to call the compare method.
func (m *DeviceManagementIntentItemRequestBuilder) CompareWithTemplateId(templateId *string)(*i4f1a4c082319bee8e877418cf5db676aa138d2c05f5586a2c19d6a0251406722.CompareWithTemplateIdRequestBuilder) {
    return i4f1a4c082319bee8e877418cf5db676aa138d2c05f5586a2c19d6a0251406722.NewCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
// NewDeviceManagementIntentItemRequestBuilderInternal instantiates a new DeviceManagementIntentItemRequestBuilder and sets the default values.
func NewDeviceManagementIntentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementIntentItemRequestBuilder) {
    m := &DeviceManagementIntentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/intents/{deviceManagementIntent_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementIntentItemRequestBuilder instantiates a new DeviceManagementIntentItemRequestBuilder and sets the default values.
func NewDeviceManagementIntentItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementIntentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementIntentItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceManagementIntentItemRequestBuilder) CreateCopy()(*ifad589ef3ce91c5f33b39c8260cfbe764d5d9bb0a50352392359f13ea1148770.CreateCopyRequestBuilder) {
    return ifad589ef3ce91c5f33b39c8260cfbe764d5d9bb0a50352392359f13ea1148770.NewCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property intents for deviceManagement
func (m *DeviceManagementIntentItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementIntentItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the device management intents
func (m *DeviceManagementIntentItemRequestBuilder) CreateGetRequestInformation(options *DeviceManagementIntentItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property intents in deviceManagement
func (m *DeviceManagementIntentItemRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementIntentItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property intents for deviceManagement
func (m *DeviceManagementIntentItemRequestBuilder) Delete(options *DeviceManagementIntentItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["deviceManagementIntentDeviceSettingStateSummary_id"] = id
    }
    return i844844c782791d7d5f694de937d8ad00ac682c52a360e46aa3f1340df417d61c.NewDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["deviceManagementIntentDeviceState_id"] = id
    }
    return ib490048065b539bfe5b6cfa49328e603b61cfae55d4624781d0df2de10463072.NewDeviceManagementIntentDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementIntentItemRequestBuilder) DeviceStateSummary()(*i8744b36b6ccd58cfb5566d8bcfaaaa5a91c2ca5a7d99218a0d837bd74414e860.DeviceStateSummaryRequestBuilder) {
    return i8744b36b6ccd58cfb5566d8bcfaaaa5a91c2ca5a7d99218a0d837bd74414e860.NewDeviceStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device management intents
func (m *DeviceManagementIntentItemRequestBuilder) Get(options *DeviceManagementIntentItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceManagementIntentFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntentable), nil
}
func (m *DeviceManagementIntentItemRequestBuilder) MigrateToTemplate()(*i8baeef7b87df0daa79bc3d6db8884f08954d45824e962bdd11e8d114daa152f2.MigrateToTemplateRequestBuilder) {
    return i8baeef7b87df0daa79bc3d6db8884f08954d45824e962bdd11e8d114daa152f2.NewMigrateToTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property intents in deviceManagement
func (m *DeviceManagementIntentItemRequestBuilder) Patch(options *DeviceManagementIntentItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["deviceManagementSettingInstance_id"] = id
    }
    return i52e44e28ad3f36bea30cd68ebb25b8657782d8b7b8c755f6add0e96288a7800a.NewDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementIntentItemRequestBuilder) UpdateSettings()(*ie911e70a120660c48ad60cfc3d6cc4ec217fdcaa92a6d7f3f220127d107b81db.UpdateSettingsRequestBuilder) {
    return ie911e70a120660c48ad60cfc3d6cc4ec217fdcaa92a6d7f3f220127d107b81db.NewUpdateSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["deviceManagementIntentUserState_id"] = id
    }
    return i292492bc5eb43f5aa2fe6644e7c97af199a20c37d91e9b0991d9fbca2bb024de.NewDeviceManagementIntentUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementIntentItemRequestBuilder) UserStateSummary()(*ie85946f822147a5126c80dde1abb6503c5a243e2db7c41afd5c15aa0236e4778.UserStateSummaryRequestBuilder) {
    return ie85946f822147a5126c80dde1abb6503c5a243e2db7c41afd5c15aa0236e4778.NewUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

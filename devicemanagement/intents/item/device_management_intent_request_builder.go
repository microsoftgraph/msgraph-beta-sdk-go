package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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

// DeviceManagementIntentRequestBuilder builds and executes requests for operations under \deviceManagement\intents\{deviceManagementIntent-id}
type DeviceManagementIntentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementIntentRequestBuilderDeleteOptions options for Delete
type DeviceManagementIntentRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementIntentRequestBuilderGetOptions options for Get
type DeviceManagementIntentRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementIntentRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementIntentRequestBuilderGetQueryParameters the device management intents
type DeviceManagementIntentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementIntentRequestBuilderPatchOptions options for Patch
type DeviceManagementIntentRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntent;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementIntentRequestBuilder) Assign()(*i77dceab723bbe034f83895d00554f2801769e77222afcf5f03698bf9f20065f0.AssignRequestBuilder) {
    return i77dceab723bbe034f83895d00554f2801769e77222afcf5f03698bf9f20065f0.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementIntentRequestBuilder) Assignments()(*id4e165bb07af7f89bf1f315b6258bc4ffa52e23f05e7b625bca9be4b284ad5b0.AssignmentsRequestBuilder) {
    return id4e165bb07af7f89bf1f315b6258bc4ffa52e23f05e7b625bca9be4b284ad5b0.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.assignments.item collection
func (m *DeviceManagementIntentRequestBuilder) AssignmentsById(id string)(*ie20d3efb9c88b7a4dececad1e9328ce5526d4cdbd23f6257aa72d18d9126ed81.DeviceManagementIntentAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentAssignment_id"] = id
    }
    return ie20d3efb9c88b7a4dececad1e9328ce5526d4cdbd23f6257aa72d18d9126ed81.NewDeviceManagementIntentAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementIntentRequestBuilder) Categories()(*ia8fb096dcfa8bad706be070aca9e18df3a577a485e053215faf223148f7c812f.CategoriesRequestBuilder) {
    return ia8fb096dcfa8bad706be070aca9e18df3a577a485e053215faf223148f7c812f.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.categories.item collection
func (m *DeviceManagementIntentRequestBuilder) CategoriesById(id string)(*i050af1e40f4b04a3234cd6eaef05c2dba88a1a4304262e163378332314ce51be.DeviceManagementIntentSettingCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentSettingCategory_id"] = id
    }
    return i050af1e40f4b04a3234cd6eaef05c2dba88a1a4304262e163378332314ce51be.NewDeviceManagementIntentSettingCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompareWithTemplateId builds and executes requests for operations under \deviceManagement\intents\{deviceManagementIntent-id}\microsoft.graph.compare(templateId='{templateId}')
func (m *DeviceManagementIntentRequestBuilder) CompareWithTemplateId(templateId *string)(*i4f1a4c082319bee8e877418cf5db676aa138d2c05f5586a2c19d6a0251406722.CompareWithTemplateIdRequestBuilder) {
    return i4f1a4c082319bee8e877418cf5db676aa138d2c05f5586a2c19d6a0251406722.NewCompareWithTemplateIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, templateId);
}
// NewDeviceManagementIntentRequestBuilderInternal instantiates a new DeviceManagementIntentRequestBuilder and sets the default values.
func NewDeviceManagementIntentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementIntentRequestBuilder) {
    m := &DeviceManagementIntentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/intents/{deviceManagementIntent_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementIntentRequestBuilder instantiates a new DeviceManagementIntentRequestBuilder and sets the default values.
func NewDeviceManagementIntentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementIntentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementIntentRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceManagementIntentRequestBuilder) CreateCopy()(*ifad589ef3ce91c5f33b39c8260cfbe764d5d9bb0a50352392359f13ea1148770.CreateCopyRequestBuilder) {
    return ifad589ef3ce91c5f33b39c8260cfbe764d5d9bb0a50352392359f13ea1148770.NewCreateCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the device management intents
func (m *DeviceManagementIntentRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementIntentRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementIntentRequestBuilder) CreateGetRequestInformation(options *DeviceManagementIntentRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the device management intents
func (m *DeviceManagementIntentRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementIntentRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the device management intents
func (m *DeviceManagementIntentRequestBuilder) Delete(options *DeviceManagementIntentRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceManagementIntentRequestBuilder) DeviceSettingStateSummaries()(*i5e948c2b0dc0310c910b94d64fbeab1a9145ed93f22976d0c6a733f47d53ba47.DeviceSettingStateSummariesRequestBuilder) {
    return i5e948c2b0dc0310c910b94d64fbeab1a9145ed93f22976d0c6a733f47d53ba47.NewDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.deviceSettingStateSummaries.item collection
func (m *DeviceManagementIntentRequestBuilder) DeviceSettingStateSummariesById(id string)(*i844844c782791d7d5f694de937d8ad00ac682c52a360e46aa3f1340df417d61c.DeviceManagementIntentDeviceSettingStateSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentDeviceSettingStateSummary_id"] = id
    }
    return i844844c782791d7d5f694de937d8ad00ac682c52a360e46aa3f1340df417d61c.NewDeviceManagementIntentDeviceSettingStateSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementIntentRequestBuilder) DeviceStates()(*i23e71622e02878eed00b3312aa65af0ceb8cee6cb6853cdeedd5ba33f3bd24d7.DeviceStatesRequestBuilder) {
    return i23e71622e02878eed00b3312aa65af0ceb8cee6cb6853cdeedd5ba33f3bd24d7.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.deviceStates.item collection
func (m *DeviceManagementIntentRequestBuilder) DeviceStatesById(id string)(*ib490048065b539bfe5b6cfa49328e603b61cfae55d4624781d0df2de10463072.DeviceManagementIntentDeviceStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentDeviceState_id"] = id
    }
    return ib490048065b539bfe5b6cfa49328e603b61cfae55d4624781d0df2de10463072.NewDeviceManagementIntentDeviceStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementIntentRequestBuilder) DeviceStateSummary()(*i8744b36b6ccd58cfb5566d8bcfaaaa5a91c2ca5a7d99218a0d837bd74414e860.DeviceStateSummaryRequestBuilder) {
    return i8744b36b6ccd58cfb5566d8bcfaaaa5a91c2ca5a7d99218a0d837bd74414e860.NewDeviceStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device management intents
func (m *DeviceManagementIntentRequestBuilder) Get(options *DeviceManagementIntentRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntent, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementIntent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntent), nil
}
func (m *DeviceManagementIntentRequestBuilder) MigrateToTemplate()(*i8baeef7b87df0daa79bc3d6db8884f08954d45824e962bdd11e8d114daa152f2.MigrateToTemplateRequestBuilder) {
    return i8baeef7b87df0daa79bc3d6db8884f08954d45824e962bdd11e8d114daa152f2.NewMigrateToTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the device management intents
func (m *DeviceManagementIntentRequestBuilder) Patch(options *DeviceManagementIntentRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceManagementIntentRequestBuilder) Settings()(*ib10ff44e65c3ab5e68ee1890f6442340272e8295671c255a8cfe146e7d186f6c.SettingsRequestBuilder) {
    return ib10ff44e65c3ab5e68ee1890f6442340272e8295671c255a8cfe146e7d186f6c.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.settings.item collection
func (m *DeviceManagementIntentRequestBuilder) SettingsById(id string)(*i52e44e28ad3f36bea30cd68ebb25b8657782d8b7b8c755f6add0e96288a7800a.DeviceManagementSettingInstanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance_id"] = id
    }
    return i52e44e28ad3f36bea30cd68ebb25b8657782d8b7b8c755f6add0e96288a7800a.NewDeviceManagementSettingInstanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementIntentRequestBuilder) UpdateSettings()(*ie911e70a120660c48ad60cfc3d6cc4ec217fdcaa92a6d7f3f220127d107b81db.UpdateSettingsRequestBuilder) {
    return ie911e70a120660c48ad60cfc3d6cc4ec217fdcaa92a6d7f3f220127d107b81db.NewUpdateSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementIntentRequestBuilder) UserStates()(*i2195331110855a9498295e2fc8c02d9e419cef49a131e88a43c884e5f36bf2df.UserStatesRequestBuilder) {
    return i2195331110855a9498295e2fc8c02d9e419cef49a131e88a43c884e5f36bf2df.NewUserStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intents.item.userStates.item collection
func (m *DeviceManagementIntentRequestBuilder) UserStatesById(id string)(*i292492bc5eb43f5aa2fe6644e7c97af199a20c37d91e9b0991d9fbca2bb024de.DeviceManagementIntentUserStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntentUserState_id"] = id
    }
    return i292492bc5eb43f5aa2fe6644e7c97af199a20c37d91e9b0991d9fbca2bb024de.NewDeviceManagementIntentUserStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementIntentRequestBuilder) UserStateSummary()(*ie85946f822147a5126c80dde1abb6503c5a243e2db7c41afd5c15aa0236e4778.UserStateSummaryRequestBuilder) {
    return ie85946f822147a5126c80dde1abb6503c5a243e2db7c41afd5c15aa0236e4778.NewUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package outlook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i767cd2a73c29879daa39ae90aebfabcd397159b1e720933714a03e3a5fceb72a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/supportedtimezoneswithtimezonestandard"
    i7832228cb9bca400c299eb3f31f40734874357edc378956e69605b0e3b39ef21 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups"
    i82b391a4153948653004ca349b6eac39e2c3e7b3926201627e2e3eea88aad251 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders"
    ia9234bf87076aa410405ec78141ae5dc5da406a4d6cc28b4a7be5092b0ee6296 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/mastercategories"
    ib55f78ab3053067f88ec9c133bd6a0d5d1ffdb2ab438361fe4ce817346e58c04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/supportedtimezones"
    ibf3b7a436ea0ea400b8c6c0590a9be4465672faf619ae54e97e418017b25ef00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks"
    ic40f6d057fc54666d8fde145e03546b6c1c12cc293b9e830bbf8cba8197b79be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/supportedlanguages"
    i84930f2750a94ee2130f035e9c9bdbcfffcb068c52a806479d0964d9d2020e40 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/mastercategories/item"
    i984fa253a0580cb2d81d93e449a106836d455f23e065de8d03edfcc8dbf44acb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskgroups/item"
    ic19e499d3e4541329156d6926f4f6db570dcce60051f4e1edd3f0b0db81cff9a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item"
    id3508c1f1aa1eff5970c7a4fe8b1bb7b16a0d219f157caf9278141e89295cd00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/tasks/item"
)

// OutlookRequestBuilder builds and executes requests for operations under \me\outlook
type OutlookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OutlookRequestBuilderDeleteOptions options for Delete
type OutlookRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutlookRequestBuilderGetOptions options for Get
type OutlookRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OutlookRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutlookRequestBuilderGetQueryParameters read-only.
type OutlookRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// OutlookRequestBuilderPatchOptions options for Patch
type OutlookRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOutlookRequestBuilderInternal instantiates a new OutlookRequestBuilder and sets the default values.
func NewOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookRequestBuilder) {
    m := &OutlookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOutlookRequestBuilder instantiates a new OutlookRequestBuilder and sets the default values.
func NewOutlookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only.
func (m *OutlookRequestBuilder) CreateDeleteRequestInformation(options *OutlookRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only.
func (m *OutlookRequestBuilder) CreateGetRequestInformation(options *OutlookRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation read-only.
func (m *OutlookRequestBuilder) CreatePatchRequestInformation(options *OutlookRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete read-only.
func (m *OutlookRequestBuilder) Delete(options *OutlookRequestBuilderDeleteOptions)(error) {
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
// Get read-only.
func (m *OutlookRequestBuilder) Get(options *OutlookRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookUser() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookUser), nil
}
func (m *OutlookRequestBuilder) MasterCategories()(*ia9234bf87076aa410405ec78141ae5dc5da406a4d6cc28b4a7be5092b0ee6296.MasterCategoriesRequestBuilder) {
    return ia9234bf87076aa410405ec78141ae5dc5da406a4d6cc28b4a7be5092b0ee6296.NewMasterCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MasterCategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.masterCategories.item collection
func (m *OutlookRequestBuilder) MasterCategoriesById(id string)(*i84930f2750a94ee2130f035e9c9bdbcfffcb068c52a806479d0964d9d2020e40.OutlookCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookCategory_id"] = id
    }
    return i84930f2750a94ee2130f035e9c9bdbcfffcb068c52a806479d0964d9d2020e40.NewOutlookCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch read-only.
func (m *OutlookRequestBuilder) Patch(options *OutlookRequestBuilderPatchOptions)(error) {
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
// SupportedLanguages builds and executes requests for operations under \me\outlook\microsoft.graph.supportedLanguages()
func (m *OutlookRequestBuilder) SupportedLanguages()(*ic40f6d057fc54666d8fde145e03546b6c1c12cc293b9e830bbf8cba8197b79be.SupportedLanguagesRequestBuilder) {
    return ic40f6d057fc54666d8fde145e03546b6c1c12cc293b9e830bbf8cba8197b79be.NewSupportedLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZones builds and executes requests for operations under \me\outlook\microsoft.graph.supportedTimeZones()
func (m *OutlookRequestBuilder) SupportedTimeZones()(*ib55f78ab3053067f88ec9c133bd6a0d5d1ffdb2ab438361fe4ce817346e58c04.SupportedTimeZonesRequestBuilder) {
    return ib55f78ab3053067f88ec9c133bd6a0d5d1ffdb2ab438361fe4ce817346e58c04.NewSupportedTimeZonesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedTimeZonesWithTimeZoneStandard builds and executes requests for operations under \me\outlook\microsoft.graph.supportedTimeZones(TimeZoneStandard={TimeZoneStandard})
func (m *OutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*i767cd2a73c29879daa39ae90aebfabcd397159b1e720933714a03e3a5fceb72a.SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return i767cd2a73c29879daa39ae90aebfabcd397159b1e720933714a03e3a5fceb72a.NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.pathParameters, m.requestAdapter, timeZoneStandard);
}
func (m *OutlookRequestBuilder) TaskFolders()(*i82b391a4153948653004ca349b6eac39e2c3e7b3926201627e2e3eea88aad251.TaskFoldersRequestBuilder) {
    return i82b391a4153948653004ca349b6eac39e2c3e7b3926201627e2e3eea88aad251.NewTaskFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item collection
func (m *OutlookRequestBuilder) TaskFoldersById(id string)(*ic19e499d3e4541329156d6926f4f6db570dcce60051f4e1edd3f0b0db81cff9a.OutlookTaskFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskFolder_id"] = id
    }
    return ic19e499d3e4541329156d6926f4f6db570dcce60051f4e1edd3f0b0db81cff9a.NewOutlookTaskFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookRequestBuilder) TaskGroups()(*i7832228cb9bca400c299eb3f31f40734874357edc378956e69605b0e3b39ef21.TaskGroupsRequestBuilder) {
    return i7832228cb9bca400c299eb3f31f40734874357edc378956e69605b0e3b39ef21.NewTaskGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskGroups.item collection
func (m *OutlookRequestBuilder) TaskGroupsById(id string)(*i984fa253a0580cb2d81d93e449a106836d455f23e065de8d03edfcc8dbf44acb.OutlookTaskGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTaskGroup_id"] = id
    }
    return i984fa253a0580cb2d81d93e449a106836d455f23e065de8d03edfcc8dbf44acb.NewOutlookTaskGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookRequestBuilder) Tasks()(*ibf3b7a436ea0ea400b8c6c0590a9be4465672faf619ae54e97e418017b25ef00.TasksRequestBuilder) {
    return ibf3b7a436ea0ea400b8c6c0590a9be4465672faf619ae54e97e418017b25ef00.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.tasks.item collection
func (m *OutlookRequestBuilder) TasksById(id string)(*id3508c1f1aa1eff5970c7a4fe8b1bb7b16a0d219f157caf9278141e89295cd00.OutlookTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return id3508c1f1aa1eff5970c7a4fe8b1bb7b16a0d219f157caf9278141e89295cd00.NewOutlookTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

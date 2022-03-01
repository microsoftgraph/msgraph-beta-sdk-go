package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0227b9d371a89beb62067a796c02c0b39fdf080a3997ff8e2e1bccd0df359555 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/driverinventories"
    i0d5104ae2038733f4a0ba840268b8369f6bc121826293eab58ba88947073f553 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/assign"
    i7d322cde9c6dc3cbc3b99c3bd99e285cfd57d7d19bb515fe2fe8f2e1362010b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/executeaction"
    idf402f71068be00c76511f89c1a60aaa1debfdf51df203d20c1f5ace9153aa77 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/syncinventory"
    ifa4b3a3995a95d8ea842462ce1688b377028489725a54208f7a79be75071c764 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/assignments"
    i34e12a4e98dbefb122831ccf78a7c140934f5de1ca516a3631404c0f4718c186 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/driverinventories/item"
    i823e1a7037c21e1cb6a062bc58c9b7f7aacc3516364a051cdbc689c661c6ff50 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/assignments/item"
)

// WindowsDriverUpdateProfileItemRequestBuilder builds and executes requests for operations under \deviceManagement\windowsDriverUpdateProfiles\{windowsDriverUpdateProfile-id}
type WindowsDriverUpdateProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsDriverUpdateProfileItemRequestBuilderDeleteOptions options for Delete
type WindowsDriverUpdateProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsDriverUpdateProfileItemRequestBuilderGetOptions options for Get
type WindowsDriverUpdateProfileItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters a collection of windows driver update profiles
type WindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsDriverUpdateProfileItemRequestBuilderPatchOptions options for Patch
type WindowsDriverUpdateProfileItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDriverUpdateProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WindowsDriverUpdateProfileItemRequestBuilder) Assign()(*i0d5104ae2038733f4a0ba840268b8369f6bc121826293eab58ba88947073f553.AssignRequestBuilder) {
    return i0d5104ae2038733f4a0ba840268b8369f6bc121826293eab58ba88947073f553.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsDriverUpdateProfileItemRequestBuilder) Assignments()(*ifa4b3a3995a95d8ea842462ce1688b377028489725a54208f7a79be75071c764.AssignmentsRequestBuilder) {
    return ifa4b3a3995a95d8ea842462ce1688b377028489725a54208f7a79be75071c764.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.windowsDriverUpdateProfiles.item.assignments.item collection
func (m *WindowsDriverUpdateProfileItemRequestBuilder) AssignmentsById(id string)(*i823e1a7037c21e1cb6a062bc58c9b7f7aacc3516364a051cdbc689c661c6ff50.WindowsDriverUpdateProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDriverUpdateProfileAssignment_id"] = id
    }
    return i823e1a7037c21e1cb6a062bc58c9b7f7aacc3516364a051cdbc689c661c6ff50.NewWindowsDriverUpdateProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWindowsDriverUpdateProfileItemRequestBuilderInternal instantiates a new WindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsDriverUpdateProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDriverUpdateProfileItemRequestBuilder) {
    m := &WindowsDriverUpdateProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsDriverUpdateProfiles/{windowsDriverUpdateProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsDriverUpdateProfileItemRequestBuilder instantiates a new WindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsDriverUpdateProfileItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDriverUpdateProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsDriverUpdateProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfileItemRequestBuilder) CreateDeleteRequestInformation(options *WindowsDriverUpdateProfileItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfileItemRequestBuilder) CreateGetRequestInformation(options *WindowsDriverUpdateProfileItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfileItemRequestBuilder) CreatePatchRequestInformation(options *WindowsDriverUpdateProfileItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfileItemRequestBuilder) Delete(options *WindowsDriverUpdateProfileItemRequestBuilderDeleteOptions)(error) {
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
func (m *WindowsDriverUpdateProfileItemRequestBuilder) DriverInventories()(*i0227b9d371a89beb62067a796c02c0b39fdf080a3997ff8e2e1bccd0df359555.DriverInventoriesRequestBuilder) {
    return i0227b9d371a89beb62067a796c02c0b39fdf080a3997ff8e2e1bccd0df359555.NewDriverInventoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DriverInventoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.windowsDriverUpdateProfiles.item.driverInventories.item collection
func (m *WindowsDriverUpdateProfileItemRequestBuilder) DriverInventoriesById(id string)(*i34e12a4e98dbefb122831ccf78a7c140934f5de1ca516a3631404c0f4718c186.WindowsDriverUpdateInventoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDriverUpdateInventory_id"] = id
    }
    return i34e12a4e98dbefb122831ccf78a7c140934f5de1ca516a3631404c0f4718c186.NewWindowsDriverUpdateInventoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WindowsDriverUpdateProfileItemRequestBuilder) ExecuteAction()(*i7d322cde9c6dc3cbc3b99c3bd99e285cfd57d7d19bb515fe2fe8f2e1362010b9.ExecuteActionRequestBuilder) {
    return i7d322cde9c6dc3cbc3b99c3bd99e285cfd57d7d19bb515fe2fe8f2e1362010b9.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfileItemRequestBuilder) Get(options *WindowsDriverUpdateProfileItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDriverUpdateProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsDriverUpdateProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDriverUpdateProfile), nil
}
// Patch a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfileItemRequestBuilder) Patch(options *WindowsDriverUpdateProfileItemRequestBuilderPatchOptions)(error) {
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
func (m *WindowsDriverUpdateProfileItemRequestBuilder) SyncInventory()(*idf402f71068be00c76511f89c1a60aaa1debfdf51df203d20c1f5ace9153aa77.SyncInventoryRequestBuilder) {
    return idf402f71068be00c76511f89c1a60aaa1debfdf51df203d20c1f5ace9153aa77.NewSyncInventoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0227b9d371a89beb62067a796c02c0b39fdf080a3997ff8e2e1bccd0df359555 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/driverinventories"
    i0d5104ae2038733f4a0ba840268b8369f6bc121826293eab58ba88947073f553 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/assign"
    i7d322cde9c6dc3cbc3b99c3bd99e285cfd57d7d19bb515fe2fe8f2e1362010b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/executeaction"
    idf402f71068be00c76511f89c1a60aaa1debfdf51df203d20c1f5ace9153aa77 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/syncinventory"
    ifa4b3a3995a95d8ea842462ce1688b377028489725a54208f7a79be75071c764 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/assignments"
    i34e12a4e98dbefb122831ccf78a7c140934f5de1ca516a3631404c0f4718c186 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/driverinventories/item"
    i823e1a7037c21e1cb6a062bc58c9b7f7aacc3516364a051cdbc689c661c6ff50 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsdriverupdateprofiles/item/assignments/item"
)

// WindowsDriverUpdateProfileItemRequestBuilder provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
type WindowsDriverUpdateProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsDriverUpdateProfileItemRequestBuilderDeleteOptions options for Delete
type WindowsDriverUpdateProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// WindowsDriverUpdateProfileItemRequestBuilderGetOptions options for Get
type WindowsDriverUpdateProfileItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// WindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters a collection of windows driver update profiles
type WindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsDriverUpdateProfileItemRequestBuilderPatchOptions options for Patch
type WindowsDriverUpdateProfileItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Assign the assign property
func (m *WindowsDriverUpdateProfileItemRequestBuilder) Assign()(*i0d5104ae2038733f4a0ba840268b8369f6bc121826293eab58ba88947073f553.AssignRequestBuilder) {
    return i0d5104ae2038733f4a0ba840268b8369f6bc121826293eab58ba88947073f553.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
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
        urlTplParams["windowsDriverUpdateProfileAssignment%2Did"] = id
    }
    return i823e1a7037c21e1cb6a062bc58c9b7f7aacc3516364a051cdbc689c661c6ff50.NewWindowsDriverUpdateProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWindowsDriverUpdateProfileItemRequestBuilderInternal instantiates a new WindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsDriverUpdateProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsDriverUpdateProfileItemRequestBuilder) {
    m := &WindowsDriverUpdateProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsDriverUpdateProfiles/{windowsDriverUpdateProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsDriverUpdateProfileItemRequestBuilder instantiates a new WindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsDriverUpdateProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsDriverUpdateProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsDriverUpdateProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property windowsDriverUpdateProfiles for deviceManagement
func (m *WindowsDriverUpdateProfileItemRequestBuilder) CreateDeleteRequestInformation(options *WindowsDriverUpdateProfileItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfileItemRequestBuilder) CreateGetRequestInformation(options *WindowsDriverUpdateProfileItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property windowsDriverUpdateProfiles in deviceManagement
func (m *WindowsDriverUpdateProfileItemRequestBuilder) CreatePatchRequestInformation(options *WindowsDriverUpdateProfileItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property windowsDriverUpdateProfiles for deviceManagement
func (m *WindowsDriverUpdateProfileItemRequestBuilder) Delete(options *WindowsDriverUpdateProfileItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DriverInventories the driverInventories property
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
        urlTplParams["windowsDriverUpdateInventory%2Did"] = id
    }
    return i34e12a4e98dbefb122831ccf78a7c140934f5de1ca516a3631404c0f4718c186.NewWindowsDriverUpdateInventoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExecuteAction the executeAction property
func (m *WindowsDriverUpdateProfileItemRequestBuilder) ExecuteAction()(*i7d322cde9c6dc3cbc3b99c3bd99e285cfd57d7d19bb515fe2fe8f2e1362010b9.ExecuteActionRequestBuilder) {
    return i7d322cde9c6dc3cbc3b99c3bd99e285cfd57d7d19bb515fe2fe8f2e1362010b9.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfileItemRequestBuilder) Get(options *WindowsDriverUpdateProfileItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateProfileFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable), nil
}
// Patch update the navigation property windowsDriverUpdateProfiles in deviceManagement
func (m *WindowsDriverUpdateProfileItemRequestBuilder) Patch(options *WindowsDriverUpdateProfileItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SyncInventory the syncInventory property
func (m *WindowsDriverUpdateProfileItemRequestBuilder) SyncInventory()(*idf402f71068be00c76511f89c1a60aaa1debfdf51df203d20c1f5ace9153aa77.SyncInventoryRequestBuilder) {
    return idf402f71068be00c76511f89c1a60aaa1debfdf51df203d20c1f5ace9153aa77.NewSyncInventoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

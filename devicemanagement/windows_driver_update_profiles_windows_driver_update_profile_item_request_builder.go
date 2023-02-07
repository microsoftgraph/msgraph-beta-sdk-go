package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
type WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters a collection of windows driver update profiles
type WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters
}
// WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsDriverUpdateProfile entity.
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) Assignments()(*WindowsDriverUpdateProfilesItemAssignmentsRequestBuilder) {
    return NewWindowsDriverUpdateProfilesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.windowsDriverUpdateProfile entity.
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) AssignmentsById(id string)(*WindowsDriverUpdateProfilesItemAssignmentsWindowsDriverUpdateProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDriverUpdateProfileAssignment%2Did"] = id
    }
    return NewWindowsDriverUpdateProfilesItemAssignmentsWindowsDriverUpdateProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderInternal instantiates a new WindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) {
    m := &WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsDriverUpdateProfiles/{windowsDriverUpdateProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder instantiates a new WindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsDriverUpdateProfiles for deviceManagement
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DriverInventories provides operations to manage the driverInventories property of the microsoft.graph.windowsDriverUpdateProfile entity.
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) DriverInventories()(*WindowsDriverUpdateProfilesItemDriverInventoriesRequestBuilder) {
    return NewWindowsDriverUpdateProfilesItemDriverInventoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DriverInventoriesById provides operations to manage the driverInventories property of the microsoft.graph.windowsDriverUpdateProfile entity.
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) DriverInventoriesById(id string)(*WindowsDriverUpdateProfilesItemDriverInventoriesWindowsDriverUpdateInventoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDriverUpdateInventory%2Did"] = id
    }
    return NewWindowsDriverUpdateProfilesItemDriverInventoriesWindowsDriverUpdateInventoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable), nil
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) MicrosoftGraphAssign()(*WindowsDriverUpdateProfilesItemMicrosoftGraphAssignRequestBuilder) {
    return NewWindowsDriverUpdateProfilesItemMicrosoftGraphAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphExecuteAction provides operations to call the executeAction method.
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) MicrosoftGraphExecuteAction()(*WindowsDriverUpdateProfilesItemMicrosoftGraphExecuteActionRequestBuilder) {
    return NewWindowsDriverUpdateProfilesItemMicrosoftGraphExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSyncInventory provides operations to call the syncInventory method.
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) MicrosoftGraphSyncInventory()(*WindowsDriverUpdateProfilesItemMicrosoftGraphSyncInventoryRequestBuilder) {
    return NewWindowsDriverUpdateProfilesItemMicrosoftGraphSyncInventoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property windowsDriverUpdateProfiles in deviceManagement
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, requestConfiguration *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable), nil
}
// ToDeleteRequestInformation delete navigation property windowsDriverUpdateProfiles for deviceManagement
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation a collection of windows driver update profiles
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsDriverUpdateProfiles in deviceManagement
func (m *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, requestConfiguration *WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
type WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters a collection of windows driver update profiles
type WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters
}
// WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *WindowsdriverupdateprofilesItemAssignRequestBuilder when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) Assign()(*WindowsdriverupdateprofilesItemAssignRequestBuilder) {
    return NewWindowsdriverupdateprofilesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsDriverUpdateProfile entity.
// returns a *WindowsdriverupdateprofilesItemAssignmentsRequestBuilder when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) Assignments()(*WindowsdriverupdateprofilesItemAssignmentsRequestBuilder) {
    return NewWindowsdriverupdateprofilesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderInternal instantiates a new WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) {
    m := &WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsDriverUpdateProfiles/{windowsDriverUpdateProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder instantiates a new WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsDriverUpdateProfiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DriverInventories provides operations to manage the driverInventories property of the microsoft.graph.windowsDriverUpdateProfile entity.
// returns a *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) DriverInventories()(*WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) {
    return NewWindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExecuteAction provides operations to call the executeAction method.
// returns a *WindowsdriverupdateprofilesItemExecuteactionExecuteActionRequestBuilder when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) ExecuteAction()(*WindowsdriverupdateprofilesItemExecuteactionExecuteActionRequestBuilder) {
    return NewWindowsdriverupdateprofilesItemExecuteactionExecuteActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of windows driver update profiles
// returns a WindowsDriverUpdateProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable), nil
}
// Patch update the navigation property windowsDriverUpdateProfiles in deviceManagement
// returns a WindowsDriverUpdateProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable), nil
}
// SyncInventory provides operations to call the syncInventory method.
// returns a *WindowsdriverupdateprofilesItemSyncinventorySyncInventoryRequestBuilder when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) SyncInventory()(*WindowsdriverupdateprofilesItemSyncinventorySyncInventoryRequestBuilder) {
    return NewWindowsdriverupdateprofilesItemSyncinventorySyncInventoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property windowsDriverUpdateProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of windows driver update profiles
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsDriverUpdateProfiles in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) WithUrl(rawUrl string)(*WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) {
    return NewWindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

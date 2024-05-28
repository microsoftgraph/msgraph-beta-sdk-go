package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder provides operations to manage the driverInventories property of the microsoft.graph.windowsDriverUpdateProfile entity.
type WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderGetQueryParameters driver inventories for this profile.
type WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderGetQueryParameters
}
// WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderInternal instantiates a new WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder and sets the default values.
func NewWindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) {
    m := &WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsDriverUpdateProfiles/{windowsDriverUpdateProfile%2Did}/driverInventories/{windowsDriverUpdateInventory%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder instantiates a new WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder and sets the default values.
func NewWindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property driverInventories for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get driver inventories for this profile.
// returns a WindowsDriverUpdateInventoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateInventoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryable), nil
}
// Patch update the navigation property driverInventories in deviceManagement
// returns a WindowsDriverUpdateInventoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryable, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateInventoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryable), nil
}
// ToDeleteRequestInformation delete navigation property driverInventories for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation driver inventories for this profile.
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property driverInventories in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryable, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder when successful
func (m *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) WithUrl(rawUrl string)(*WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) {
    return NewWindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder provides operations to manage the driverInventories property of the microsoft.graph.windowsDriverUpdateProfile entity.
type WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderGetQueryParameters driver inventories for this profile.
type WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderGetQueryParameters
}
// WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsDriverUpdateInventoryId provides operations to manage the driverInventories property of the microsoft.graph.windowsDriverUpdateProfile entity.
// returns a *WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder when successful
func (m *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) ByWindowsDriverUpdateInventoryId(windowsDriverUpdateInventoryId string)(*WindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsDriverUpdateInventoryId != "" {
        urlTplParams["windowsDriverUpdateInventory%2Did"] = windowsDriverUpdateInventoryId
    }
    return NewWindowsdriverupdateprofilesItemDriverinventoriesWindowsDriverUpdateInventoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderInternal instantiates a new WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder and sets the default values.
func NewWindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) {
    m := &WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsDriverUpdateProfiles/{windowsDriverUpdateProfile%2Did}/driverInventories{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder instantiates a new WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder and sets the default values.
func NewWindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsdriverupdateprofilesItemDriverinventoriesCountRequestBuilder when successful
func (m *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) Count()(*WindowsdriverupdateprofilesItemDriverinventoriesCountRequestBuilder) {
    return NewWindowsdriverupdateprofilesItemDriverinventoriesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get driver inventories for this profile.
// returns a WindowsDriverUpdateInventoryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateInventoryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryCollectionResponseable), nil
}
// Post create new navigation property to driverInventories for deviceManagement
// returns a WindowsDriverUpdateInventoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryable, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation driver inventories for this profile.
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to driverInventories for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateInventoryable, requestConfiguration *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder when successful
func (m *WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) WithUrl(rawUrl string)(*WindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder) {
    return NewWindowsdriverupdateprofilesItemDriverinventoriesDriverInventoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

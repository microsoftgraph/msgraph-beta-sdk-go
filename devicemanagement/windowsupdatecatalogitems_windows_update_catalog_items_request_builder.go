package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
type WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderGetQueryParameters a collection of windows update catalog items (fetaure updates item , quality updates item)
type WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderGetQueryParameters struct {
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
// WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderGetQueryParameters
}
// WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsUpdateCatalogItemId provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder when successful
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) ByWindowsUpdateCatalogItemId(windowsUpdateCatalogItemId string)(*WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsUpdateCatalogItemId != "" {
        urlTplParams["windowsUpdateCatalogItem%2Did"] = windowsUpdateCatalogItemId
    }
    return NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderInternal instantiates a new WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder and sets the default values.
func NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) {
    m := &WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsUpdateCatalogItems{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder instantiates a new WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder and sets the default values.
func NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsupdatecatalogitemsCountRequestBuilder when successful
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) Count()(*WindowsupdatecatalogitemsCountRequestBuilder) {
    return NewWindowsupdatecatalogitemsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of windows update catalog items (fetaure updates item , quality updates item)
// returns a WindowsUpdateCatalogItemCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsUpdateCatalogItemCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemCollectionResponseable), nil
}
// Post create new navigation property to windowsUpdateCatalogItems for deviceManagement
// returns a WindowsUpdateCatalogItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsUpdateCatalogItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable), nil
}
// ToGetRequestInformation a collection of windows update catalog items (fetaure updates item , quality updates item)
// returns a *RequestInformation when successful
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to windowsUpdateCatalogItems for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder when successful
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) WithUrl(rawUrl string)(*WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) {
    return NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

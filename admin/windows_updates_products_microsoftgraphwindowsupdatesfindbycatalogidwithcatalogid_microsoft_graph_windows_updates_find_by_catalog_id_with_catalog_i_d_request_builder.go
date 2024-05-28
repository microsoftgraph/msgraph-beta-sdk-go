package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder provides operations to call the findByCatalogId method.
type WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderGetQueryParameters invoke function findByCatalogId
type WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderGetQueryParameters struct {
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
// WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderGetQueryParameters
}
// NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderInternal instantiates a new WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, catalogID *string)(*WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder) {
    m := &WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/products/microsoft.graph.windowsUpdates.findByCatalogId(catalogID='{catalogID}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if catalogID != nil {
        m.BaseRequestBuilder.PathParameters["catalogID"] = *catalogID
    }
    return m
}
// NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder instantiates a new WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function findByCatalogId
// Deprecated: This method is obsolete. Use GetAsFindByCatalogIdWithCatalogIDGetResponse instead.
// returns a WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidFindByCatalogIdWithCatalogIDResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderGetRequestConfiguration)(WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidFindByCatalogIdWithCatalogIDResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidFindByCatalogIdWithCatalogIDResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidFindByCatalogIdWithCatalogIDResponseable), nil
}
// GetAsFindByCatalogIdWithCatalogIDGetResponse invoke function findByCatalogId
// returns a WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidFindByCatalogIdWithCatalogIDGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder) GetAsFindByCatalogIdWithCatalogIDGetResponse(ctx context.Context, requestConfiguration *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderGetRequestConfiguration)(WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidFindByCatalogIdWithCatalogIDGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidFindByCatalogIdWithCatalogIDGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidFindByCatalogIdWithCatalogIDGetResponseable), nil
}
// ToGetRequestInformation invoke function findByCatalogId
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder when successful
func (m *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder) {
    return NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbycatalogidwithcatalogidMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

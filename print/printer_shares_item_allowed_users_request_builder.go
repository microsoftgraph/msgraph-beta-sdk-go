package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrinterSharesItemAllowedUsersRequestBuilder provides operations to manage the allowedUsers property of the microsoft.graph.printerShare entity.
type PrinterSharesItemAllowedUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrinterSharesItemAllowedUsersRequestBuilderGetQueryParameters the users who have access to print using the printer.
type PrinterSharesItemAllowedUsersRequestBuilderGetQueryParameters struct {
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
// PrinterSharesItemAllowedUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrinterSharesItemAllowedUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrinterSharesItemAllowedUsersRequestBuilderGetQueryParameters
}
// ByUserId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printerShares.item.allowedUsers.item collection
// returns a *PrinterSharesItemAllowedUsersUserItemRequestBuilder when successful
func (m *PrinterSharesItemAllowedUsersRequestBuilder) ByUserId(userId string)(*PrinterSharesItemAllowedUsersUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userId != "" {
        urlTplParams["user%2Did"] = userId
    }
    return NewPrinterSharesItemAllowedUsersUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrinterSharesItemAllowedUsersRequestBuilderInternal instantiates a new PrinterSharesItemAllowedUsersRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedUsersRequestBuilder) {
    m := &PrinterSharesItemAllowedUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedUsers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrinterSharesItemAllowedUsersRequestBuilder instantiates a new PrinterSharesItemAllowedUsersRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterSharesItemAllowedUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrinterSharesItemAllowedUsersCountRequestBuilder when successful
func (m *PrinterSharesItemAllowedUsersRequestBuilder) Count()(*PrinterSharesItemAllowedUsersCountRequestBuilder) {
    return NewPrinterSharesItemAllowedUsersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the users who have access to print using the printer.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a UserCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrinterSharesItemAllowedUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *PrinterSharesItemAllowedUsersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCollectionResponseable), nil
}
// Ref provides operations to manage the collection of print entities.
// returns a *PrinterSharesItemAllowedUsersRefRequestBuilder when successful
func (m *PrinterSharesItemAllowedUsersRequestBuilder) Ref()(*PrinterSharesItemAllowedUsersRefRequestBuilder) {
    return NewPrinterSharesItemAllowedUsersRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the users who have access to print using the printer.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *PrinterSharesItemAllowedUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrinterSharesItemAllowedUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *PrinterSharesItemAllowedUsersRequestBuilder when successful
func (m *PrinterSharesItemAllowedUsersRequestBuilder) WithUrl(rawUrl string)(*PrinterSharesItemAllowedUsersRequestBuilder) {
    return NewPrinterSharesItemAllowedUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

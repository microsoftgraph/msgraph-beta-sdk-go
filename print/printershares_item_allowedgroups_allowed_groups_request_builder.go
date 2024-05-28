package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder provides operations to manage the allowedGroups property of the microsoft.graph.printerShare entity.
type PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderGetQueryParameters the groups whose users have access to print using the printer.
type PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderGetQueryParameters struct {
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
// PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderGetQueryParameters
}
// ByGroupId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printerShares.item.allowedGroups.item collection
// returns a *PrintersharesItemAllowedgroupsGroupItemRequestBuilder when successful
func (m *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder) ByGroupId(groupId string)(*PrintersharesItemAllowedgroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupId != "" {
        urlTplParams["group%2Did"] = groupId
    }
    return NewPrintersharesItemAllowedgroupsGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderInternal instantiates a new PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder and sets the default values.
func NewPrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder) {
    m := &PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder instantiates a new PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder and sets the default values.
func NewPrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrintersharesItemAllowedgroupsCountRequestBuilder when successful
func (m *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder) Count()(*PrintersharesItemAllowedgroupsCountRequestBuilder) {
    return NewPrintersharesItemAllowedgroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the groups whose users have access to print using the printer.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a GroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable), nil
}
// Ref provides operations to manage the collection of print entities.
// returns a *PrintersharesItemAllowedgroupsRefRequestBuilder when successful
func (m *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder) Ref()(*PrintersharesItemAllowedgroupsRefRequestBuilder) {
    return NewPrintersharesItemAllowedgroupsRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the groups whose users have access to print using the printer.
// Deprecated: The printerShares navigation property is deprecated and will stop returning data on July 31, 2023. Please use the shares navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder when successful
func (m *PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder) WithUrl(rawUrl string)(*PrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder) {
    return NewPrintersharesItemAllowedgroupsAllowedGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

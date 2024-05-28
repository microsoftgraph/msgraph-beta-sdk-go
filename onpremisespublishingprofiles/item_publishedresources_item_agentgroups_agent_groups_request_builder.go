package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder provides operations to manage the agentGroups property of the microsoft.graph.publishedResource entity.
type ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderGetQueryParameters list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
type ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderGetQueryParameters struct {
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
// ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderGetQueryParameters
}
// ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOnPremisesAgentGroupId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.publishedResources.item.agentGroups.item collection
// returns a *ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder when successful
func (m *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) ByOnPremisesAgentGroupId(onPremisesAgentGroupId string)(*ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if onPremisesAgentGroupId != "" {
        urlTplParams["onPremisesAgentGroup%2Did"] = onPremisesAgentGroupId
    }
    return NewItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderInternal instantiates a new ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder and sets the default values.
func NewItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) {
    m := &ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/publishedResources/{publishedResource%2Did}/agentGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder instantiates a new ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder and sets the default values.
func NewItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemPublishedresourcesItemAgentgroupsCountRequestBuilder when successful
func (m *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) Count()(*ItemPublishedresourcesItemAgentgroupsCountRequestBuilder) {
    return NewItemPublishedresourcesItemAgentgroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
// returns a OnPremisesAgentGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesAgentGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupCollectionResponseable), nil
}
// Post create new navigation property to agentGroups for onPremisesPublishingProfiles
// returns a OnPremisesAgentGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable, requestConfiguration *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnPremisesAgentGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable), nil
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemPublishedresourcesItemAgentgroupsRefRequestBuilder when successful
func (m *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) Ref()(*ItemPublishedresourcesItemAgentgroupsRefRequestBuilder) {
    return NewItemPublishedresourcesItemAgentgroupsRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to agentGroups for onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupable, requestConfiguration *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder when successful
func (m *ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder) {
    return NewItemPublishedresourcesItemAgentgroupsAgentGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder provides operations to manage the agentGroups property of the microsoft.graph.onPremisesAgent entity.
type ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderGetQueryParameters list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
type ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderGetQueryParameters struct {
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
// ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderGetQueryParameters
}
// ByOnPremisesAgentGroupId1 gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agentGroups.item.agents.item.agentGroups.item collection
// returns a *ItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder when successful
func (m *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) ByOnPremisesAgentGroupId1(onPremisesAgentGroupId1 string)(*ItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if onPremisesAgentGroupId1 != "" {
        urlTplParams["onPremisesAgentGroup%2Did1"] = onPremisesAgentGroupId1
    }
    return NewItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderInternal instantiates a new ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder and sets the default values.
func NewItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) {
    m := &ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/agents/{onPremisesAgent%2Did}/agentGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder instantiates a new ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder and sets the default values.
func NewItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAgentGroupsItemAgentsItemAgentGroupsCountRequestBuilder when successful
func (m *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) Count()(*ItemAgentGroupsItemAgentsItemAgentGroupsCountRequestBuilder) {
    return NewItemAgentGroupsItemAgentsItemAgentGroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
// returns a OnPremisesAgentGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnPremisesAgentGroupCollectionResponseable, error) {
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
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemAgentGroupsItemAgentsItemAgentGroupsRefRequestBuilder when successful
func (m *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) Ref()(*ItemAgentGroupsItemAgentsItemAgentGroupsRefRequestBuilder) {
    return NewItemAgentGroupsItemAgentsItemAgentGroupsRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder when successful
func (m *ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder) {
    return NewItemAgentGroupsItemAgentsItemAgentGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

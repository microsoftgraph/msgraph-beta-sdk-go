package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\publishedResources\{publishedResource-id}\agentGroups\{onPremisesAgentGroup-id}
type ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    m := &ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/publishedResources/{publishedResource%2Did}/agentGroups/{onPremisesAgentGroup%2Did}", pathParameters),
    }
    return m
}
// NewItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder instantiates a new ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemPublishedresourcesItemAgentgroupsItemRefRequestBuilder when successful
func (m *ItemPublishedresourcesItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) Ref()(*ItemPublishedresourcesItemAgentgroupsItemRefRequestBuilder) {
    return NewItemPublishedresourcesItemAgentgroupsItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

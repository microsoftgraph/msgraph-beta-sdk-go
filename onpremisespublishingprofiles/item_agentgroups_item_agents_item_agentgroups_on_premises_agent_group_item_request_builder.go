package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\agentGroups\{onPremisesAgentGroup-id}\agents\{onPremisesAgent-id}\agentGroups\{onPremisesAgentGroup-id1}
type ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    m := &ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/agents/{onPremisesAgent%2Did}/agentGroups/{onPremisesAgentGroup%2Did1}", pathParameters),
    }
    return m
}
// NewItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder instantiates a new ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemAgentgroupsItemAgentsItemAgentgroupsItemRefRequestBuilder when successful
func (m *ItemAgentgroupsItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) Ref()(*ItemAgentgroupsItemAgentsItemAgentgroupsItemRefRequestBuilder) {
    return NewItemAgentgroupsItemAgentsItemAgentgroupsItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

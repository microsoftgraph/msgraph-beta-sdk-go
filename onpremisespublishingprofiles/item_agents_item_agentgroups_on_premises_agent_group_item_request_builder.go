package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\agents\{onPremisesAgent-id}\agentGroups\{onPremisesAgentGroup-id}
type ItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new ItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    m := &ItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agents/{onPremisesAgent%2Did}/agentGroups/{onPremisesAgentGroup%2Did}", pathParameters),
    }
    return m
}
// NewItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder instantiates a new ItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemAgentsItemAgentgroupsItemRefRequestBuilder when successful
func (m *ItemAgentsItemAgentgroupsOnPremisesAgentGroupItemRequestBuilder) Ref()(*ItemAgentsItemAgentgroupsItemRefRequestBuilder) {
    return NewItemAgentsItemAgentgroupsItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

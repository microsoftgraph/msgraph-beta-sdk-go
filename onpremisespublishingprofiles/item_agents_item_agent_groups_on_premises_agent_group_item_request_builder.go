package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\agents\{onPremisesAgent-id}\agentGroups\{onPremisesAgentGroup-id}
type ItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    m := &ItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agents/{onPremisesAgent%2Did}/agentGroups/{onPremisesAgentGroup%2Did}", pathParameters),
    }
    return m
}
// NewItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *ItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) Ref()(*ItemAgentsItemAgentGroupsItemRefRequestBuilder) {
    return NewItemAgentsItemAgentGroupsItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

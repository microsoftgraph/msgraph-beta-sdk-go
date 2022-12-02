package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\agentGroups\{onPremisesAgentGroup-id}\agents\{onPremisesAgent-id}\agentGroups\{onPremisesAgentGroup-id1}
type OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/agents/{onPremisesAgent%2Did}/agentGroups/{onPremisesAgentGroup%2Did1}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) Ref()(*OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder) {
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

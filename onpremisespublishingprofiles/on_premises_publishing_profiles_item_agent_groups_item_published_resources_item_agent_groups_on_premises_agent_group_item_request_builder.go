package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\agentGroups\{onPremisesAgentGroup-id}\publishedResources\{publishedResource-id}\agentGroups\{onPremisesAgentGroup-id1}
type OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/publishedResources/{publishedResource%2Did}/agentGroups/{onPremisesAgentGroup%2Did1}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) Ref()(*OnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsItemRefRequestBuilder) {
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemPublishedResourcesItemAgentGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

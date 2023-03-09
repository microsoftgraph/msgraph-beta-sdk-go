package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\agentGroups\{onPremisesAgentGroup-id}\publishedResources\{publishedResource-id}\agentGroups\{onPremisesAgentGroup-id1}
type ItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    m := &ItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/publishedResources/{publishedResource%2Did}/agentGroups/{onPremisesAgentGroup%2Did1}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *ItemAgentGroupsItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) Ref()(*ItemAgentGroupsItemPublishedResourcesItemAgentGroupsItemRefRequestBuilder) {
    return NewItemAgentGroupsItemPublishedResourcesItemAgentGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

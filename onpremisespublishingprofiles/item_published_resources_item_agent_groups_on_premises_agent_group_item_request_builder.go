package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\publishedResources\{publishedResource-id}\agentGroups\{onPremisesAgentGroup-id}
type ItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    m := &ItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/publishedResources/{publishedResource%2Did}/agentGroups/{onPremisesAgentGroup%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder instantiates a new OnPremisesAgentGroupItemRequestBuilder and sets the default values.
func NewItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *ItemPublishedResourcesItemAgentGroupsOnPremisesAgentGroupItemRequestBuilder) Ref()(*ItemPublishedResourcesItemAgentGroupsItemRefRequestBuilder) {
    return NewItemPublishedResourcesItemAgentGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

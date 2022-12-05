package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectorGroups\{connectorGroup-id}\members\{connector-id}
type OnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewOnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal instantiates a new ConnectorItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}/members/{connector%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilder instantiates a new ConnectorItemRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *OnPremisesPublishingProfilesItemConnectorGroupsItemMembersConnectorItemRequestBuilder) Ref()(*OnPremisesPublishingProfilesItemConnectorGroupsItemMembersItemRefRequestBuilder) {
    return NewOnPremisesPublishingProfilesItemConnectorGroupsItemMembersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

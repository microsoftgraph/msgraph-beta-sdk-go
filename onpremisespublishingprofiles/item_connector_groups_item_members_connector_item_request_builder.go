package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemConnectorGroupsItemMembersConnectorItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectorGroups\{connectorGroup-id}\members\{connector-id}
type ItemConnectorGroupsItemMembersConnectorItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal instantiates a new ConnectorItemRequestBuilder and sets the default values.
func NewItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorGroupsItemMembersConnectorItemRequestBuilder) {
    m := &ItemConnectorGroupsItemMembersConnectorItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}/members/{connector%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemConnectorGroupsItemMembersConnectorItemRequestBuilder instantiates a new ConnectorItemRequestBuilder and sets the default values.
func NewItemConnectorGroupsItemMembersConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorGroupsItemMembersConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *ItemConnectorGroupsItemMembersConnectorItemRequestBuilder) Ref()(*ItemConnectorGroupsItemMembersItemRefRequestBuilder) {
    return NewItemConnectorGroupsItemMembersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

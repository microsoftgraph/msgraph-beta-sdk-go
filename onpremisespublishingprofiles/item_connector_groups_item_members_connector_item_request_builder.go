package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemConnectorGroupsItemMembersConnectorItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectorGroups\{connectorGroup-id}\members\{connector-id}
type ItemConnectorGroupsItemMembersConnectorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal instantiates a new ItemConnectorGroupsItemMembersConnectorItemRequestBuilder and sets the default values.
func NewItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorGroupsItemMembersConnectorItemRequestBuilder) {
    m := &ItemConnectorGroupsItemMembersConnectorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}/members/{connector%2Did}", pathParameters),
    }
    return m
}
// NewItemConnectorGroupsItemMembersConnectorItemRequestBuilder instantiates a new ItemConnectorGroupsItemMembersConnectorItemRequestBuilder and sets the default values.
func NewItemConnectorGroupsItemMembersConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorGroupsItemMembersConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorGroupsItemMembersConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemConnectorGroupsItemMembersItemRefRequestBuilder when successful
func (m *ItemConnectorGroupsItemMembersConnectorItemRequestBuilder) Ref()(*ItemConnectorGroupsItemMembersItemRefRequestBuilder) {
    return NewItemConnectorGroupsItemMembersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

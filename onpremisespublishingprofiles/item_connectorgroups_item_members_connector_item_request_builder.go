package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemConnectorgroupsItemMembersConnectorItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectorGroups\{connectorGroup-id}\members\{connector-id}
type ItemConnectorgroupsItemMembersConnectorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemConnectorgroupsItemMembersConnectorItemRequestBuilderInternal instantiates a new ItemConnectorgroupsItemMembersConnectorItemRequestBuilder and sets the default values.
func NewItemConnectorgroupsItemMembersConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorgroupsItemMembersConnectorItemRequestBuilder) {
    m := &ItemConnectorgroupsItemMembersConnectorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectorGroups/{connectorGroup%2Did}/members/{connector%2Did}", pathParameters),
    }
    return m
}
// NewItemConnectorgroupsItemMembersConnectorItemRequestBuilder instantiates a new ItemConnectorgroupsItemMembersConnectorItemRequestBuilder and sets the default values.
func NewItemConnectorgroupsItemMembersConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorgroupsItemMembersConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorgroupsItemMembersConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemConnectorgroupsItemMembersItemRefRequestBuilder when successful
func (m *ItemConnectorgroupsItemMembersConnectorItemRequestBuilder) Ref()(*ItemConnectorgroupsItemMembersItemRefRequestBuilder) {
    return NewItemConnectorgroupsItemMembersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

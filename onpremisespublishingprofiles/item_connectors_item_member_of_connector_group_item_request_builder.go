package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectors\{connector-id}\memberOf\{connectorGroup-id}
type ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilderInternal instantiates a new ConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder) {
    m := &ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectors/{connector%2Did}/memberOf/{connectorGroup%2Did}", pathParameters),
    }
    return m
}
// NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder instantiates a new ConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder) Ref()(*ItemConnectorsItemMemberOfItemRefRequestBuilder) {
    return NewItemConnectorsItemMemberOfItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

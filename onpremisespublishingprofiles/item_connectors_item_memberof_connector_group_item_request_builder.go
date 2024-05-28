package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectors\{connector-id}\memberOf\{connectorGroup-id}
type ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemConnectorsItemMemberofConnectorGroupItemRequestBuilderInternal instantiates a new ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberofConnectorGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder) {
    m := &ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectors/{connector%2Did}/memberOf/{connectorGroup%2Did}", pathParameters),
    }
    return m
}
// NewItemConnectorsItemMemberofConnectorGroupItemRequestBuilder instantiates a new ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberofConnectorGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorsItemMemberofConnectorGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemConnectorsItemMemberofItemRefRequestBuilder when successful
func (m *ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder) Ref()(*ItemConnectorsItemMemberofItemRefRequestBuilder) {
    return NewItemConnectorsItemMemberofItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

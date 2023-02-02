package onpremisespublishingprofiles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectors\{connector-id}\memberOf\{connectorGroup-id}
type ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilderInternal instantiates a new ConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, connectorGroupId *string)(*ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder) {
    m := &ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectors/{connector%2Did}/memberOf/{connectorGroup%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if connectorGroupId != nil {
        urlTplParams["connectorGroup%2Did"] = *connectorGroupId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder instantiates a new ConnectorGroupItemRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorsItemMemberOfConnectorGroupItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *ItemConnectorsItemMemberOfConnectorGroupItemRequestBuilder) Ref()(*ItemConnectorsItemMemberOfItemRefRequestBuilder) {
    return NewItemConnectorsItemMemberOfItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

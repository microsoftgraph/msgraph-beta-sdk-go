package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i07a3a2dae6578a77d572f0d2887fd37652ddd1c1e8fd473314791b42b3f4dfa5 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectors/item/memberof/item/ref"
)

// ConnectorGroupItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectors\{connector-id}\memberOf\{connectorGroup-id}
type ConnectorGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewConnectorGroupItemRequestBuilderInternal instantiates a new ConnectorGroupItemRequestBuilder and sets the default values.
func NewConnectorGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectorGroupItemRequestBuilder) {
    m := &ConnectorGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectors/{connector%2Did}/memberOf/{connectorGroup%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewConnectorGroupItemRequestBuilder instantiates a new ConnectorGroupItemRequestBuilder and sets the default values.
func NewConnectorGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectorGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectorGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *ConnectorGroupItemRequestBuilder) Ref()(*i07a3a2dae6578a77d572f0d2887fd37652ddd1c1e8fd473314791b42b3f4dfa5.RefRequestBuilder) {
    return i07a3a2dae6578a77d572f0d2887fd37652ddd1c1e8fd473314791b42b3f4dfa5.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

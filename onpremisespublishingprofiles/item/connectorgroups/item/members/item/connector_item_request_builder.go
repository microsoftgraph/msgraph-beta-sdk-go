package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i827f6d4387b03580057ce94a8e380577fc5238383ad0c9d34c6d8868e4a8d15e "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups/item/members/item/ref"
)

// ConnectorItemRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectorGroups\{connectorGroup-id}\members\{connector-id}
type ConnectorItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewConnectorItemRequestBuilderInternal instantiates a new ConnectorItemRequestBuilder and sets the default values.
func NewConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectorItemRequestBuilder) {
    m := &ConnectorItemRequestBuilder{
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
// NewConnectorItemRequestBuilder instantiates a new ConnectorItemRequestBuilder and sets the default values.
func NewConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref the Ref property
func (m *ConnectorItemRequestBuilder) Ref()(*i827f6d4387b03580057ce94a8e380577fc5238383ad0c9d34c6d8868e4a8d15e.RefRequestBuilder) {
    return i827f6d4387b03580057ce94a8e380577fc5238383ad0c9d34c6d8868e4a8d15e.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConnectorsItemMemberofMemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.connector entity.
type ItemConnectorsItemMemberofMemberOfRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConnectorsItemMemberofMemberOfRequestBuilderGetQueryParameters the connectorGroup that the connector is a member of. Read-only.
type ItemConnectorsItemMemberofMemberOfRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemConnectorsItemMemberofMemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorsItemMemberofMemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConnectorsItemMemberofMemberOfRequestBuilderGetQueryParameters
}
// ByConnectorGroupId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.connectors.item.memberOf.item collection
// returns a *ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder when successful
func (m *ItemConnectorsItemMemberofMemberOfRequestBuilder) ByConnectorGroupId(connectorGroupId string)(*ItemConnectorsItemMemberofConnectorGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if connectorGroupId != "" {
        urlTplParams["connectorGroup%2Did"] = connectorGroupId
    }
    return NewItemConnectorsItemMemberofConnectorGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemConnectorsItemMemberofMemberOfRequestBuilderInternal instantiates a new ItemConnectorsItemMemberofMemberOfRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberofMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorsItemMemberofMemberOfRequestBuilder) {
    m := &ItemConnectorsItemMemberofMemberOfRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectors/{connector%2Did}/memberOf{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemConnectorsItemMemberofMemberOfRequestBuilder instantiates a new ItemConnectorsItemMemberofMemberOfRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberofMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorsItemMemberofMemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorsItemMemberofMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemConnectorsItemMemberofCountRequestBuilder when successful
func (m *ItemConnectorsItemMemberofMemberOfRequestBuilder) Count()(*ItemConnectorsItemMemberofCountRequestBuilder) {
    return NewItemConnectorsItemMemberofCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the connectorGroup that the connector is a member of. Read-only.
// returns a ConnectorGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemConnectorsItemMemberofMemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConnectorsItemMemberofMemberOfRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConnectorGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectorGroupCollectionResponseable), nil
}
// Ref provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *ItemConnectorsItemMemberofRefRequestBuilder when successful
func (m *ItemConnectorsItemMemberofMemberOfRequestBuilder) Ref()(*ItemConnectorsItemMemberofRefRequestBuilder) {
    return NewItemConnectorsItemMemberofRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the connectorGroup that the connector is a member of. Read-only.
// returns a *RequestInformation when successful
func (m *ItemConnectorsItemMemberofMemberOfRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorsItemMemberofMemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemConnectorsItemMemberofMemberOfRequestBuilder when successful
func (m *ItemConnectorsItemMemberofMemberOfRequestBuilder) WithUrl(rawUrl string)(*ItemConnectorsItemMemberofMemberOfRequestBuilder) {
    return NewItemConnectorsItemMemberofMemberOfRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

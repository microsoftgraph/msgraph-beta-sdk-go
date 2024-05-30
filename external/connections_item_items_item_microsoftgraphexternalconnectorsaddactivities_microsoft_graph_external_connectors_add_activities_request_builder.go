package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder provides operations to call the addActivities method.
type ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderInternal instantiates a new ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder and sets the default values.
func NewConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) {
    m := &ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/connections/{externalConnection%2Did}/items/{externalItem%2Did}/microsoft.graph.externalConnectors.addActivities", pathParameters),
    }
    return m
}
// NewConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder instantiates a new ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder and sets the default values.
func NewConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addActivities
// Deprecated: This method is obsolete. Use PostAsAddActivitiesPostResponse instead.
// returns a ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) Post(ctx context.Context, body ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostRequestBodyable, requestConfiguration *ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration)(ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesResponseable), nil
}
// PostAsAddActivitiesPostResponse invoke action addActivities
// returns a ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) PostAsAddActivitiesPostResponse(ctx context.Context, body ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostRequestBodyable, requestConfiguration *ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration)(ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostResponseable), nil
}
// ToPostRequestInformation invoke action addActivities
// returns a *RequestInformation when successful
func (m *ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesAddActivitiesPostRequestBodyable, requestConfiguration *ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder when successful
func (m *ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) WithUrl(rawUrl string)(*ConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder) {
    return NewConnectionsItemItemsItemMicrosoftgraphexternalconnectorsaddactivitiesMicrosoftGraphExternalConnectorsAddActivitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder provides operations to call the delta method.
type DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderGetQueryParameters get the list of messages from all chats in which a user is a participant, including one-on-one chats, group chats, and meeting chats. When you use delta query, you can obtain new or updated messages. To get the replies for a message, use the list message replies or the get message reply operations. A GET request with the delta function returns one of the following: State tokens are opaque to the client. To proceed with a round of change tracking, copy and apply the @odata.nextLink or @odata.deltaLink URL returned from the last GET request to the next delta function call. An @odata.deltaLink returned in a response signifies that the current round of change tracking is complete. You can save and use the @odata.deltaLink URL when you begin to retrieve more changes (messages changed or posted after you acquire @odata.deltaLink). For more information, see the delta query documentation.
type DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderGetQueryParameters struct {
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
// DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderGetQueryParameters
}
// NewDeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderInternal instantiates a new DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder) {
    m := &DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/messages/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder instantiates a new DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the list of messages from all chats in which a user is a participant, including one-on-one chats, group chats, and meeting chats. When you use delta query, you can obtain new or updated messages. To get the replies for a message, use the list message replies or the get message reply operations. A GET request with the delta function returns one of the following: State tokens are opaque to the client. To proceed with a round of change tracking, copy and apply the @odata.nextLink or @odata.deltaLink URL returned from the last GET request to the next delta function call. An @odata.deltaLink returned in a response signifies that the current round of change tracking is complete. You can save and use the @odata.deltaLink URL when you begin to retrieve more changes (messages changed or posted after you acquire @odata.deltaLink). For more information, see the delta query documentation.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a DeletedTeamsItemChannelsItemMessagesDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-delta?view=graph-rest-beta
func (m *DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration)(DeletedTeamsItemChannelsItemMessagesDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletedTeamsItemChannelsItemMessagesDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletedTeamsItemChannelsItemMessagesDeltaResponseable), nil
}
// GetAsDeltaGetResponse get the list of messages from all chats in which a user is a participant, including one-on-one chats, group chats, and meeting chats. When you use delta query, you can obtain new or updated messages. To get the replies for a message, use the list message replies or the get message reply operations. A GET request with the delta function returns one of the following: State tokens are opaque to the client. To proceed with a round of change tracking, copy and apply the @odata.nextLink or @odata.deltaLink URL returned from the last GET request to the next delta function call. An @odata.deltaLink returned in a response signifies that the current round of change tracking is complete. You can save and use the @odata.deltaLink URL when you begin to retrieve more changes (messages changed or posted after you acquire @odata.deltaLink). For more information, see the delta query documentation.
// returns a DeletedTeamsItemChannelsItemMessagesDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-delta?view=graph-rest-beta
func (m *DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration)(DeletedTeamsItemChannelsItemMessagesDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeletedTeamsItemChannelsItemMessagesDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeletedTeamsItemChannelsItemMessagesDeltaGetResponseable), nil
}
// ToGetRequestInformation get the list of messages from all chats in which a user is a participant, including one-on-one chats, group chats, and meeting chats. When you use delta query, you can obtain new or updated messages. To get the replies for a message, use the list message replies or the get message reply operations. A GET request with the delta function returns one of the following: State tokens are opaque to the client. To proceed with a round of change tracking, copy and apply the @odata.nextLink or @odata.deltaLink URL returned from the last GET request to the next delta function call. An @odata.deltaLink returned in a response signifies that the current round of change tracking is complete. You can save and use the @odata.deltaLink URL when you begin to retrieve more changes (messages changed or posted after you acquire @odata.deltaLink). For more information, see the delta query documentation.
// returns a *RequestInformation when successful
func (m *DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder when successful
func (m *DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder) WithUrl(rawUrl string)(*DeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemMessagesDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

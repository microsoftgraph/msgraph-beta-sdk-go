package places

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GraphRoomRequestBuilder casts the previous resource to room.
type GraphRoomRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GraphRoomRequestBuilderGetQueryParameters get a collection of the specified type of place objects defined in the tenant.  You can do the following for a given tenant:- List all the rooms.- List all the workspaces.- List all the room lists.- List rooms in a specific room list.- List workspaces in a specific room list. A place object can be one of the following types: The room, workspace and roomList resources are derived from the place object. By default, this operation returns up to 100 places per page.  Compared with the findRooms and findRoomLists functions, this operation returns a richer payload for rooms and room lists. For details about how they compare, see Using the places API.
type GraphRoomRequestBuilderGetQueryParameters struct {
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
// GraphRoomRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GraphRoomRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GraphRoomRequestBuilderGetQueryParameters
}
// NewGraphRoomRequestBuilderInternal instantiates a new GraphRoomRequestBuilder and sets the default values.
func NewGraphRoomRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GraphRoomRequestBuilder) {
    m := &GraphRoomRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/places/graph.room{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGraphRoomRequestBuilder instantiates a new GraphRoomRequestBuilder and sets the default values.
func NewGraphRoomRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GraphRoomRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGraphRoomRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GraphRoomCountRequestBuilder when successful
func (m *GraphRoomRequestBuilder) Count()(*GraphRoomCountRequestBuilder) {
    return NewGraphRoomCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a collection of the specified type of place objects defined in the tenant.  You can do the following for a given tenant:- List all the rooms.- List all the workspaces.- List all the room lists.- List rooms in a specific room list.- List workspaces in a specific room list. A place object can be one of the following types: The room, workspace and roomList resources are derived from the place object. By default, this operation returns up to 100 places per page.  Compared with the findRooms and findRoomLists functions, this operation returns a richer payload for rooms and room lists. For details about how they compare, see Using the places API.
// returns a RoomCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/place-list?view=graph-rest-beta
func (m *GraphRoomRequestBuilder) Get(ctx context.Context, requestConfiguration *GraphRoomRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoomCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoomCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoomCollectionResponseable), nil
}
// ToGetRequestInformation get a collection of the specified type of place objects defined in the tenant.  You can do the following for a given tenant:- List all the rooms.- List all the workspaces.- List all the room lists.- List rooms in a specific room list.- List workspaces in a specific room list. A place object can be one of the following types: The room, workspace and roomList resources are derived from the place object. By default, this operation returns up to 100 places per page.  Compared with the findRooms and findRoomLists functions, this operation returns a richer payload for rooms and room lists. For details about how they compare, see Using the places API.
// returns a *RequestInformation when successful
func (m *GraphRoomRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GraphRoomRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GraphRoomRequestBuilder when successful
func (m *GraphRoomRequestBuilder) WithUrl(rawUrl string)(*GraphRoomRequestBuilder) {
    return NewGraphRoomRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

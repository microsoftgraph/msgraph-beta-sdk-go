package places

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGraphRoomRequestBuilder casts the previous resource to room.
type ItemGraphRoomRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGraphRoomRequestBuilderGetQueryParameters get a collection of the specified type of place objects defined in the tenant.  You can do the following for a given tenant:- List all the rooms.- List all the workspaces.- List all the room lists.- List rooms in a specific room list.- List workspaces in a specific room list. A place object can be one of the following types: The room, workspace and roomList resources are derived from the place object. By default, this operation returns up to 100 places per page.  Compared with the findRooms and findRoomLists functions, this operation returns a richer payload for rooms and room lists. For details about how they compare, see Using the places API.
type ItemGraphRoomRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemGraphRoomRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGraphRoomRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGraphRoomRequestBuilderGetQueryParameters
}
// NewItemGraphRoomRequestBuilderInternal instantiates a new ItemGraphRoomRequestBuilder and sets the default values.
func NewItemGraphRoomRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGraphRoomRequestBuilder) {
    m := &ItemGraphRoomRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/places/{place%2Did}/graph.room{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemGraphRoomRequestBuilder instantiates a new ItemGraphRoomRequestBuilder and sets the default values.
func NewItemGraphRoomRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGraphRoomRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGraphRoomRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a collection of the specified type of place objects defined in the tenant.  You can do the following for a given tenant:- List all the rooms.- List all the workspaces.- List all the room lists.- List rooms in a specific room list.- List workspaces in a specific room list. A place object can be one of the following types: The room, workspace and roomList resources are derived from the place object. By default, this operation returns up to 100 places per page.  Compared with the findRooms and findRoomLists functions, this operation returns a richer payload for rooms and room lists. For details about how they compare, see Using the places API.
// returns a Roomable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/place-list?view=graph-rest-beta
func (m *ItemGraphRoomRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGraphRoomRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Roomable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoomFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Roomable), nil
}
// ToGetRequestInformation get a collection of the specified type of place objects defined in the tenant.  You can do the following for a given tenant:- List all the rooms.- List all the workspaces.- List all the room lists.- List rooms in a specific room list.- List workspaces in a specific room list. A place object can be one of the following types: The room, workspace and roomList resources are derived from the place object. By default, this operation returns up to 100 places per page.  Compared with the findRooms and findRoomLists functions, this operation returns a richer payload for rooms and room lists. For details about how they compare, see Using the places API.
// returns a *RequestInformation when successful
func (m *ItemGraphRoomRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGraphRoomRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGraphRoomRequestBuilder when successful
func (m *ItemGraphRoomRequestBuilder) WithUrl(rawUrl string)(*ItemGraphRoomRequestBuilder) {
    return NewItemGraphRoomRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

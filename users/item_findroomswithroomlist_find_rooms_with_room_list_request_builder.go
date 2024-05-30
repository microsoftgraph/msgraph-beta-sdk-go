package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder provides operations to call the findRooms method.
type ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderGetQueryParameters invoke function findRooms
type ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderGetQueryParameters
}
// NewItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderInternal instantiates a new ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder and sets the default values.
func NewItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, roomList *string)(*ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder) {
    m := &ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/findRooms(RoomList='{RoomList}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if roomList != nil {
        m.BaseRequestBuilder.PathParameters["RoomList"] = *roomList
    }
    return m
}
// NewItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder instantiates a new ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder and sets the default values.
func NewItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function findRooms
// Deprecated: This method is obsolete. Use GetAsFindRoomsWithRoomListGetResponse instead.
// returns a ItemFindroomswithroomlistFindRoomsWithRoomListResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderGetRequestConfiguration)(ItemFindroomswithroomlistFindRoomsWithRoomListResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemFindroomswithroomlistFindRoomsWithRoomListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemFindroomswithroomlistFindRoomsWithRoomListResponseable), nil
}
// GetAsFindRoomsWithRoomListGetResponse invoke function findRooms
// returns a ItemFindroomswithroomlistFindRoomsWithRoomListGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder) GetAsFindRoomsWithRoomListGetResponse(ctx context.Context, requestConfiguration *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderGetRequestConfiguration)(ItemFindroomswithroomlistFindRoomsWithRoomListGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemFindroomswithroomlistFindRoomsWithRoomListGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemFindroomswithroomlistFindRoomsWithRoomListGetResponseable), nil
}
// ToGetRequestInformation invoke function findRooms
// returns a *RequestInformation when successful
func (m *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder when successful
func (m *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder) WithUrl(rawUrl string)(*ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder) {
    return NewItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

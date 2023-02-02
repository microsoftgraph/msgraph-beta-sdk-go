package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilder provides operations to call the findRooms method.
type MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderGetQueryParameters invoke function findRooms
type MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderGetQueryParameters struct {
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
// MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderGetQueryParameters
}
// NewMicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderInternal instantiates a new FindRoomsWithRoomListRequestBuilder and sets the default values.
func NewMicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, roomList *string)(*MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilder) {
    m := &MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.findRooms(RoomList='{RoomList}'){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if roomList != nil {
        urlTplParams["RoomList"] = *roomList
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilder instantiates a new FindRoomsWithRoomListRequestBuilder and sets the default values.
func NewMicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function findRooms
func (m *MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderGetRequestConfiguration)(MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListResponseable), nil
}
// ToGetRequestInformation invoke function findRooms
func (m *MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphFindRoomsWithRoomListFindRoomsWithRoomListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

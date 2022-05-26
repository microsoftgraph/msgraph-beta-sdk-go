package findroomswithroomlist

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FindRoomsWithRoomListRequestBuilder provides operations to call the findRooms method.
type FindRoomsWithRoomListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FindRoomsWithRoomListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FindRoomsWithRoomListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFindRoomsWithRoomListRequestBuilderInternal instantiates a new FindRoomsWithRoomListRequestBuilder and sets the default values.
func NewFindRoomsWithRoomListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, roomList *string)(*FindRoomsWithRoomListRequestBuilder) {
    m := &FindRoomsWithRoomListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.user/microsoft.graph.findRooms(RoomList='{RoomList}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if roomList != nil {
        urlTplParams["RoomList"] = *roomList
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFindRoomsWithRoomListRequestBuilder instantiates a new FindRoomsWithRoomListRequestBuilder and sets the default values.
func NewFindRoomsWithRoomListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FindRoomsWithRoomListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFindRoomsWithRoomListRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function findRooms
func (m *FindRoomsWithRoomListRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function findRooms
func (m *FindRoomsWithRoomListRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *FindRoomsWithRoomListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function findRooms
func (m *FindRoomsWithRoomListRequestBuilder) Get()(FindRoomsWithRoomListResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function findRooms
func (m *FindRoomsWithRoomListRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *FindRoomsWithRoomListRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(FindRoomsWithRoomListResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateFindRoomsWithRoomListResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(FindRoomsWithRoomListResponseable), nil
}

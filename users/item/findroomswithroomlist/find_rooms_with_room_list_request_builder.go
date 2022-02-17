package findroomswithroomlist

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FindRoomsWithRoomListRequestBuilder builds and executes requests for operations under \users\{user-id}\microsoft.graph.findRooms(RoomList='{RoomList}')
type FindRoomsWithRoomListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// FindRoomsWithRoomListRequestBuilderGetOptions options for Get
type FindRoomsWithRoomListRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewFindRoomsWithRoomListRequestBuilderInternal instantiates a new FindRoomsWithRoomListRequestBuilder and sets the default values.
func NewFindRoomsWithRoomListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, roomList *string)(*FindRoomsWithRoomListRequestBuilder) {
    m := &FindRoomsWithRoomListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/microsoft.graph.findRooms(RoomList='{RoomList}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if roomList != nil {
        urlTplParams["RoomList"] = *roomList
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFindRoomsWithRoomListRequestBuilder instantiates a new FindRoomsWithRoomListRequestBuilder and sets the default values.
func NewFindRoomsWithRoomListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FindRoomsWithRoomListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFindRoomsWithRoomListRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function findRooms
func (m *FindRoomsWithRoomListRequestBuilder) CreateGetRequestInformation(options *FindRoomsWithRoomListRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function findRooms
func (m *FindRoomsWithRoomListRequestBuilder) Get(options *FindRoomsWithRoomListRequestBuilderGetOptions)([]FindRoomsWithRoomList, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFindRoomsWithRoomList() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]FindRoomsWithRoomList, len(res))
    for i, v := range res {
        val[i] = *(v.(*FindRoomsWithRoomList))
    }
    return val, nil
}

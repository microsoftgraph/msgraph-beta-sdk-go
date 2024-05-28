package places

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder provides operations to manage the rooms property of the microsoft.graph.roomList entity.
type ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderGetQueryParameters get rooms from places
type ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderGetQueryParameters
}
// ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderInternal instantiates a new ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder and sets the default values.
func NewItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, placeId *string)(*ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) {
    m := &ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/places/{place%2Did}/graph.roomList/rooms(placeId='{placeId}'){?%24expand,%24select}", pathParameters),
    }
    if placeId != nil {
        m.BaseRequestBuilder.PathParameters["placeId"] = *placeId
    }
    return m
}
// NewItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder instantiates a new ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder and sets the default values.
func NewItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property rooms for places
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get rooms from places
// returns a Roomable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Roomable, error) {
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
// Patch update the navigation property rooms in places
// returns a Roomable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Roomable, requestConfiguration *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Roomable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property rooms for places
// returns a *RequestInformation when successful
func (m *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get rooms from places
// returns a *RequestInformation when successful
func (m *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property rooms in places
// returns a *RequestInformation when successful
func (m *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Roomable, requestConfiguration *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder when successful
func (m *ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) WithUrl(rawUrl string)(*ItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder) {
    return NewItemGraphroomlistRoomswithplaceidRoomsWithPlaceIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

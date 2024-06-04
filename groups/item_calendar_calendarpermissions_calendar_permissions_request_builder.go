package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
type ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderGetQueryParameters the permissions of the users with whom the calendar is shared.
type ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderGetQueryParameters struct {
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
// ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderGetQueryParameters
}
// ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCalendarPermissionId provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
// returns a *ItemCalendarCalendarpermissionsCalendarPermissionItemRequestBuilder when successful
func (m *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) ByCalendarPermissionId(calendarPermissionId string)(*ItemCalendarCalendarpermissionsCalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if calendarPermissionId != "" {
        urlTplParams["calendarPermission%2Did"] = calendarPermissionId
    }
    return NewItemCalendarCalendarpermissionsCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderInternal instantiates a new ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder and sets the default values.
func NewItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) {
    m := &ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/calendarPermissions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder instantiates a new ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder and sets the default values.
func NewItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemCalendarCalendarpermissionsCountRequestBuilder when successful
func (m *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) Count()(*ItemCalendarCalendarpermissionsCountRequestBuilder) {
    return NewItemCalendarCalendarpermissionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the permissions of the users with whom the calendar is shared.
// returns a CalendarPermissionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CalendarPermissionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCalendarPermissionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CalendarPermissionCollectionResponseable), nil
}
// Post create new navigation property to calendarPermissions for groups
// returns a CalendarPermissionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CalendarPermissionable, requestConfiguration *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CalendarPermissionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCalendarPermissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CalendarPermissionable), nil
}
// ToGetRequestInformation the permissions of the users with whom the calendar is shared.
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to calendarPermissions for groups
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CalendarPermissionable, requestConfiguration *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder when successful
func (m *ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder) {
    return NewItemCalendarCalendarpermissionsCalendarPermissionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderGetQueryParameters get exceptionOccurrences from users
type UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderGetQueryParameters
}
// NewUsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderInternal instantiates a new ExceptionOccurrencesRequestBuilder and sets the default values.
func NewUsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder) {
    m := &UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder instantiates a new ExceptionOccurrencesRequestBuilder and sets the default values.
func NewUsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder) Count()(*UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesCountRequestBuilder) {
    return NewUsersItemCalendarsItemCalendarViewItemExceptionOccurrencesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get exceptionOccurrences from users
func (m *UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder) Delta()(*UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesDeltaRequestBuilder) {
    return NewUsersItemCalendarsItemCalendarViewItemExceptionOccurrencesDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemCalendarsItemCalendarViewItemExceptionOccurrencesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable), nil
}

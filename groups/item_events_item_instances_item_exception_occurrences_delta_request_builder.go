package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilder provides operations to call the delta method.
type ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderGetQueryParameters get a set of event resources that have been added, deleted, or updated in one or more calendars.  You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a **calendarView** (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar of the user's. In the case of getting incremental changes on **calendarView**, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or **calendarView** in a local store entails a round of multiple **delta** function calls. The initial call is a full synchronization, and every subsequent **delta** call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified calendar, without having to fetch all the events of that calendar from the server every time. The following table lists the differences between the **delta** function on events and the **delta** function on a **calendarView** in a calendar.
type ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderGetQueryParameters
}
// NewItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderInternal instantiates a new DeltaRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilder) {
    m := &ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilder instantiates a new DeltaRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of event resources that have been added, deleted, or updated in one or more calendars.  You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a **calendarView** (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar of the user's. In the case of getting incremental changes on **calendarView**, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or **calendarView** in a local store entails a round of multiple **delta** function calls. The initial call is a full synchronization, and every subsequent **delta** call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified calendar, without having to fetch all the events of that calendar from the server every time. The following table lists the differences between the **delta** function on events and the **delta** function on a **calendarView** in a calendar.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-delta?view=graph-rest-1.0
func (m *ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderGetRequestConfiguration)(ItemEventsItemInstancesItemExceptionOccurrencesDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemEventsItemInstancesItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemEventsItemInstancesItemExceptionOccurrencesDeltaResponseable), nil
}
// ToGetRequestInformation get a set of event resources that have been added, deleted, or updated in one or more calendars.  You can get specific types of these incremental changes in the events in all the calendars of a mailbox or in a specific calendar, or in an event collection of a **calendarView** (range of events defined by start and end dates) of a calendar. The calendar can be the default calendar or some other specified calendar of the user's. In the case of getting incremental changes on **calendarView**, the calendar can be a group calendar as well. Typically, synchronizing events in a calendar or **calendarView** in a local store entails a round of multiple **delta** function calls. The initial call is a full synchronization, and every subsequent **delta** call in the same round gets the incremental changes (additions, deletions, or updates). This allows you to maintain and synchronize a local store of events in the specified calendar, without having to fetch all the events of that calendar from the server every time. The following table lists the differences between the **delta** function on events and the **delta** function on a **calendarView** in a calendar.
func (m *ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEventsItemInstancesItemExceptionOccurrencesDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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

package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5d489a322600c399fdedc6aced62ac05f02e4768b34e03c0d0c43a0d3f91f2d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/allowedcalendarsharingroleswithuser"
    iba1f8c99094d51cddebb2bf20a5f731030766c1dd2f5b451f8bd627ea9655eab "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/ref"
    id0adbea957932b06dd78ef7b88a847a18c39958be814e6f3cee44354d4cacad6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/getschedule"
)

// CalendarRequestBuilder builds and executes requests for operations under \groups\{group-id}\calendarView\{event-id}\calendar
type CalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CalendarRequestBuilderGetOptions options for Get
type CalendarRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CalendarRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CalendarRequestBuilderGetQueryParameters the calendar that contains the event. Navigation property. Read-only.
type CalendarRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \groups\{group-id}\calendarView\{event-id}\calendar\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i5d489a322600c399fdedc6aced62ac05f02e4768b34e03c0d0c43a0d3f91f2d3.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i5d489a322600c399fdedc6aced62ac05f02e4768b34e03c0d0c43a0d3f91f2d3.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendarView/{event_id}/calendar{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarRequestBuilder instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) CreateGetRequestInformation(options *CalendarRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// Get the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) Get(options *CalendarRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCalendar() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar), nil
}
func (m *CalendarRequestBuilder) GetSchedule()(*id0adbea957932b06dd78ef7b88a847a18c39958be814e6f3cee44354d4cacad6.GetScheduleRequestBuilder) {
    return id0adbea957932b06dd78ef7b88a847a18c39958be814e6f3cee44354d4cacad6.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) Ref()(*iba1f8c99094d51cddebb2bf20a5f731030766c1dd2f5b451f8bd627ea9655eab.RefRequestBuilder) {
    return iba1f8c99094d51cddebb2bf20a5f731030766c1dd2f5b451f8bd627ea9655eab.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

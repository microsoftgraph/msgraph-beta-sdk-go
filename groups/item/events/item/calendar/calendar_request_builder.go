package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3949123b651b9af052c8de6e7f5c595e080112b647c9a3626918cdcb55bc5902 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/singlevalueextendedproperties"
    i469e106165832f085e3fd273b20eefa6a88349191e203975179d82cd8fffe28b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/getschedule"
    i7b33c072b2a7ba3d82b9f69335f559cda585a5fab49a46127f175c6415fdb6c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/events"
    id04b107f7faaecce66eb6592f5f4f31f2f7b40b9bc6676aba5a06edbd4ee0fb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/multivalueextendedproperties"
    ie2cb151ce95182b548e67f27c8742fa555088b1e0e11ff7c1ce9d68f8b9e4880 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/calendarpermissions"
    ieda5ed6170b9c595db5a4f0e03158aac1856d1fd011e0472cd153361781f8835 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/calendarview"
    if6b857c575d61df1976f20edc812574d146eb925d961b7b3e2ba89bb97a44457 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/allowedcalendarsharingroleswithuser"
    i21c7ac2dc439414814348cc02de537ca481a6879b19eb08e0d4a93b78a1c9d60 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/calendarpermissions/item"
    i447ef94e57e8b820a1b16f1c03a188af03e71ba0d09d8a47dc42a19de9f764b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/singlevalueextendedproperties/item"
    i7035dcea3e0c498885a3febbcb76e699752f3550a3dffebcf26ee83bb5595f35 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/events/item"
    idd097aa2b0081a4a295f78cfc1bc963a100844ea07f6226dad11082cf1e15242 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/calendarview/item"
    ie14fdde3d6c5df3883fe4da02fa316e51f128cbca18042d14a0105153c540a5e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/multivalueextendedproperties/item"
)

// CalendarRequestBuilder builds and executes requests for operations under \groups\{group-id}\events\{event-id}\calendar
type CalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CalendarRequestBuilderDeleteOptions options for Delete
type CalendarRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
// CalendarRequestBuilderPatchOptions options for Patch
type CalendarRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \groups\{group-id}\events\{event-id}\calendar\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*if6b857c575d61df1976f20edc812574d146eb925d961b7b3e2ba89bb97a44457.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return if6b857c575d61df1976f20edc812574d146eb925d961b7b3e2ba89bb97a44457.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*ie2cb151ce95182b548e67f27c8742fa555088b1e0e11ff7c1ce9d68f8b9e4880.CalendarPermissionsRequestBuilder) {
    return ie2cb151ce95182b548e67f27c8742fa555088b1e0e11ff7c1ce9d68f8b9e4880.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.calendar.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i21c7ac2dc439414814348cc02de537ca481a6879b19eb08e0d4a93b78a1c9d60.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i21c7ac2dc439414814348cc02de537ca481a6879b19eb08e0d4a93b78a1c9d60.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ieda5ed6170b9c595db5a4f0e03158aac1856d1fd011e0472cd153361781f8835.CalendarViewRequestBuilder) {
    return ieda5ed6170b9c595db5a4f0e03158aac1856d1fd011e0472cd153361781f8835.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.calendar.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*idd097aa2b0081a4a295f78cfc1bc963a100844ea07f6226dad11082cf1e15242.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return idd097aa2b0081a4a295f78cfc1bc963a100844ea07f6226dad11082cf1e15242.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/events/{event_id}/calendar{?select}";
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
// CreateDeleteRequestInformation the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) CreateDeleteRequestInformation(options *CalendarRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreatePatchRequestInformation the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(options *CalendarRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Delete the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) Delete(options *CalendarRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarRequestBuilder) Events()(*i7b33c072b2a7ba3d82b9f69335f559cda585a5fab49a46127f175c6415fdb6c1.EventsRequestBuilder) {
    return i7b33c072b2a7ba3d82b9f69335f559cda585a5fab49a46127f175c6415fdb6c1.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.calendar.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*i7035dcea3e0c498885a3febbcb76e699752f3550a3dffebcf26ee83bb5595f35.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i7035dcea3e0c498885a3febbcb76e699752f3550a3dffebcf26ee83bb5595f35.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) Get(options *CalendarRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCalendar() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar), nil
}
func (m *CalendarRequestBuilder) GetSchedule()(*i469e106165832f085e3fd273b20eefa6a88349191e203975179d82cd8fffe28b.GetScheduleRequestBuilder) {
    return i469e106165832f085e3fd273b20eefa6a88349191e203975179d82cd8fffe28b.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*id04b107f7faaecce66eb6592f5f4f31f2f7b40b9bc6676aba5a06edbd4ee0fb3.MultiValueExtendedPropertiesRequestBuilder) {
    return id04b107f7faaecce66eb6592f5f4f31f2f7b40b9bc6676aba5a06edbd4ee0fb3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.calendar.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie14fdde3d6c5df3883fe4da02fa316e51f128cbca18042d14a0105153c540a5e.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ie14fdde3d6c5df3883fe4da02fa316e51f128cbca18042d14a0105153c540a5e.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) Patch(options *CalendarRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i3949123b651b9af052c8de6e7f5c595e080112b647c9a3626918cdcb55bc5902.SingleValueExtendedPropertiesRequestBuilder) {
    return i3949123b651b9af052c8de6e7f5c595e080112b647c9a3626918cdcb55bc5902.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.calendar.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i447ef94e57e8b820a1b16f1c03a188af03e71ba0d09d8a47dc42a19de9f764b7.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i447ef94e57e8b820a1b16f1c03a188af03e71ba0d09d8a47dc42a19de9f764b7.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

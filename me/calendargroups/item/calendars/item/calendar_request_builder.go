package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i21d12d245d27df0693994a8bc5d5cb129c8c438d8f0d501c089b2742dd07da63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/singlevalueextendedproperties"
    i2bbc53a9d42f80d48bde303e47a6edf7fd4514d20543476fc96d7654f44459c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/getschedule"
    i3a99b227ed3c41b86c6ed7d567ee4be24b191378e5c0a25d4ee01a213bae8cc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarpermissions"
    i4cbc0b47a2f4d174f3aacee057ce86418a0d927f3a1682711f7f56d4fa432dc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/multivalueextendedproperties"
    ie4148df7c524cadb5341a6d6e61b55f25127deff83b2a0c89cf4131e1fa474f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview"
    ie7a0d904d236f9db50b3140c9590b5afdb51f2d8f4f57f0bac0dd742c01a94ce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events"
    ifdf0230f7f25a33effd6a753eaabc0e03fdd6933c1b808f4077742b51f7e8438 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/allowedcalendarsharingroleswithuser"
    i0cf8829e9a39acb3b6c6c71c0e9e7092ef194c89704827f0cb9ce7c26fd01fc9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/singlevalueextendedproperties/item"
    i209ddc6ff230debede7f258040e209b82083649dc4aba8bc8b32c8345af01a3f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item"
    i3014ad4ca2bdf0c9ac65d0252d7f6da0745d1c8fdfb5591c641e2eb39b4f775d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item"
    i52c8e49a23935de33330e44b9a94cf9f4bf113557ba760732e561ac40734cfd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarpermissions/item"
    i97497d8b531a6c7061e0336a8740e7cb8cb0d35b3de77bbc36a82dff33a8e672 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/multivalueextendedproperties/item"
)

// CalendarRequestBuilder builds and executes requests for operations under \me\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}
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
// CalendarRequestBuilderGetQueryParameters the calendars in the calendar group. Navigation property. Read-only. Nullable.
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
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \me\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ifdf0230f7f25a33effd6a753eaabc0e03fdd6933c1b808f4077742b51f7e8438.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ifdf0230f7f25a33effd6a753eaabc0e03fdd6933c1b808f4077742b51f7e8438.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i3a99b227ed3c41b86c6ed7d567ee4be24b191378e5c0a25d4ee01a213bae8cc2.CalendarPermissionsRequestBuilder) {
    return i3a99b227ed3c41b86c6ed7d567ee4be24b191378e5c0a25d4ee01a213bae8cc2.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i52c8e49a23935de33330e44b9a94cf9f4bf113557ba760732e561ac40734cfd9.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i52c8e49a23935de33330e44b9a94cf9f4bf113557ba760732e561ac40734cfd9.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ie4148df7c524cadb5341a6d6e61b55f25127deff83b2a0c89cf4131e1fa474f5.CalendarViewRequestBuilder) {
    return ie4148df7c524cadb5341a6d6e61b55f25127deff83b2a0c89cf4131e1fa474f5.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i209ddc6ff230debede7f258040e209b82083649dc4aba8bc8b32c8345af01a3f.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i209ddc6ff230debede7f258040e209b82083649dc4aba8bc8b32c8345af01a3f.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}{?select}";
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
// CreateDeleteRequestInformation the calendars in the calendar group. Navigation property. Read-only. Nullable.
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
// CreateGetRequestInformation the calendars in the calendar group. Navigation property. Read-only. Nullable.
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
// CreatePatchRequestInformation the calendars in the calendar group. Navigation property. Read-only. Nullable.
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
// Delete the calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *CalendarRequestBuilder) Delete(options *CalendarRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarRequestBuilder) Events()(*ie7a0d904d236f9db50b3140c9590b5afdb51f2d8f4f57f0bac0dd742c01a94ce.EventsRequestBuilder) {
    return ie7a0d904d236f9db50b3140c9590b5afdb51f2d8f4f57f0bac0dd742c01a94ce.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*i3014ad4ca2bdf0c9ac65d0252d7f6da0745d1c8fdfb5591c641e2eb39b4f775d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i3014ad4ca2bdf0c9ac65d0252d7f6da0745d1c8fdfb5591c641e2eb39b4f775d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the calendars in the calendar group. Navigation property. Read-only. Nullable.
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
func (m *CalendarRequestBuilder) GetSchedule()(*i2bbc53a9d42f80d48bde303e47a6edf7fd4514d20543476fc96d7654f44459c7.GetScheduleRequestBuilder) {
    return i2bbc53a9d42f80d48bde303e47a6edf7fd4514d20543476fc96d7654f44459c7.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i4cbc0b47a2f4d174f3aacee057ce86418a0d927f3a1682711f7f56d4fa432dc6.MultiValueExtendedPropertiesRequestBuilder) {
    return i4cbc0b47a2f4d174f3aacee057ce86418a0d927f3a1682711f7f56d4fa432dc6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i97497d8b531a6c7061e0336a8740e7cb8cb0d35b3de77bbc36a82dff33a8e672.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i97497d8b531a6c7061e0336a8740e7cb8cb0d35b3de77bbc36a82dff33a8e672.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *CalendarRequestBuilder) Patch(options *CalendarRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i21d12d245d27df0693994a8bc5d5cb129c8c438d8f0d501c089b2742dd07da63.SingleValueExtendedPropertiesRequestBuilder) {
    return i21d12d245d27df0693994a8bc5d5cb129c8c438d8f0d501c089b2742dd07da63.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0cf8829e9a39acb3b6c6c71c0e9e7092ef194c89704827f0cb9ce7c26fd01fc9.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i0cf8829e9a39acb3b6c6c71c0e9e7092ef194c89704827f0cb9ce7c26fd01fc9.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

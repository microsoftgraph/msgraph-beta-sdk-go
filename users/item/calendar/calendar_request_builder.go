package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0d974efbbcde965f86cd784750b0d398768a58c4f0c31e8aa6eebab38714cfb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview"
    i5f2b87b4b8ae2f7ac0f87df9716f56e4739dd16a5886917392d33a88c18c770c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events"
    i7f99ab356098627c58430897ce6db68b26eb164759df0d5c099175b468c773cd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/singlevalueextendedproperties"
    ia046aa77ed3651894558d377e9f2faf10b90f9fa6ea4b8d58fa81d2c0bf352b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarpermissions"
    ib11c18a31cdd7cea653dae6c94f9cb3c8880c5d2ff9d71db4311df3cb8ef6acb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/allowedcalendarsharingroleswithuser"
    ic0da2257b14ab755d200b83c581b125370f630995453afb2a431a7b9a660b4fb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/getschedule"
    ie4f67284862d68d4348bd087a24d1359caa53a620bec669a57587285e12ce7b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/multivalueextendedproperties"
    i35fa4f4f3eb3bdb367848efc1a7e0f1f004ad8371e5de0506e46fa00e172f96a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item"
    i646e54f4392ea806f149eff3b9db0f4c56b5fd5bf2023a287a9335a93382fc50 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/singlevalueextendedproperties/item"
    i6acf51572e8a0bbb972c27da38b03c2fa4c3df3e13192ade1cbb983349a39367 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item"
    i94baab4924959ccc5f343b01737d73626b9abc213582a99069414c9ca8dd6fc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarpermissions/item"
    ibe6df2514f80166133d14155ec9a4ce76559c25b036bd39eee4eb72b15967d12 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/multivalueextendedproperties/item"
)

// CalendarRequestBuilder builds and executes requests for operations under \users\{user-id}\calendar
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
// CalendarRequestBuilderGetQueryParameters the user's primary calendar. Read-only.
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
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \users\{user-id}\calendar\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ib11c18a31cdd7cea653dae6c94f9cb3c8880c5d2ff9d71db4311df3cb8ef6acb.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ib11c18a31cdd7cea653dae6c94f9cb3c8880c5d2ff9d71db4311df3cb8ef6acb.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*ia046aa77ed3651894558d377e9f2faf10b90f9fa6ea4b8d58fa81d2c0bf352b7.CalendarPermissionsRequestBuilder) {
    return ia046aa77ed3651894558d377e9f2faf10b90f9fa6ea4b8d58fa81d2c0bf352b7.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i94baab4924959ccc5f343b01737d73626b9abc213582a99069414c9ca8dd6fc2.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i94baab4924959ccc5f343b01737d73626b9abc213582a99069414c9ca8dd6fc2.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i0d974efbbcde965f86cd784750b0d398768a58c4f0c31e8aa6eebab38714cfb3.CalendarViewRequestBuilder) {
    return i0d974efbbcde965f86cd784750b0d398768a58c4f0c31e8aa6eebab38714cfb3.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i35fa4f4f3eb3bdb367848efc1a7e0f1f004ad8371e5de0506e46fa00e172f96a.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i35fa4f4f3eb3bdb367848efc1a7e0f1f004ad8371e5de0506e46fa00e172f96a.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar{?select}";
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
// CreateDeleteRequestInformation the user's primary calendar. Read-only.
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
// CreateGetRequestInformation the user's primary calendar. Read-only.
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
// CreatePatchRequestInformation the user's primary calendar. Read-only.
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
// Delete the user's primary calendar. Read-only.
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
func (m *CalendarRequestBuilder) Events()(*i5f2b87b4b8ae2f7ac0f87df9716f56e4739dd16a5886917392d33a88c18c770c.EventsRequestBuilder) {
    return i5f2b87b4b8ae2f7ac0f87df9716f56e4739dd16a5886917392d33a88c18c770c.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*i6acf51572e8a0bbb972c27da38b03c2fa4c3df3e13192ade1cbb983349a39367.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i6acf51572e8a0bbb972c27da38b03c2fa4c3df3e13192ade1cbb983349a39367.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the user's primary calendar. Read-only.
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
func (m *CalendarRequestBuilder) GetSchedule()(*ic0da2257b14ab755d200b83c581b125370f630995453afb2a431a7b9a660b4fb.GetScheduleRequestBuilder) {
    return ic0da2257b14ab755d200b83c581b125370f630995453afb2a431a7b9a660b4fb.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*ie4f67284862d68d4348bd087a24d1359caa53a620bec669a57587285e12ce7b3.MultiValueExtendedPropertiesRequestBuilder) {
    return ie4f67284862d68d4348bd087a24d1359caa53a620bec669a57587285e12ce7b3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibe6df2514f80166133d14155ec9a4ce76559c25b036bd39eee4eb72b15967d12.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ibe6df2514f80166133d14155ec9a4ce76559c25b036bd39eee4eb72b15967d12.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's primary calendar. Read-only.
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i7f99ab356098627c58430897ce6db68b26eb164759df0d5c099175b468c773cd.SingleValueExtendedPropertiesRequestBuilder) {
    return i7f99ab356098627c58430897ce6db68b26eb164759df0d5c099175b468c773cd.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i646e54f4392ea806f149eff3b9db0f4c56b5fd5bf2023a287a9335a93382fc50.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i646e54f4392ea806f149eff3b9db0f4c56b5fd5bf2023a287a9335a93382fc50.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

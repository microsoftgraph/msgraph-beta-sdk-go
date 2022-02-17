package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i06cb71e1758c94dc508ff81b35d70e6b87e40dc3c43fc277a55fc5b07c51c025 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events"
    i4e394637a163af92a315f34da04f62ec4dc1d045e17f3716acaced9ea35ad045 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/getschedule"
    ib7ee09a2619abb2c78701384def8fcdab068e606d667dea6587234b16d250e88 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview"
    ic79a06981e01676a5b41be5ce0766fd4aeeb0f4a22bbea5fca16431948ae127a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarpermissions"
    ic906c3aa3cff8041006e1db1afa8c5ecce79b293f1b9649a48c05726e6f788f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/multivalueextendedproperties"
    icd7a045db01e44db897afd8ba5343e7d38c1599413b6a17400da01b527973e57 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/allowedcalendarsharingroleswithuser"
    ie357553e3f2994fd5e3a6ab96e22c5c927c07c3e3c999ee4350aca9434117b9c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/singlevalueextendedproperties"
    i2877c9d97209d99b4d909de67ba8cd762e6a506898bb79d3f55b0659900ca31c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item"
    i2b0a7eca3f232e3b28e09ce2398345a2762596c5fc5f1b3b203b447ca9d62dcd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/multivalueextendedproperties/item"
    i6389000acb4a221ddef87343d30894dc16987ef37595f24421ad461fe767083f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarpermissions/item"
    i7135f234b9b77ac7d1f38566664eb079da2d1647d2b78e7e1558a59a4f6a8ae5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/singlevalueextendedproperties/item"
    iafdc3a5c49cb6bb0133e40dda5d4c139576d35278bf232fa62de70f01714f934 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item"
)

// CalendarRequestBuilder builds and executes requests for operations under \groups\{group-id}\calendar
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
// CalendarRequestBuilderGetQueryParameters the group's calendar. Read-only.
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
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \groups\{group-id}\calendar\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*icd7a045db01e44db897afd8ba5343e7d38c1599413b6a17400da01b527973e57.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return icd7a045db01e44db897afd8ba5343e7d38c1599413b6a17400da01b527973e57.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*ic79a06981e01676a5b41be5ce0766fd4aeeb0f4a22bbea5fca16431948ae127a.CalendarPermissionsRequestBuilder) {
    return ic79a06981e01676a5b41be5ce0766fd4aeeb0f4a22bbea5fca16431948ae127a.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i6389000acb4a221ddef87343d30894dc16987ef37595f24421ad461fe767083f.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i6389000acb4a221ddef87343d30894dc16987ef37595f24421ad461fe767083f.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ib7ee09a2619abb2c78701384def8fcdab068e606d667dea6587234b16d250e88.CalendarViewRequestBuilder) {
    return ib7ee09a2619abb2c78701384def8fcdab068e606d667dea6587234b16d250e88.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i2877c9d97209d99b4d909de67ba8cd762e6a506898bb79d3f55b0659900ca31c.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i2877c9d97209d99b4d909de67ba8cd762e6a506898bb79d3f55b0659900ca31c.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar{?select}";
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
// CreateDeleteRequestInformation the group's calendar. Read-only.
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
// CreateGetRequestInformation the group's calendar. Read-only.
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
// CreatePatchRequestInformation the group's calendar. Read-only.
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
// Delete the group's calendar. Read-only.
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
func (m *CalendarRequestBuilder) Events()(*i06cb71e1758c94dc508ff81b35d70e6b87e40dc3c43fc277a55fc5b07c51c025.EventsRequestBuilder) {
    return i06cb71e1758c94dc508ff81b35d70e6b87e40dc3c43fc277a55fc5b07c51c025.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*iafdc3a5c49cb6bb0133e40dda5d4c139576d35278bf232fa62de70f01714f934.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return iafdc3a5c49cb6bb0133e40dda5d4c139576d35278bf232fa62de70f01714f934.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the group's calendar. Read-only.
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
func (m *CalendarRequestBuilder) GetSchedule()(*i4e394637a163af92a315f34da04f62ec4dc1d045e17f3716acaced9ea35ad045.GetScheduleRequestBuilder) {
    return i4e394637a163af92a315f34da04f62ec4dc1d045e17f3716acaced9ea35ad045.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*ic906c3aa3cff8041006e1db1afa8c5ecce79b293f1b9649a48c05726e6f788f8.MultiValueExtendedPropertiesRequestBuilder) {
    return ic906c3aa3cff8041006e1db1afa8c5ecce79b293f1b9649a48c05726e6f788f8.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2b0a7eca3f232e3b28e09ce2398345a2762596c5fc5f1b3b203b447ca9d62dcd.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i2b0a7eca3f232e3b28e09ce2398345a2762596c5fc5f1b3b203b447ca9d62dcd.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the group's calendar. Read-only.
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*ie357553e3f2994fd5e3a6ab96e22c5c927c07c3e3c999ee4350aca9434117b9c.SingleValueExtendedPropertiesRequestBuilder) {
    return ie357553e3f2994fd5e3a6ab96e22c5c927c07c3e3c999ee4350aca9434117b9c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7135f234b9b77ac7d1f38566664eb079da2d1647d2b78e7e1558a59a4f6a8ae5.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i7135f234b9b77ac7d1f38566664eb079da2d1647d2b78e7e1558a59a4f6a8ae5.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

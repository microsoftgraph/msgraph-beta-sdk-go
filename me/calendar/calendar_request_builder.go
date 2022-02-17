package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2f1046c95e29e9dc0a0e26221e25babc15231d926fd21540bc4f8be55206acba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarpermissions"
    i310a0918126f2fd1a54cb6d55a6aa248b29b49ce4b8f6944a5a18dfe96111178 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i84f41c2d8a59ad72a2557098f3b161d03f5d8ac8e3987507b0fa7ecdf05601bf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/multivalueextendedproperties"
    i85085bfff31dfdfaf3940c26b7428d4dd4f5f214e1afb26f8ff92de42666214c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/allowedcalendarsharingroleswithuser"
    ia8c270bfd48992dd6d8586b4eb88c58708cf3fbfc388c4660cde8b22341c2d61 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview"
    iafafdd8f3e81db193ebe1f176611ba8d070c598779572146f18a7a2320d2b8fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/singlevalueextendedproperties"
    id1ef0f108571f18aefac2c61a51615b16187fb53daeb8778a7d3fc99ee5b1a30 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/getschedule"
    i1339b57f4311b596f1e163d07c2af7ba9105f71396c487bac3b1051937b9a343 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/singlevalueextendedproperties/item"
    i32007971daba2b6bfd608251e3fa3037f0c6076b83bf11d1cebc0be6d191b58f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item"
    i7923064a9f4c4b75d44193b15f31eea4e3de5a03a690127e925545ca3fcb8ce5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item"
    icc49a296c7df21a7cc4befcd6f1a34f299f25134015701387f39f2be77ec1b62 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarpermissions/item"
    iec1f5e677b964978263004630240f212d437a10091543ccef2397757ae2b5c2e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/multivalueextendedproperties/item"
)

// CalendarRequestBuilder builds and executes requests for operations under \me\calendar
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
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \me\calendar\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i85085bfff31dfdfaf3940c26b7428d4dd4f5f214e1afb26f8ff92de42666214c.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i85085bfff31dfdfaf3940c26b7428d4dd4f5f214e1afb26f8ff92de42666214c.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i2f1046c95e29e9dc0a0e26221e25babc15231d926fd21540bc4f8be55206acba.CalendarPermissionsRequestBuilder) {
    return i2f1046c95e29e9dc0a0e26221e25babc15231d926fd21540bc4f8be55206acba.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*icc49a296c7df21a7cc4befcd6f1a34f299f25134015701387f39f2be77ec1b62.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return icc49a296c7df21a7cc4befcd6f1a34f299f25134015701387f39f2be77ec1b62.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ia8c270bfd48992dd6d8586b4eb88c58708cf3fbfc388c4660cde8b22341c2d61.CalendarViewRequestBuilder) {
    return ia8c270bfd48992dd6d8586b4eb88c58708cf3fbfc388c4660cde8b22341c2d61.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i32007971daba2b6bfd608251e3fa3037f0c6076b83bf11d1cebc0be6d191b58f.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i32007971daba2b6bfd608251e3fa3037f0c6076b83bf11d1cebc0be6d191b58f.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar{?select}";
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
func (m *CalendarRequestBuilder) Events()(*i310a0918126f2fd1a54cb6d55a6aa248b29b49ce4b8f6944a5a18dfe96111178.EventsRequestBuilder) {
    return i310a0918126f2fd1a54cb6d55a6aa248b29b49ce4b8f6944a5a18dfe96111178.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*i7923064a9f4c4b75d44193b15f31eea4e3de5a03a690127e925545ca3fcb8ce5.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i7923064a9f4c4b75d44193b15f31eea4e3de5a03a690127e925545ca3fcb8ce5.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*id1ef0f108571f18aefac2c61a51615b16187fb53daeb8778a7d3fc99ee5b1a30.GetScheduleRequestBuilder) {
    return id1ef0f108571f18aefac2c61a51615b16187fb53daeb8778a7d3fc99ee5b1a30.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i84f41c2d8a59ad72a2557098f3b161d03f5d8ac8e3987507b0fa7ecdf05601bf.MultiValueExtendedPropertiesRequestBuilder) {
    return i84f41c2d8a59ad72a2557098f3b161d03f5d8ac8e3987507b0fa7ecdf05601bf.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iec1f5e677b964978263004630240f212d437a10091543ccef2397757ae2b5c2e.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iec1f5e677b964978263004630240f212d437a10091543ccef2397757ae2b5c2e.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*iafafdd8f3e81db193ebe1f176611ba8d070c598779572146f18a7a2320d2b8fc.SingleValueExtendedPropertiesRequestBuilder) {
    return iafafdd8f3e81db193ebe1f176611ba8d070c598779572146f18a7a2320d2b8fc.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i1339b57f4311b596f1e163d07c2af7ba9105f71396c487bac3b1051937b9a343.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i1339b57f4311b596f1e163d07c2af7ba9105f71396c487bac3b1051937b9a343.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1e3d1947b13aae3f970d296ef71a4efc9e040742b37371ed6c138c0e253b1bf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarpermissions"
    i472232d154eda28ab8efd7fd9d2cc15ad8f3f68df3d90d6a2763b48e0a0fb5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/getschedule"
    i583738b29cef42ce350c32e40b7980bf07bc18ac8bde87fb209297b7485e7533 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview"
    ia34d5d2f8afdb98de85287585fbda3c10ce8763f60cf66a9c81a18f05aa37a17 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/singlevalueextendedproperties"
    ieb8d5dbdf577b51a4d75e40f7dc370bdfc82658f15e893321d1709a963ed78e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/multivalueextendedproperties"
    ied8df54815613d8a9050657a30d013ad8f4a43920a2fecbb829a53e69f7ef01a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events"
    ife78aa1675e428b35acac9c3d51963dd3aa40c64cd0973840b3ba47dc64b5df6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/allowedcalendarsharingroleswithuser"
    i00da9d0223281912e3c155074179db42ccbaa07b1ce54f8636bc10641c22665f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item"
    i285827c984f3b6bedeb07253a041706200e08a011dabe8ea34a88c14953200bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/singlevalueextendedproperties/item"
    i4e566b1ce953f539e8901716dfe774ae3102813087a55a1aa73d94bf3d010178 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item"
    i5a21f9744b199c1e3416254b970e7ad0211ea0e94c465251fdcf1915a650f97e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/multivalueextendedproperties/item"
    id86c8a2a7201c8af159055ea250d6a414ff1975b2f7042526a50a2700d86486d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarpermissions/item"
)

// CalendarRequestBuilder builds and executes requests for operations under \me\calendars\{calendar-id}
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
// CalendarRequestBuilderGetQueryParameters the user's calendars. Read-only. Nullable.
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
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \me\calendars\{calendar-id}\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ife78aa1675e428b35acac9c3d51963dd3aa40c64cd0973840b3ba47dc64b5df6.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ife78aa1675e428b35acac9c3d51963dd3aa40c64cd0973840b3ba47dc64b5df6.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i1e3d1947b13aae3f970d296ef71a4efc9e040742b37371ed6c138c0e253b1bf4.CalendarPermissionsRequestBuilder) {
    return i1e3d1947b13aae3f970d296ef71a4efc9e040742b37371ed6c138c0e253b1bf4.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*id86c8a2a7201c8af159055ea250d6a414ff1975b2f7042526a50a2700d86486d.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return id86c8a2a7201c8af159055ea250d6a414ff1975b2f7042526a50a2700d86486d.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i583738b29cef42ce350c32e40b7980bf07bc18ac8bde87fb209297b7485e7533.CalendarViewRequestBuilder) {
    return i583738b29cef42ce350c32e40b7980bf07bc18ac8bde87fb209297b7485e7533.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i4e566b1ce953f539e8901716dfe774ae3102813087a55a1aa73d94bf3d010178.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i4e566b1ce953f539e8901716dfe774ae3102813087a55a1aa73d94bf3d010178.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}{?select}";
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
// CreateDeleteRequestInformation the user's calendars. Read-only. Nullable.
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
// CreateGetRequestInformation the user's calendars. Read-only. Nullable.
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
// CreatePatchRequestInformation the user's calendars. Read-only. Nullable.
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
// Delete the user's calendars. Read-only. Nullable.
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
func (m *CalendarRequestBuilder) Events()(*ied8df54815613d8a9050657a30d013ad8f4a43920a2fecbb829a53e69f7ef01a.EventsRequestBuilder) {
    return ied8df54815613d8a9050657a30d013ad8f4a43920a2fecbb829a53e69f7ef01a.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*i00da9d0223281912e3c155074179db42ccbaa07b1ce54f8636bc10641c22665f.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i00da9d0223281912e3c155074179db42ccbaa07b1ce54f8636bc10641c22665f.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the user's calendars. Read-only. Nullable.
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
func (m *CalendarRequestBuilder) GetSchedule()(*i472232d154eda28ab8efd7fd9d2cc15ad8f3f68df3d90d6a2763b48e0a0fb5d2.GetScheduleRequestBuilder) {
    return i472232d154eda28ab8efd7fd9d2cc15ad8f3f68df3d90d6a2763b48e0a0fb5d2.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*ieb8d5dbdf577b51a4d75e40f7dc370bdfc82658f15e893321d1709a963ed78e2.MultiValueExtendedPropertiesRequestBuilder) {
    return ieb8d5dbdf577b51a4d75e40f7dc370bdfc82658f15e893321d1709a963ed78e2.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5a21f9744b199c1e3416254b970e7ad0211ea0e94c465251fdcf1915a650f97e.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5a21f9744b199c1e3416254b970e7ad0211ea0e94c465251fdcf1915a650f97e.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's calendars. Read-only. Nullable.
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*ia34d5d2f8afdb98de85287585fbda3c10ce8763f60cf66a9c81a18f05aa37a17.SingleValueExtendedPropertiesRequestBuilder) {
    return ia34d5d2f8afdb98de85287585fbda3c10ce8763f60cf66a9c81a18f05aa37a17.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i285827c984f3b6bedeb07253a041706200e08a011dabe8ea34a88c14953200bc.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i285827c984f3b6bedeb07253a041706200e08a011dabe8ea34a88c14953200bc.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
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

// CalendarItemRequestBuilder provides operations to manage the calendars property of the microsoft.graph.user entity.
type CalendarItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CalendarItemRequestBuilderDeleteOptions options for Delete
type CalendarItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CalendarItemRequestBuilderGetOptions options for Get
type CalendarItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CalendarItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CalendarItemRequestBuilderGetQueryParameters the user's calendars. Read-only. Nullable.
type CalendarItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// CalendarItemRequestBuilderPatchOptions options for Patch
type CalendarItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendarable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *CalendarItemRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ife78aa1675e428b35acac9c3d51963dd3aa40c64cd0973840b3ba47dc64b5df6.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ife78aa1675e428b35acac9c3d51963dd3aa40c64cd0973840b3ba47dc64b5df6.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarItemRequestBuilder) CalendarPermissions()(*i1e3d1947b13aae3f970d296ef71a4efc9e040742b37371ed6c138c0e253b1bf4.CalendarPermissionsRequestBuilder) {
    return i1e3d1947b13aae3f970d296ef71a4efc9e040742b37371ed6c138c0e253b1bf4.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarPermissions.item collection
func (m *CalendarItemRequestBuilder) CalendarPermissionsById(id string)(*id86c8a2a7201c8af159055ea250d6a414ff1975b2f7042526a50a2700d86486d.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return id86c8a2a7201c8af159055ea250d6a414ff1975b2f7042526a50a2700d86486d.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarItemRequestBuilder) CalendarView()(*i583738b29cef42ce350c32e40b7980bf07bc18ac8bde87fb209297b7485e7533.CalendarViewRequestBuilder) {
    return i583738b29cef42ce350c32e40b7980bf07bc18ac8bde87fb209297b7485e7533.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item collection
func (m *CalendarItemRequestBuilder) CalendarViewById(id string)(*i4e566b1ce953f539e8901716dfe774ae3102813087a55a1aa73d94bf3d010178.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i4e566b1ce953f539e8901716dfe774ae3102813087a55a1aa73d94bf3d010178.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarItemRequestBuilderInternal instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarItemRequestBuilder) {
    m := &CalendarItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarItemRequestBuilder instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) CreateDeleteRequestInformation(options *CalendarItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CalendarItemRequestBuilder) CreateGetRequestInformation(options *CalendarItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) CreatePatchRequestInformation(options *CalendarItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) Delete(options *CalendarItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarItemRequestBuilder) Events()(*ied8df54815613d8a9050657a30d013ad8f4a43920a2fecbb829a53e69f7ef01a.EventsRequestBuilder) {
    return ied8df54815613d8a9050657a30d013ad8f4a43920a2fecbb829a53e69f7ef01a.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item collection
func (m *CalendarItemRequestBuilder) EventsById(id string)(*i00da9d0223281912e3c155074179db42ccbaa07b1ce54f8636bc10641c22665f.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i00da9d0223281912e3c155074179db42ccbaa07b1ce54f8636bc10641c22665f.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the user's calendars. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) Get(options *CalendarItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendarable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateCalendarFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendarable), nil
}
func (m *CalendarItemRequestBuilder) GetSchedule()(*i472232d154eda28ab8efd7fd9d2cc15ad8f3f68df3d90d6a2763b48e0a0fb5d2.GetScheduleRequestBuilder) {
    return i472232d154eda28ab8efd7fd9d2cc15ad8f3f68df3d90d6a2763b48e0a0fb5d2.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarItemRequestBuilder) MultiValueExtendedProperties()(*ieb8d5dbdf577b51a4d75e40f7dc370bdfc82658f15e893321d1709a963ed78e2.MultiValueExtendedPropertiesRequestBuilder) {
    return ieb8d5dbdf577b51a4d75e40f7dc370bdfc82658f15e893321d1709a963ed78e2.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5a21f9744b199c1e3416254b970e7ad0211ea0e94c465251fdcf1915a650f97e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5a21f9744b199c1e3416254b970e7ad0211ea0e94c465251fdcf1915a650f97e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) Patch(options *CalendarItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarItemRequestBuilder) SingleValueExtendedProperties()(*ia34d5d2f8afdb98de85287585fbda3c10ce8763f60cf66a9c81a18f05aa37a17.SingleValueExtendedPropertiesRequestBuilder) {
    return ia34d5d2f8afdb98de85287585fbda3c10ce8763f60cf66a9c81a18f05aa37a17.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i285827c984f3b6bedeb07253a041706200e08a011dabe8ea34a88c14953200bc.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i285827c984f3b6bedeb07253a041706200e08a011dabe8ea34a88c14953200bc.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i107fef4eaf68aad98710c6ed5741a779aedb5359cc720332178dd533bf580859 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/multivalueextendedproperties"
    i19fbc325e46c8fffe47fa7226b4ad58a693aaa589bbf75c4b07e1d3df1e99237 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/allowedcalendarsharingroleswithuser"
    i1b213689741b1f5d2af2c287095b8e888f197fda7151c363c7e8a900f41b8c4a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events"
    i3744e1cd491347d480225f68a03302b535f0954ff443e4e58c022d15ffb6482d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/getschedule"
    i66de381a5b76495277b06584cd6039afa4769e36332757edaddd7d9e9b7aedb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarpermissions"
    i695392ea00eecabe1dcafb3489509e32336d2c87b43f2671c7934d387868992c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview"
    idd4feb523349c82feb1d6eef10e6276e441747df121268ee1ee289b42c4bb203 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/singlevalueextendedproperties"
    i1bf21f289996cae587c88ab3b232bd6c4afd875f12ad3e51adc0ca65805604d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item"
    i4ad3556e21fc36cd0eafb1a2fcc702ef4d28a32659279b8cdab60dacdc8da4b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/multivalueextendedproperties/item"
    icdf931cdb535bbf176c98a81c0de99ad4873545f5ec174c0e729a04b68633698 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item"
    iec879e43c2dcb4958281731f34b9c04a6adb22cd3d5fbc5b82f417f98dd09207 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarpermissions/item"
    ied63794e427bbe49231644e80f6559ae5dded695472591f4790b3801b63d6ca0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/singlevalueextendedproperties/item"
)

// CalendarRequestBuilder builds and executes requests for operations under \users\{user-id}\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}
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
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \users\{user-id}\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i19fbc325e46c8fffe47fa7226b4ad58a693aaa589bbf75c4b07e1d3df1e99237.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i19fbc325e46c8fffe47fa7226b4ad58a693aaa589bbf75c4b07e1d3df1e99237.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i66de381a5b76495277b06584cd6039afa4769e36332757edaddd7d9e9b7aedb2.CalendarPermissionsRequestBuilder) {
    return i66de381a5b76495277b06584cd6039afa4769e36332757edaddd7d9e9b7aedb2.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*iec879e43c2dcb4958281731f34b9c04a6adb22cd3d5fbc5b82f417f98dd09207.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return iec879e43c2dcb4958281731f34b9c04a6adb22cd3d5fbc5b82f417f98dd09207.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i695392ea00eecabe1dcafb3489509e32336d2c87b43f2671c7934d387868992c.CalendarViewRequestBuilder) {
    return i695392ea00eecabe1dcafb3489509e32336d2c87b43f2671c7934d387868992c.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*icdf931cdb535bbf176c98a81c0de99ad4873545f5ec174c0e729a04b68633698.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return icdf931cdb535bbf176c98a81c0de99ad4873545f5ec174c0e729a04b68633698.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}{?select}";
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
func (m *CalendarRequestBuilder) Events()(*i1b213689741b1f5d2af2c287095b8e888f197fda7151c363c7e8a900f41b8c4a.EventsRequestBuilder) {
    return i1b213689741b1f5d2af2c287095b8e888f197fda7151c363c7e8a900f41b8c4a.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*i1bf21f289996cae587c88ab3b232bd6c4afd875f12ad3e51adc0ca65805604d0.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i1bf21f289996cae587c88ab3b232bd6c4afd875f12ad3e51adc0ca65805604d0.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*i3744e1cd491347d480225f68a03302b535f0954ff443e4e58c022d15ffb6482d.GetScheduleRequestBuilder) {
    return i3744e1cd491347d480225f68a03302b535f0954ff443e4e58c022d15ffb6482d.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i107fef4eaf68aad98710c6ed5741a779aedb5359cc720332178dd533bf580859.MultiValueExtendedPropertiesRequestBuilder) {
    return i107fef4eaf68aad98710c6ed5741a779aedb5359cc720332178dd533bf580859.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4ad3556e21fc36cd0eafb1a2fcc702ef4d28a32659279b8cdab60dacdc8da4b0.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i4ad3556e21fc36cd0eafb1a2fcc702ef4d28a32659279b8cdab60dacdc8da4b0.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*idd4feb523349c82feb1d6eef10e6276e441747df121268ee1ee289b42c4bb203.SingleValueExtendedPropertiesRequestBuilder) {
    return idd4feb523349c82feb1d6eef10e6276e441747df121268ee1ee289b42c4bb203.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ied63794e427bbe49231644e80f6559ae5dded695472591f4790b3801b63d6ca0.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ied63794e427bbe49231644e80f6559ae5dded695472591f4790b3801b63d6ca0.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

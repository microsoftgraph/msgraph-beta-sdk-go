package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1398a734602b7083a166d3c39a433c5553621ad4c8e0216a4b839e8b831f3eba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/allowedcalendarsharingroleswithuser"
    i52969290f7fe959d6344cb99c3dbf36de567ba883a7f8e793cceb01560a3eee9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events"
    i5a8dc824626d608596468c8c2569c127c82227913446611aeb14a5e561ec3e81 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarpermissions"
    i6c20f628b7c7685400a144e3f1edfefacde585389763aae0916a903d35705372 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/singlevalueextendedproperties"
    iaf1e8a15a3c01f2204f94ecbf4ae1f9afb88aec6321e88eaa168ff240bde37da "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview"
    ic1e82b25ac9b6b26f8e7f1f932493eac94e12e158eb4a3df9f19f06d54b19114 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/getschedule"
    idddd37927b0549f9c2cae4baabd2894cfc4dcd36e790df408df463dac8093c8c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/multivalueextendedproperties"
    i2a1c5db47ed1e6fa6a0bf1fb730272391cf3e2ee9087749a96f77a6b17e1ab1d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item"
    i4812b126fc61d4810964e5652d96c94e77dd60c5b9f95babbea76c4661719e02 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/singlevalueextendedproperties/item"
    i5a19cfec81df5566ec7256cc31fb50b53bf9621a2d781bb4adcbea8d2db8b1fe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/multivalueextendedproperties/item"
    ib146c972bd8d23cbe643e701115fcd46bd4a5f95b1996f20e4aecca166cb54d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarpermissions/item"
    icef90dffc294a1de51c4ddd94b32ddd5c9eb7a31d60ee1da59d680c5abf58aac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item"
)

// CalendarRequestBuilder builds and executes requests for operations under \users\{user-id}\calendars\{calendar-id}
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
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \users\{user-id}\calendars\{calendar-id}\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i1398a734602b7083a166d3c39a433c5553621ad4c8e0216a4b839e8b831f3eba.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i1398a734602b7083a166d3c39a433c5553621ad4c8e0216a4b839e8b831f3eba.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i5a8dc824626d608596468c8c2569c127c82227913446611aeb14a5e561ec3e81.CalendarPermissionsRequestBuilder) {
    return i5a8dc824626d608596468c8c2569c127c82227913446611aeb14a5e561ec3e81.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*ib146c972bd8d23cbe643e701115fcd46bd4a5f95b1996f20e4aecca166cb54d8.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return ib146c972bd8d23cbe643e701115fcd46bd4a5f95b1996f20e4aecca166cb54d8.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*iaf1e8a15a3c01f2204f94ecbf4ae1f9afb88aec6321e88eaa168ff240bde37da.CalendarViewRequestBuilder) {
    return iaf1e8a15a3c01f2204f94ecbf4ae1f9afb88aec6321e88eaa168ff240bde37da.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i2a1c5db47ed1e6fa6a0bf1fb730272391cf3e2ee9087749a96f77a6b17e1ab1d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i2a1c5db47ed1e6fa6a0bf1fb730272391cf3e2ee9087749a96f77a6b17e1ab1d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}{?select}";
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
func (m *CalendarRequestBuilder) Events()(*i52969290f7fe959d6344cb99c3dbf36de567ba883a7f8e793cceb01560a3eee9.EventsRequestBuilder) {
    return i52969290f7fe959d6344cb99c3dbf36de567ba883a7f8e793cceb01560a3eee9.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*icef90dffc294a1de51c4ddd94b32ddd5c9eb7a31d60ee1da59d680c5abf58aac.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return icef90dffc294a1de51c4ddd94b32ddd5c9eb7a31d60ee1da59d680c5abf58aac.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*ic1e82b25ac9b6b26f8e7f1f932493eac94e12e158eb4a3df9f19f06d54b19114.GetScheduleRequestBuilder) {
    return ic1e82b25ac9b6b26f8e7f1f932493eac94e12e158eb4a3df9f19f06d54b19114.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*idddd37927b0549f9c2cae4baabd2894cfc4dcd36e790df408df463dac8093c8c.MultiValueExtendedPropertiesRequestBuilder) {
    return idddd37927b0549f9c2cae4baabd2894cfc4dcd36e790df408df463dac8093c8c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5a19cfec81df5566ec7256cc31fb50b53bf9621a2d781bb4adcbea8d2db8b1fe.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5a19cfec81df5566ec7256cc31fb50b53bf9621a2d781bb4adcbea8d2db8b1fe.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i6c20f628b7c7685400a144e3f1edfefacde585389763aae0916a903d35705372.SingleValueExtendedPropertiesRequestBuilder) {
    return i6c20f628b7c7685400a144e3f1edfefacde585389763aae0916a903d35705372.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4812b126fc61d4810964e5652d96c94e77dd60c5b9f95babbea76c4661719e02.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i4812b126fc61d4810964e5652d96c94e77dd60c5b9f95babbea76c4661719e02.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

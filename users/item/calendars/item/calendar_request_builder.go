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

type CalendarRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type CalendarRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i1398a734602b7083a166d3c39a433c5553621ad4c8e0216a4b839e8b831f3eba.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i1398a734602b7083a166d3c39a433c5553621ad4c8e0216a4b839e8b831f3eba.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i5a8dc824626d608596468c8c2569c127c82227913446611aeb14a5e561ec3e81.CalendarPermissionsRequestBuilder) {
    return i5a8dc824626d608596468c8c2569c127c82227913446611aeb14a5e561ec3e81.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/calendars/{calendar_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewCalendarRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *CalendarRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *CalendarRequestBuilder) CreateGetRequestInformation(q func (value *CalendarRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(CalendarRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *CalendarRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarRequestBuilder) Events()(*i52969290f7fe959d6344cb99c3dbf36de567ba883a7f8e793cceb01560a3eee9.EventsRequestBuilder) {
    return i52969290f7fe959d6344cb99c3dbf36de567ba883a7f8e793cceb01560a3eee9.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
func (m *CalendarRequestBuilder) Get(q func (value *CalendarRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCalendar() }, responseHandler)
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
func (m *CalendarRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i6c20f628b7c7685400a144e3f1edfefacde585389763aae0916a903d35705372.SingleValueExtendedPropertiesRequestBuilder) {
    return i6c20f628b7c7685400a144e3f1edfefacde585389763aae0916a903d35705372.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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

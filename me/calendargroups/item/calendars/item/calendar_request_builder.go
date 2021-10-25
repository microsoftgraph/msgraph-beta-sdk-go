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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ifdf0230f7f25a33effd6a753eaabc0e03fdd6933c1b808f4077742b51f7e8438.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ifdf0230f7f25a33effd6a753eaabc0e03fdd6933c1b808f4077742b51f7e8438.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i3a99b227ed3c41b86c6ed7d567ee4be24b191378e5c0a25d4ee01a213bae8cc2.CalendarPermissionsRequestBuilder) {
    return i3a99b227ed3c41b86c6ed7d567ee4be24b191378e5c0a25d4ee01a213bae8cc2.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}{?select,expand}";
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
func (m *CalendarRequestBuilder) Events()(*ie7a0d904d236f9db50b3140c9590b5afdb51f2d8f4f57f0bac0dd742c01a94ce.EventsRequestBuilder) {
    return ie7a0d904d236f9db50b3140c9590b5afdb51f2d8f4f57f0bac0dd742c01a94ce.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
func (m *CalendarRequestBuilder) GetSchedule()(*i2bbc53a9d42f80d48bde303e47a6edf7fd4514d20543476fc96d7654f44459c7.GetScheduleRequestBuilder) {
    return i2bbc53a9d42f80d48bde303e47a6edf7fd4514d20543476fc96d7654f44459c7.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i4cbc0b47a2f4d174f3aacee057ce86418a0d927f3a1682711f7f56d4fa432dc6.MultiValueExtendedPropertiesRequestBuilder) {
    return i4cbc0b47a2f4d174f3aacee057ce86418a0d927f3a1682711f7f56d4fa432dc6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i21d12d245d27df0693994a8bc5d5cb129c8c438d8f0d501c089b2742dd07da63.SingleValueExtendedPropertiesRequestBuilder) {
    return i21d12d245d27df0693994a8bc5d5cb129c8c438d8f0d501c089b2742dd07da63.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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

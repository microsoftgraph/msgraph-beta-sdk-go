package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3949123b651b9af052c8de6e7f5c595e080112b647c9a3626918cdcb55bc5902 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/singlevalueextendedproperties"
    i469e106165832f085e3fd273b20eefa6a88349191e203975179d82cd8fffe28b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/getschedule"
    i7b33c072b2a7ba3d82b9f69335f559cda585a5fab49a46127f175c6415fdb6c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/events"
    id04b107f7faaecce66eb6592f5f4f31f2f7b40b9bc6676aba5a06edbd4ee0fb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/multivalueextendedproperties"
    ie2cb151ce95182b548e67f27c8742fa555088b1e0e11ff7c1ce9d68f8b9e4880 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/calendarpermissions"
    ieda5ed6170b9c595db5a4f0e03158aac1856d1fd011e0472cd153361781f8835 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/calendarview"
    if6b857c575d61df1976f20edc812574d146eb925d961b7b3e2ba89bb97a44457 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/allowedcalendarsharingroleswithuser"
    i21c7ac2dc439414814348cc02de537ca481a6879b19eb08e0d4a93b78a1c9d60 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/calendarpermissions/item"
    i447ef94e57e8b820a1b16f1c03a188af03e71ba0d09d8a47dc42a19de9f764b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/singlevalueextendedproperties/item"
    i7035dcea3e0c498885a3febbcb76e699752f3550a3dffebcf26ee83bb5595f35 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/events/item"
    idd097aa2b0081a4a295f78cfc1bc963a100844ea07f6226dad11082cf1e15242 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/calendarview/item"
    ie14fdde3d6c5df3883fe4da02fa316e51f128cbca18042d14a0105153c540a5e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar/multivalueextendedproperties/item"
)

type CalendarRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type CalendarRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*if6b857c575d61df1976f20edc812574d146eb925d961b7b3e2ba89bb97a44457.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return if6b857c575d61df1976f20edc812574d146eb925d961b7b3e2ba89bb97a44457.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*ie2cb151ce95182b548e67f27c8742fa555088b1e0e11ff7c1ce9d68f8b9e4880.CalendarPermissionsRequestBuilder) {
    return ie2cb151ce95182b548e67f27c8742fa555088b1e0e11ff7c1ce9d68f8b9e4880.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i21c7ac2dc439414814348cc02de537ca481a6879b19eb08e0d4a93b78a1c9d60.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i21c7ac2dc439414814348cc02de537ca481a6879b19eb08e0d4a93b78a1c9d60.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ieda5ed6170b9c595db5a4f0e03158aac1856d1fd011e0472cd153361781f8835.CalendarViewRequestBuilder) {
    return ieda5ed6170b9c595db5a4f0e03158aac1856d1fd011e0472cd153361781f8835.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*idd097aa2b0081a4a295f78cfc1bc963a100844ea07f6226dad11082cf1e15242.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return idd097aa2b0081a4a295f78cfc1bc963a100844ea07f6226dad11082cf1e15242.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/groups/{group_id}/events/{event_id}/calendar{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
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
func (m *CalendarRequestBuilder) Events()(*i7b33c072b2a7ba3d82b9f69335f559cda585a5fab49a46127f175c6415fdb6c1.EventsRequestBuilder) {
    return i7b33c072b2a7ba3d82b9f69335f559cda585a5fab49a46127f175c6415fdb6c1.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*i7035dcea3e0c498885a3febbcb76e699752f3550a3dffebcf26ee83bb5595f35.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i7035dcea3e0c498885a3febbcb76e699752f3550a3dffebcf26ee83bb5595f35.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*i469e106165832f085e3fd273b20eefa6a88349191e203975179d82cd8fffe28b.GetScheduleRequestBuilder) {
    return i469e106165832f085e3fd273b20eefa6a88349191e203975179d82cd8fffe28b.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*id04b107f7faaecce66eb6592f5f4f31f2f7b40b9bc6676aba5a06edbd4ee0fb3.MultiValueExtendedPropertiesRequestBuilder) {
    return id04b107f7faaecce66eb6592f5f4f31f2f7b40b9bc6676aba5a06edbd4ee0fb3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie14fdde3d6c5df3883fe4da02fa316e51f128cbca18042d14a0105153c540a5e.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ie14fdde3d6c5df3883fe4da02fa316e51f128cbca18042d14a0105153c540a5e.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i3949123b651b9af052c8de6e7f5c595e080112b647c9a3626918cdcb55bc5902.SingleValueExtendedPropertiesRequestBuilder) {
    return i3949123b651b9af052c8de6e7f5c595e080112b647c9a3626918cdcb55bc5902.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i447ef94e57e8b820a1b16f1c03a188af03e71ba0d09d8a47dc42a19de9f764b7.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i447ef94e57e8b820a1b16f1c03a188af03e71ba0d09d8a47dc42a19de9f764b7.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i07589e5c53764a8b7f9186999e5ac2c7a82f10c525e08d08b4157884b2ec1ae7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/events"
    i08608bf5b96da051cc1cc910acc8d3f634f017d82416f3e0adc9fcdba93b638e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/calendarview"
    i5d60b0b9061864cba5e24638f8664aa037585cf8cf4d5a9b6d2817286b8360bb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/calendarpermissions"
    i708121932b406762e22ea2b912cff7368a77bfa2e098a26b39975ed68ed4f923 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/multivalueextendedproperties"
    i7ed43b8cd9450ed6934b53fc81f9bae4e596485336d0800d5048b39d1f0b818f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/allowedcalendarsharingroleswithuser"
    i91486446fa0c746740a36c7ee18fe9d0753613790613ae6d57ca6f46ec698e04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/singlevalueextendedproperties"
    ia6a37ec4851228eee682bf9233b63ffc75268a9c56a0de16b6dbe42ce35e6f69 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/getschedule"
    i6353afd79e46720776b44376eb3b69718a792316be79700910fb25c375d12e79 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/calendarview/item"
    i868667428ad5882f78e6367e4dccbe0b4c19d0496b30e61dbae9c19af6f2dbcd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/calendarpermissions/item"
    i8cc29c5d543c11fa04c48202f6c14bf7e322f71e93a34dc29da60756bcc2dc55 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/events/item"
    ie0b6d9b9e34ca499fb78a7851a04a0ced05d5dc963d6681022540387ecff4bfd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/multivalueextendedproperties/item"
    if2693a76fbddc151dce428226d837447faac4e95095ccd6cc890c54d60e436cc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar/singlevalueextendedproperties/item"
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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i7ed43b8cd9450ed6934b53fc81f9bae4e596485336d0800d5048b39d1f0b818f.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i7ed43b8cd9450ed6934b53fc81f9bae4e596485336d0800d5048b39d1f0b818f.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i5d60b0b9061864cba5e24638f8664aa037585cf8cf4d5a9b6d2817286b8360bb.CalendarPermissionsRequestBuilder) {
    return i5d60b0b9061864cba5e24638f8664aa037585cf8cf4d5a9b6d2817286b8360bb.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i868667428ad5882f78e6367e4dccbe0b4c19d0496b30e61dbae9c19af6f2dbcd.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i868667428ad5882f78e6367e4dccbe0b4c19d0496b30e61dbae9c19af6f2dbcd.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i08608bf5b96da051cc1cc910acc8d3f634f017d82416f3e0adc9fcdba93b638e.CalendarViewRequestBuilder) {
    return i08608bf5b96da051cc1cc910acc8d3f634f017d82416f3e0adc9fcdba93b638e.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i6353afd79e46720776b44376eb3b69718a792316be79700910fb25c375d12e79.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i6353afd79e46720776b44376eb3b69718a792316be79700910fb25c375d12e79.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/calendarView/{event_id}/calendar{?select,expand}";
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
func (m *CalendarRequestBuilder) Events()(*i07589e5c53764a8b7f9186999e5ac2c7a82f10c525e08d08b4157884b2ec1ae7.EventsRequestBuilder) {
    return i07589e5c53764a8b7f9186999e5ac2c7a82f10c525e08d08b4157884b2ec1ae7.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*i8cc29c5d543c11fa04c48202f6c14bf7e322f71e93a34dc29da60756bcc2dc55.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i8cc29c5d543c11fa04c48202f6c14bf7e322f71e93a34dc29da60756bcc2dc55.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*ia6a37ec4851228eee682bf9233b63ffc75268a9c56a0de16b6dbe42ce35e6f69.GetScheduleRequestBuilder) {
    return ia6a37ec4851228eee682bf9233b63ffc75268a9c56a0de16b6dbe42ce35e6f69.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i708121932b406762e22ea2b912cff7368a77bfa2e098a26b39975ed68ed4f923.MultiValueExtendedPropertiesRequestBuilder) {
    return i708121932b406762e22ea2b912cff7368a77bfa2e098a26b39975ed68ed4f923.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie0b6d9b9e34ca499fb78a7851a04a0ced05d5dc963d6681022540387ecff4bfd.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ie0b6d9b9e34ca499fb78a7851a04a0ced05d5dc963d6681022540387ecff4bfd.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i91486446fa0c746740a36c7ee18fe9d0753613790613ae6d57ca6f46ec698e04.SingleValueExtendedPropertiesRequestBuilder) {
    return i91486446fa0c746740a36c7ee18fe9d0753613790613ae6d57ca6f46ec698e04.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if2693a76fbddc151dce428226d837447faac4e95095ccd6cc890c54d60e436cc.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if2693a76fbddc151dce428226d837447faac4e95095ccd6cc890c54d60e436cc.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

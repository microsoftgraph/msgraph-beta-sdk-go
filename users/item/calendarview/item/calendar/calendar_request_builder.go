package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i03c699ff99af05c8da0ccabc9b19ac621aac2a5873b349ee6f5480c668fd73c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/calendarview"
    i1e600023cf502b3dfd6d10ca042857664c588f8ec60bbe2e838ee55f9b95cf93 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/allowedcalendarsharingroleswithuser"
    i4dd6f8353430d638cd57a1eb19aa48937408b9867bfddb61c7f2bdb5b2738636 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/events"
    i95c5be2233e98359982251375ae28585777d763ed57fa996880290fcedcaa756 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/getschedule"
    i992f1057f123c0d722110180b800607089aa869c03c5e9ae9506c8678db67e76 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/multivalueextendedproperties"
    ia5aa10f7cef8d361b8951c3b40f0e79dd11fc56fbdcd55e92914e81c4a403268 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/singlevalueextendedproperties"
    if55e49d8c63e6c8587409da7e5769e3a25dfd13c83c7e20f8b29a30a2f4cd879 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/calendarpermissions"
    i01b4addf39e3591c77ec5d603f94fc505bb0cfe1ef1092b38153eb846a649b09 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/multivalueextendedproperties/item"
    i9a80fde80fcace7ec1120c763d7e1a98154f015989d68eb8f8672f7f79e40409 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/events/item"
    iaa27aa128f1d6fb31d237f805be6637b09b964cb8e95ce74e4b03800d29a3247 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/calendarview/item"
    ie60200d99b8ebf1cbc7f60b560dbf254786c129fedc2719898483c37cd38b4c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/singlevalueextendedproperties/item"
    if2987591b8eda4899a6947c290e40c6d935de860b156eeff4e2d63f9768cb48f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar/calendarpermissions/item"
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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i1e600023cf502b3dfd6d10ca042857664c588f8ec60bbe2e838ee55f9b95cf93.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i1e600023cf502b3dfd6d10ca042857664c588f8ec60bbe2e838ee55f9b95cf93.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*if55e49d8c63e6c8587409da7e5769e3a25dfd13c83c7e20f8b29a30a2f4cd879.CalendarPermissionsRequestBuilder) {
    return if55e49d8c63e6c8587409da7e5769e3a25dfd13c83c7e20f8b29a30a2f4cd879.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*if2987591b8eda4899a6947c290e40c6d935de860b156eeff4e2d63f9768cb48f.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return if2987591b8eda4899a6947c290e40c6d935de860b156eeff4e2d63f9768cb48f.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i03c699ff99af05c8da0ccabc9b19ac621aac2a5873b349ee6f5480c668fd73c5.CalendarViewRequestBuilder) {
    return i03c699ff99af05c8da0ccabc9b19ac621aac2a5873b349ee6f5480c668fd73c5.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*iaa27aa128f1d6fb31d237f805be6637b09b964cb8e95ce74e4b03800d29a3247.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return iaa27aa128f1d6fb31d237f805be6637b09b964cb8e95ce74e4b03800d29a3247.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/calendarView/{event_id}/calendar{?select,expand}";
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
func (m *CalendarRequestBuilder) Events()(*i4dd6f8353430d638cd57a1eb19aa48937408b9867bfddb61c7f2bdb5b2738636.EventsRequestBuilder) {
    return i4dd6f8353430d638cd57a1eb19aa48937408b9867bfddb61c7f2bdb5b2738636.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*i9a80fde80fcace7ec1120c763d7e1a98154f015989d68eb8f8672f7f79e40409.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i9a80fde80fcace7ec1120c763d7e1a98154f015989d68eb8f8672f7f79e40409.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*i95c5be2233e98359982251375ae28585777d763ed57fa996880290fcedcaa756.GetScheduleRequestBuilder) {
    return i95c5be2233e98359982251375ae28585777d763ed57fa996880290fcedcaa756.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i992f1057f123c0d722110180b800607089aa869c03c5e9ae9506c8678db67e76.MultiValueExtendedPropertiesRequestBuilder) {
    return i992f1057f123c0d722110180b800607089aa869c03c5e9ae9506c8678db67e76.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i01b4addf39e3591c77ec5d603f94fc505bb0cfe1ef1092b38153eb846a649b09.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i01b4addf39e3591c77ec5d603f94fc505bb0cfe1ef1092b38153eb846a649b09.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*ia5aa10f7cef8d361b8951c3b40f0e79dd11fc56fbdcd55e92914e81c4a403268.SingleValueExtendedPropertiesRequestBuilder) {
    return ia5aa10f7cef8d361b8951c3b40f0e79dd11fc56fbdcd55e92914e81c4a403268.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie60200d99b8ebf1cbc7f60b560dbf254786c129fedc2719898483c37cd38b4c1.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ie60200d99b8ebf1cbc7f60b560dbf254786c129fedc2719898483c37cd38b4c1.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

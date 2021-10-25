package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1c81543176715430dfed16ca4ba9ce71f4a6e72d2a9ee15480a796ca6e7a6977 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/events"
    i3febf8aab239205529e4612d72d4ba51cbc93f924b9e0f199fad77e200c25d6b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/singlevalueextendedproperties"
    i50a83845117764d4eb7ed3f06acbd99842db483efae783e698585aa795f3629e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/multivalueextendedproperties"
    i59d29005bc3aeab89089800060ae7f2ffd2151ebeea580d29c83283938c421f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/calendarview"
    i5d489a322600c399fdedc6aced62ac05f02e4768b34e03c0d0c43a0d3f91f2d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/allowedcalendarsharingroleswithuser"
    id0adbea957932b06dd78ef7b88a847a18c39958be814e6f3cee44354d4cacad6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/getschedule"
    ieaa8ae04529471400ed51d7de011f1b52752acc45b5924cc2ec787a0fd0f1b33 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/calendarpermissions"
    i3f19bf3c70a5c1160aa8ae45500026d34af6fa6d9dffdf0a70094e72f7ce4f75 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/calendarpermissions/item"
    i4efcc695e7e8cc865246a06e13ced3836ecc141719585fa8a74c3580bda95c1f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/singlevalueextendedproperties/item"
    ibf776992d599fce46a5bb6e19c142fc6251910f387e00b377fc2cd55c54d9d2f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/multivalueextendedproperties/item"
    ifb79d12c8edcf0543eecdfc1eb98edff8345b031e7f4952db919e598f288b9cf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/calendarview/item"
    ifc8eaf472d534a6acd5be77a68646ed09c1ccc758d10bedf88b714a8e6826b2a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar/events/item"
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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i5d489a322600c399fdedc6aced62ac05f02e4768b34e03c0d0c43a0d3f91f2d3.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i5d489a322600c399fdedc6aced62ac05f02e4768b34e03c0d0c43a0d3f91f2d3.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*ieaa8ae04529471400ed51d7de011f1b52752acc45b5924cc2ec787a0fd0f1b33.CalendarPermissionsRequestBuilder) {
    return ieaa8ae04529471400ed51d7de011f1b52752acc45b5924cc2ec787a0fd0f1b33.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i3f19bf3c70a5c1160aa8ae45500026d34af6fa6d9dffdf0a70094e72f7ce4f75.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i3f19bf3c70a5c1160aa8ae45500026d34af6fa6d9dffdf0a70094e72f7ce4f75.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i59d29005bc3aeab89089800060ae7f2ffd2151ebeea580d29c83283938c421f8.CalendarViewRequestBuilder) {
    return i59d29005bc3aeab89089800060ae7f2ffd2151ebeea580d29c83283938c421f8.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*ifb79d12c8edcf0543eecdfc1eb98edff8345b031e7f4952db919e598f288b9cf.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ifb79d12c8edcf0543eecdfc1eb98edff8345b031e7f4952db919e598f288b9cf.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/groups/{group_id}/calendarView/{event_id}/calendar{?select,expand}";
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
func (m *CalendarRequestBuilder) Events()(*i1c81543176715430dfed16ca4ba9ce71f4a6e72d2a9ee15480a796ca6e7a6977.EventsRequestBuilder) {
    return i1c81543176715430dfed16ca4ba9ce71f4a6e72d2a9ee15480a796ca6e7a6977.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*ifc8eaf472d534a6acd5be77a68646ed09c1ccc758d10bedf88b714a8e6826b2a.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ifc8eaf472d534a6acd5be77a68646ed09c1ccc758d10bedf88b714a8e6826b2a.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*id0adbea957932b06dd78ef7b88a847a18c39958be814e6f3cee44354d4cacad6.GetScheduleRequestBuilder) {
    return id0adbea957932b06dd78ef7b88a847a18c39958be814e6f3cee44354d4cacad6.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i50a83845117764d4eb7ed3f06acbd99842db483efae783e698585aa795f3629e.MultiValueExtendedPropertiesRequestBuilder) {
    return i50a83845117764d4eb7ed3f06acbd99842db483efae783e698585aa795f3629e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibf776992d599fce46a5bb6e19c142fc6251910f387e00b377fc2cd55c54d9d2f.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ibf776992d599fce46a5bb6e19c142fc6251910f387e00b377fc2cd55c54d9d2f.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i3febf8aab239205529e4612d72d4ba51cbc93f924b9e0f199fad77e200c25d6b.SingleValueExtendedPropertiesRequestBuilder) {
    return i3febf8aab239205529e4612d72d4ba51cbc93f924b9e0f199fad77e200c25d6b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4efcc695e7e8cc865246a06e13ced3836ecc141719585fa8a74c3580bda95c1f.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i4efcc695e7e8cc865246a06e13ced3836ecc141719585fa8a74c3580bda95c1f.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

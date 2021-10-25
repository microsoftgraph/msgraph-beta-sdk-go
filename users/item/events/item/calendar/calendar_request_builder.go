package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1277f92db9d06c42c5bf077928d90d27b913b76b52f4926ba80f9c9159159571 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/allowedcalendarsharingroleswithuser"
    i2a640f4d8b1b637316403ce5017b5cb441486636e1254ecc4096951c18cf3dec "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/calendarview"
    ib8a3bd4ce0a0955f270fa1af2a6cc2ff74d752a50925a561aaa671972382537d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/getschedule"
    ic8ec8c7a307549d429d164cd83af7a075bad3055f7666a72b13a808d8b3a4974 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/events"
    icfb5984ee4c9dd8c169336826f2529a9bcfca6e088b552d64ff1f724ee578d4d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/calendarpermissions"
    ifc4f69cc19a3d35b50a62ece0b60b56feaf0515e0d3ddbb065a2a5154160ac6b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/singlevalueextendedproperties"
    ifdb40715f9d351a34d410c9794bd43c98db227777f19185f449ce2e395129a6d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/multivalueextendedproperties"
    i73df58d0f07876f22883f182a8fe518c598db69bd0e807b6d694a7a8bacce2a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/calendarpermissions/item"
    i8b4003c8c76f553df434b05216802452394e598cef225538b6ce11a3323f4143 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/singlevalueextendedproperties/item"
    ia9f9037fb7272d3f731c18f4a2c1be0897485e8be8ba63dcf4c71ca3baf0c5ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/events/item"
    ic1d4fe87963568a2be5f429ee535423517a4faf3472d4a8cc41a1510b6c6e6b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/multivalueextendedproperties/item"
    if346fb3637637d595c322048c3ba6ca6ecd20fa853f8b6eb533c49838ddc5613 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar/calendarview/item"
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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i1277f92db9d06c42c5bf077928d90d27b913b76b52f4926ba80f9c9159159571.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i1277f92db9d06c42c5bf077928d90d27b913b76b52f4926ba80f9c9159159571.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*icfb5984ee4c9dd8c169336826f2529a9bcfca6e088b552d64ff1f724ee578d4d.CalendarPermissionsRequestBuilder) {
    return icfb5984ee4c9dd8c169336826f2529a9bcfca6e088b552d64ff1f724ee578d4d.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i73df58d0f07876f22883f182a8fe518c598db69bd0e807b6d694a7a8bacce2a1.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i73df58d0f07876f22883f182a8fe518c598db69bd0e807b6d694a7a8bacce2a1.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i2a640f4d8b1b637316403ce5017b5cb441486636e1254ecc4096951c18cf3dec.CalendarViewRequestBuilder) {
    return i2a640f4d8b1b637316403ce5017b5cb441486636e1254ecc4096951c18cf3dec.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*if346fb3637637d595c322048c3ba6ca6ecd20fa853f8b6eb533c49838ddc5613.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return if346fb3637637d595c322048c3ba6ca6ecd20fa853f8b6eb533c49838ddc5613.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/events/{event_id}/calendar{?select,expand}";
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
func (m *CalendarRequestBuilder) Events()(*ic8ec8c7a307549d429d164cd83af7a075bad3055f7666a72b13a808d8b3a4974.EventsRequestBuilder) {
    return ic8ec8c7a307549d429d164cd83af7a075bad3055f7666a72b13a808d8b3a4974.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*ia9f9037fb7272d3f731c18f4a2c1be0897485e8be8ba63dcf4c71ca3baf0c5ed.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ia9f9037fb7272d3f731c18f4a2c1be0897485e8be8ba63dcf4c71ca3baf0c5ed.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*ib8a3bd4ce0a0955f270fa1af2a6cc2ff74d752a50925a561aaa671972382537d.GetScheduleRequestBuilder) {
    return ib8a3bd4ce0a0955f270fa1af2a6cc2ff74d752a50925a561aaa671972382537d.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*ifdb40715f9d351a34d410c9794bd43c98db227777f19185f449ce2e395129a6d.MultiValueExtendedPropertiesRequestBuilder) {
    return ifdb40715f9d351a34d410c9794bd43c98db227777f19185f449ce2e395129a6d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic1d4fe87963568a2be5f429ee535423517a4faf3472d4a8cc41a1510b6c6e6b5.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ic1d4fe87963568a2be5f429ee535423517a4faf3472d4a8cc41a1510b6c6e6b5.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*ifc4f69cc19a3d35b50a62ece0b60b56feaf0515e0d3ddbb065a2a5154160ac6b.SingleValueExtendedPropertiesRequestBuilder) {
    return ifc4f69cc19a3d35b50a62ece0b60b56feaf0515e0d3ddbb065a2a5154160ac6b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8b4003c8c76f553df434b05216802452394e598cef225538b6ce11a3323f4143.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i8b4003c8c76f553df434b05216802452394e598cef225538b6ce11a3323f4143.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

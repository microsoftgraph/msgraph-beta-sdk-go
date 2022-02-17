package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0c676d38d8314bffa49bea812001d84f18c4ddb055b033e0a539defb0d02d3e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences"
    i2532f3a30e33eb4bacd99e3aa1c95a819dd346720654a0e1a4d4eaaf4bbc1cc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/extensions"
    i37429c30f02551bdfbdd0e08ac54baeb24952b7e175ddbe2f32d40a682662585 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/cancel"
    i4473d5928afebc2f59c1862599eba34dccfa7ac885758ac0100c30a1311a5e28 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/tentativelyaccept"
    i5e035fd82444687e4bae951ea055b43f91b2243e639417476a58d1442016efe4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/attachments"
    i60bcc5d65b7dd984a48a299550c0724c142e8b89518fcaf945e7962c685b324d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/decline"
    i75652b3bea02403190e9330d7ef1a0545efd7dfb073dd0e29d08ac3107dd576c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/dismissreminder"
    ib86b35f0cebf04ca1f7dde01684b755ade9ddcc99c6acb7dd3cbab1d4570a2fa "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/forward"
    ic16f8f96aec796872d7677667f67931d92b34925f52a478d00ab14861c57bb18 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/accept"
    id6cd6a7760d192f1351760416ed540e8f199c56460676a76373cbeeb33f3b215 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/multivalueextendedproperties"
    id70b4efcde1a5decae4a08cc5d5790500ffea34b41906e83163eb824ff13b712 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/calendar"
    iddc90c75deab5f42a94960c7eac7853d7cd3219829ca8639ec70e07e51d69098 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/snoozereminder"
    ide9f97b833afcd44e28729097af5dd4aff8ee9cdbbbe131f31586941cf8e8bf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/singlevalueextendedproperties"
    if4e2cba388021fd91c0806e55e00dba53bcdbad5f977f815eb3b4503577c9775 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances"
    i0d3b3b0127a9ead0a16d0c73c03be7961006d9126aec22490405c42e3308133d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/attachments/item"
    i4ba2bdf718016a844aea70ab7664666d537cf068b189ade66bde1e59a27d8096 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/extensions/item"
    i90330af2351b81bdbcacfa3fb1bcabef2f1c565caf2362f15a9dbbbed185d4f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item"
    ib9db5026335e4341b909d92fdc380a8586c07ba447eed939427e2201a04d7ba3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/multivalueextendedproperties/item"
    ibda0295406a6a730f4c00e20649c362abe39faf50be42b8c74d6a751f1c248da "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/singlevalueextendedproperties/item"
    if0013914c12ae268d15c1d7063d1f2cff1b2a8a89c2edf5a9493e6d52a62f2ef "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item"
)

// EventRequestBuilder builds and executes requests for operations under \groups\{group-id}\calendar\calendarView\{event-id}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventRequestBuilderDeleteOptions options for Delete
type EventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetOptions options for Get
type EventRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
}
// EventRequestBuilderPatchOptions options for Patch
type EventRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventRequestBuilder) Accept()(*ic16f8f96aec796872d7677667f67931d92b34925f52a478d00ab14861c57bb18.AcceptRequestBuilder) {
    return ic16f8f96aec796872d7677667f67931d92b34925f52a478d00ab14861c57bb18.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i5e035fd82444687e4bae951ea055b43f91b2243e639417476a58d1442016efe4.AttachmentsRequestBuilder) {
    return i5e035fd82444687e4bae951ea055b43f91b2243e639417476a58d1442016efe4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*i0d3b3b0127a9ead0a16d0c73c03be7961006d9126aec22490405c42e3308133d.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i0d3b3b0127a9ead0a16d0c73c03be7961006d9126aec22490405c42e3308133d.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*id70b4efcde1a5decae4a08cc5d5790500ffea34b41906e83163eb824ff13b712.CalendarRequestBuilder) {
    return id70b4efcde1a5decae4a08cc5d5790500ffea34b41906e83163eb824ff13b712.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i37429c30f02551bdfbdd0e08ac54baeb24952b7e175ddbe2f32d40a682662585.CancelRequestBuilder) {
    return i37429c30f02551bdfbdd0e08ac54baeb24952b7e175ddbe2f32d40a682662585.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventRequestBuilder instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) CreateDeleteRequestInformation(options *EventRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) CreateGetRequestInformation(options *EventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) CreatePatchRequestInformation(options *EventRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*i60bcc5d65b7dd984a48a299550c0724c142e8b89518fcaf945e7962c685b324d.DeclineRequestBuilder) {
    return i60bcc5d65b7dd984a48a299550c0724c142e8b89518fcaf945e7962c685b324d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the calendar view for the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) Delete(options *EventRequestBuilderDeleteOptions)(error) {
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
func (m *EventRequestBuilder) DismissReminder()(*i75652b3bea02403190e9330d7ef1a0545efd7dfb073dd0e29d08ac3107dd576c.DismissReminderRequestBuilder) {
    return i75652b3bea02403190e9330d7ef1a0545efd7dfb073dd0e29d08ac3107dd576c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i0c676d38d8314bffa49bea812001d84f18c4ddb055b033e0a539defb0d02d3e9.ExceptionOccurrencesRequestBuilder) {
    return i0c676d38d8314bffa49bea812001d84f18c4ddb055b033e0a539defb0d02d3e9.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i90330af2351b81bdbcacfa3fb1bcabef2f1c565caf2362f15a9dbbbed185d4f2.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i90330af2351b81bdbcacfa3fb1bcabef2f1c565caf2362f15a9dbbbed185d4f2.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i2532f3a30e33eb4bacd99e3aa1c95a819dd346720654a0e1a4d4eaaf4bbc1cc6.ExtensionsRequestBuilder) {
    return i2532f3a30e33eb4bacd99e3aa1c95a819dd346720654a0e1a4d4eaaf4bbc1cc6.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i4ba2bdf718016a844aea70ab7664666d537cf068b189ade66bde1e59a27d8096.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i4ba2bdf718016a844aea70ab7664666d537cf068b189ade66bde1e59a27d8096.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ib86b35f0cebf04ca1f7dde01684b755ade9ddcc99c6acb7dd3cbab1d4570a2fa.ForwardRequestBuilder) {
    return ib86b35f0cebf04ca1f7dde01684b755ade9ddcc99c6acb7dd3cbab1d4570a2fa.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) Get(options *EventRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event), nil
}
func (m *EventRequestBuilder) Instances()(*if4e2cba388021fd91c0806e55e00dba53bcdbad5f977f815eb3b4503577c9775.InstancesRequestBuilder) {
    return if4e2cba388021fd91c0806e55e00dba53bcdbad5f977f815eb3b4503577c9775.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*if0013914c12ae268d15c1d7063d1f2cff1b2a8a89c2edf5a9493e6d52a62f2ef.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return if0013914c12ae268d15c1d7063d1f2cff1b2a8a89c2edf5a9493e6d52a62f2ef.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*id6cd6a7760d192f1351760416ed540e8f199c56460676a76373cbeeb33f3b215.MultiValueExtendedPropertiesRequestBuilder) {
    return id6cd6a7760d192f1351760416ed540e8f199c56460676a76373cbeeb33f3b215.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib9db5026335e4341b909d92fdc380a8586c07ba447eed939427e2201a04d7ba3.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib9db5026335e4341b909d92fdc380a8586c07ba447eed939427e2201a04d7ba3.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar view for the calendar. Navigation property. Read-only.
func (m *EventRequestBuilder) Patch(options *EventRequestBuilderPatchOptions)(error) {
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ide9f97b833afcd44e28729097af5dd4aff8ee9cdbbbe131f31586941cf8e8bf2.SingleValueExtendedPropertiesRequestBuilder) {
    return ide9f97b833afcd44e28729097af5dd4aff8ee9cdbbbe131f31586941cf8e8bf2.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ibda0295406a6a730f4c00e20649c362abe39faf50be42b8c74d6a751f1c248da.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ibda0295406a6a730f4c00e20649c362abe39faf50be42b8c74d6a751f1c248da.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*iddc90c75deab5f42a94960c7eac7853d7cd3219829ca8639ec70e07e51d69098.SnoozeReminderRequestBuilder) {
    return iddc90c75deab5f42a94960c7eac7853d7cd3219829ca8639ec70e07e51d69098.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i4473d5928afebc2f59c1862599eba34dccfa7ac885758ac0100c30a1311a5e28.TentativelyAcceptRequestBuilder) {
    return i4473d5928afebc2f59c1862599eba34dccfa7ac885758ac0100c30a1311a5e28.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

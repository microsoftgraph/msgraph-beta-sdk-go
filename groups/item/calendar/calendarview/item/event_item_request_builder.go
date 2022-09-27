package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type EventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string
}
// EventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*ic16f8f96aec796872d7677667f67931d92b34925f52a478d00ab14861c57bb18.AcceptRequestBuilder) {
    return ic16f8f96aec796872d7677667f67931d92b34925f52a478d00ab14861c57bb18.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i5e035fd82444687e4bae951ea055b43f91b2243e639417476a58d1442016efe4.AttachmentsRequestBuilder) {
    return i5e035fd82444687e4bae951ea055b43f91b2243e639417476a58d1442016efe4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i0d3b3b0127a9ead0a16d0c73c03be7961006d9126aec22490405c42e3308133d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i0d3b3b0127a9ead0a16d0c73c03be7961006d9126aec22490405c42e3308133d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*id70b4efcde1a5decae4a08cc5d5790500ffea34b41906e83163eb824ff13b712.CalendarRequestBuilder) {
    return id70b4efcde1a5decae4a08cc5d5790500ffea34b41906e83163eb824ff13b712.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i37429c30f02551bdfbdd0e08ac54baeb24952b7e175ddbe2f32d40a682662585.CancelRequestBuilder) {
    return i37429c30f02551bdfbdd0e08ac54baeb24952b7e175ddbe2f32d40a682662585.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}{?startDateTime*,endDateTime*,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i60bcc5d65b7dd984a48a299550c0724c142e8b89518fcaf945e7962c685b324d.DeclineRequestBuilder) {
    return i60bcc5d65b7dd984a48a299550c0724c142e8b89518fcaf945e7962c685b324d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i75652b3bea02403190e9330d7ef1a0545efd7dfb073dd0e29d08ac3107dd576c.DismissReminderRequestBuilder) {
    return i75652b3bea02403190e9330d7ef1a0545efd7dfb073dd0e29d08ac3107dd576c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i0c676d38d8314bffa49bea812001d84f18c4ddb055b033e0a539defb0d02d3e9.ExceptionOccurrencesRequestBuilder) {
    return i0c676d38d8314bffa49bea812001d84f18c4ddb055b033e0a539defb0d02d3e9.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i90330af2351b81bdbcacfa3fb1bcabef2f1c565caf2362f15a9dbbbed185d4f2.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i90330af2351b81bdbcacfa3fb1bcabef2f1c565caf2362f15a9dbbbed185d4f2.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i2532f3a30e33eb4bacd99e3aa1c95a819dd346720654a0e1a4d4eaaf4bbc1cc6.ExtensionsRequestBuilder) {
    return i2532f3a30e33eb4bacd99e3aa1c95a819dd346720654a0e1a4d4eaaf4bbc1cc6.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i4ba2bdf718016a844aea70ab7664666d537cf068b189ade66bde1e59a27d8096.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i4ba2bdf718016a844aea70ab7664666d537cf068b189ade66bde1e59a27d8096.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ib86b35f0cebf04ca1f7dde01684b755ade9ddcc99c6acb7dd3cbab1d4570a2fa.ForwardRequestBuilder) {
    return ib86b35f0cebf04ca1f7dde01684b755ade9ddcc99c6acb7dd3cbab1d4570a2fa.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*if4e2cba388021fd91c0806e55e00dba53bcdbad5f977f815eb3b4503577c9775.InstancesRequestBuilder) {
    return if4e2cba388021fd91c0806e55e00dba53bcdbad5f977f815eb3b4503577c9775.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*if0013914c12ae268d15c1d7063d1f2cff1b2a8a89c2edf5a9493e6d52a62f2ef.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return if0013914c12ae268d15c1d7063d1f2cff1b2a8a89c2edf5a9493e6d52a62f2ef.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*id6cd6a7760d192f1351760416ed540e8f199c56460676a76373cbeeb33f3b215.MultiValueExtendedPropertiesRequestBuilder) {
    return id6cd6a7760d192f1351760416ed540e8f199c56460676a76373cbeeb33f3b215.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib9db5026335e4341b909d92fdc380a8586c07ba447eed939427e2201a04d7ba3.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ib9db5026335e4341b909d92fdc380a8586c07ba447eed939427e2201a04d7ba3.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ide9f97b833afcd44e28729097af5dd4aff8ee9cdbbbe131f31586941cf8e8bf2.SingleValueExtendedPropertiesRequestBuilder) {
    return ide9f97b833afcd44e28729097af5dd4aff8ee9cdbbbe131f31586941cf8e8bf2.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ibda0295406a6a730f4c00e20649c362abe39faf50be42b8c74d6a751f1c248da.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ibda0295406a6a730f4c00e20649c362abe39faf50be42b8c74d6a751f1c248da.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*iddc90c75deab5f42a94960c7eac7853d7cd3219829ca8639ec70e07e51d69098.SnoozeReminderRequestBuilder) {
    return iddc90c75deab5f42a94960c7eac7853d7cd3219829ca8639ec70e07e51d69098.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i4473d5928afebc2f59c1862599eba34dccfa7ac885758ac0100c30a1311a5e28.TentativelyAcceptRequestBuilder) {
    return i4473d5928afebc2f59c1862599eba34dccfa7ac885758ac0100c30a1311a5e28.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

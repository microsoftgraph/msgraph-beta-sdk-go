package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07b0f5c126d3fd8e32f7c54a9d6f30a2c41a47efc37fbc9e8e49bdae5c2fd9e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/dismissreminder"
    i1a941c95f5627fa057d3dd8f69e7a76f5d91a7980d086a7888814dfd3e470968 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences"
    i296b15e1aad1ec85f3f8b45773230524b885d2afe41813071aa2b54c663963a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/forward"
    i4198a050c6fc737c2f34080e3829e36d908a821a9d7591fb186578ed061ee36d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/accept"
    i45997cdc7a813a982ecba7b8d3a586df5f2d88422358d670e2214804e3b53d00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/attachments"
    i4d8785343bb818131232baf9fc68c1d519ffd7bdf81da717ad76b235e251b9c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/multivalueextendedproperties"
    i535cf1cace249b664640330f949a728d46b263902138f84849c221b9abe894c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/calendar"
    i9a40e362aa3e36ede13cbac3bcd8bfe04901a57f8fb7b3971b67b06bf4459783 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances"
    ibe4fece240dca74238f03c2fa5a7bbcdd654c947e8b78e285b49ab9c7dfea505 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/tentativelyaccept"
    ibf789e58fb65fce33e2cae00e3c293e6ed19d7976b391f80d636389ef81f352c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/decline"
    icf3fd4bf51eebd15a020fc5021462c2e0a5518f7116ed3243ab0a1db3517572c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/cancel"
    id520c21d1389ffd38b7be74ec9a55a88e91677de395badf00b7ec8d8fa5400be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/extensions"
    if239cd3f23a0eda2f101672148a42c42a844663475919c4ecf0efcbe29330f19 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/snoozereminder"
    ifbc956c1dc35585d5237ea921bfac64de28ea44bcca4be498442c82a7a1f0a1e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/singlevalueextendedproperties"
    ia87adf0c3cc4103016c1134c4bb596fb8685d587b22a450dd7199fe89763d0ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item"
    iadba4b231f387f8389de2e976dffa453258e027a2e2a2e9b08e4283bd5bc7c45 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/extensions/item"
    ib8a97a009d82433874d4147f1a44d955f32a3435ad75792434ce1d59aaac5ae0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/attachments/item"
    ic0c0093229c77d1529a83c452145260e3b4924f194817de9810cfd2218f001e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/multivalueextendedproperties/item"
    if12270c62b6873dd5445fa0f6853dc371ddec134c8027ed30812ed5c01cbd2a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item"
    ifae272c2a7c5a0611ecb939bd341b17d95d1884ae7c2800fcad35ac9b30a8fcb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i4198a050c6fc737c2f34080e3829e36d908a821a9d7591fb186578ed061ee36d.AcceptRequestBuilder) {
    return i4198a050c6fc737c2f34080e3829e36d908a821a9d7591fb186578ed061ee36d.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i45997cdc7a813a982ecba7b8d3a586df5f2d88422358d670e2214804e3b53d00.AttachmentsRequestBuilder) {
    return i45997cdc7a813a982ecba7b8d3a586df5f2d88422358d670e2214804e3b53d00.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib8a97a009d82433874d4147f1a44d955f32a3435ad75792434ce1d59aaac5ae0.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ib8a97a009d82433874d4147f1a44d955f32a3435ad75792434ce1d59aaac5ae0.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i535cf1cace249b664640330f949a728d46b263902138f84849c221b9abe894c5.CalendarRequestBuilder) {
    return i535cf1cace249b664640330f949a728d46b263902138f84849c221b9abe894c5.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*icf3fd4bf51eebd15a020fc5021462c2e0a5518f7116ed3243ab0a1db3517572c.CancelRequestBuilder) {
    return icf3fd4bf51eebd15a020fc5021462c2e0a5518f7116ed3243ab0a1db3517572c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event%2Did}{?startDateTime*,endDateTime*,%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*ibf789e58fb65fce33e2cae00e3c293e6ed19d7976b391f80d636389ef81f352c.DeclineRequestBuilder) {
    return ibf789e58fb65fce33e2cae00e3c293e6ed19d7976b391f80d636389ef81f352c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i07b0f5c126d3fd8e32f7c54a9d6f30a2c41a47efc37fbc9e8e49bdae5c2fd9e0.DismissReminderRequestBuilder) {
    return i07b0f5c126d3fd8e32f7c54a9d6f30a2c41a47efc37fbc9e8e49bdae5c2fd9e0.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i1a941c95f5627fa057d3dd8f69e7a76f5d91a7980d086a7888814dfd3e470968.ExceptionOccurrencesRequestBuilder) {
    return i1a941c95f5627fa057d3dd8f69e7a76f5d91a7980d086a7888814dfd3e470968.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*ia87adf0c3cc4103016c1134c4bb596fb8685d587b22a450dd7199fe89763d0ad.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return ia87adf0c3cc4103016c1134c4bb596fb8685d587b22a450dd7199fe89763d0ad.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*id520c21d1389ffd38b7be74ec9a55a88e91677de395badf00b7ec8d8fa5400be.ExtensionsRequestBuilder) {
    return id520c21d1389ffd38b7be74ec9a55a88e91677de395badf00b7ec8d8fa5400be.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iadba4b231f387f8389de2e976dffa453258e027a2e2a2e9b08e4283bd5bc7c45.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return iadba4b231f387f8389de2e976dffa453258e027a2e2a2e9b08e4283bd5bc7c45.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i296b15e1aad1ec85f3f8b45773230524b885d2afe41813071aa2b54c663963a6.ForwardRequestBuilder) {
    return i296b15e1aad1ec85f3f8b45773230524b885d2afe41813071aa2b54c663963a6.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*i9a40e362aa3e36ede13cbac3bcd8bfe04901a57f8fb7b3971b67b06bf4459783.InstancesRequestBuilder) {
    return i9a40e362aa3e36ede13cbac3bcd8bfe04901a57f8fb7b3971b67b06bf4459783.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*if12270c62b6873dd5445fa0f6853dc371ddec134c8027ed30812ed5c01cbd2a1.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return if12270c62b6873dd5445fa0f6853dc371ddec134c8027ed30812ed5c01cbd2a1.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i4d8785343bb818131232baf9fc68c1d519ffd7bdf81da717ad76b235e251b9c0.MultiValueExtendedPropertiesRequestBuilder) {
    return i4d8785343bb818131232baf9fc68c1d519ffd7bdf81da717ad76b235e251b9c0.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic0c0093229c77d1529a83c452145260e3b4924f194817de9810cfd2218f001e2.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ic0c0093229c77d1529a83c452145260e3b4924f194817de9810cfd2218f001e2.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ifbc956c1dc35585d5237ea921bfac64de28ea44bcca4be498442c82a7a1f0a1e.SingleValueExtendedPropertiesRequestBuilder) {
    return ifbc956c1dc35585d5237ea921bfac64de28ea44bcca4be498442c82a7a1f0a1e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifae272c2a7c5a0611ecb939bd341b17d95d1884ae7c2800fcad35ac9b30a8fcb.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ifae272c2a7c5a0611ecb939bd341b17d95d1884ae7c2800fcad35ac9b30a8fcb.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*if239cd3f23a0eda2f101672148a42c42a844663475919c4ecf0efcbe29330f19.SnoozeReminderRequestBuilder) {
    return if239cd3f23a0eda2f101672148a42c42a844663475919c4ecf0efcbe29330f19.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibe4fece240dca74238f03c2fa5a7bbcdd654c947e8b78e285b49ab9c7dfea505.TentativelyAcceptRequestBuilder) {
    return ibe4fece240dca74238f03c2fa5a7bbcdd654c947e8b78e285b49ab9c7dfea505.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i276be7fad31bd680ac1411035eef4f3a00776a60d1125d214d24b184f7469220 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances"
    i3145dece064d967ee1c16e3e936c47ea8011ed2847cccb8dedaaa5c1eedaa013 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/cancel"
    i5986037676fe795074e810e4246ff753605940169fbbeed27b2aefa176b11751 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/calendar"
    i6caf8586c74aeb53bd8b7a86089cf9aa13e361aa8d0dd14f84486a74e16e0222 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/dismissreminder"
    i88425e8a17893a7e305a95729332a64c30460d352b0bf5845f5d7090f92c8f2b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/snoozereminder"
    i88ecb5106cc0bb521ce42c551cf99da444f6f2720897e5c4464c18a087b9fa03 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/decline"
    ia22ed1a7b043dc5283f91ddeb06193ce26ef17e32a8a4abd851e6155127ed218 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/tentativelyaccept"
    iaa58f3091eded3bb8f8022ff383adf6893ae04638de3510fb6f50e685897d70d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/accept"
    iab57ad64a283353ed19ad1d9c7814f25835c7ab4a0d769767a6db218d4703ee7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties"
    ib0aa93c8bad1e825bd5cdfdc4fa22e31b84830d4f2814d7e3ccfe19465129484 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/attachments"
    ib53b9735ba043d622c998452495258595cc62dbf7581e5a3605dfc29007b5997 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/forward"
    ic53f6b01f6b3da1d49fabe4f45413f9d94236e324fbc5f83450593600644550a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    ic99712d0b0d1392d14e91b9fb4b95688da8938561196adae18f98062d303b315 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences"
    ie128b633259361fd2a945ba339cd6061e92d56f2cdc9422dd860b687dc19ebfb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/extensions"
    i3e68bb94a10239ef0489773e8e20a5d33a576d00b16ded587e7df49c942e261a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/extensions/item"
    ib8afe549ec5fafbb94854314181e9eac838b20474be31174184b3169d64d40a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/attachments/item"
    ic9ed65a5fbed6446506612533b9a4560f521e68e07c8a6e70fc6eb495cf13716 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    id21bb72bf65597e6d902b0acc913993108ac302ca65b2d1f3c3fcaaada16028d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item"
    id97c351c3207d208122c8654eec8fa4fae5d90b3136fc2c79c3605d5a1fee25d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item"
    if2b4feb3f0e58d1cd206b56ee504e531df602a972bcdd86373ae72f7f5884e6f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
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
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
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
// Accept provides operations to call the accept method.
func (m *EventItemRequestBuilder) Accept()(*iaa58f3091eded3bb8f8022ff383adf6893ae04638de3510fb6f50e685897d70d.AcceptRequestBuilder) {
    return iaa58f3091eded3bb8f8022ff383adf6893ae04638de3510fb6f50e685897d70d.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*ib0aa93c8bad1e825bd5cdfdc4fa22e31b84830d4f2814d7e3ccfe19465129484.AttachmentsRequestBuilder) {
    return ib0aa93c8bad1e825bd5cdfdc4fa22e31b84830d4f2814d7e3ccfe19465129484.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib8afe549ec5fafbb94854314181e9eac838b20474be31174184b3169d64d40a6.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ib8afe549ec5fafbb94854314181e9eac838b20474be31174184b3169d64d40a6.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i5986037676fe795074e810e4246ff753605940169fbbeed27b2aefa176b11751.CalendarRequestBuilder) {
    return i5986037676fe795074e810e4246ff753605940169fbbeed27b2aefa176b11751.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*i3145dece064d967ee1c16e3e936c47ea8011ed2847cccb8dedaaa5c1eedaa013.CancelRequestBuilder) {
    return i3145dece064d967ee1c16e3e936c47ea8011ed2847cccb8dedaaa5c1eedaa013.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}{?%24select}";
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
func (m *EventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Decline provides operations to call the decline method.
func (m *EventItemRequestBuilder) Decline()(*i88ecb5106cc0bb521ce42c551cf99da444f6f2720897e5c4464c18a087b9fa03.DeclineRequestBuilder) {
    return i88ecb5106cc0bb521ce42c551cf99da444f6f2720897e5c4464c18a087b9fa03.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i6caf8586c74aeb53bd8b7a86089cf9aa13e361aa8d0dd14f84486a74e16e0222.DismissReminderRequestBuilder) {
    return i6caf8586c74aeb53bd8b7a86089cf9aa13e361aa8d0dd14f84486a74e16e0222.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ic99712d0b0d1392d14e91b9fb4b95688da8938561196adae18f98062d303b315.ExceptionOccurrencesRequestBuilder) {
    return ic99712d0b0d1392d14e91b9fb4b95688da8938561196adae18f98062d303b315.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*id97c351c3207d208122c8654eec8fa4fae5d90b3136fc2c79c3605d5a1fee25d.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return id97c351c3207d208122c8654eec8fa4fae5d90b3136fc2c79c3605d5a1fee25d.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*ie128b633259361fd2a945ba339cd6061e92d56f2cdc9422dd860b687dc19ebfb.ExtensionsRequestBuilder) {
    return ie128b633259361fd2a945ba339cd6061e92d56f2cdc9422dd860b687dc19ebfb.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i3e68bb94a10239ef0489773e8e20a5d33a576d00b16ded587e7df49c942e261a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i3e68bb94a10239ef0489773e8e20a5d33a576d00b16ded587e7df49c942e261a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*ib53b9735ba043d622c998452495258595cc62dbf7581e5a3605dfc29007b5997.ForwardRequestBuilder) {
    return ib53b9735ba043d622c998452495258595cc62dbf7581e5a3605dfc29007b5997.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Instances()(*i276be7fad31bd680ac1411035eef4f3a00776a60d1125d214d24b184f7469220.InstancesRequestBuilder) {
    return i276be7fad31bd680ac1411035eef4f3a00776a60d1125d214d24b184f7469220.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) InstancesById(id string)(*id21bb72bf65597e6d902b0acc913993108ac302ca65b2d1f3c3fcaaada16028d.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return id21bb72bf65597e6d902b0acc913993108ac302ca65b2d1f3c3fcaaada16028d.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iab57ad64a283353ed19ad1d9c7814f25835c7ab4a0d769767a6db218d4703ee7.MultiValueExtendedPropertiesRequestBuilder) {
    return iab57ad64a283353ed19ad1d9c7814f25835c7ab4a0d769767a6db218d4703ee7.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic9ed65a5fbed6446506612533b9a4560f521e68e07c8a6e70fc6eb495cf13716.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ic9ed65a5fbed6446506612533b9a4560f521e68e07c8a6e70fc6eb495cf13716.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ic53f6b01f6b3da1d49fabe4f45413f9d94236e324fbc5f83450593600644550a.SingleValueExtendedPropertiesRequestBuilder) {
    return ic53f6b01f6b3da1d49fabe4f45413f9d94236e324fbc5f83450593600644550a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if2b4feb3f0e58d1cd206b56ee504e531df602a972bcdd86373ae72f7f5884e6f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return if2b4feb3f0e58d1cd206b56ee504e531df602a972bcdd86373ae72f7f5884e6f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i88425e8a17893a7e305a95729332a64c30460d352b0bf5845f5d7090f92c8f2b.SnoozeReminderRequestBuilder) {
    return i88425e8a17893a7e305a95729332a64c30460d352b0bf5845f5d7090f92c8f2b.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*ia22ed1a7b043dc5283f91ddeb06193ce26ef17e32a8a4abd851e6155127ed218.TentativelyAcceptRequestBuilder) {
    return ia22ed1a7b043dc5283f91ddeb06193ce26ef17e32a8a4abd851e6155127ed218.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

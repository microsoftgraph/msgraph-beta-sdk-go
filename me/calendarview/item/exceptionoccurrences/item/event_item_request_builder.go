package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i099061384b6b2643b6ab93cb8d9bffb5afc89505e03e785e8165ef788bd636f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/accept"
    i103c78b9a83408fbdee7c5c5de0c2c26989785618e904c39762a8734811ba449 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/attachments"
    i24965cde6828d872d3c97cc064e3452c181d204ae1bdd9a5eef17e64adb44f14 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/calendar"
    i3fd6d2a1f574a9810dccae420fdd454a912df52f30df96f4a185d78cfd6f9c05 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i613860500ed7e5d220cc7aa468f340f68057e89254c0752dd1d5b23406d22c1c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances"
    i6c15a86bafa44851c86fe349ed155295840701f27d3d114abdaedaacd1b5f3e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/extensions"
    i70f37a3c126075fa8649885fae8b9269c2b414b14c355d5f56c6bacda4a2f642 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/forward"
    iaa0ad74adb35429933f3f3327c4b5922c01827bbda95187f603f9d4e9affa9d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    ib0a65b527edcae53c938ad49f5fb1239de546ce9a7e9ddc6ab30fdaa6b0c812b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/decline"
    icf1905e5a6e127a34411d19fe37e4528b1694e36035d0deea29c8d1bd69d4a4f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/snoozereminder"
    if1131a08ca4381b46da8854282509db2688d2d9b02903846bb32439175f3662c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    if5065ddf175b0f84cd6627753587da3589344bf0cd9a5e77e04e33a69206241f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/cancel"
    ifc0d3784d5ad28351ddcb223fd135e46f7d84b4158b7b72d356bdba776e96014 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/dismissreminder"
    i0cb8c1819c8dc655be3b43bd10af595d89a37c562af5ed1ad4e4058fefcfa02d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/attachments/item"
    i2964545c5adf79646b25d3619484630d27e775d018d943673dcad63b7066cd28 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item"
    i404080d5ce3bee4160c8d92d96ad2ce6f1bcc29995158139ca80da35ef3bc34f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    iaa4104cd8dc9111c099a2d6f38da15af0b5a77c014003598bb9647382389c86c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ib94d6c12e63d710d8b26d01371f9e1111e2c7adb20498d7e3e46721ccdae8166 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
func (m *EventItemRequestBuilder) Accept()(*i099061384b6b2643b6ab93cb8d9bffb5afc89505e03e785e8165ef788bd636f4.AcceptRequestBuilder) {
    return i099061384b6b2643b6ab93cb8d9bffb5afc89505e03e785e8165ef788bd636f4.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i103c78b9a83408fbdee7c5c5de0c2c26989785618e904c39762a8734811ba449.AttachmentsRequestBuilder) {
    return i103c78b9a83408fbdee7c5c5de0c2c26989785618e904c39762a8734811ba449.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i0cb8c1819c8dc655be3b43bd10af595d89a37c562af5ed1ad4e4058fefcfa02d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i0cb8c1819c8dc655be3b43bd10af595d89a37c562af5ed1ad4e4058fefcfa02d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i24965cde6828d872d3c97cc064e3452c181d204ae1bdd9a5eef17e64adb44f14.CalendarRequestBuilder) {
    return i24965cde6828d872d3c97cc064e3452c181d204ae1bdd9a5eef17e64adb44f14.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*if5065ddf175b0f84cd6627753587da3589344bf0cd9a5e77e04e33a69206241f.CancelRequestBuilder) {
    return if5065ddf175b0f84cd6627753587da3589344bf0cd9a5e77e04e33a69206241f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
// CreateGetRequestInformation get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Decline()(*ib0a65b527edcae53c938ad49f5fb1239de546ce9a7e9ddc6ab30fdaa6b0c812b.DeclineRequestBuilder) {
    return ib0a65b527edcae53c938ad49f5fb1239de546ce9a7e9ddc6ab30fdaa6b0c812b.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*ifc0d3784d5ad28351ddcb223fd135e46f7d84b4158b7b72d356bdba776e96014.DismissReminderRequestBuilder) {
    return ifc0d3784d5ad28351ddcb223fd135e46f7d84b4158b7b72d356bdba776e96014.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i6c15a86bafa44851c86fe349ed155295840701f27d3d114abdaedaacd1b5f3e4.ExtensionsRequestBuilder) {
    return i6c15a86bafa44851c86fe349ed155295840701f27d3d114abdaedaacd1b5f3e4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ib94d6c12e63d710d8b26d01371f9e1111e2c7adb20498d7e3e46721ccdae8166.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ib94d6c12e63d710d8b26d01371f9e1111e2c7adb20498d7e3e46721ccdae8166.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i70f37a3c126075fa8649885fae8b9269c2b414b14c355d5f56c6bacda4a2f642.ForwardRequestBuilder) {
    return i70f37a3c126075fa8649885fae8b9269c2b414b14c355d5f56c6bacda4a2f642.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Instances()(*i613860500ed7e5d220cc7aa468f340f68057e89254c0752dd1d5b23406d22c1c.InstancesRequestBuilder) {
    return i613860500ed7e5d220cc7aa468f340f68057e89254c0752dd1d5b23406d22c1c.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) InstancesById(id string)(*i2964545c5adf79646b25d3619484630d27e775d018d943673dcad63b7066cd28.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i2964545c5adf79646b25d3619484630d27e775d018d943673dcad63b7066cd28.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iaa0ad74adb35429933f3f3327c4b5922c01827bbda95187f603f9d4e9affa9d6.MultiValueExtendedPropertiesRequestBuilder) {
    return iaa0ad74adb35429933f3f3327c4b5922c01827bbda95187f603f9d4e9affa9d6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iaa4104cd8dc9111c099a2d6f38da15af0b5a77c014003598bb9647382389c86c.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iaa4104cd8dc9111c099a2d6f38da15af0b5a77c014003598bb9647382389c86c.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*if1131a08ca4381b46da8854282509db2688d2d9b02903846bb32439175f3662c.SingleValueExtendedPropertiesRequestBuilder) {
    return if1131a08ca4381b46da8854282509db2688d2d9b02903846bb32439175f3662c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i404080d5ce3bee4160c8d92d96ad2ce6f1bcc29995158139ca80da35ef3bc34f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i404080d5ce3bee4160c8d92d96ad2ce6f1bcc29995158139ca80da35ef3bc34f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*icf1905e5a6e127a34411d19fe37e4528b1694e36035d0deea29c8d1bd69d4a4f.SnoozeReminderRequestBuilder) {
    return icf1905e5a6e127a34411d19fe37e4528b1694e36035d0deea29c8d1bd69d4a4f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3fd6d2a1f574a9810dccae420fdd454a912df52f30df96f4a185d78cfd6f9c05.TentativelyAcceptRequestBuilder) {
    return i3fd6d2a1f574a9810dccae420fdd454a912df52f30df96f4a185d78cfd6f9c05.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

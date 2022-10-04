package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i21b91d21ace7dd784e511382cb621ebdf0cbebf09aa8603b2cc348973ff01e70 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/decline"
    i39a03fa97d00434f29d398e2c2abb5e4be5c135ddb6abb839904edceeb9c244e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/snoozereminder"
    i3a6c4f950f62a51604639dcff12cfeabc7ab5d2e72c62068656996c2718b66ac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances"
    i436ed00ce497758bc2d302d40575176733302a96d48aa967cc1acf8d11ed08ec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/attachments"
    i49178e35ea57efbebab8a8dd17c001a020367588ecf95239a5f8a51b8a6c7374 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/dismissreminder"
    i5411e7b3bbd860759f4df6028b0d077464624f8253497696a85bdb6530697eef "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/extensions"
    i5ab191c17f5dd0321843fcfe12d686bf2381580aa1cc2dba88c0b05e154c20c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i65bb3fb2883a932af85b792e43fe2181fb85465f76e9f5e4f096a9e5c0ad1e75 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/cancel"
    iab7313e3364628e59702f770b05edf7c5ae589b20a27a8c188cc01e82aa4a1cf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/accept"
    ib546f656eece3f6df3f8e7eb880b81e0e7451e3ba35e3125d828ce91ff67fa90 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/calendar"
    iddebd1e3f87433578f55e1ff78fe817a1a29e5b44f4b341356e679afa0a79530 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    ie552d7a0d127881bba6ab68162ee9abe40366795fb9d6dac98e574dbe69a95c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/tentativelyaccept"
    ieba73b19bb251fc9546081f5ea196883f6ea665606e2bffddf2bd0a47da43157 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/forward"
    i29699ba46743f920f16253ed8919bb6b466bc8071c7f7618def67908deac3616 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i65146039de459592459d94da453a296ef9ed6b8da487d70d37ecbc4999fc67bf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/attachments/item"
    i6fdb5177d22fd54431bb4fc9863f8f1d848124e48f10154dc76632a1c3227083 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i9f7a02a0074dccf8bef8f7977d3f5f045ed21b86d4cc5dd8bf437ff1443b3979 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/instances/item"
    ib2b505e37b2e23b27deeb67a2397766ae750669797c69980154851e357645580 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/exceptionoccurrences/item/extensions/item"
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
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*iab7313e3364628e59702f770b05edf7c5ae589b20a27a8c188cc01e82aa4a1cf.AcceptRequestBuilder) {
    return iab7313e3364628e59702f770b05edf7c5ae589b20a27a8c188cc01e82aa4a1cf.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i436ed00ce497758bc2d302d40575176733302a96d48aa967cc1acf8d11ed08ec.AttachmentsRequestBuilder) {
    return i436ed00ce497758bc2d302d40575176733302a96d48aa967cc1acf8d11ed08ec.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i65146039de459592459d94da453a296ef9ed6b8da487d70d37ecbc4999fc67bf.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i65146039de459592459d94da453a296ef9ed6b8da487d70d37ecbc4999fc67bf.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ib546f656eece3f6df3f8e7eb880b81e0e7451e3ba35e3125d828ce91ff67fa90.CalendarRequestBuilder) {
    return ib546f656eece3f6df3f8e7eb880b81e0e7451e3ba35e3125d828ce91ff67fa90.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i65bb3fb2883a932af85b792e43fe2181fb85465f76e9f5e4f096a9e5c0ad1e75.CancelRequestBuilder) {
    return i65bb3fb2883a932af85b792e43fe2181fb85465f76e9f5e4f096a9e5c0ad1e75.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i21b91d21ace7dd784e511382cb621ebdf0cbebf09aa8603b2cc348973ff01e70.DeclineRequestBuilder) {
    return i21b91d21ace7dd784e511382cb621ebdf0cbebf09aa8603b2cc348973ff01e70.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i49178e35ea57efbebab8a8dd17c001a020367588ecf95239a5f8a51b8a6c7374.DismissReminderRequestBuilder) {
    return i49178e35ea57efbebab8a8dd17c001a020367588ecf95239a5f8a51b8a6c7374.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i5411e7b3bbd860759f4df6028b0d077464624f8253497696a85bdb6530697eef.ExtensionsRequestBuilder) {
    return i5411e7b3bbd860759f4df6028b0d077464624f8253497696a85bdb6530697eef.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ib2b505e37b2e23b27deeb67a2397766ae750669797c69980154851e357645580.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ib2b505e37b2e23b27deeb67a2397766ae750669797c69980154851e357645580.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ieba73b19bb251fc9546081f5ea196883f6ea665606e2bffddf2bd0a47da43157.ForwardRequestBuilder) {
    return ieba73b19bb251fc9546081f5ea196883f6ea665606e2bffddf2bd0a47da43157.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*i3a6c4f950f62a51604639dcff12cfeabc7ab5d2e72c62068656996c2718b66ac.InstancesRequestBuilder) {
    return i3a6c4f950f62a51604639dcff12cfeabc7ab5d2e72c62068656996c2718b66ac.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i9f7a02a0074dccf8bef8f7977d3f5f045ed21b86d4cc5dd8bf437ff1443b3979.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i9f7a02a0074dccf8bef8f7977d3f5f045ed21b86d4cc5dd8bf437ff1443b3979.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iddebd1e3f87433578f55e1ff78fe817a1a29e5b44f4b341356e679afa0a79530.MultiValueExtendedPropertiesRequestBuilder) {
    return iddebd1e3f87433578f55e1ff78fe817a1a29e5b44f4b341356e679afa0a79530.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i29699ba46743f920f16253ed8919bb6b466bc8071c7f7618def67908deac3616.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i29699ba46743f920f16253ed8919bb6b466bc8071c7f7618def67908deac3616.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i5ab191c17f5dd0321843fcfe12d686bf2381580aa1cc2dba88c0b05e154c20c0.SingleValueExtendedPropertiesRequestBuilder) {
    return i5ab191c17f5dd0321843fcfe12d686bf2381580aa1cc2dba88c0b05e154c20c0.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i6fdb5177d22fd54431bb4fc9863f8f1d848124e48f10154dc76632a1c3227083.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i6fdb5177d22fd54431bb4fc9863f8f1d848124e48f10154dc76632a1c3227083.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i39a03fa97d00434f29d398e2c2abb5e4be5c135ddb6abb839904edceeb9c244e.SnoozeReminderRequestBuilder) {
    return i39a03fa97d00434f29d398e2c2abb5e4be5c135ddb6abb839904edceeb9c244e.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ie552d7a0d127881bba6ab68162ee9abe40366795fb9d6dac98e574dbe69a95c1.TentativelyAcceptRequestBuilder) {
    return ie552d7a0d127881bba6ab68162ee9abe40366795fb9d6dac98e574dbe69a95c1.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1b1111ea62e03ac11db57d95ca87c9209685cd4264e0f3a9b46b5dc367addc7c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/forward"
    i1c24408faa90db0103ca708bd0ab4df935472eb8c986cbcd6d6208b52a263610 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/dismissreminder"
    i3d7185caf42cd55d24ee58b9541cdac1afe68384d9bd99cd760e60149cb274e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/calendar"
    i7fe445f0475d401867c3b5e8986506ebc22a260df1a668ae410be96791d8db07 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances"
    i8267f6fabca0d1ba6ec4fce1ec64e717eeaa8676c36381a988eb901bcf20eb53 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    i9464a1ae547d95204b43b8f0b34a1b7e20857d76823e43e9b20c1b8079d0f58a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/extensions"
    ib18bded0ca532dde1fd1bc015965fe2ca043fc9511ba2020a5710243de2891f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/tentativelyaccept"
    ib5e3e2d47b4d9fe71ee0778722dd20d1288ff4686ddd8812ffb7505c387937b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/cancel"
    ib936006b84c4a0901664d0f12ed8fa925842c58ff8d7c9877f495fbae3cae93d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ic3c36e0f03e323474e209765f675e955a880426b5b930fea691bc394b146255e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/accept"
    iedcf7ca58e290b5af941bec45e0651818ef3f253888d0d4df377f57c1cbcdac6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/decline"
    if5e944aeb623a0c93576fab9b79580c8d8d3906a670d3cd546a45efed62b8da9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/snoozereminder"
    if97ab9caa5b781ef607163fcdf41cf8b3d35b2f76a0cd71bfb09be4f9e567aee "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/attachments"
    i69fa1e1bc399c5056a808fc71f80d2f56d8052e5b97beee8bb160391f90112be "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i7fb3d03a7e4471336529d13f85bfe81c94e9ca5d0de8a606f0417766e180b379 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i94d8e275817a71e29f28b7a6a0803e8ffd1e29cb43d37125e80d10435c40479d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item"
    iaac9073c68fbfc629f20be3985f9dbdcbadd8f401729a91b6263ce7c3dda2d8e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/extensions/item"
    ifae9f8a69cc08020158aef5e056f5ea120a02f1c2ff70534b52af33908863219 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/attachments/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Accept()(*ic3c36e0f03e323474e209765f675e955a880426b5b930fea691bc394b146255e.AcceptRequestBuilder) {
    return ic3c36e0f03e323474e209765f675e955a880426b5b930fea691bc394b146255e.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*if97ab9caa5b781ef607163fcdf41cf8b3d35b2f76a0cd71bfb09be4f9e567aee.AttachmentsRequestBuilder) {
    return if97ab9caa5b781ef607163fcdf41cf8b3d35b2f76a0cd71bfb09be4f9e567aee.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ifae9f8a69cc08020158aef5e056f5ea120a02f1c2ff70534b52af33908863219.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ifae9f8a69cc08020158aef5e056f5ea120a02f1c2ff70534b52af33908863219.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i3d7185caf42cd55d24ee58b9541cdac1afe68384d9bd99cd760e60149cb274e2.CalendarRequestBuilder) {
    return i3d7185caf42cd55d24ee58b9541cdac1afe68384d9bd99cd760e60149cb274e2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*ib5e3e2d47b4d9fe71ee0778722dd20d1288ff4686ddd8812ffb7505c387937b1.CancelRequestBuilder) {
    return ib5e3e2d47b4d9fe71ee0778722dd20d1288ff4686ddd8812ffb7505c387937b1.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
// CreateGetRequestInformation get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Decline()(*iedcf7ca58e290b5af941bec45e0651818ef3f253888d0d4df377f57c1cbcdac6.DeclineRequestBuilder) {
    return iedcf7ca58e290b5af941bec45e0651818ef3f253888d0d4df377f57c1cbcdac6.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i1c24408faa90db0103ca708bd0ab4df935472eb8c986cbcd6d6208b52a263610.DismissReminderRequestBuilder) {
    return i1c24408faa90db0103ca708bd0ab4df935472eb8c986cbcd6d6208b52a263610.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i9464a1ae547d95204b43b8f0b34a1b7e20857d76823e43e9b20c1b8079d0f58a.ExtensionsRequestBuilder) {
    return i9464a1ae547d95204b43b8f0b34a1b7e20857d76823e43e9b20c1b8079d0f58a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iaac9073c68fbfc629f20be3985f9dbdcbadd8f401729a91b6263ce7c3dda2d8e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return iaac9073c68fbfc629f20be3985f9dbdcbadd8f401729a91b6263ce7c3dda2d8e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i1b1111ea62e03ac11db57d95ca87c9209685cd4264e0f3a9b46b5dc367addc7c.ForwardRequestBuilder) {
    return i1b1111ea62e03ac11db57d95ca87c9209685cd4264e0f3a9b46b5dc367addc7c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Instances()(*i7fe445f0475d401867c3b5e8986506ebc22a260df1a668ae410be96791d8db07.InstancesRequestBuilder) {
    return i7fe445f0475d401867c3b5e8986506ebc22a260df1a668ae410be96791d8db07.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) InstancesById(id string)(*i94d8e275817a71e29f28b7a6a0803e8ffd1e29cb43d37125e80d10435c40479d.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i94d8e275817a71e29f28b7a6a0803e8ffd1e29cb43d37125e80d10435c40479d.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i8267f6fabca0d1ba6ec4fce1ec64e717eeaa8676c36381a988eb901bcf20eb53.MultiValueExtendedPropertiesRequestBuilder) {
    return i8267f6fabca0d1ba6ec4fce1ec64e717eeaa8676c36381a988eb901bcf20eb53.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i69fa1e1bc399c5056a808fc71f80d2f56d8052e5b97beee8bb160391f90112be.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i69fa1e1bc399c5056a808fc71f80d2f56d8052e5b97beee8bb160391f90112be.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib936006b84c4a0901664d0f12ed8fa925842c58ff8d7c9877f495fbae3cae93d.SingleValueExtendedPropertiesRequestBuilder) {
    return ib936006b84c4a0901664d0f12ed8fa925842c58ff8d7c9877f495fbae3cae93d.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7fb3d03a7e4471336529d13f85bfe81c94e9ca5d0de8a606f0417766e180b379.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i7fb3d03a7e4471336529d13f85bfe81c94e9ca5d0de8a606f0417766e180b379.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*if5e944aeb623a0c93576fab9b79580c8d8d3906a670d3cd546a45efed62b8da9.SnoozeReminderRequestBuilder) {
    return if5e944aeb623a0c93576fab9b79580c8d8d3906a670d3cd546a45efed62b8da9.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*ib18bded0ca532dde1fd1bc015965fe2ca043fc9511ba2020a5710243de2891f2.TentativelyAcceptRequestBuilder) {
    return ib18bded0ca532dde1fd1bc015965fe2ca043fc9511ba2020a5710243de2891f2.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

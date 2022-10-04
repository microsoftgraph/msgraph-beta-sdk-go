package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i17c54c3c2b19949bf0cbb0584162969422f54561f9dad78ae7d3e62f33e48eb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences"
    i3ed7421f446cad9ac11811d16d726d5f45763d9359d0a5055c35a0eececc2925 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/forward"
    i53f0fa0c5f3ba4917ce7348779fefd76e5c680dfe3e842fae8635a0d6a2dc0a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/snoozereminder"
    i675da67808cafe510e5e48bf8643155cad91ef964341494ab62fd141c604f253 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/dismissreminder"
    i6d050d3ce78ee9e35dcc70dc0ecbf6679e683ebe3e79de83e85710c7c5812dfd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/calendar"
    i70af343a303395c5a7b32914bcb8dd245cb9c68a08ba0057ad0bb9891f2f6dc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/decline"
    i7be662f58bcb608d18c2c743d57f5254d27e1dc748ae07da637f2993a0721e17 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/accept"
    i890a3cb447c32a6d47c4e737d9329ffca90841d490a71bd45cebe47994057e87 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/singlevalueextendedproperties"
    ia3da9e30021d43b632b9db43a8eccfae5bf5ca9e93c3a1ed67055bc6a32622c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/multivalueextendedproperties"
    iba406b1990c023050334f6c1671fe86ec3d69d7d022a18f0e750f0f74d7fefab "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/tentativelyaccept"
    ic2aca819f825d9c17ae3dc918c051897a82985726f8ef22ddb89aabed82293b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/cancel"
    id0a4bd7189b538d05956f396f1a157726db0f75815b29a9f870aea9c4c7b617d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/extensions"
    id4dcd79f762248fc30a62ad15e31adbcec15e2606c53275ef659a938740e2d18 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/attachments"
    i01a264d04855ac708e3451b590e7c6d8b312c5b96d61c6bcf49602b2aa5d80b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/singlevalueextendedproperties/item"
    i27d10484d39c10806952603b865770c2c4c14c65be466642912ef2cd2f170e1d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/multivalueextendedproperties/item"
    i5449aae33e844dcddcd22f22ea45bb9d7d058f9316b1ca1c659d5b512dd76609 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item"
    ic837e06f1daafaac2a2ea9f846424d09a61e6a52a9abeedd0a917fa1e6bf5f46 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/attachments/item"
    idbe496dae4cc76f6dfe06942068f13329ef5d48632a19bc696e24b3d7c7d1897 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i7be662f58bcb608d18c2c743d57f5254d27e1dc748ae07da637f2993a0721e17.AcceptRequestBuilder) {
    return i7be662f58bcb608d18c2c743d57f5254d27e1dc748ae07da637f2993a0721e17.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*id4dcd79f762248fc30a62ad15e31adbcec15e2606c53275ef659a938740e2d18.AttachmentsRequestBuilder) {
    return id4dcd79f762248fc30a62ad15e31adbcec15e2606c53275ef659a938740e2d18.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ic837e06f1daafaac2a2ea9f846424d09a61e6a52a9abeedd0a917fa1e6bf5f46.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ic837e06f1daafaac2a2ea9f846424d09a61e6a52a9abeedd0a917fa1e6bf5f46.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i6d050d3ce78ee9e35dcc70dc0ecbf6679e683ebe3e79de83e85710c7c5812dfd.CalendarRequestBuilder) {
    return i6d050d3ce78ee9e35dcc70dc0ecbf6679e683ebe3e79de83e85710c7c5812dfd.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ic2aca819f825d9c17ae3dc918c051897a82985726f8ef22ddb89aabed82293b7.CancelRequestBuilder) {
    return ic2aca819f825d9c17ae3dc918c051897a82985726f8ef22ddb89aabed82293b7.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Decline()(*i70af343a303395c5a7b32914bcb8dd245cb9c68a08ba0057ad0bb9891f2f6dc5.DeclineRequestBuilder) {
    return i70af343a303395c5a7b32914bcb8dd245cb9c68a08ba0057ad0bb9891f2f6dc5.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i675da67808cafe510e5e48bf8643155cad91ef964341494ab62fd141c604f253.DismissReminderRequestBuilder) {
    return i675da67808cafe510e5e48bf8643155cad91ef964341494ab62fd141c604f253.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i17c54c3c2b19949bf0cbb0584162969422f54561f9dad78ae7d3e62f33e48eb8.ExceptionOccurrencesRequestBuilder) {
    return i17c54c3c2b19949bf0cbb0584162969422f54561f9dad78ae7d3e62f33e48eb8.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i5449aae33e844dcddcd22f22ea45bb9d7d058f9316b1ca1c659d5b512dd76609.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i5449aae33e844dcddcd22f22ea45bb9d7d058f9316b1ca1c659d5b512dd76609.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*id0a4bd7189b538d05956f396f1a157726db0f75815b29a9f870aea9c4c7b617d.ExtensionsRequestBuilder) {
    return id0a4bd7189b538d05956f396f1a157726db0f75815b29a9f870aea9c4c7b617d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*idbe496dae4cc76f6dfe06942068f13329ef5d48632a19bc696e24b3d7c7d1897.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return idbe496dae4cc76f6dfe06942068f13329ef5d48632a19bc696e24b3d7c7d1897.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i3ed7421f446cad9ac11811d16d726d5f45763d9359d0a5055c35a0eececc2925.ForwardRequestBuilder) {
    return i3ed7421f446cad9ac11811d16d726d5f45763d9359d0a5055c35a0eececc2925.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ia3da9e30021d43b632b9db43a8eccfae5bf5ca9e93c3a1ed67055bc6a32622c9.MultiValueExtendedPropertiesRequestBuilder) {
    return ia3da9e30021d43b632b9db43a8eccfae5bf5ca9e93c3a1ed67055bc6a32622c9.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i27d10484d39c10806952603b865770c2c4c14c65be466642912ef2cd2f170e1d.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i27d10484d39c10806952603b865770c2c4c14c65be466642912ef2cd2f170e1d.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i890a3cb447c32a6d47c4e737d9329ffca90841d490a71bd45cebe47994057e87.SingleValueExtendedPropertiesRequestBuilder) {
    return i890a3cb447c32a6d47c4e737d9329ffca90841d490a71bd45cebe47994057e87.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i01a264d04855ac708e3451b590e7c6d8b312c5b96d61c6bcf49602b2aa5d80b7.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i01a264d04855ac708e3451b590e7c6d8b312c5b96d61c6bcf49602b2aa5d80b7.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i53f0fa0c5f3ba4917ce7348779fefd76e5c680dfe3e842fae8635a0d6a2dc0a2.SnoozeReminderRequestBuilder) {
    return i53f0fa0c5f3ba4917ce7348779fefd76e5c680dfe3e842fae8635a0d6a2dc0a2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*iba406b1990c023050334f6c1671fe86ec3d69d7d022a18f0e750f0f74d7fefab.TentativelyAcceptRequestBuilder) {
    return iba406b1990c023050334f6c1671fe86ec3d69d7d022a18f0e750f0f74d7fefab.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

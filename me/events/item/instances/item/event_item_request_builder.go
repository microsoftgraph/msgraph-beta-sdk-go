package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventItemRequestBuilder) Accept()(*i7be662f58bcb608d18c2c743d57f5254d27e1dc748ae07da637f2993a0721e17.AcceptRequestBuilder) {
    return i7be662f58bcb608d18c2c743d57f5254d27e1dc748ae07da637f2993a0721e17.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["attachment_id"] = id
    }
    return ic837e06f1daafaac2a2ea9f846424d09a61e6a52a9abeedd0a917fa1e6bf5f46.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i6d050d3ce78ee9e35dcc70dc0ecbf6679e683ebe3e79de83e85710c7c5812dfd.CalendarRequestBuilder) {
    return i6d050d3ce78ee9e35dcc70dc0ecbf6679e683ebe3e79de83e85710c7c5812dfd.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ic2aca819f825d9c17ae3dc918c051897a82985726f8ef22ddb89aabed82293b7.CancelRequestBuilder) {
    return ic2aca819f825d9c17ae3dc918c051897a82985726f8ef22ddb89aabed82293b7.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event_id}/instances/{event_id1}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property instances for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property instances in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventItemRequestBuilder) Decline()(*i70af343a303395c5a7b32914bcb8dd245cb9c68a08ba0057ad0bb9891f2f6dc5.DeclineRequestBuilder) {
    return i70af343a303395c5a7b32914bcb8dd245cb9c68a08ba0057ad0bb9891f2f6dc5.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*i675da67808cafe510e5e48bf8643155cad91ef964341494ab62fd141c604f253.DismissReminderRequestBuilder) {
    return i675da67808cafe510e5e48bf8643155cad91ef964341494ab62fd141c604f253.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["event_id2"] = id
    }
    return i5449aae33e844dcddcd22f22ea45bb9d7d058f9316b1ca1c659d5b512dd76609.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["extension_id"] = id
    }
    return idbe496dae4cc76f6dfe06942068f13329ef5d48632a19bc696e24b3d7c7d1897.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i3ed7421f446cad9ac11811d16d726d5f45763d9359d0a5055c35a0eececc2925.ForwardRequestBuilder) {
    return i3ed7421f446cad9ac11811d16d726d5f45763d9359d0a5055c35a0eececc2925.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable), nil
}
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
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i27d10484d39c10806952603b865770c2c4c14c65be466642912ef2cd2f170e1d.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i01a264d04855ac708e3451b590e7c6d8b312c5b96d61c6bcf49602b2aa5d80b7.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i53f0fa0c5f3ba4917ce7348779fefd76e5c680dfe3e842fae8635a0d6a2dc0a2.SnoozeReminderRequestBuilder) {
    return i53f0fa0c5f3ba4917ce7348779fefd76e5c680dfe3e842fae8635a0d6a2dc0a2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*iba406b1990c023050334f6c1671fe86ec3d69d7d022a18f0e750f0f74d7fefab.TentativelyAcceptRequestBuilder) {
    return iba406b1990c023050334f6c1671fe86ec3d69d7d022a18f0e750f0f74d7fefab.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

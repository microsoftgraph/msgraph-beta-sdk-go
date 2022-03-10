package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i0b36e7fbc97a2d8370891893c3c221c35d5170abe34df872e61110332d277b91 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/singlevalueextendedproperties"
    i1b9574753683cc8d202bd6ebe4d47660d0c4af6c9b6e916baae1434101fdc2cf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/extensions"
    i1f5e3d65a5cf771aca552b386199739fe6316961b86a8d245c7ee77e728e3042 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/calendar"
    i2c5a6d255f388fbe9f458474c3d7b8375b48a7a3427dcf464b1a2d6f3d8a0a5a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/multivalueextendedproperties"
    i59cd209bb367a3ab914363b6be8f5fb451216cf21050e0d5a896fa664fc34276 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/decline"
    i5b094fdc86d92954de5d4355ef2be265bc48f21e126689ede44ba2eee10350e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/attachments"
    i8e02c23b1d6b8b21ee7ec54c3b7d52509b13487c70ba47e4d17cd88e1f51cec4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/cancel"
    i99342ad7888c3b1b3de1ebe64b60c40c790ddd95c47e43ccecef6ef4862cfbf1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/forward"
    ia0961fac91912f701735f9967d33ee12b5c710fb6ded8ba0937d6628742ca96f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences"
    ib5416c6361187b3a360d9a310bfeae26f774901100a1d92c2d62e00cf601b6fe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/accept"
    ib9b9b411c5dd682c0d1ff5f4867c9cd6f857e0f0132e9460a7153ca0b4713af3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/dismissreminder"
    ic5a32a2b62a0bf106cf3425ad8844f8aff7ec2c750471a939a605ee262f4bce0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/tentativelyaccept"
    id825ca3df52f5a3f6c09527f1b9c03b850d3da5afd13f533f8d806dd8a94f968 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/snoozereminder"
    i28aca109418a72a7f76736bcf0260278de4d6157ed8f33f3f1c3c23942637431 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/singlevalueextendedproperties/item"
    i3a75008e9baf8acede4ca3413c6986388307a21f4074dba1b6f99b21d130eb68 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/multivalueextendedproperties/item"
    i89d4ea45299649e9b18492aaa78288497d03472775e3df598917d86ed072b972 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item"
    ia0eebe8027560a990309ae2f1973622881ab677fc092d3f9a9c72a48ae18c04b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/attachments/item"
    ied016554a0be1e56a06c14b4213fff5b2a35322c0b4bd752e9f5a25a7a40e738 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*ib5416c6361187b3a360d9a310bfeae26f774901100a1d92c2d62e00cf601b6fe.AcceptRequestBuilder) {
    return ib5416c6361187b3a360d9a310bfeae26f774901100a1d92c2d62e00cf601b6fe.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i5b094fdc86d92954de5d4355ef2be265bc48f21e126689ede44ba2eee10350e4.AttachmentsRequestBuilder) {
    return i5b094fdc86d92954de5d4355ef2be265bc48f21e126689ede44ba2eee10350e4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia0eebe8027560a990309ae2f1973622881ab677fc092d3f9a9c72a48ae18c04b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia0eebe8027560a990309ae2f1973622881ab677fc092d3f9a9c72a48ae18c04b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i1f5e3d65a5cf771aca552b386199739fe6316961b86a8d245c7ee77e728e3042.CalendarRequestBuilder) {
    return i1f5e3d65a5cf771aca552b386199739fe6316961b86a8d245c7ee77e728e3042.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i8e02c23b1d6b8b21ee7ec54c3b7d52509b13487c70ba47e4d17cd88e1f51cec4.CancelRequestBuilder) {
    return i8e02c23b1d6b8b21ee7ec54c3b7d52509b13487c70ba47e4d17cd88e1f51cec4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/events/{event_id}/instances/{event_id1}{?select}";
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
// CreateDeleteRequestInformation delete navigation property instances for users
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
// CreatePatchRequestInformation update the navigation property instances in users
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
func (m *EventItemRequestBuilder) Decline()(*i59cd209bb367a3ab914363b6be8f5fb451216cf21050e0d5a896fa664fc34276.DeclineRequestBuilder) {
    return i59cd209bb367a3ab914363b6be8f5fb451216cf21050e0d5a896fa664fc34276.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*ib9b9b411c5dd682c0d1ff5f4867c9cd6f857e0f0132e9460a7153ca0b4713af3.DismissReminderRequestBuilder) {
    return ib9b9b411c5dd682c0d1ff5f4867c9cd6f857e0f0132e9460a7153ca0b4713af3.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ia0961fac91912f701735f9967d33ee12b5c710fb6ded8ba0937d6628742ca96f.ExceptionOccurrencesRequestBuilder) {
    return ia0961fac91912f701735f9967d33ee12b5c710fb6ded8ba0937d6628742ca96f.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i89d4ea45299649e9b18492aaa78288497d03472775e3df598917d86ed072b972.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id2"] = id
    }
    return i89d4ea45299649e9b18492aaa78288497d03472775e3df598917d86ed072b972.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i1b9574753683cc8d202bd6ebe4d47660d0c4af6c9b6e916baae1434101fdc2cf.ExtensionsRequestBuilder) {
    return i1b9574753683cc8d202bd6ebe4d47660d0c4af6c9b6e916baae1434101fdc2cf.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ied016554a0be1e56a06c14b4213fff5b2a35322c0b4bd752e9f5a25a7a40e738.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ied016554a0be1e56a06c14b4213fff5b2a35322c0b4bd752e9f5a25a7a40e738.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i99342ad7888c3b1b3de1ebe64b60c40c790ddd95c47e43ccecef6ef4862cfbf1.ForwardRequestBuilder) {
    return i99342ad7888c3b1b3de1ebe64b60c40c790ddd95c47e43ccecef6ef4862cfbf1.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable), nil
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i2c5a6d255f388fbe9f458474c3d7b8375b48a7a3427dcf464b1a2d6f3d8a0a5a.MultiValueExtendedPropertiesRequestBuilder) {
    return i2c5a6d255f388fbe9f458474c3d7b8375b48a7a3427dcf464b1a2d6f3d8a0a5a.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i3a75008e9baf8acede4ca3413c6986388307a21f4074dba1b6f99b21d130eb68.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i3a75008e9baf8acede4ca3413c6986388307a21f4074dba1b6f99b21d130eb68.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i0b36e7fbc97a2d8370891893c3c221c35d5170abe34df872e61110332d277b91.SingleValueExtendedPropertiesRequestBuilder) {
    return i0b36e7fbc97a2d8370891893c3c221c35d5170abe34df872e61110332d277b91.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i28aca109418a72a7f76736bcf0260278de4d6157ed8f33f3f1c3c23942637431.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i28aca109418a72a7f76736bcf0260278de4d6157ed8f33f3f1c3c23942637431.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*id825ca3df52f5a3f6c09527f1b9c03b850d3da5afd13f533f8d806dd8a94f968.SnoozeReminderRequestBuilder) {
    return id825ca3df52f5a3f6c09527f1b9c03b850d3da5afd13f533f8d806dd8a94f968.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ic5a32a2b62a0bf106cf3425ad8844f8aff7ec2c750471a939a605ee262f4bce0.TentativelyAcceptRequestBuilder) {
    return ic5a32a2b62a0bf106cf3425ad8844f8aff7ec2c750471a939a605ee262f4bce0.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

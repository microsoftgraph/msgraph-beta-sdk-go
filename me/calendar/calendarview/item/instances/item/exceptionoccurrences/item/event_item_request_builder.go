package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i0aadbce778f36aa513868f516e5ffad3025dbe27f86d71a9d9147ec3bf3e1ad8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i541ac8bb3a9b9aa90c702d3089526f11eaa92bd7ed73703e0e937a8db0dea8f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i60c9ddb58a648e8f18cc0fb256cb3daf66e65facbdef5cdf1c1217d6e5a41f6e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i6ebb67bb2eaef9e8084d6969a59ba242663a177470d4f11ef3781a8c4154faff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    i7287dc284706d53453c76f84574ec8f968dffd9293efd4a11db069de780128f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i80e9536fb378174df68b546107be33706d34b7be2c79b6a8765cb147a788a144 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    iaa9c5e99699dc6c2ad9e713147782c977af2ef1d4bc78266b05cc510789c41fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    id1838a07822173eb48b071d2d003462d8a24197c8257227d39ad9d3e8a95b419 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    id840bf6fc51988f9de96f5b0e025c35bd6e23985a54427b145444ee887bacc38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    ieeda8b047771491bb8d8031b02f5631157d7f011c42c9426c307e60b041fb8ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    if15d10785ea5361129ee94f27ce2f941856cb340377fef7cff9416e0d9998492 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    ifcd5acbc04a2ee5d19368c443db0facfddec0e107b233f994bd5e37f124a9151 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i1381fe9be1b65482ff4df4b8fa880949f9ed8eb0bf61cc3e19c5270683c5c31b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    i536a15e5e1ca2b5878e0dbfe0363e0db4952eb3e3d9060788822293da1715ab7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    i5779a6defff10b2cdcd8fe04d42ad26909d0f6d1ce129f0ae36d20b669622101 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ic7eec9126a7d11c93e4978f0741e10116423345123847974101f52db6a3047b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
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
func (m *EventItemRequestBuilder) Accept()(*i6ebb67bb2eaef9e8084d6969a59ba242663a177470d4f11ef3781a8c4154faff.AcceptRequestBuilder) {
    return i6ebb67bb2eaef9e8084d6969a59ba242663a177470d4f11ef3781a8c4154faff.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*id1838a07822173eb48b071d2d003462d8a24197c8257227d39ad9d3e8a95b419.AttachmentsRequestBuilder) {
    return id1838a07822173eb48b071d2d003462d8a24197c8257227d39ad9d3e8a95b419.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i536a15e5e1ca2b5878e0dbfe0363e0db4952eb3e3d9060788822293da1715ab7.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i536a15e5e1ca2b5878e0dbfe0363e0db4952eb3e3d9060788822293da1715ab7.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*id840bf6fc51988f9de96f5b0e025c35bd6e23985a54427b145444ee887bacc38.CalendarRequestBuilder) {
    return id840bf6fc51988f9de96f5b0e025c35bd6e23985a54427b145444ee887bacc38.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i60c9ddb58a648e8f18cc0fb256cb3daf66e65facbdef5cdf1c1217d6e5a41f6e.CancelRequestBuilder) {
    return i60c9ddb58a648e8f18cc0fb256cb3daf66e65facbdef5cdf1c1217d6e5a41f6e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for me
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
// CreateGetRequestInformation get exceptionOccurrences from me
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) Decline()(*ieeda8b047771491bb8d8031b02f5631157d7f011c42c9426c307e60b041fb8ba.DeclineRequestBuilder) {
    return ieeda8b047771491bb8d8031b02f5631157d7f011c42c9426c307e60b041fb8ba.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i80e9536fb378174df68b546107be33706d34b7be2c79b6a8765cb147a788a144.DismissReminderRequestBuilder) {
    return i80e9536fb378174df68b546107be33706d34b7be2c79b6a8765cb147a788a144.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*ifcd5acbc04a2ee5d19368c443db0facfddec0e107b233f994bd5e37f124a9151.ExtensionsRequestBuilder) {
    return ifcd5acbc04a2ee5d19368c443db0facfddec0e107b233f994bd5e37f124a9151.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i1381fe9be1b65482ff4df4b8fa880949f9ed8eb0bf61cc3e19c5270683c5c31b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i1381fe9be1b65482ff4df4b8fa880949f9ed8eb0bf61cc3e19c5270683c5c31b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*iaa9c5e99699dc6c2ad9e713147782c977af2ef1d4bc78266b05cc510789c41fc.ForwardRequestBuilder) {
    return iaa9c5e99699dc6c2ad9e713147782c977af2ef1d4bc78266b05cc510789c41fc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*if15d10785ea5361129ee94f27ce2f941856cb340377fef7cff9416e0d9998492.MultiValueExtendedPropertiesRequestBuilder) {
    return if15d10785ea5361129ee94f27ce2f941856cb340377fef7cff9416e0d9998492.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic7eec9126a7d11c93e4978f0741e10116423345123847974101f52db6a3047b3.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ic7eec9126a7d11c93e4978f0741e10116423345123847974101f52db6a3047b3.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i0aadbce778f36aa513868f516e5ffad3025dbe27f86d71a9d9147ec3bf3e1ad8.SingleValueExtendedPropertiesRequestBuilder) {
    return i0aadbce778f36aa513868f516e5ffad3025dbe27f86d71a9d9147ec3bf3e1ad8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5779a6defff10b2cdcd8fe04d42ad26909d0f6d1ce129f0ae36d20b669622101.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i5779a6defff10b2cdcd8fe04d42ad26909d0f6d1ce129f0ae36d20b669622101.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i541ac8bb3a9b9aa90c702d3089526f11eaa92bd7ed73703e0e937a8db0dea8f3.SnoozeReminderRequestBuilder) {
    return i541ac8bb3a9b9aa90c702d3089526f11eaa92bd7ed73703e0e937a8db0dea8f3.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i7287dc284706d53453c76f84574ec8f968dffd9293efd4a11db069de780128f3.TentativelyAcceptRequestBuilder) {
    return i7287dc284706d53453c76f84574ec8f968dffd9293efd4a11db069de780128f3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

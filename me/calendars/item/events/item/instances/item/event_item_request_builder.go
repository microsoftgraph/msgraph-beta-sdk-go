package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i048cae98afed2bcfc52124c5a73b9724614574aa8636518cd4d1321fda7bfc07 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/tentativelyaccept"
    i09bbc4cb000c31d7d0dcc2d8d2a19d825cc947d91157459d7037271c58d21cf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/multivalueextendedproperties"
    i3d9b649c8ca7312d3ab05eed2b11021d247349a0d6684c22e352a75aa4b85780 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/extensions"
    i628be3bc640d9f7ad97770af1fb243685ad0495937c654c8bf3b409bbe2206fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/decline"
    i6bfd098884ab420a6c8080f2ea836f7a50f00fb47a739fcf71f3da14fad8daa6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/calendar"
    i6d79d0caec33d046d9854c5df0a59a36dd43c13a8931767a705d1709b177c1fd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/dismissreminder"
    i8e37990dab8ad7d081bb17c5a28ae24aab95bd004c238f26afbbc7ba02f88102 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/singlevalueextendedproperties"
    iccc31c4251a5fd741c25404a7d87955e7597506a851e61a35c3e472804baf2db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/cancel"
    id83e39dd29a8e3c2e9469f256dc18b44ff99d8b8ff288cabe59845d7f6f5f434 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/snoozereminder"
    id9f64ebcc1f739b750d6541599359d7f296e285ad3ee55ce90c2e2cf891d96d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/forward"
    ie914662dfee924e7297194e72446e1746c92203c8f0dea99255cdbdc185d3879 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/attachments"
    if11bd22064e35e4b39c13be59c77670bcc446ce0c6fbf2201c1d2c1246f4b44f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/accept"
    if7c0420b86216b7cf8e153fed9f0be790a80daf529bca27757ba6a5276a1762c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences"
    i1b4e4991f3d8bd757f3f6adaba892d558311427a930d1d93a3e84446a9c57fc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/singlevalueextendedproperties/item"
    i207bb0827311b7ceeaa1ab3b1ce079beb84fe98b8884787b8d0d2d54a5e7ee25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/multivalueextendedproperties/item"
    i94f4d1e66dfeaf32c3c9b4151d666ac69c7f00e8fa0a68c672c447c6009326c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/attachments/item"
    i9816b27bba149eee28e30c1ef89d862feb5f58d24fc827c4bcbde62ee5999e7c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/extensions/item"
    id22bc80b512b5084443604a788667020ca343bef03de2bf08262a8a317d05750 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*if11bd22064e35e4b39c13be59c77670bcc446ce0c6fbf2201c1d2c1246f4b44f.AcceptRequestBuilder) {
    return if11bd22064e35e4b39c13be59c77670bcc446ce0c6fbf2201c1d2c1246f4b44f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ie914662dfee924e7297194e72446e1746c92203c8f0dea99255cdbdc185d3879.AttachmentsRequestBuilder) {
    return ie914662dfee924e7297194e72446e1746c92203c8f0dea99255cdbdc185d3879.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i94f4d1e66dfeaf32c3c9b4151d666ac69c7f00e8fa0a68c672c447c6009326c0.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i94f4d1e66dfeaf32c3c9b4151d666ac69c7f00e8fa0a68c672c447c6009326c0.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i6bfd098884ab420a6c8080f2ea836f7a50f00fb47a739fcf71f3da14fad8daa6.CalendarRequestBuilder) {
    return i6bfd098884ab420a6c8080f2ea836f7a50f00fb47a739fcf71f3da14fad8daa6.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*iccc31c4251a5fd741c25404a7d87955e7597506a851e61a35c3e472804baf2db.CancelRequestBuilder) {
    return iccc31c4251a5fd741c25404a7d87955e7597506a851e61a35c3e472804baf2db.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/events/{event_id}/instances/{event_id1}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i628be3bc640d9f7ad97770af1fb243685ad0495937c654c8bf3b409bbe2206fa.DeclineRequestBuilder) {
    return i628be3bc640d9f7ad97770af1fb243685ad0495937c654c8bf3b409bbe2206fa.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i6d79d0caec33d046d9854c5df0a59a36dd43c13a8931767a705d1709b177c1fd.DismissReminderRequestBuilder) {
    return i6d79d0caec33d046d9854c5df0a59a36dd43c13a8931767a705d1709b177c1fd.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*if7c0420b86216b7cf8e153fed9f0be790a80daf529bca27757ba6a5276a1762c.ExceptionOccurrencesRequestBuilder) {
    return if7c0420b86216b7cf8e153fed9f0be790a80daf529bca27757ba6a5276a1762c.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*id22bc80b512b5084443604a788667020ca343bef03de2bf08262a8a317d05750.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id2"] = id
    }
    return id22bc80b512b5084443604a788667020ca343bef03de2bf08262a8a317d05750.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i3d9b649c8ca7312d3ab05eed2b11021d247349a0d6684c22e352a75aa4b85780.ExtensionsRequestBuilder) {
    return i3d9b649c8ca7312d3ab05eed2b11021d247349a0d6684c22e352a75aa4b85780.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9816b27bba149eee28e30c1ef89d862feb5f58d24fc827c4bcbde62ee5999e7c.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i9816b27bba149eee28e30c1ef89d862feb5f58d24fc827c4bcbde62ee5999e7c.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*id9f64ebcc1f739b750d6541599359d7f296e285ad3ee55ce90c2e2cf891d96d6.ForwardRequestBuilder) {
    return id9f64ebcc1f739b750d6541599359d7f296e285ad3ee55ce90c2e2cf891d96d6.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i09bbc4cb000c31d7d0dcc2d8d2a19d825cc947d91157459d7037271c58d21cf2.MultiValueExtendedPropertiesRequestBuilder) {
    return i09bbc4cb000c31d7d0dcc2d8d2a19d825cc947d91157459d7037271c58d21cf2.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i207bb0827311b7ceeaa1ab3b1ce079beb84fe98b8884787b8d0d2d54a5e7ee25.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i207bb0827311b7ceeaa1ab3b1ce079beb84fe98b8884787b8d0d2d54a5e7ee25.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i8e37990dab8ad7d081bb17c5a28ae24aab95bd004c238f26afbbc7ba02f88102.SingleValueExtendedPropertiesRequestBuilder) {
    return i8e37990dab8ad7d081bb17c5a28ae24aab95bd004c238f26afbbc7ba02f88102.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i1b4e4991f3d8bd757f3f6adaba892d558311427a930d1d93a3e84446a9c57fc5.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i1b4e4991f3d8bd757f3f6adaba892d558311427a930d1d93a3e84446a9c57fc5.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*id83e39dd29a8e3c2e9469f256dc18b44ff99d8b8ff288cabe59845d7f6f5f434.SnoozeReminderRequestBuilder) {
    return id83e39dd29a8e3c2e9469f256dc18b44ff99d8b8ff288cabe59845d7f6f5f434.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i048cae98afed2bcfc52124c5a73b9724614574aa8636518cd4d1321fda7bfc07.TentativelyAcceptRequestBuilder) {
    return i048cae98afed2bcfc52124c5a73b9724614574aa8636518cd4d1321fda7bfc07.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

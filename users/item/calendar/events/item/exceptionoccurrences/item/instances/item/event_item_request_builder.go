package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i008218bdd230e5c440d7e31c947f7a75f4a78eef5ae9c22dc03d02809c34aef9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/cancel"
    i2f9ecc515b427c27ce4a395cc9d19a130371a6cace2dc1c245f8cdd9acab8861 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i337144d43899bb837ba5777513dd1213a0ceb9f567fe71e4a0803ef764282995 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/extensions"
    i69b0effd541fa454e4dd7148f7e67e85a4163bb1037d60491c46dd3c05237538 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/decline"
    i6febbfb4afafa61cee72b6edd1b1b21f25a31004dc0f39c764b654663c004975 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/calendar"
    iace4d5272ade5fcb1ef3183556a2d6aec1c3ee25f9a8d8d52fcaeb5963356dd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/forward"
    ib683a8e220442458a142614a7c65ade726f53ee8614c850e794244390d2a1164 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/attachments"
    ibdbf84868b12f074e181c74cbcda6b4f165381d9adda54a16724cb41608a50ad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    ic0a6c4181b18e7fd1956ee26bbaa5ea870df1ba4459f26a9dfa53ca845b7c696 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    idb1613c815bc09b5374d87d34ae7cf66b46eaa861749a77720aac79223c08ad0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    idbe5dde1f6681ed18fed5872224c8c4b63ff83fb7ad4904bc45ba0d8f8ae85c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/accept"
    if2b8f75077d51fcce6c62feed81f91d9f34621eaecad7bbea439865d382d1f49 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i47ae5cd8dad46835d256384083c050ae7709917466645f5a7cf49a8aea896ca2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i6e1791c50c24ae04fbc968ed68d75f8579b4c9e0eab8a25d6c82cae53091a75f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    ib11f08b9334ad3e1181d8975fd0ce95e618ec443ad05d06d74558e69daae4bb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    ic4009be5e9259939a7c21fcd8f7bb24fadcf99569caa4f9b61bcfabdafdc3fb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*idbe5dde1f6681ed18fed5872224c8c4b63ff83fb7ad4904bc45ba0d8f8ae85c7.AcceptRequestBuilder) {
    return idbe5dde1f6681ed18fed5872224c8c4b63ff83fb7ad4904bc45ba0d8f8ae85c7.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ib683a8e220442458a142614a7c65ade726f53ee8614c850e794244390d2a1164.AttachmentsRequestBuilder) {
    return ib683a8e220442458a142614a7c65ade726f53ee8614c850e794244390d2a1164.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i47ae5cd8dad46835d256384083c050ae7709917466645f5a7cf49a8aea896ca2.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i47ae5cd8dad46835d256384083c050ae7709917466645f5a7cf49a8aea896ca2.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i6febbfb4afafa61cee72b6edd1b1b21f25a31004dc0f39c764b654663c004975.CalendarRequestBuilder) {
    return i6febbfb4afafa61cee72b6edd1b1b21f25a31004dc0f39c764b654663c004975.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i008218bdd230e5c440d7e31c947f7a75f4a78eef5ae9c22dc03d02809c34aef9.CancelRequestBuilder) {
    return i008218bdd230e5c440d7e31c947f7a75f4a78eef5ae9c22dc03d02809c34aef9.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/events/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*i69b0effd541fa454e4dd7148f7e67e85a4163bb1037d60491c46dd3c05237538.DeclineRequestBuilder) {
    return i69b0effd541fa454e4dd7148f7e67e85a4163bb1037d60491c46dd3c05237538.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*ibdbf84868b12f074e181c74cbcda6b4f165381d9adda54a16724cb41608a50ad.DismissReminderRequestBuilder) {
    return ibdbf84868b12f074e181c74cbcda6b4f165381d9adda54a16724cb41608a50ad.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i337144d43899bb837ba5777513dd1213a0ceb9f567fe71e4a0803ef764282995.ExtensionsRequestBuilder) {
    return i337144d43899bb837ba5777513dd1213a0ceb9f567fe71e4a0803ef764282995.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ib11f08b9334ad3e1181d8975fd0ce95e618ec443ad05d06d74558e69daae4bb5.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ib11f08b9334ad3e1181d8975fd0ce95e618ec443ad05d06d74558e69daae4bb5.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*iace4d5272ade5fcb1ef3183556a2d6aec1c3ee25f9a8d8d52fcaeb5963356dd0.ForwardRequestBuilder) {
    return iace4d5272ade5fcb1ef3183556a2d6aec1c3ee25f9a8d8d52fcaeb5963356dd0.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*if2b8f75077d51fcce6c62feed81f91d9f34621eaecad7bbea439865d382d1f49.MultiValueExtendedPropertiesRequestBuilder) {
    return if2b8f75077d51fcce6c62feed81f91d9f34621eaecad7bbea439865d382d1f49.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i6e1791c50c24ae04fbc968ed68d75f8579b4c9e0eab8a25d6c82cae53091a75f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i6e1791c50c24ae04fbc968ed68d75f8579b4c9e0eab8a25d6c82cae53091a75f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ic0a6c4181b18e7fd1956ee26bbaa5ea870df1ba4459f26a9dfa53ca845b7c696.SingleValueExtendedPropertiesRequestBuilder) {
    return ic0a6c4181b18e7fd1956ee26bbaa5ea870df1ba4459f26a9dfa53ca845b7c696.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic4009be5e9259939a7c21fcd8f7bb24fadcf99569caa4f9b61bcfabdafdc3fb8.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ic4009be5e9259939a7c21fcd8f7bb24fadcf99569caa4f9b61bcfabdafdc3fb8.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i2f9ecc515b427c27ce4a395cc9d19a130371a6cace2dc1c245f8cdd9acab8861.SnoozeReminderRequestBuilder) {
    return i2f9ecc515b427c27ce4a395cc9d19a130371a6cace2dc1c245f8cdd9acab8861.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*idb1613c815bc09b5374d87d34ae7cf66b46eaa861749a77720aac79223c08ad0.TentativelyAcceptRequestBuilder) {
    return idb1613c815bc09b5374d87d34ae7cf66b46eaa861749a77720aac79223c08ad0.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

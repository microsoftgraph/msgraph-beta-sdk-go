package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i02138ebcd8c54ad66fcc1b83c5aba691862df535a3f8cd0e5677dcddeb5322a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i0d5d0b4ad13238336e9693fd8966907ac00388d551ab79ac0d593835617b34ab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/cancel"
    i0fc36d4a45386d570212af1896e142e6704d07e98752280e60a7661f7f5ecc49 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i16df64c8d20d87a6609bcf56f4455b0f9c8fd3cdb27d9a6169b3a2ef60ec8360 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i5cf6d8c5f4f353d9d389158ac22f9dcaae558a31c2f6c99650d644edb641ddb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i684131f8a9a189e3f27ad66b5d5356f66131fe0d85d8b4a4597886a0d28f0fa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/forward"
    i6be21aff574b2becf273f1fabc83ec951e8ef79c132598c72836600ba7d2c962 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i8b24e9964b7c61504f543f7577e033a90854a5aaee3f167b1dfc0eb75a657124 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments"
    i938486f4acbe372b7aadcf4078813716807e26d4204315e7c962fd0761bd7013 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/accept"
    i9c3a2ce49501f8a32b329eacfb9c472ce25f422938a806c3a50cca37a16ed465 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/calendar"
    id6fc161013bc8ec2332f0381c2efc6de32567458d77dc23487e740c0776c344c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions"
    id79b50e5381bfe03a214a9398ec00b575ecceb92f3af4ffa0278cef9b3f2c4d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/decline"
    i1f332daed746484fef47995403c11dfd599d87f9e0f48d7eb0195e1f0d78461d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i4d1d20ce6486a01a14fbf938b2b458c7c73ba14533b80dcbf2eff8451f3d26ec "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    i9cef18a06a3c57f209883112eedc73f1b6ad564b50bb98be203409d2bd7b36d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    iab4639258e6785c5299a99811e24d3487e45d5c5b1f7831007604d0fad936c90 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) Accept()(*i938486f4acbe372b7aadcf4078813716807e26d4204315e7c962fd0761bd7013.AcceptRequestBuilder) {
    return i938486f4acbe372b7aadcf4078813716807e26d4204315e7c962fd0761bd7013.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i8b24e9964b7c61504f543f7577e033a90854a5aaee3f167b1dfc0eb75a657124.AttachmentsRequestBuilder) {
    return i8b24e9964b7c61504f543f7577e033a90854a5aaee3f167b1dfc0eb75a657124.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i9cef18a06a3c57f209883112eedc73f1b6ad564b50bb98be203409d2bd7b36d1.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i9cef18a06a3c57f209883112eedc73f1b6ad564b50bb98be203409d2bd7b36d1.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i9c3a2ce49501f8a32b329eacfb9c472ce25f422938a806c3a50cca37a16ed465.CalendarRequestBuilder) {
    return i9c3a2ce49501f8a32b329eacfb9c472ce25f422938a806c3a50cca37a16ed465.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i0d5d0b4ad13238336e9693fd8966907ac00388d551ab79ac0d593835617b34ab.CancelRequestBuilder) {
    return i0d5d0b4ad13238336e9693fd8966907ac00388d551ab79ac0d593835617b34ab.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/events/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for users
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
// CreateGetRequestInformation get exceptionOccurrences from users
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) Decline()(*id79b50e5381bfe03a214a9398ec00b575ecceb92f3af4ffa0278cef9b3f2c4d3.DeclineRequestBuilder) {
    return id79b50e5381bfe03a214a9398ec00b575ecceb92f3af4ffa0278cef9b3f2c4d3.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i0fc36d4a45386d570212af1896e142e6704d07e98752280e60a7661f7f5ecc49.DismissReminderRequestBuilder) {
    return i0fc36d4a45386d570212af1896e142e6704d07e98752280e60a7661f7f5ecc49.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*id6fc161013bc8ec2332f0381c2efc6de32567458d77dc23487e740c0776c344c.ExtensionsRequestBuilder) {
    return id6fc161013bc8ec2332f0381c2efc6de32567458d77dc23487e740c0776c344c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i4d1d20ce6486a01a14fbf938b2b458c7c73ba14533b80dcbf2eff8451f3d26ec.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i4d1d20ce6486a01a14fbf938b2b458c7c73ba14533b80dcbf2eff8451f3d26ec.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i684131f8a9a189e3f27ad66b5d5356f66131fe0d85d8b4a4597886a0d28f0fa0.ForwardRequestBuilder) {
    return i684131f8a9a189e3f27ad66b5d5356f66131fe0d85d8b4a4597886a0d28f0fa0.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i02138ebcd8c54ad66fcc1b83c5aba691862df535a3f8cd0e5677dcddeb5322a8.MultiValueExtendedPropertiesRequestBuilder) {
    return i02138ebcd8c54ad66fcc1b83c5aba691862df535a3f8cd0e5677dcddeb5322a8.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i1f332daed746484fef47995403c11dfd599d87f9e0f48d7eb0195e1f0d78461d.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i1f332daed746484fef47995403c11dfd599d87f9e0f48d7eb0195e1f0d78461d.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i6be21aff574b2becf273f1fabc83ec951e8ef79c132598c72836600ba7d2c962.SingleValueExtendedPropertiesRequestBuilder) {
    return i6be21aff574b2becf273f1fabc83ec951e8ef79c132598c72836600ba7d2c962.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iab4639258e6785c5299a99811e24d3487e45d5c5b1f7831007604d0fad936c90.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return iab4639258e6785c5299a99811e24d3487e45d5c5b1f7831007604d0fad936c90.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i5cf6d8c5f4f353d9d389158ac22f9dcaae558a31c2f6c99650d644edb641ddb8.SnoozeReminderRequestBuilder) {
    return i5cf6d8c5f4f353d9d389158ac22f9dcaae558a31c2f6c99650d644edb641ddb8.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i16df64c8d20d87a6609bcf56f4455b0f9c8fd3cdb27d9a6169b3a2ef60ec8360.TentativelyAcceptRequestBuilder) {
    return i16df64c8d20d87a6609bcf56f4455b0f9c8fd3cdb27d9a6169b3a2ef60ec8360.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

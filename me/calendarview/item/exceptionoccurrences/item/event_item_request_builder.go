package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
func (m *EventItemRequestBuilder) Accept()(*i099061384b6b2643b6ab93cb8d9bffb5afc89505e03e785e8165ef788bd636f4.AcceptRequestBuilder) {
    return i099061384b6b2643b6ab93cb8d9bffb5afc89505e03e785e8165ef788bd636f4.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i103c78b9a83408fbdee7c5c5de0c2c26989785618e904c39762a8734811ba449.AttachmentsRequestBuilder) {
    return i103c78b9a83408fbdee7c5c5de0c2c26989785618e904c39762a8734811ba449.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i0cb8c1819c8dc655be3b43bd10af595d89a37c562af5ed1ad4e4058fefcfa02d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i0cb8c1819c8dc655be3b43bd10af595d89a37c562af5ed1ad4e4058fefcfa02d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i24965cde6828d872d3c97cc064e3452c181d204ae1bdd9a5eef17e64adb44f14.CalendarRequestBuilder) {
    return i24965cde6828d872d3c97cc064e3452c181d204ae1bdd9a5eef17e64adb44f14.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*if5065ddf175b0f84cd6627753587da3589344bf0cd9a5e77e04e33a69206241f.CancelRequestBuilder) {
    return if5065ddf175b0f84cd6627753587da3589344bf0cd9a5e77e04e33a69206241f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event_id}/exceptionOccurrences/{event_id1}{?select,expand}";
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
func (m *EventItemRequestBuilder) Decline()(*ib0a65b527edcae53c938ad49f5fb1239de546ce9a7e9ddc6ab30fdaa6b0c812b.DeclineRequestBuilder) {
    return ib0a65b527edcae53c938ad49f5fb1239de546ce9a7e9ddc6ab30fdaa6b0c812b.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*ifc0d3784d5ad28351ddcb223fd135e46f7d84b4158b7b72d356bdba776e96014.DismissReminderRequestBuilder) {
    return ifc0d3784d5ad28351ddcb223fd135e46f7d84b4158b7b72d356bdba776e96014.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i6c15a86bafa44851c86fe349ed155295840701f27d3d114abdaedaacd1b5f3e4.ExtensionsRequestBuilder) {
    return i6c15a86bafa44851c86fe349ed155295840701f27d3d114abdaedaacd1b5f3e4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ib94d6c12e63d710d8b26d01371f9e1111e2c7adb20498d7e3e46721ccdae8166.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ib94d6c12e63d710d8b26d01371f9e1111e2c7adb20498d7e3e46721ccdae8166.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i70f37a3c126075fa8649885fae8b9269c2b414b14c355d5f56c6bacda4a2f642.ForwardRequestBuilder) {
    return i70f37a3c126075fa8649885fae8b9269c2b414b14c355d5f56c6bacda4a2f642.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Instances()(*i613860500ed7e5d220cc7aa468f340f68057e89254c0752dd1d5b23406d22c1c.InstancesRequestBuilder) {
    return i613860500ed7e5d220cc7aa468f340f68057e89254c0752dd1d5b23406d22c1c.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i2964545c5adf79646b25d3619484630d27e775d018d943673dcad63b7066cd28.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id2"] = id
    }
    return i2964545c5adf79646b25d3619484630d27e775d018d943673dcad63b7066cd28.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iaa0ad74adb35429933f3f3327c4b5922c01827bbda95187f603f9d4e9affa9d6.MultiValueExtendedPropertiesRequestBuilder) {
    return iaa0ad74adb35429933f3f3327c4b5922c01827bbda95187f603f9d4e9affa9d6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iaa4104cd8dc9111c099a2d6f38da15af0b5a77c014003598bb9647382389c86c.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iaa4104cd8dc9111c099a2d6f38da15af0b5a77c014003598bb9647382389c86c.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*if1131a08ca4381b46da8854282509db2688d2d9b02903846bb32439175f3662c.SingleValueExtendedPropertiesRequestBuilder) {
    return if1131a08ca4381b46da8854282509db2688d2d9b02903846bb32439175f3662c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i404080d5ce3bee4160c8d92d96ad2ce6f1bcc29995158139ca80da35ef3bc34f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i404080d5ce3bee4160c8d92d96ad2ce6f1bcc29995158139ca80da35ef3bc34f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*icf1905e5a6e127a34411d19fe37e4528b1694e36035d0deea29c8d1bd69d4a4f.SnoozeReminderRequestBuilder) {
    return icf1905e5a6e127a34411d19fe37e4528b1694e36035d0deea29c8d1bd69d4a4f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3fd6d2a1f574a9810dccae420fdd454a912df52f30df96f4a185d78cfd6f9c05.TentativelyAcceptRequestBuilder) {
    return i3fd6d2a1f574a9810dccae420fdd454a912df52f30df96f4a185d78cfd6f9c05.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

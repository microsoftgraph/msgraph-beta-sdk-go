package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i04547fc25f9a19b56375a31791a646bb9a7291300dab821083dcd9dff492eec5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i0cc9b39109a903e3c7493bc95907d6b99e23de8914847bebc5b2b7718aa48e85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    i1b93fb71e20e87d0c94ea6fbd6fb298fc08fb5d0e4a2e8f9c34137b47ee5e9d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    i2061516be4b73f307ef27998a5436c0d6263ff0f2dfe896d758d303753de3249 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    i3c0518909ebf9cd483b89bfb6b30484ba04b2be163059af0f1fb66f6749d7f13 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    iaaf46f7a2ead40b95e11dc338f68f4cc605581625ff09115cf3dfa212c70c55d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    iba1ffd0464b70bb45422cece7b13a6e10464f2457707c0bf58770857e0fccd33 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    ibbcd29199f49a1e2c6067cb3d21b38496815d0537ce16e70bbffe8a002c63ac3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    ibdcbb5181e62a6c318e7e080ec9929125febe7b6feb6ff10777c7a7626f2ccf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    ic8e9ae378ae88a52d0b547ce11c9cff604379b3f87d4a0da0c87d617167b5c54 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    ica549a59bb07a4b59d1250ec60c1c7f20cc7b2c95841d2af82d8b77c130ebf03 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    if9d4d521baacac2476f774dd755d2919846176c52baa2d6a33807091c4f57372 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    i64550b35c048770208de1ceba65993bad9d9150195e269fb6b146ca69ce8afc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    i9b67f74d1e1e9cf957ac05505c98e12f20babaf600449186355ccc1ef198c643 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    ib7178ba408e138ce3cdd93a42a6b5a070b0b033794f27a9ebb32ed3651994b60 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    id6bbec9cfa87470cc5313b0414204e4191e208f07a93b995e7462d3584214ebb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i1b93fb71e20e87d0c94ea6fbd6fb298fc08fb5d0e4a2e8f9c34137b47ee5e9d3.AcceptRequestBuilder) {
    return i1b93fb71e20e87d0c94ea6fbd6fb298fc08fb5d0e4a2e8f9c34137b47ee5e9d3.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*iba1ffd0464b70bb45422cece7b13a6e10464f2457707c0bf58770857e0fccd33.AttachmentsRequestBuilder) {
    return iba1ffd0464b70bb45422cece7b13a6e10464f2457707c0bf58770857e0fccd33.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i64550b35c048770208de1ceba65993bad9d9150195e269fb6b146ca69ce8afc7.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i64550b35c048770208de1ceba65993bad9d9150195e269fb6b146ca69ce8afc7.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i2061516be4b73f307ef27998a5436c0d6263ff0f2dfe896d758d303753de3249.CalendarRequestBuilder) {
    return i2061516be4b73f307ef27998a5436c0d6263ff0f2dfe896d758d303753de3249.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*ibdcbb5181e62a6c318e7e080ec9929125febe7b6feb6ff10777c7a7626f2ccf2.CancelRequestBuilder) {
    return ibdcbb5181e62a6c318e7e080ec9929125febe7b6feb6ff10777c7a7626f2ccf2.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}/calendarView/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
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
func (m *EventItemRequestBuilder) Decline()(*ica549a59bb07a4b59d1250ec60c1c7f20cc7b2c95841d2af82d8b77c130ebf03.DeclineRequestBuilder) {
    return ica549a59bb07a4b59d1250ec60c1c7f20cc7b2c95841d2af82d8b77c130ebf03.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i04547fc25f9a19b56375a31791a646bb9a7291300dab821083dcd9dff492eec5.DismissReminderRequestBuilder) {
    return i04547fc25f9a19b56375a31791a646bb9a7291300dab821083dcd9dff492eec5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*if9d4d521baacac2476f774dd755d2919846176c52baa2d6a33807091c4f57372.ExtensionsRequestBuilder) {
    return if9d4d521baacac2476f774dd755d2919846176c52baa2d6a33807091c4f57372.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9b67f74d1e1e9cf957ac05505c98e12f20babaf600449186355ccc1ef198c643.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i9b67f74d1e1e9cf957ac05505c98e12f20babaf600449186355ccc1ef198c643.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i0cc9b39109a903e3c7493bc95907d6b99e23de8914847bebc5b2b7718aa48e85.ForwardRequestBuilder) {
    return i0cc9b39109a903e3c7493bc95907d6b99e23de8914847bebc5b2b7718aa48e85.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic8e9ae378ae88a52d0b547ce11c9cff604379b3f87d4a0da0c87d617167b5c54.MultiValueExtendedPropertiesRequestBuilder) {
    return ic8e9ae378ae88a52d0b547ce11c9cff604379b3f87d4a0da0c87d617167b5c54.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib7178ba408e138ce3cdd93a42a6b5a070b0b033794f27a9ebb32ed3651994b60.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib7178ba408e138ce3cdd93a42a6b5a070b0b033794f27a9ebb32ed3651994b60.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i3c0518909ebf9cd483b89bfb6b30484ba04b2be163059af0f1fb66f6749d7f13.SingleValueExtendedPropertiesRequestBuilder) {
    return i3c0518909ebf9cd483b89bfb6b30484ba04b2be163059af0f1fb66f6749d7f13.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id6bbec9cfa87470cc5313b0414204e4191e208f07a93b995e7462d3584214ebb.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return id6bbec9cfa87470cc5313b0414204e4191e208f07a93b995e7462d3584214ebb.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*iaaf46f7a2ead40b95e11dc338f68f4cc605581625ff09115cf3dfa212c70c55d.SnoozeReminderRequestBuilder) {
    return iaaf46f7a2ead40b95e11dc338f68f4cc605581625ff09115cf3dfa212c70c55d.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibbcd29199f49a1e2c6067cb3d21b38496815d0537ce16e70bbffe8a002c63ac3.TentativelyAcceptRequestBuilder) {
    return ibbcd29199f49a1e2c6067cb3d21b38496815d0537ce16e70bbffe8a002c63ac3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

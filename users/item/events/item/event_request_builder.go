package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0ab944a433b01a980614ef06794522864fa910f39b21002e379b9f028d4ef2a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/extensions"
    i40c370fc1a9ee7a2ac8d30f9aa169f816be63cf6bb6459bac521a9d9be80072e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/calendar"
    i40ef9dbec34f6598bd5e64b074eb1fc738c4d6cfce9140c3875a60bb9e84d1ae "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences"
    i4d0b018ae4445247bb45375ac01fa18bc25be862015d0b7e8e61b61823000bb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/attachments"
    i524e360741674ea3a3843aeb78c0b77691aefc0965cff77654c66140620074b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/cancel"
    i6bd313e9caaa2477a34a0c6921300c24c8c818a1eb8a6c928fd1bdb89fc45c5b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/forward"
    i6ea56bd8b3c0842a7cf1539390d5b5246fc5bbdac93b7c7013f5526d92380fa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/decline"
    i70dc1b9e9bb4a1ea662f4589c936f3283d5737ebed09cdecf22904f981b14bf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/tentativelyaccept"
    i77a0682874bb4a37002c7e4f677064e4d44e24e3c7124a7acf8c5425bad2f6ab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/multivalueextendedproperties"
    i83e4bb983efc91a00ad33b55d31911173647d449f3f7bf7875dc2345f7d320e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/snoozereminder"
    i93843b5c69826696a18bb7d5621c6e98c4bfdc37c80dc70356c2d9c6e6ac8908 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/dismissreminder"
    iab097ad642bee606b7266ec0e626e72fd150acb2506817557369fe0cb5773e87 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/accept"
    iaceed976ce69c3a839e3e0dfb726138753535f432f9e42707e946bb3282ba0a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances"
    iec1b43a554ce0caa02e1ae756bf1ee3ff1eb08b4a250481ac8423b614f6af107 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/singlevalueextendedproperties"
    i3b6f4b55b6a0f7c3d67c1c2d0128cef539616d6fe442226352eb6f0541025401 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item"
    i680e72a8871bfb354a7d740cfe6334270689df7bd9c7773514f13b397b147673 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/extensions/item"
    i7bec747178558ca6942fd148663b71eecd9c1332bc0dda2fbabcce59ff4bf809 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/multivalueextendedproperties/item"
    ic2f36ebf4800cd886805d7095eacef0d70ab09aea70f9f65bfbffa7743ebeed3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/attachments/item"
    ic6a9a2d098cefb19cb3a8450d4ebad5f9826aa605e840b30d7fe0de5e82fce36 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/instances/item"
    ifdcda3bbafdd32c31e3d590c71e85f7bdbca6a9d3893208625188afb3ab40a08 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/singlevalueextendedproperties/item"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\events\{event-id}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventRequestBuilderDeleteOptions options for Delete
type EventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetOptions options for Get
type EventRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetQueryParameters the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
type EventRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventRequestBuilderPatchOptions options for Patch
type EventRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventRequestBuilder) Accept()(*iab097ad642bee606b7266ec0e626e72fd150acb2506817557369fe0cb5773e87.AcceptRequestBuilder) {
    return iab097ad642bee606b7266ec0e626e72fd150acb2506817557369fe0cb5773e87.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i4d0b018ae4445247bb45375ac01fa18bc25be862015d0b7e8e61b61823000bb7.AttachmentsRequestBuilder) {
    return i4d0b018ae4445247bb45375ac01fa18bc25be862015d0b7e8e61b61823000bb7.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ic2f36ebf4800cd886805d7095eacef0d70ab09aea70f9f65bfbffa7743ebeed3.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ic2f36ebf4800cd886805d7095eacef0d70ab09aea70f9f65bfbffa7743ebeed3.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i40c370fc1a9ee7a2ac8d30f9aa169f816be63cf6bb6459bac521a9d9be80072e.CalendarRequestBuilder) {
    return i40c370fc1a9ee7a2ac8d30f9aa169f816be63cf6bb6459bac521a9d9be80072e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i524e360741674ea3a3843aeb78c0b77691aefc0965cff77654c66140620074b1.CancelRequestBuilder) {
    return i524e360741674ea3a3843aeb78c0b77691aefc0965cff77654c66140620074b1.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/events/{event_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventRequestBuilder instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventRequestBuilder) CreateDeleteRequestInformation(options *EventRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventRequestBuilder) CreateGetRequestInformation(options *EventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventRequestBuilder) CreatePatchRequestInformation(options *EventRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*i6ea56bd8b3c0842a7cf1539390d5b5246fc5bbdac93b7c7013f5526d92380fa0.DeclineRequestBuilder) {
    return i6ea56bd8b3c0842a7cf1539390d5b5246fc5bbdac93b7c7013f5526d92380fa0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventRequestBuilder) Delete(options *EventRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) DismissReminder()(*i93843b5c69826696a18bb7d5621c6e98c4bfdc37c80dc70356c2d9c6e6ac8908.DismissReminderRequestBuilder) {
    return i93843b5c69826696a18bb7d5621c6e98c4bfdc37c80dc70356c2d9c6e6ac8908.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i40ef9dbec34f6598bd5e64b074eb1fc738c4d6cfce9140c3875a60bb9e84d1ae.ExceptionOccurrencesRequestBuilder) {
    return i40ef9dbec34f6598bd5e64b074eb1fc738c4d6cfce9140c3875a60bb9e84d1ae.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i3b6f4b55b6a0f7c3d67c1c2d0128cef539616d6fe442226352eb6f0541025401.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i3b6f4b55b6a0f7c3d67c1c2d0128cef539616d6fe442226352eb6f0541025401.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i0ab944a433b01a980614ef06794522864fa910f39b21002e379b9f028d4ef2a9.ExtensionsRequestBuilder) {
    return i0ab944a433b01a980614ef06794522864fa910f39b21002e379b9f028d4ef2a9.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i680e72a8871bfb354a7d740cfe6334270689df7bd9c7773514f13b397b147673.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i680e72a8871bfb354a7d740cfe6334270689df7bd9c7773514f13b397b147673.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i6bd313e9caaa2477a34a0c6921300c24c8c818a1eb8a6c928fd1bdb89fc45c5b.ForwardRequestBuilder) {
    return i6bd313e9caaa2477a34a0c6921300c24c8c818a1eb8a6c928fd1bdb89fc45c5b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventRequestBuilder) Get(options *EventRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event), nil
}
func (m *EventRequestBuilder) Instances()(*iaceed976ce69c3a839e3e0dfb726138753535f432f9e42707e946bb3282ba0a5.InstancesRequestBuilder) {
    return iaceed976ce69c3a839e3e0dfb726138753535f432f9e42707e946bb3282ba0a5.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*ic6a9a2d098cefb19cb3a8450d4ebad5f9826aa605e840b30d7fe0de5e82fce36.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ic6a9a2d098cefb19cb3a8450d4ebad5f9826aa605e840b30d7fe0de5e82fce36.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i77a0682874bb4a37002c7e4f677064e4d44e24e3c7124a7acf8c5425bad2f6ab.MultiValueExtendedPropertiesRequestBuilder) {
    return i77a0682874bb4a37002c7e4f677064e4d44e24e3c7124a7acf8c5425bad2f6ab.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7bec747178558ca6942fd148663b71eecd9c1332bc0dda2fbabcce59ff4bf809.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i7bec747178558ca6942fd148663b71eecd9c1332bc0dda2fbabcce59ff4bf809.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's events. Default is to show Events under the Default Calendar. Read-only. Nullable.
func (m *EventRequestBuilder) Patch(options *EventRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*iec1b43a554ce0caa02e1ae756bf1ee3ff1eb08b4a250481ac8423b614f6af107.SingleValueExtendedPropertiesRequestBuilder) {
    return iec1b43a554ce0caa02e1ae756bf1ee3ff1eb08b4a250481ac8423b614f6af107.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifdcda3bbafdd32c31e3d590c71e85f7bdbca6a9d3893208625188afb3ab40a08.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ifdcda3bbafdd32c31e3d590c71e85f7bdbca6a9d3893208625188afb3ab40a08.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i83e4bb983efc91a00ad33b55d31911173647d449f3f7bf7875dc2345f7d320e2.SnoozeReminderRequestBuilder) {
    return i83e4bb983efc91a00ad33b55d31911173647d449f3f7bf7875dc2345f7d320e2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i70dc1b9e9bb4a1ea662f4589c936f3283d5737ebed09cdecf22904f981b14bf7.TentativelyAcceptRequestBuilder) {
    return i70dc1b9e9bb4a1ea662f4589c936f3283d5737ebed09cdecf22904f981b14bf7.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

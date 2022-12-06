package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.user entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EventItemRequestBuilderGetQueryParameters the user's events. Default is to show events under the Default Calendar. Read-only. Nullable.
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
// EventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept provides operations to call the accept method.
func (m *EventItemRequestBuilder) Accept()(*iab097ad642bee606b7266ec0e626e72fd150acb2506817557369fe0cb5773e87.AcceptRequestBuilder) {
    return iab097ad642bee606b7266ec0e626e72fd150acb2506817557369fe0cb5773e87.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i4d0b018ae4445247bb45375ac01fa18bc25be862015d0b7e8e61b61823000bb7.AttachmentsRequestBuilder) {
    return i4d0b018ae4445247bb45375ac01fa18bc25be862015d0b7e8e61b61823000bb7.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ic2f36ebf4800cd886805d7095eacef0d70ab09aea70f9f65bfbffa7743ebeed3.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ic2f36ebf4800cd886805d7095eacef0d70ab09aea70f9f65bfbffa7743ebeed3.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i40c370fc1a9ee7a2ac8d30f9aa169f816be63cf6bb6459bac521a9d9be80072e.CalendarRequestBuilder) {
    return i40c370fc1a9ee7a2ac8d30f9aa169f816be63cf6bb6459bac521a9d9be80072e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*i524e360741674ea3a3843aeb78c0b77691aefc0965cff77654c66140620074b1.CancelRequestBuilder) {
    return i524e360741674ea3a3843aeb78c0b77691aefc0965cff77654c66140620074b1.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/events/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property events for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the user's events. Default is to show events under the Default Calendar. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property events in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline provides operations to call the decline method.
func (m *EventItemRequestBuilder) Decline()(*i6ea56bd8b3c0842a7cf1539390d5b5246fc5bbdac93b7c7013f5526d92380fa0.DeclineRequestBuilder) {
    return i6ea56bd8b3c0842a7cf1539390d5b5246fc5bbdac93b7c7013f5526d92380fa0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for users
func (m *EventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i93843b5c69826696a18bb7d5621c6e98c4bfdc37c80dc70356c2d9c6e6ac8908.DismissReminderRequestBuilder) {
    return i93843b5c69826696a18bb7d5621c6e98c4bfdc37c80dc70356c2d9c6e6ac8908.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i40ef9dbec34f6598bd5e64b074eb1fc738c4d6cfce9140c3875a60bb9e84d1ae.ExceptionOccurrencesRequestBuilder) {
    return i40ef9dbec34f6598bd5e64b074eb1fc738c4d6cfce9140c3875a60bb9e84d1ae.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i3b6f4b55b6a0f7c3d67c1c2d0128cef539616d6fe442226352eb6f0541025401.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i3b6f4b55b6a0f7c3d67c1c2d0128cef539616d6fe442226352eb6f0541025401.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i0ab944a433b01a980614ef06794522864fa910f39b21002e379b9f028d4ef2a9.ExtensionsRequestBuilder) {
    return i0ab944a433b01a980614ef06794522864fa910f39b21002e379b9f028d4ef2a9.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i680e72a8871bfb354a7d740cfe6334270689df7bd9c7773514f13b397b147673.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i680e72a8871bfb354a7d740cfe6334270689df7bd9c7773514f13b397b147673.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i6bd313e9caaa2477a34a0c6921300c24c8c818a1eb8a6c928fd1bdb89fc45c5b.ForwardRequestBuilder) {
    return i6bd313e9caaa2477a34a0c6921300c24c8c818a1eb8a6c928fd1bdb89fc45c5b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user's events. Default is to show events under the Default Calendar. Read-only. Nullable.
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
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Instances()(*iaceed976ce69c3a839e3e0dfb726138753535f432f9e42707e946bb3282ba0a5.InstancesRequestBuilder) {
    return iaceed976ce69c3a839e3e0dfb726138753535f432f9e42707e946bb3282ba0a5.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) InstancesById(id string)(*ic6a9a2d098cefb19cb3a8450d4ebad5f9826aa605e840b30d7fe0de5e82fce36.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return ic6a9a2d098cefb19cb3a8450d4ebad5f9826aa605e840b30d7fe0de5e82fce36.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i77a0682874bb4a37002c7e4f677064e4d44e24e3c7124a7acf8c5425bad2f6ab.MultiValueExtendedPropertiesRequestBuilder) {
    return i77a0682874bb4a37002c7e4f677064e4d44e24e3c7124a7acf8c5425bad2f6ab.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7bec747178558ca6942fd148663b71eecd9c1332bc0dda2fbabcce59ff4bf809.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i7bec747178558ca6942fd148663b71eecd9c1332bc0dda2fbabcce59ff4bf809.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in users
func (m *EventItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*iec1b43a554ce0caa02e1ae756bf1ee3ff1eb08b4a250481ac8423b614f6af107.SingleValueExtendedPropertiesRequestBuilder) {
    return iec1b43a554ce0caa02e1ae756bf1ee3ff1eb08b4a250481ac8423b614f6af107.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifdcda3bbafdd32c31e3d590c71e85f7bdbca6a9d3893208625188afb3ab40a08.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ifdcda3bbafdd32c31e3d590c71e85f7bdbca6a9d3893208625188afb3ab40a08.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i83e4bb983efc91a00ad33b55d31911173647d449f3f7bf7875dc2345f7d320e2.SnoozeReminderRequestBuilder) {
    return i83e4bb983efc91a00ad33b55d31911173647d449f3f7bf7875dc2345f7d320e2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i70dc1b9e9bb4a1ea662f4589c936f3283d5737ebed09cdecf22904f981b14bf7.TentativelyAcceptRequestBuilder) {
    return i70dc1b9e9bb4a1ea662f4589c936f3283d5737ebed09cdecf22904f981b14bf7.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

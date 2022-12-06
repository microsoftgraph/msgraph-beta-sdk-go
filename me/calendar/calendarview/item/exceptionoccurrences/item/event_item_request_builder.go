package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i37e2aac2ec2d8c105189ca2544e370e7fc0958f16a63ef7de65a5f6363b206db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/attachments"
    i5a0c16c3dd020cab0d66bb0e5457ffe049a18b1655e51c79b1febbcf80a7fa9f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i5d50c89c24e48235b3bc5d9be555672bf2919de218fc1437404f57b8ebb727b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/calendar"
    i67974d6836dc46285b14edeb89a6cc14ccd949a6fda096c0adcadf3a3a15709a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i751db39f334dace98a9ac095cf5bc28353e5caaaac494c73c2bcb1b383091028 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/accept"
    i7b24fd421c3a73c7375fff8de767f5fc8cf4fa8c1d28e602d4390bf7ffc8ba00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/cancel"
    i871fd36829eb2d510d8627e6bd03349fa5e1585a675bde4cd8312dea1bd43014 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/dismissreminder"
    i8892cc3c5a2b94aaed280c21a486a17a49a084b6e6940a07d12cb57a6b56e12c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i8a38ea508feccfcfcf3c7bfac456b88c116947fb3c571d820ea465aad727913b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances"
    ia62999cf5bdd3039a2f3126c03534c97c92913a0042ec0650874a0be8cd118e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    icd9b4667a22080225ade5263bbe7395f8ef95e8b8fcddb409a37a2be3dbbc40f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/decline"
    iea9143168d67bb3f6044912351b9db33c5f3a8394890d049ec2b5f502e2ceac3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/extensions"
    if0c39aa7dd56f7c86762937d8ad5836cec244f2c81e0d3030ce5d03ea91584d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/forward"
    i41381b35a2f4715a5c613010acbd06e17b7a91f09eabde1664b733013c5c0675 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/extensions/item"
    i55eaabdba1cd872cabd6649de30f2e194b91f5b6660b923367186e4b5f4462e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i5dd315f30a93cfa6800379001ffdae29d17cb7a9867a6e913bc7db4d3d36a2a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item"
    ib20bfdf08ca95540961a4fa084926c8d262e903dce1f9c87549a640f5545d8ee "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/attachments/item"
    icf624d7f78f2c83a7ce505c2bb775394615d59f66b35d57871bf1425b0dff3b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
// Accept provides operations to call the accept method.
func (m *EventItemRequestBuilder) Accept()(*i751db39f334dace98a9ac095cf5bc28353e5caaaac494c73c2bcb1b383091028.AcceptRequestBuilder) {
    return i751db39f334dace98a9ac095cf5bc28353e5caaaac494c73c2bcb1b383091028.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i37e2aac2ec2d8c105189ca2544e370e7fc0958f16a63ef7de65a5f6363b206db.AttachmentsRequestBuilder) {
    return i37e2aac2ec2d8c105189ca2544e370e7fc0958f16a63ef7de65a5f6363b206db.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib20bfdf08ca95540961a4fa084926c8d262e903dce1f9c87549a640f5545d8ee.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ib20bfdf08ca95540961a4fa084926c8d262e903dce1f9c87549a640f5545d8ee.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i5d50c89c24e48235b3bc5d9be555672bf2919de218fc1437404f57b8ebb727b9.CalendarRequestBuilder) {
    return i5d50c89c24e48235b3bc5d9be555672bf2919de218fc1437404f57b8ebb727b9.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*i7b24fd421c3a73c7375fff8de767f5fc8cf4fa8c1d28e602d4390bf7ffc8ba00.CancelRequestBuilder) {
    return i7b24fd421c3a73c7375fff8de767f5fc8cf4fa8c1d28e602d4390bf7ffc8ba00.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
// CreateGetRequestInformation get exceptionOccurrences from me
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
// Decline provides operations to call the decline method.
func (m *EventItemRequestBuilder) Decline()(*icd9b4667a22080225ade5263bbe7395f8ef95e8b8fcddb409a37a2be3dbbc40f.DeclineRequestBuilder) {
    return icd9b4667a22080225ade5263bbe7395f8ef95e8b8fcddb409a37a2be3dbbc40f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i871fd36829eb2d510d8627e6bd03349fa5e1585a675bde4cd8312dea1bd43014.DismissReminderRequestBuilder) {
    return i871fd36829eb2d510d8627e6bd03349fa5e1585a675bde4cd8312dea1bd43014.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*iea9143168d67bb3f6044912351b9db33c5f3a8394890d049ec2b5f502e2ceac3.ExtensionsRequestBuilder) {
    return iea9143168d67bb3f6044912351b9db33c5f3a8394890d049ec2b5f502e2ceac3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i41381b35a2f4715a5c613010acbd06e17b7a91f09eabde1664b733013c5c0675.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i41381b35a2f4715a5c613010acbd06e17b7a91f09eabde1664b733013c5c0675.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*if0c39aa7dd56f7c86762937d8ad5836cec244f2c81e0d3030ce5d03ea91584d7.ForwardRequestBuilder) {
    return if0c39aa7dd56f7c86762937d8ad5836cec244f2c81e0d3030ce5d03ea91584d7.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Instances()(*i8a38ea508feccfcfcf3c7bfac456b88c116947fb3c571d820ea465aad727913b.InstancesRequestBuilder) {
    return i8a38ea508feccfcfcf3c7bfac456b88c116947fb3c571d820ea465aad727913b.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) InstancesById(id string)(*i5dd315f30a93cfa6800379001ffdae29d17cb7a9867a6e913bc7db4d3d36a2a8.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i5dd315f30a93cfa6800379001ffdae29d17cb7a9867a6e913bc7db4d3d36a2a8.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ia62999cf5bdd3039a2f3126c03534c97c92913a0042ec0650874a0be8cd118e5.MultiValueExtendedPropertiesRequestBuilder) {
    return ia62999cf5bdd3039a2f3126c03534c97c92913a0042ec0650874a0be8cd118e5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i55eaabdba1cd872cabd6649de30f2e194b91f5b6660b923367186e4b5f4462e5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i55eaabdba1cd872cabd6649de30f2e194b91f5b6660b923367186e4b5f4462e5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i8892cc3c5a2b94aaed280c21a486a17a49a084b6e6940a07d12cb57a6b56e12c.SingleValueExtendedPropertiesRequestBuilder) {
    return i8892cc3c5a2b94aaed280c21a486a17a49a084b6e6940a07d12cb57a6b56e12c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*icf624d7f78f2c83a7ce505c2bb775394615d59f66b35d57871bf1425b0dff3b6.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return icf624d7f78f2c83a7ce505c2bb775394615d59f66b35d57871bf1425b0dff3b6.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i5a0c16c3dd020cab0d66bb0e5457ffe049a18b1655e51c79b1febbcf80a7fa9f.SnoozeReminderRequestBuilder) {
    return i5a0c16c3dd020cab0d66bb0e5457ffe049a18b1655e51c79b1febbcf80a7fa9f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i67974d6836dc46285b14edeb89a6cc14ccd949a6fda096c0adcadf3a3a15709a.TentativelyAcceptRequestBuilder) {
    return i67974d6836dc46285b14edeb89a6cc14ccd949a6fda096c0adcadf3a3a15709a.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

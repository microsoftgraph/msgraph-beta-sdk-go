package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i123a73c0a760ecc25a7b1ca61a9a699251ba6fc529bc128b778d54b88f9397b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i133164ac21d2f57dfbbd5a936164a3a5664ee6979a0000f97e432c272ec7587a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/calendar"
    i691b21a2c6b0f7b7c63547de1feb2b796f0d068520df55438d061c80252b38ab "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/accept"
    i716e8e4608791f76528c4904d850dfc9f1785ac89f45ad89ce12b25a3cde8ce4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/attachments"
    i72d5efdd4b81236842f9e42f9f1a19d811bf7a2ff3f5ffdbf527d1fb1102241e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i786d0ae7226986f2f6f85a2036f206054aa568cc0107e444d6bf6ad47912217a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/cancel"
    i87df222fedae8d0f4073c9038d792110f7dd63b37fb4953580ff1473bef550c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    icaefefe70b727e216ce2655a680726e0fe5f42db6cba80d1dffd4bf8f1c58cf1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    id352bb6dd2b2c6d5f51fac6a2baec56f92db4840599c5ccd75f167f4385286b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/extensions"
    ie66a7bbf092517952e59aa40bf043a8656b11b97c8f400dc8c7b832c879f6268 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/forward"
    if71fc266cb9b6edd9fd1a04f08bfff770834966f6d0b04bf516e465c6b391869 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/decline"
    ifa8eb554da16d23e17ac8a851efbe08a33b655674ef2213dfe30db6e9b50d448 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i42bdf21a5ac363d91c4ac0c4c0f51534c2eb8ed64be8d020e64a37f9f40e6e9a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i7a1640dcaf30d6e5028a3777cea0921953c2bbf85fabc1522df0ec57e70a06a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    ia73d52ba0b8eab5d650e420014d87ffe5f9c18a2b6c8a2987fe7768925382078 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    ic622fe1ef834c61e5e405144d06693dfee887e43dcb64f73402519b96876c70d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i691b21a2c6b0f7b7c63547de1feb2b796f0d068520df55438d061c80252b38ab.AcceptRequestBuilder) {
    return i691b21a2c6b0f7b7c63547de1feb2b796f0d068520df55438d061c80252b38ab.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i716e8e4608791f76528c4904d850dfc9f1785ac89f45ad89ce12b25a3cde8ce4.AttachmentsRequestBuilder) {
    return i716e8e4608791f76528c4904d850dfc9f1785ac89f45ad89ce12b25a3cde8ce4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia73d52ba0b8eab5d650e420014d87ffe5f9c18a2b6c8a2987fe7768925382078.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ia73d52ba0b8eab5d650e420014d87ffe5f9c18a2b6c8a2987fe7768925382078.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i133164ac21d2f57dfbbd5a936164a3a5664ee6979a0000f97e432c272ec7587a.CalendarRequestBuilder) {
    return i133164ac21d2f57dfbbd5a936164a3a5664ee6979a0000f97e432c272ec7587a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i786d0ae7226986f2f6f85a2036f206054aa568cc0107e444d6bf6ad47912217a.CancelRequestBuilder) {
    return i786d0ae7226986f2f6f85a2036f206054aa568cc0107e444d6bf6ad47912217a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get exceptionOccurrences from groups
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*if71fc266cb9b6edd9fd1a04f08bfff770834966f6d0b04bf516e465c6b391869.DeclineRequestBuilder) {
    return if71fc266cb9b6edd9fd1a04f08bfff770834966f6d0b04bf516e465c6b391869.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for groups
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*icaefefe70b727e216ce2655a680726e0fe5f42db6cba80d1dffd4bf8f1c58cf1.DismissReminderRequestBuilder) {
    return icaefefe70b727e216ce2655a680726e0fe5f42db6cba80d1dffd4bf8f1c58cf1.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*id352bb6dd2b2c6d5f51fac6a2baec56f92db4840599c5ccd75f167f4385286b2.ExtensionsRequestBuilder) {
    return id352bb6dd2b2c6d5f51fac6a2baec56f92db4840599c5ccd75f167f4385286b2.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i7a1640dcaf30d6e5028a3777cea0921953c2bbf85fabc1522df0ec57e70a06a1.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i7a1640dcaf30d6e5028a3777cea0921953c2bbf85fabc1522df0ec57e70a06a1.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ie66a7bbf092517952e59aa40bf043a8656b11b97c8f400dc8c7b832c879f6268.ForwardRequestBuilder) {
    return ie66a7bbf092517952e59aa40bf043a8656b11b97c8f400dc8c7b832c879f6268.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i87df222fedae8d0f4073c9038d792110f7dd63b37fb4953580ff1473bef550c2.MultiValueExtendedPropertiesRequestBuilder) {
    return i87df222fedae8d0f4073c9038d792110f7dd63b37fb4953580ff1473bef550c2.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic622fe1ef834c61e5e405144d06693dfee887e43dcb64f73402519b96876c70d.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ic622fe1ef834c61e5e405144d06693dfee887e43dcb64f73402519b96876c70d.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in groups
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i123a73c0a760ecc25a7b1ca61a9a699251ba6fc529bc128b778d54b88f9397b3.SingleValueExtendedPropertiesRequestBuilder) {
    return i123a73c0a760ecc25a7b1ca61a9a699251ba6fc529bc128b778d54b88f9397b3.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i42bdf21a5ac363d91c4ac0c4c0f51534c2eb8ed64be8d020e64a37f9f40e6e9a.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i42bdf21a5ac363d91c4ac0c4c0f51534c2eb8ed64be8d020e64a37f9f40e6e9a.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i72d5efdd4b81236842f9e42f9f1a19d811bf7a2ff3f5ffdbf527d1fb1102241e.SnoozeReminderRequestBuilder) {
    return i72d5efdd4b81236842f9e42f9f1a19d811bf7a2ff3f5ffdbf527d1fb1102241e.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ifa8eb554da16d23e17ac8a851efbe08a33b655674ef2213dfe30db6e9b50d448.TentativelyAcceptRequestBuilder) {
    return ifa8eb554da16d23e17ac8a851efbe08a33b655674ef2213dfe30db6e9b50d448.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

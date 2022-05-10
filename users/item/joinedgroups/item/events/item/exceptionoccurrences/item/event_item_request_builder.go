package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i059c854347fecc394e94a42319d56bd62fd44fb3f3f6484c5c3a1bc932a250d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/cancel"
    i2d46b0cb97e0d91364aaeba5702c90ca725e7cf2b3cdc65df891823c71b5f4f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/dismissreminder"
    i2ea4f8742f570efb9d1d8824977d7bcc1c579e58a83829d8bdf0d02c9deb0b6a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/decline"
    i467a3aa52d76f3fa7144de39c22d8a16731b06ac425e3f7b91efaa9a076cd99e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/attachments"
    i5954268c4b5bd889386c970755c154acb55edc10cc1f58312c622b9d15329173 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/tentativelyaccept"
    i6994b4f0cc9c62a1720cb7bfbbf2217e59c08fbbb44f1440b655f4694890aada "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/calendar"
    i6c6f5f623b19e6fd756050150bbf19c6ae56c4188b232a70051b2a978ea5c18d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/extensions"
    i7f4f932996b3d1bd19e531679cde389dc8583ec2a919946186d0ca11e7561797 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/forward"
    ia59525281800669f154e1ca89f7730c54391fc80ac4c75a4009e966c7177acad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/accept"
    ibe4eea52da1e12f4e1e97acde86491498af335157b9b453a63b100077457de89 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/snoozereminder"
    ie5a70941d98729e86afa508a623c4c051e8eab19db9f03fd6f2e697a6f7528a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ieea2bbfd462325949474e2ff74906e43d2d888cb8d275d14263911ad3d6396d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances"
    ifab68c88ec30a44942d7c30dba807b443a91a6400db45385964fb07a1b58dc76 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    i0342ef3b3bd40e96a047379d7409325a2b5a9c8212e527d27db0f47d4e0312fe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/attachments/item"
    i09f0dc12d9ea7195d55b551f6e98fc79cdccabb693c69f600af8292f65255d30 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i1de57ddd3e0d53cb05dca29df23129977faad19a406d96864c69efab73f7c874 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/extensions/item"
    i3a94527f4d703baf0572b895616ba4276057054ad700b3296c61de3cb408388c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ic274bcedecc04e074f9a74e41a3b5173f10e6a361e49db70f82f62a22bcd218a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item"
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
// EventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
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
// EventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*ia59525281800669f154e1ca89f7730c54391fc80ac4c75a4009e966c7177acad.AcceptRequestBuilder) {
    return ia59525281800669f154e1ca89f7730c54391fc80ac4c75a4009e966c7177acad.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i467a3aa52d76f3fa7144de39c22d8a16731b06ac425e3f7b91efaa9a076cd99e.AttachmentsRequestBuilder) {
    return i467a3aa52d76f3fa7144de39c22d8a16731b06ac425e3f7b91efaa9a076cd99e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i0342ef3b3bd40e96a047379d7409325a2b5a9c8212e527d27db0f47d4e0312fe.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i0342ef3b3bd40e96a047379d7409325a2b5a9c8212e527d27db0f47d4e0312fe.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i6994b4f0cc9c62a1720cb7bfbbf2217e59c08fbbb44f1440b655f4694890aada.CalendarRequestBuilder) {
    return i6994b4f0cc9c62a1720cb7bfbbf2217e59c08fbbb44f1440b655f4694890aada.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i059c854347fecc394e94a42319d56bd62fd44fb3f3f6484c5c3a1bc932a250d4.CancelRequestBuilder) {
    return i059c854347fecc394e94a42319d56bd62fd44fb3f3f6484c5c3a1bc932a250d4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get exceptionOccurrences from users
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from users
func (m *EventItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i2ea4f8742f570efb9d1d8824977d7bcc1c579e58a83829d8bdf0d02c9deb0b6a.DeclineRequestBuilder) {
    return i2ea4f8742f570efb9d1d8824977d7bcc1c579e58a83829d8bdf0d02c9deb0b6a.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i2d46b0cb97e0d91364aaeba5702c90ca725e7cf2b3cdc65df891823c71b5f4f3.DismissReminderRequestBuilder) {
    return i2d46b0cb97e0d91364aaeba5702c90ca725e7cf2b3cdc65df891823c71b5f4f3.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i6c6f5f623b19e6fd756050150bbf19c6ae56c4188b232a70051b2a978ea5c18d.ExtensionsRequestBuilder) {
    return i6c6f5f623b19e6fd756050150bbf19c6ae56c4188b232a70051b2a978ea5c18d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i1de57ddd3e0d53cb05dca29df23129977faad19a406d96864c69efab73f7c874.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i1de57ddd3e0d53cb05dca29df23129977faad19a406d96864c69efab73f7c874.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i7f4f932996b3d1bd19e531679cde389dc8583ec2a919946186d0ca11e7561797.ForwardRequestBuilder) {
    return i7f4f932996b3d1bd19e531679cde389dc8583ec2a919946186d0ca11e7561797.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get exceptionOccurrences from users
func (m *EventItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*ieea2bbfd462325949474e2ff74906e43d2d888cb8d275d14263911ad3d6396d2.InstancesRequestBuilder) {
    return ieea2bbfd462325949474e2ff74906e43d2d888cb8d275d14263911ad3d6396d2.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ic274bcedecc04e074f9a74e41a3b5173f10e6a361e49db70f82f62a22bcd218a.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return ic274bcedecc04e074f9a74e41a3b5173f10e6a361e49db70f82f62a22bcd218a.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ifab68c88ec30a44942d7c30dba807b443a91a6400db45385964fb07a1b58dc76.MultiValueExtendedPropertiesRequestBuilder) {
    return ifab68c88ec30a44942d7c30dba807b443a91a6400db45385964fb07a1b58dc76.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i3a94527f4d703baf0572b895616ba4276057054ad700b3296c61de3cb408388c.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i3a94527f4d703baf0572b895616ba4276057054ad700b3296c61de3cb408388c.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ie5a70941d98729e86afa508a623c4c051e8eab19db9f03fd6f2e697a6f7528a1.SingleValueExtendedPropertiesRequestBuilder) {
    return ie5a70941d98729e86afa508a623c4c051e8eab19db9f03fd6f2e697a6f7528a1.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i09f0dc12d9ea7195d55b551f6e98fc79cdccabb693c69f600af8292f65255d30.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i09f0dc12d9ea7195d55b551f6e98fc79cdccabb693c69f600af8292f65255d30.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ibe4eea52da1e12f4e1e97acde86491498af335157b9b453a63b100077457de89.SnoozeReminderRequestBuilder) {
    return ibe4eea52da1e12f4e1e97acde86491498af335157b9b453a63b100077457de89.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i5954268c4b5bd889386c970755c154acb55edc10cc1f58312c622b9d15329173.TentativelyAcceptRequestBuilder) {
    return i5954268c4b5bd889386c970755c154acb55edc10cc1f58312c622b9d15329173.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

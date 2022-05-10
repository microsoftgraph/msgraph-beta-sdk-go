package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i24592c823a3778a932eca7aab6573ec12cc93c9acb82875ef40c473cd3ae9707 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/singlevalueextendedproperties"
    i2592ec1ea397d7c23d4df2230686d48c78ae77d36c157865cc10f0a48c394611 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/accept"
    i2683d7c5f69cb719666c8706ef773c5150e9ddcfcf0cfcee3abf1a12983f6f4b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/tentativelyaccept"
    i2760b16dc78b5f146f5fa955761c2ec480765a3beab6c8bbf1d74231eacc58bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/decline"
    i621dc0c449efd94f3c7730b0e70df6b5414665c624f454bba8699fac4c113350 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/calendar"
    i6ad67a6ede111d32e87b5624cb1901d4f5c4462aea19949d51e071de416890fe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances"
    i70e6ae2ca4459ca154ea3d3c32a5929fe61f7adbcbf06d0cff1fdad042eb27ed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/forward"
    i742887614f40fe47b60b57533928802ff1695186f8f025a1bed53de12ac54fcc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/extensions"
    i7b67463d6bcaad8c5521cc75162fd4d51f66ffce0cd825f009cc16bb94184c56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/snoozereminder"
    iac7fa72dcbe2041024fc70d740345923864485e9e749e56a7ebc02fd9dfd2a84 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/attachments"
    ib1f8171871dddf44148799ac785fcade6025f6b8e90f548220c468c98c0460f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/cancel"
    ib29683d43f27c9486997cd27678e4b01a82b75812ca5657ba17862bb47ccb8f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/multivalueextendedproperties"
    ie7d1545c4815dd320ff14f9b3caf7257105c6e07e7781edf067ea10c97b7f7c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences"
    ied1891a14240b96bb8df1830942597f8077d37ae27374cda6f463cb96f282462 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/dismissreminder"
    i05207a29e50b1b87173f3e9edab789be34eb3622b4b494ee841fd942c42504f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/singlevalueextendedproperties/item"
    i0996d70461c6e1ca1898322eeda52d17b3163ff9fc9b7d5a2ff3e34be4b3b5ec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/exceptionoccurrences/item"
    i22015e9546a34674b69ca595d86ba0db5db91994a2bfc1ca4fc454181565a8f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/attachments/item"
    i31b5127fa28c8e19016e1b464a69e683a5d0faba104a4f45731ef020d0b8b721 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/multivalueextendedproperties/item"
    i8a49b42e65617ed1e66173d1db86b4fbbb936508f21f4b00ffe620b2ae0cd5cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item"
    i9b8aa411f87289d031dcccd2a1a876db160f4f144e0fd08564ea16b2c0d3f017 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.group entity.
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
// EventItemRequestBuilderGetQueryParameters the group's events.
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
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i2592ec1ea397d7c23d4df2230686d48c78ae77d36c157865cc10f0a48c394611.AcceptRequestBuilder) {
    return i2592ec1ea397d7c23d4df2230686d48c78ae77d36c157865cc10f0a48c394611.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*iac7fa72dcbe2041024fc70d740345923864485e9e749e56a7ebc02fd9dfd2a84.AttachmentsRequestBuilder) {
    return iac7fa72dcbe2041024fc70d740345923864485e9e749e56a7ebc02fd9dfd2a84.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i22015e9546a34674b69ca595d86ba0db5db91994a2bfc1ca4fc454181565a8f4.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i22015e9546a34674b69ca595d86ba0db5db91994a2bfc1ca4fc454181565a8f4.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i621dc0c449efd94f3c7730b0e70df6b5414665c624f454bba8699fac4c113350.CalendarRequestBuilder) {
    return i621dc0c449efd94f3c7730b0e70df6b5414665c624f454bba8699fac4c113350.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ib1f8171871dddf44148799ac785fcade6025f6b8e90f548220c468c98c0460f5.CancelRequestBuilder) {
    return ib1f8171871dddf44148799ac785fcade6025f6b8e90f548220c468c98c0460f5.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/events/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property events for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property events for me
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
// CreateGetRequestInformation the group's events.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the group's events.
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
// CreatePatchRequestInformation update the navigation property events in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property events in me
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
func (m *EventItemRequestBuilder) Decline()(*i2760b16dc78b5f146f5fa955761c2ec480765a3beab6c8bbf1d74231eacc58bd.DeclineRequestBuilder) {
    return i2760b16dc78b5f146f5fa955761c2ec480765a3beab6c8bbf1d74231eacc58bd.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property events for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*ied1891a14240b96bb8df1830942597f8077d37ae27374cda6f463cb96f282462.DismissReminderRequestBuilder) {
    return ied1891a14240b96bb8df1830942597f8077d37ae27374cda6f463cb96f282462.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ie7d1545c4815dd320ff14f9b3caf7257105c6e07e7781edf067ea10c97b7f7c4.ExceptionOccurrencesRequestBuilder) {
    return ie7d1545c4815dd320ff14f9b3caf7257105c6e07e7781edf067ea10c97b7f7c4.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i0996d70461c6e1ca1898322eeda52d17b3163ff9fc9b7d5a2ff3e34be4b3b5ec.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i0996d70461c6e1ca1898322eeda52d17b3163ff9fc9b7d5a2ff3e34be4b3b5ec.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i742887614f40fe47b60b57533928802ff1695186f8f025a1bed53de12ac54fcc.ExtensionsRequestBuilder) {
    return i742887614f40fe47b60b57533928802ff1695186f8f025a1bed53de12ac54fcc.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9b8aa411f87289d031dcccd2a1a876db160f4f144e0fd08564ea16b2c0d3f017.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9b8aa411f87289d031dcccd2a1a876db160f4f144e0fd08564ea16b2c0d3f017.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i70e6ae2ca4459ca154ea3d3c32a5929fe61f7adbcbf06d0cff1fdad042eb27ed.ForwardRequestBuilder) {
    return i70e6ae2ca4459ca154ea3d3c32a5929fe61f7adbcbf06d0cff1fdad042eb27ed.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the group's events.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the group's events.
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
func (m *EventItemRequestBuilder) Instances()(*i6ad67a6ede111d32e87b5624cb1901d4f5c4462aea19949d51e071de416890fe.InstancesRequestBuilder) {
    return i6ad67a6ede111d32e87b5624cb1901d4f5c4462aea19949d51e071de416890fe.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i8a49b42e65617ed1e66173d1db86b4fbbb936508f21f4b00ffe620b2ae0cd5cd.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i8a49b42e65617ed1e66173d1db86b4fbbb936508f21f4b00ffe620b2ae0cd5cd.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ib29683d43f27c9486997cd27678e4b01a82b75812ca5657ba17862bb47ccb8f7.MultiValueExtendedPropertiesRequestBuilder) {
    return ib29683d43f27c9486997cd27678e4b01a82b75812ca5657ba17862bb47ccb8f7.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i31b5127fa28c8e19016e1b464a69e683a5d0faba104a4f45731ef020d0b8b721.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i31b5127fa28c8e19016e1b464a69e683a5d0faba104a4f45731ef020d0b8b721.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property events in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i24592c823a3778a932eca7aab6573ec12cc93c9acb82875ef40c473cd3ae9707.SingleValueExtendedPropertiesRequestBuilder) {
    return i24592c823a3778a932eca7aab6573ec12cc93c9acb82875ef40c473cd3ae9707.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i05207a29e50b1b87173f3e9edab789be34eb3622b4b494ee841fd942c42504f6.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i05207a29e50b1b87173f3e9edab789be34eb3622b4b494ee841fd942c42504f6.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i7b67463d6bcaad8c5521cc75162fd4d51f66ffce0cd825f009cc16bb94184c56.SnoozeReminderRequestBuilder) {
    return i7b67463d6bcaad8c5521cc75162fd4d51f66ffce0cd825f009cc16bb94184c56.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i2683d7c5f69cb719666c8706ef773c5150e9ddcfcf0cfcee3abf1a12983f6f4b.TentativelyAcceptRequestBuilder) {
    return i2683d7c5f69cb719666c8706ef773c5150e9ddcfcf0cfcee3abf1a12983f6f4b.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

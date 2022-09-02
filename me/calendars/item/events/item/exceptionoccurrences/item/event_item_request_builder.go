package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2e35467dedb09abc3f789b9492c7a7345aa1bbc47271131cd72f0c1cf32416ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/calendar"
    i37d29e90b678b9968642fb73a5ac7d10f486d6d16a90be81949d96897dd78f57 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/cancel"
    i5e9f805037cadcad48f7c7e245ac583a7f26272356cfc766d5e16ce1753b6e0d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/extensions"
    i5fe2d5d4221d7f70d4c72668fdc59fe7ae4218ec27b4d1109c27533327b7f55f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/accept"
    i7398f913d0b1d7f89fb886ab526b0bff45af288b1a77cb667228377dfb3b5654 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances"
    i73a558d6c49919e4e6bc753489cc2377b8bd6ae90a8dc5cbf62363115c1c8e5a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/attachments"
    i7daf4fdd37c9751a56b2863942fe2dbedd34f68a4efb078f8e3f3969ee8006e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/snoozereminder"
    i856c909b65889957a2bedb9bb98d91a58847ed04666e81dcab98806c27f8efe7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i87e1561c37eab795bb4810566f3d99bc0a26ea5061e38e1f2389122dcb930165 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/tentativelyaccept"
    ia1d675ace72efc2ee6d06f7972e11380d8529e248041c9740277d5dce2c676e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/forward"
    ia9d65d6431c43dcbe5294d975681cddbebc4df1f0474890119ea2be5a35c004d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/dismissreminder"
    ibb23d60c65c4b2cd40bb61e164329196f4bea1c483e786a13500f523b78caebc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    ie71f9c5564cd7668e9932138beadb6e8287d61e8d4ff49844167b6887a3e836d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/decline"
    i4ab94683bca463f7130220aa3e406baffefc0889228f548c19bbe36afbb73015 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/extensions/item"
    i55ede7043af2f7be625ab18c3dfbf56db2db1378c2d47b16a417a64c7b0d7bd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i5b6ccf20f21e40a5587355d022d743875bde8c8be78010cdb34a488eda8bb700 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i945a6029ddd8b34dbe08094d0c9a822be4c0659cfb31cff961c8890677afb696 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/instances/item"
    ib91584c821a775a386dfa5fbb8202d58f17569ef606c6af6e9b0bf781ec73431 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item/attachments/item"
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
// EventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i5fe2d5d4221d7f70d4c72668fdc59fe7ae4218ec27b4d1109c27533327b7f55f.AcceptRequestBuilder) {
    return i5fe2d5d4221d7f70d4c72668fdc59fe7ae4218ec27b4d1109c27533327b7f55f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i73a558d6c49919e4e6bc753489cc2377b8bd6ae90a8dc5cbf62363115c1c8e5a.AttachmentsRequestBuilder) {
    return i73a558d6c49919e4e6bc753489cc2377b8bd6ae90a8dc5cbf62363115c1c8e5a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib91584c821a775a386dfa5fbb8202d58f17569ef606c6af6e9b0bf781ec73431.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ib91584c821a775a386dfa5fbb8202d58f17569ef606c6af6e9b0bf781ec73431.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i2e35467dedb09abc3f789b9492c7a7345aa1bbc47271131cd72f0c1cf32416ba.CalendarRequestBuilder) {
    return i2e35467dedb09abc3f789b9492c7a7345aa1bbc47271131cd72f0c1cf32416ba.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i37d29e90b678b9968642fb73a5ac7d10f486d6d16a90be81949d96897dd78f57.CancelRequestBuilder) {
    return i37d29e90b678b9968642fb73a5ac7d10f486d6d16a90be81949d96897dd78f57.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property exceptionOccurrences for me
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
// CreateGetRequestInformation get exceptionOccurrences from me
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from me
func (m *EventItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) Decline()(*ie71f9c5564cd7668e9932138beadb6e8287d61e8d4ff49844167b6887a3e836d.DeclineRequestBuilder) {
    return ie71f9c5564cd7668e9932138beadb6e8287d61e8d4ff49844167b6887a3e836d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
func (m *EventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*ia9d65d6431c43dcbe5294d975681cddbebc4df1f0474890119ea2be5a35c004d.DismissReminderRequestBuilder) {
    return ia9d65d6431c43dcbe5294d975681cddbebc4df1f0474890119ea2be5a35c004d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i5e9f805037cadcad48f7c7e245ac583a7f26272356cfc766d5e16ce1753b6e0d.ExtensionsRequestBuilder) {
    return i5e9f805037cadcad48f7c7e245ac583a7f26272356cfc766d5e16ce1753b6e0d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i4ab94683bca463f7130220aa3e406baffefc0889228f548c19bbe36afbb73015.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i4ab94683bca463f7130220aa3e406baffefc0889228f548c19bbe36afbb73015.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ia1d675ace72efc2ee6d06f7972e11380d8529e248041c9740277d5dce2c676e9.ForwardRequestBuilder) {
    return ia1d675ace72efc2ee6d06f7972e11380d8529e248041c9740277d5dce2c676e9.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*i7398f913d0b1d7f89fb886ab526b0bff45af288b1a77cb667228377dfb3b5654.InstancesRequestBuilder) {
    return i7398f913d0b1d7f89fb886ab526b0bff45af288b1a77cb667228377dfb3b5654.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i945a6029ddd8b34dbe08094d0c9a822be4c0659cfb31cff961c8890677afb696.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i945a6029ddd8b34dbe08094d0c9a822be4c0659cfb31cff961c8890677afb696.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ibb23d60c65c4b2cd40bb61e164329196f4bea1c483e786a13500f523b78caebc.MultiValueExtendedPropertiesRequestBuilder) {
    return ibb23d60c65c4b2cd40bb61e164329196f4bea1c483e786a13500f523b78caebc.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5b6ccf20f21e40a5587355d022d743875bde8c8be78010cdb34a488eda8bb700.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i5b6ccf20f21e40a5587355d022d743875bde8c8be78010cdb34a488eda8bb700.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in me
func (m *EventItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i856c909b65889957a2bedb9bb98d91a58847ed04666e81dcab98806c27f8efe7.SingleValueExtendedPropertiesRequestBuilder) {
    return i856c909b65889957a2bedb9bb98d91a58847ed04666e81dcab98806c27f8efe7.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i55ede7043af2f7be625ab18c3dfbf56db2db1378c2d47b16a417a64c7b0d7bd3.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i55ede7043af2f7be625ab18c3dfbf56db2db1378c2d47b16a417a64c7b0d7bd3.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i7daf4fdd37c9751a56b2863942fe2dbedd34f68a4efb078f8e3f3969ee8006e0.SnoozeReminderRequestBuilder) {
    return i7daf4fdd37c9751a56b2863942fe2dbedd34f68a4efb078f8e3f3969ee8006e0.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i87e1561c37eab795bb4810566f3d99bc0a26ea5061e38e1f2389122dcb930165.TentativelyAcceptRequestBuilder) {
    return i87e1561c37eab795bb4810566f3d99bc0a26ea5061e38e1f2389122dcb930165.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

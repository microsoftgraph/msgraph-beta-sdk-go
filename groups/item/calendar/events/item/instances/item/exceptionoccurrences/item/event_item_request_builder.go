package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i15f92df896b379c185f794e9ffab67e334dedbec9e88d7b86bd84d68d577ae5f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i286f1c31f8fd38525ee7d192dffca61b2a6863707a278d98bdd7935381861683 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/forward"
    i35573439ed685eddef23590bc0a1c9fb4f2ea0675471fd4af0af13251f5a7c46 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/accept"
    i36bd95c8b35fc004367c0ffab7c4b0a838fb9f35ac8297d9992fa9c7e6915b1a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/cancel"
    i541aa50f9ab36177d08a293614f2097ab3bbac67379aa06d892a8c89f9b9d945 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/decline"
    i5c0221b58587f947d01c8d30e7024054af1cad4538de5d1c62ffb60e9f974153 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i5d6da724adab0c2686d22abc37fe4b97f4fbfc7f69783572b96133b75798ef15 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions"
    i72787732b27ce6fe3b41b3ec1fd7bad4c374f6c4d97c34101ad274f1c9783b72 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i83564140ebf2e1f802106b61ba96f9a9d9d439ec86f9f2a08676beea206b2ade "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    ia561bc68d074f8a902de97a01b770979657ec027086a50f6155f19e112a56d6a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    icaf70b8c2e4eb4b4ea516ec92c86f703b2e1b0b35c32339faec7df8abf5aa7ba "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/calendar"
    icde6de5a6d7e4df71858b7b4d0a200e3583cbf4f2ec6a50dd173713f3f76c79c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments"
    i188c5cbea292c3ad74ec0251c7d640496b3f2d7301aa020f91325adea45a4427 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i308a595453d756eba27a16accae5d44920f86478ad7c30f3173b004296a51dc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    i632edb3f4dacb5abe9443ab6f3ac4264eff3b76a60b4802a9594b7dc718c4bef "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    iba8898c88c80074167c6661c5f80875a66bb01111f106d0509b6aeb775567ac7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Accept()(*i35573439ed685eddef23590bc0a1c9fb4f2ea0675471fd4af0af13251f5a7c46.AcceptRequestBuilder) {
    return i35573439ed685eddef23590bc0a1c9fb4f2ea0675471fd4af0af13251f5a7c46.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*icde6de5a6d7e4df71858b7b4d0a200e3583cbf4f2ec6a50dd173713f3f76c79c.AttachmentsRequestBuilder) {
    return icde6de5a6d7e4df71858b7b4d0a200e3583cbf4f2ec6a50dd173713f3f76c79c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i632edb3f4dacb5abe9443ab6f3ac4264eff3b76a60b4802a9594b7dc718c4bef.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i632edb3f4dacb5abe9443ab6f3ac4264eff3b76a60b4802a9594b7dc718c4bef.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*icaf70b8c2e4eb4b4ea516ec92c86f703b2e1b0b35c32339faec7df8abf5aa7ba.CalendarRequestBuilder) {
    return icaf70b8c2e4eb4b4ea516ec92c86f703b2e1b0b35c32339faec7df8abf5aa7ba.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i36bd95c8b35fc004367c0ffab7c4b0a838fb9f35ac8297d9992fa9c7e6915b1a.CancelRequestBuilder) {
    return i36bd95c8b35fc004367c0ffab7c4b0a838fb9f35ac8297d9992fa9c7e6915b1a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property exceptionOccurrences for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property exceptionOccurrences for groups
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
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from groups
func (m *EventItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from groups
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property exceptionOccurrences in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property exceptionOccurrences in groups
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
func (m *EventItemRequestBuilder) Decline()(*i541aa50f9ab36177d08a293614f2097ab3bbac67379aa06d892a8c89f9b9d945.DeclineRequestBuilder) {
    return i541aa50f9ab36177d08a293614f2097ab3bbac67379aa06d892a8c89f9b9d945.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeleteWithResponseHandler delete navigation property exceptionOccurrences for groups
func (m *EventItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property exceptionOccurrences for groups
func (m *EventItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *EventItemRequestBuilder) DismissReminder()(*i72787732b27ce6fe3b41b3ec1fd7bad4c374f6c4d97c34101ad274f1c9783b72.DismissReminderRequestBuilder) {
    return i72787732b27ce6fe3b41b3ec1fd7bad4c374f6c4d97c34101ad274f1c9783b72.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i5d6da724adab0c2686d22abc37fe4b97f4fbfc7f69783572b96133b75798ef15.ExtensionsRequestBuilder) {
    return i5d6da724adab0c2686d22abc37fe4b97f4fbfc7f69783572b96133b75798ef15.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i308a595453d756eba27a16accae5d44920f86478ad7c30f3173b004296a51dc3.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i308a595453d756eba27a16accae5d44920f86478ad7c30f3173b004296a51dc3.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i286f1c31f8fd38525ee7d192dffca61b2a6863707a278d98bdd7935381861683.ForwardRequestBuilder) {
    return i286f1c31f8fd38525ee7d192dffca61b2a6863707a278d98bdd7935381861683.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler get exceptionOccurrences from groups
func (m *EventItemRequestBuilder) GetWithResponseHandler(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler get exceptionOccurrences from groups
func (m *EventItemRequestBuilder) GetWithResponseHandler(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i83564140ebf2e1f802106b61ba96f9a9d9d439ec86f9f2a08676beea206b2ade.MultiValueExtendedPropertiesRequestBuilder) {
    return i83564140ebf2e1f802106b61ba96f9a9d9d439ec86f9f2a08676beea206b2ade.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iba8898c88c80074167c6661c5f80875a66bb01111f106d0509b6aeb775567ac7.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iba8898c88c80074167c6661c5f80875a66bb01111f106d0509b6aeb775567ac7.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property exceptionOccurrences in groups
func (m *EventItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property exceptionOccurrences in groups
func (m *EventItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia561bc68d074f8a902de97a01b770979657ec027086a50f6155f19e112a56d6a.SingleValueExtendedPropertiesRequestBuilder) {
    return ia561bc68d074f8a902de97a01b770979657ec027086a50f6155f19e112a56d6a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i188c5cbea292c3ad74ec0251c7d640496b3f2d7301aa020f91325adea45a4427.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i188c5cbea292c3ad74ec0251c7d640496b3f2d7301aa020f91325adea45a4427.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i15f92df896b379c185f794e9ffab67e334dedbec9e88d7b86bd84d68d577ae5f.SnoozeReminderRequestBuilder) {
    return i15f92df896b379c185f794e9ffab67e334dedbec9e88d7b86bd84d68d577ae5f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i5c0221b58587f947d01c8d30e7024054af1cad4538de5d1c62ffb60e9f974153.TentativelyAcceptRequestBuilder) {
    return i5c0221b58587f947d01c8d30e7024054af1cad4538de5d1c62ffb60e9f974153.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

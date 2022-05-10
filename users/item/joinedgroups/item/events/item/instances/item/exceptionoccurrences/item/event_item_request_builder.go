package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i057e6e58a74943858a2bf0e13fa820806f2dee15b95a2a88a4df4b73e394da98 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/calendar"
    i0ad7581118e310395a496bd9305fa845ef5ff4d74e3492e4a3236889d551ba7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i57c19bc154bad4e6c4d801a68b21f7ee6a5dee1df2d8480e6328b9c045f0ff7a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/extensions"
    i63c57db24ef01b3fb2cb6813ed60b5fee766e28dd299eb00c083a8ff6bfbdaa5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i6f8be33649105143e051e4cb4415e357448959580b79648a3c8c81ca564cc40d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/forward"
    i88df442027326cf28201f3d9314f9f4ff30db2dcf48d013f661720e6f176ecf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/decline"
    ia25bd9810ea64616a787184ff484b8649063c47257d12de1922a4e05b99833ea "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/cancel"
    iab85839181f30dac4bb494873fb1ad1182471b15f7f6d9e03884139f59db5612 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ibcbf1eaf0a7312089a74a4fd6fe4652059f5b4e9a0d398ace81c3a74b4703e68 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ic589778e355b02d8600bb4f4c25d699ba5b2f550d418d22f6b84ac01057fdda9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    id5d31d7a98e69ac0396f495762c5dbe1511eb8c35578aa675b19e7500b05c678 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/accept"
    ieb45871c7a747210bb4dc7c3e033175153a481dfe230896a78fc30991adea699 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/attachments"
    i0eece99216054a9d1d1fa3c7730b1b9f87c21108422748d6826d950b2db5b8eb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    i25b3e03d49e4b71cebce7948a8d294e65982eed31d9648f150b78985ef45ad19 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i2ce013fe40d4afbc662e5b6b9b9198b033276ebdfd3fe20206f006cb3de19d5f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i9ffc7a7bcde4e5963644d6876897eb99c5de44b8013dea9f59ed9a3303d63866 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*id5d31d7a98e69ac0396f495762c5dbe1511eb8c35578aa675b19e7500b05c678.AcceptRequestBuilder) {
    return id5d31d7a98e69ac0396f495762c5dbe1511eb8c35578aa675b19e7500b05c678.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ieb45871c7a747210bb4dc7c3e033175153a481dfe230896a78fc30991adea699.AttachmentsRequestBuilder) {
    return ieb45871c7a747210bb4dc7c3e033175153a481dfe230896a78fc30991adea699.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i0eece99216054a9d1d1fa3c7730b1b9f87c21108422748d6826d950b2db5b8eb.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i0eece99216054a9d1d1fa3c7730b1b9f87c21108422748d6826d950b2db5b8eb.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i057e6e58a74943858a2bf0e13fa820806f2dee15b95a2a88a4df4b73e394da98.CalendarRequestBuilder) {
    return i057e6e58a74943858a2bf0e13fa820806f2dee15b95a2a88a4df4b73e394da98.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ia25bd9810ea64616a787184ff484b8649063c47257d12de1922a4e05b99833ea.CancelRequestBuilder) {
    return ia25bd9810ea64616a787184ff484b8649063c47257d12de1922a4e05b99833ea.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i88df442027326cf28201f3d9314f9f4ff30db2dcf48d013f661720e6f176ecf5.DeclineRequestBuilder) {
    return i88df442027326cf28201f3d9314f9f4ff30db2dcf48d013f661720e6f176ecf5.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*ic589778e355b02d8600bb4f4c25d699ba5b2f550d418d22f6b84ac01057fdda9.DismissReminderRequestBuilder) {
    return ic589778e355b02d8600bb4f4c25d699ba5b2f550d418d22f6b84ac01057fdda9.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i57c19bc154bad4e6c4d801a68b21f7ee6a5dee1df2d8480e6328b9c045f0ff7a.ExtensionsRequestBuilder) {
    return i57c19bc154bad4e6c4d801a68b21f7ee6a5dee1df2d8480e6328b9c045f0ff7a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9ffc7a7bcde4e5963644d6876897eb99c5de44b8013dea9f59ed9a3303d63866.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9ffc7a7bcde4e5963644d6876897eb99c5de44b8013dea9f59ed9a3303d63866.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i6f8be33649105143e051e4cb4415e357448959580b79648a3c8c81ca564cc40d.ForwardRequestBuilder) {
    return i6f8be33649105143e051e4cb4415e357448959580b79648a3c8c81ca564cc40d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i0ad7581118e310395a496bd9305fa845ef5ff4d74e3492e4a3236889d551ba7c.MultiValueExtendedPropertiesRequestBuilder) {
    return i0ad7581118e310395a496bd9305fa845ef5ff4d74e3492e4a3236889d551ba7c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2ce013fe40d4afbc662e5b6b9b9198b033276ebdfd3fe20206f006cb3de19d5f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i2ce013fe40d4afbc662e5b6b9b9198b033276ebdfd3fe20206f006cb3de19d5f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*iab85839181f30dac4bb494873fb1ad1182471b15f7f6d9e03884139f59db5612.SingleValueExtendedPropertiesRequestBuilder) {
    return iab85839181f30dac4bb494873fb1ad1182471b15f7f6d9e03884139f59db5612.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i25b3e03d49e4b71cebce7948a8d294e65982eed31d9648f150b78985ef45ad19.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i25b3e03d49e4b71cebce7948a8d294e65982eed31d9648f150b78985ef45ad19.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i63c57db24ef01b3fb2cb6813ed60b5fee766e28dd299eb00c083a8ff6bfbdaa5.SnoozeReminderRequestBuilder) {
    return i63c57db24ef01b3fb2cb6813ed60b5fee766e28dd299eb00c083a8ff6bfbdaa5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibcbf1eaf0a7312089a74a4fd6fe4652059f5b4e9a0d398ace81c3a74b4703e68.TentativelyAcceptRequestBuilder) {
    return ibcbf1eaf0a7312089a74a4fd6fe4652059f5b4e9a0d398ace81c3a74b4703e68.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

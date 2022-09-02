package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5ed92fd148bdc06f62d723a519e11294b027fcab06932cc8229c55a913c6d94d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    i609c8658afb6e1fba2a557f53d6922337c0f9edf1dc75831dc139a3415191766 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i69894920f14453469f795f7cb40de228ab7f836fef5b2f2dbf9cdf8137bf284e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/dismissreminder"
    i7e71e3dad7a2820575af73abb842a6c0c224884ee8c682df81731917da0f8da9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/extensions"
    i871ce287a05516d39b1f49b3ae1e2a426f83722d180a385081df74ec3f7f251f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/decline"
    i9a1e5ca4fe5eaefa5e57e0bebdb944f28dbd393bd7b739088c5631bbd042b7d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/forward"
    iafee5f26af0fe975979dcf074057d5af18d9c42f6d619d78d60fa98753a1feec "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances"
    ib588a5ab51723c7f9ec45e85751073701622073e769d43b7de88f39ccba4a2a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/attachments"
    ic56ac1871236ae48e7b3fc5930be8d69ca1f924edc3b7980ee1cf3625730c2a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/snoozereminder"
    ic6df926a0a70a22f8a92765ed1fbd717a506827bff7f077706931fbfc6c5586e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/tentativelyaccept"
    id94d3975c3adcd09c892191657b78b9f76d18068014f2d9041c68423ba454f1a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/calendar"
    idb4996d355bc53d050e503e90737ec4f9e704071425773a85950c05a5af8723f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/accept"
    ifae52fa81feb2ed838bc694548d4274a40630d46213e3718c5486e1722f97ae1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/cancel"
    i37f69ed4054bde17a26093e3e3fc1545828da5a735e599e3b7eb7d95ee8130ac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i5b74d94d488b72814b9eb6e3d92ba3a76ab31a93851b39df43dcd4d95e1c2b2f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item"
    i77a7938a37296865d6779e08919b34b4a3a2a6cbbf298846630ecb5cd4747b81 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/attachments/item"
    i82ccdb27f5e15d7a4371b49aef10fe2c92f85394db41a6ec469f5b0f2cd6ca4f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ie4617bc8c6d518ee081e0b25e8c845c2b91541cb95d8f03a9fc51ba5bca569f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*idb4996d355bc53d050e503e90737ec4f9e704071425773a85950c05a5af8723f.AcceptRequestBuilder) {
    return idb4996d355bc53d050e503e90737ec4f9e704071425773a85950c05a5af8723f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ib588a5ab51723c7f9ec45e85751073701622073e769d43b7de88f39ccba4a2a4.AttachmentsRequestBuilder) {
    return ib588a5ab51723c7f9ec45e85751073701622073e769d43b7de88f39ccba4a2a4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i77a7938a37296865d6779e08919b34b4a3a2a6cbbf298846630ecb5cd4747b81.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i77a7938a37296865d6779e08919b34b4a3a2a6cbbf298846630ecb5cd4747b81.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*id94d3975c3adcd09c892191657b78b9f76d18068014f2d9041c68423ba454f1a.CalendarRequestBuilder) {
    return id94d3975c3adcd09c892191657b78b9f76d18068014f2d9041c68423ba454f1a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ifae52fa81feb2ed838bc694548d4274a40630d46213e3718c5486e1722f97ae1.CancelRequestBuilder) {
    return ifae52fa81feb2ed838bc694548d4274a40630d46213e3718c5486e1722f97ae1.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i871ce287a05516d39b1f49b3ae1e2a426f83722d180a385081df74ec3f7f251f.DeclineRequestBuilder) {
    return i871ce287a05516d39b1f49b3ae1e2a426f83722d180a385081df74ec3f7f251f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i69894920f14453469f795f7cb40de228ab7f836fef5b2f2dbf9cdf8137bf284e.DismissReminderRequestBuilder) {
    return i69894920f14453469f795f7cb40de228ab7f836fef5b2f2dbf9cdf8137bf284e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i7e71e3dad7a2820575af73abb842a6c0c224884ee8c682df81731917da0f8da9.ExtensionsRequestBuilder) {
    return i7e71e3dad7a2820575af73abb842a6c0c224884ee8c682df81731917da0f8da9.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ie4617bc8c6d518ee081e0b25e8c845c2b91541cb95d8f03a9fc51ba5bca569f0.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ie4617bc8c6d518ee081e0b25e8c845c2b91541cb95d8f03a9fc51ba5bca569f0.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i9a1e5ca4fe5eaefa5e57e0bebdb944f28dbd393bd7b739088c5631bbd042b7d3.ForwardRequestBuilder) {
    return i9a1e5ca4fe5eaefa5e57e0bebdb944f28dbd393bd7b739088c5631bbd042b7d3.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) Instances()(*iafee5f26af0fe975979dcf074057d5af18d9c42f6d619d78d60fa98753a1feec.InstancesRequestBuilder) {
    return iafee5f26af0fe975979dcf074057d5af18d9c42f6d619d78d60fa98753a1feec.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i5b74d94d488b72814b9eb6e3d92ba3a76ab31a93851b39df43dcd4d95e1c2b2f.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i5b74d94d488b72814b9eb6e3d92ba3a76ab31a93851b39df43dcd4d95e1c2b2f.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i5ed92fd148bdc06f62d723a519e11294b027fcab06932cc8229c55a913c6d94d.MultiValueExtendedPropertiesRequestBuilder) {
    return i5ed92fd148bdc06f62d723a519e11294b027fcab06932cc8229c55a913c6d94d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i37f69ed4054bde17a26093e3e3fc1545828da5a735e599e3b7eb7d95ee8130ac.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i37f69ed4054bde17a26093e3e3fc1545828da5a735e599e3b7eb7d95ee8130ac.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i609c8658afb6e1fba2a557f53d6922337c0f9edf1dc75831dc139a3415191766.SingleValueExtendedPropertiesRequestBuilder) {
    return i609c8658afb6e1fba2a557f53d6922337c0f9edf1dc75831dc139a3415191766.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i82ccdb27f5e15d7a4371b49aef10fe2c92f85394db41a6ec469f5b0f2cd6ca4f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i82ccdb27f5e15d7a4371b49aef10fe2c92f85394db41a6ec469f5b0f2cd6ca4f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ic56ac1871236ae48e7b3fc5930be8d69ca1f924edc3b7980ee1cf3625730c2a7.SnoozeReminderRequestBuilder) {
    return ic56ac1871236ae48e7b3fc5930be8d69ca1f924edc3b7980ee1cf3625730c2a7.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ic6df926a0a70a22f8a92765ed1fbd717a506827bff7f077706931fbfc6c5586e.TentativelyAcceptRequestBuilder) {
    return ic6df926a0a70a22f8a92765ed1fbd717a506827bff7f077706931fbfc6c5586e.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

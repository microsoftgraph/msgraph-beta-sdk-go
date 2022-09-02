package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1bcb02c01f74f11e3aa24ae414b4c02fbc829ed1cb109d46b654b3389426216e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/multivalueextendedproperties"
    i31a3b37c9dbd3939c543d5983439186da29622278bf71c6087acbd816f5c9239 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/dismissreminder"
    i3b113a779a2e1c26df62a5fc2779e15b76ed3e1090178152bfb73a05102fd7ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/extensions"
    i42ac3ded08c709492e9cc73bc81cf7f4f1449cd1dfae36ce915fb987d3483119 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/decline"
    i478bc546e67483a8a03a15808e207d1d211bbbfe1f2859bbad91fd68f0e9d11a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/forward"
    i4c47294298a9edb565053aaf021f3f8a03f8dd2db6441a9594f3c81e71a84ca2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/snoozereminder"
    i68776d8a0abd76d8055c38aaa88c6dfe00f5914bab2339676865bca7101b4c45 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances"
    i68a37fe1f53b4e518739b68cfbe9c468a06abeb302d9309775d0fb5864c0f17c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar"
    i7c71d80390df615835b393e48dd37596b7a2b2840beddb58f10ed47068f5db1a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/tentativelyaccept"
    i7e83e7737d083b5a247d4b212e9838745f73ee0044cef7495db101a4d27d0473 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/cancel"
    i990077bc9239ae4abf654300201198e18c2649d863293f51f3f9660e7dd48615 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/singlevalueextendedproperties"
    ibeac8ac69135ca43836cbc7e86a3296df23382a77085df4916e1d8db90d1d1b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences"
    id7b295f54e3c066e9aee590c7e1cf7df8f692457387f19fc085cb1eaf48eebe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/accept"
    ide5e10ee2a6673c0fe49d9fdac7898a497e239b5500fc38a3178a52ff622bf83 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/attachments"
    i0aff26eaed51b1fe4dcefaf7fdc9082ff77390914e54b94471bbabd9a4585874 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item"
    i0fdb512d350c6ee756976fb95fad12a002ad0d7b0d3fcf13590e89792367c733 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/multivalueextendedproperties/item"
    i53332243af1a23b575f7314ef56a041e8aad0a06d071d02238c1b61aeb8cefd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item"
    iab96f1b678646231464bd3b7fb938638c604d6f548e5c414d63590125b8591b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/singlevalueextendedproperties/item"
    ic351771b8e3695ea47a713295cb3355c9bcc6d5ad38b45b39b527beb6475782b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/attachments/item"
    ie035c10affe323f0959d54cad2fac3bd37b1a2e6c88bcaece5284fe0542447db "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.group entity.
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
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Read-only.
type EventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string
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
func (m *EventItemRequestBuilder) Accept()(*id7b295f54e3c066e9aee590c7e1cf7df8f692457387f19fc085cb1eaf48eebe5.AcceptRequestBuilder) {
    return id7b295f54e3c066e9aee590c7e1cf7df8f692457387f19fc085cb1eaf48eebe5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ide5e10ee2a6673c0fe49d9fdac7898a497e239b5500fc38a3178a52ff622bf83.AttachmentsRequestBuilder) {
    return ide5e10ee2a6673c0fe49d9fdac7898a497e239b5500fc38a3178a52ff622bf83.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ic351771b8e3695ea47a713295cb3355c9bcc6d5ad38b45b39b527beb6475782b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ic351771b8e3695ea47a713295cb3355c9bcc6d5ad38b45b39b527beb6475782b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i68a37fe1f53b4e518739b68cfbe9c468a06abeb302d9309775d0fb5864c0f17c.CalendarRequestBuilder) {
    return i68a37fe1f53b4e518739b68cfbe9c468a06abeb302d9309775d0fb5864c0f17c.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i7e83e7737d083b5a247d4b212e9838745f73ee0044cef7495db101a4d27d0473.CancelRequestBuilder) {
    return i7e83e7737d083b5a247d4b212e9838745f73ee0044cef7495db101a4d27d0473.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}{?startDateTime,endDateTime,%24select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calendarView for groups
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
// CreateGetRequestInformation the calendar view for the calendar. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the calendar view for the calendar. Read-only.
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
// CreatePatchRequestInformation update the navigation property calendarView in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calendarView in groups
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
func (m *EventItemRequestBuilder) Decline()(*i42ac3ded08c709492e9cc73bc81cf7f4f1449cd1dfae36ce915fb987d3483119.DeclineRequestBuilder) {
    return i42ac3ded08c709492e9cc73bc81cf7f4f1449cd1dfae36ce915fb987d3483119.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*i31a3b37c9dbd3939c543d5983439186da29622278bf71c6087acbd816f5c9239.DismissReminderRequestBuilder) {
    return i31a3b37c9dbd3939c543d5983439186da29622278bf71c6087acbd816f5c9239.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ibeac8ac69135ca43836cbc7e86a3296df23382a77085df4916e1d8db90d1d1b2.ExceptionOccurrencesRequestBuilder) {
    return ibeac8ac69135ca43836cbc7e86a3296df23382a77085df4916e1d8db90d1d1b2.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i53332243af1a23b575f7314ef56a041e8aad0a06d071d02238c1b61aeb8cefd9.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i53332243af1a23b575f7314ef56a041e8aad0a06d071d02238c1b61aeb8cefd9.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i3b113a779a2e1c26df62a5fc2779e15b76ed3e1090178152bfb73a05102fd7ea.ExtensionsRequestBuilder) {
    return i3b113a779a2e1c26df62a5fc2779e15b76ed3e1090178152bfb73a05102fd7ea.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ie035c10affe323f0959d54cad2fac3bd37b1a2e6c88bcaece5284fe0542447db.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ie035c10affe323f0959d54cad2fac3bd37b1a2e6c88bcaece5284fe0542447db.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i478bc546e67483a8a03a15808e207d1d211bbbfe1f2859bbad91fd68f0e9d11a.ForwardRequestBuilder) {
    return i478bc546e67483a8a03a15808e207d1d211bbbfe1f2859bbad91fd68f0e9d11a.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*i68776d8a0abd76d8055c38aaa88c6dfe00f5914bab2339676865bca7101b4c45.InstancesRequestBuilder) {
    return i68776d8a0abd76d8055c38aaa88c6dfe00f5914bab2339676865bca7101b4c45.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i0aff26eaed51b1fe4dcefaf7fdc9082ff77390914e54b94471bbabd9a4585874.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i0aff26eaed51b1fe4dcefaf7fdc9082ff77390914e54b94471bbabd9a4585874.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i1bcb02c01f74f11e3aa24ae414b4c02fbc829ed1cb109d46b654b3389426216e.MultiValueExtendedPropertiesRequestBuilder) {
    return i1bcb02c01f74f11e3aa24ae414b4c02fbc829ed1cb109d46b654b3389426216e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0fdb512d350c6ee756976fb95fad12a002ad0d7b0d3fcf13590e89792367c733.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i0fdb512d350c6ee756976fb95fad12a002ad0d7b0d3fcf13590e89792367c733.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i990077bc9239ae4abf654300201198e18c2649d863293f51f3f9660e7dd48615.SingleValueExtendedPropertiesRequestBuilder) {
    return i990077bc9239ae4abf654300201198e18c2649d863293f51f3f9660e7dd48615.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iab96f1b678646231464bd3b7fb938638c604d6f548e5c414d63590125b8591b1.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return iab96f1b678646231464bd3b7fb938638c604d6f548e5c414d63590125b8591b1.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i4c47294298a9edb565053aaf021f3f8a03f8dd2db6441a9594f3c81e71a84ca2.SnoozeReminderRequestBuilder) {
    return i4c47294298a9edb565053aaf021f3f8a03f8dd2db6441a9594f3c81e71a84ca2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i7c71d80390df615835b393e48dd37596b7a2b2840beddb58f10ed47068f5db1a.TentativelyAcceptRequestBuilder) {
    return i7c71d80390df615835b393e48dd37596b7a2b2840beddb58f10ed47068f5db1a.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0310d8e855046989511651998bfe413522a3b459da159e0dc8784492431eb0fd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/accept"
    i10ff157410f4caba25fb64973dc0bb7ae659f5fe13810eb1836d9162885361dd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/decline"
    i187fdd445cd15d72d49a6258cc263c479ab0ef074c16c41f06832236319aeda6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/multivalueextendedproperties"
    i1f2a173b514e2e8969d0480de050d812781d92a32bbb934d4ac0270956807244 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences"
    i43a5239102111d4464dbbb07b5bcf0014843328a1968b977e0d975e261f1451f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/attachments"
    i49f1a7a50f35d1b30614456acde18d6362a691d9fa1b19f1f1a5cdaf9630cf41 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/forward"
    i5173f2ed401da55284f4db7a1b92f8aa505899fb87300956a6c85fa555e2bf83 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/cancel"
    i906d194cc3e617587f1c5efdad7d4f9ce0116df9b58cb88ddf8eaa65db1ab68f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances"
    i9736266d4318336b5c22cc28b3cd10ab3880d0e5e193c8310e680bc3cbffc9ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/tentativelyaccept"
    ia30e244ddfb0349dc37fbb76184996e3c242ee83ce5a65f1e3218e46fd482cdf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/calendar"
    iafe89c1fa9e7893a20c5fa81e1c7c68d13d2c8d48fbaa71745f6ac7fbd32a8be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/extensions"
    ie18c38c268443ce8e13c854391fffd06b1cb9ebd1893e0fe5a38148d0224945c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/snoozereminder"
    ie86c1a49c7e108efe09ea7e41f8d9fdb50ddf77614df2d4ebba5986a99658451 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/singlevalueextendedproperties"
    if58b8ff23b1df293bb8163498f8ad9867c5e0f41fbb9919b2f159d1f98d7d594 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/dismissreminder"
    i407fa14099f8f21355cc4393cbbf7ba504cf2369dedf20f03995a6ab44e9f2fd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/multivalueextendedproperties/item"
    i7b7e805ab81cde8a98356318b401680f3a8979a4b70d5264cfadbb4dc16ee816 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/singlevalueextendedproperties/item"
    i8061af62edbf559165f06c66fd4b438cb6b765baed71f5405a783be5430a36ea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/extensions/item"
    i86e5a9473e59a14b11dc31fe541d37ff458d77ca44cbe124a193022f349b705d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item"
    ic19b6138a9c12c2143d07b3c037c9fd55f8ecc82e0c24e7da37f8921af6f4d60 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/attachments/item"
    ifb5134ff3ade3b59a17a0916d04699b9cd1e1442a8321299681180d57d88bfd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/instances/item"
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
func (m *EventItemRequestBuilder) Accept()(*i0310d8e855046989511651998bfe413522a3b459da159e0dc8784492431eb0fd.AcceptRequestBuilder) {
    return i0310d8e855046989511651998bfe413522a3b459da159e0dc8784492431eb0fd.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i43a5239102111d4464dbbb07b5bcf0014843328a1968b977e0d975e261f1451f.AttachmentsRequestBuilder) {
    return i43a5239102111d4464dbbb07b5bcf0014843328a1968b977e0d975e261f1451f.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ic19b6138a9c12c2143d07b3c037c9fd55f8ecc82e0c24e7da37f8921af6f4d60.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ic19b6138a9c12c2143d07b3c037c9fd55f8ecc82e0c24e7da37f8921af6f4d60.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ia30e244ddfb0349dc37fbb76184996e3c242ee83ce5a65f1e3218e46fd482cdf.CalendarRequestBuilder) {
    return ia30e244ddfb0349dc37fbb76184996e3c242ee83ce5a65f1e3218e46fd482cdf.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i5173f2ed401da55284f4db7a1b92f8aa505899fb87300956a6c85fa555e2bf83.CancelRequestBuilder) {
    return i5173f2ed401da55284f4db7a1b92f8aa505899fb87300956a6c85fa555e2bf83.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendarView/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calendarView for me
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
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property calendarView in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calendarView in me
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
func (m *EventItemRequestBuilder) Decline()(*i10ff157410f4caba25fb64973dc0bb7ae659f5fe13810eb1836d9162885361dd.DeclineRequestBuilder) {
    return i10ff157410f4caba25fb64973dc0bb7ae659f5fe13810eb1836d9162885361dd.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property calendarView for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*if58b8ff23b1df293bb8163498f8ad9867c5e0f41fbb9919b2f159d1f98d7d594.DismissReminderRequestBuilder) {
    return if58b8ff23b1df293bb8163498f8ad9867c5e0f41fbb9919b2f159d1f98d7d594.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i1f2a173b514e2e8969d0480de050d812781d92a32bbb934d4ac0270956807244.ExceptionOccurrencesRequestBuilder) {
    return i1f2a173b514e2e8969d0480de050d812781d92a32bbb934d4ac0270956807244.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i86e5a9473e59a14b11dc31fe541d37ff458d77ca44cbe124a193022f349b705d.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i86e5a9473e59a14b11dc31fe541d37ff458d77ca44cbe124a193022f349b705d.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*iafe89c1fa9e7893a20c5fa81e1c7c68d13d2c8d48fbaa71745f6ac7fbd32a8be.ExtensionsRequestBuilder) {
    return iafe89c1fa9e7893a20c5fa81e1c7c68d13d2c8d48fbaa71745f6ac7fbd32a8be.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8061af62edbf559165f06c66fd4b438cb6b765baed71f5405a783be5430a36ea.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8061af62edbf559165f06c66fd4b438cb6b765baed71f5405a783be5430a36ea.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i49f1a7a50f35d1b30614456acde18d6362a691d9fa1b19f1f1a5cdaf9630cf41.ForwardRequestBuilder) {
    return i49f1a7a50f35d1b30614456acde18d6362a691d9fa1b19f1f1a5cdaf9630cf41.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Read-only.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the calendar view for the calendar. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*i906d194cc3e617587f1c5efdad7d4f9ce0116df9b58cb88ddf8eaa65db1ab68f.InstancesRequestBuilder) {
    return i906d194cc3e617587f1c5efdad7d4f9ce0116df9b58cb88ddf8eaa65db1ab68f.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ifb5134ff3ade3b59a17a0916d04699b9cd1e1442a8321299681180d57d88bfd3.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return ifb5134ff3ade3b59a17a0916d04699b9cd1e1442a8321299681180d57d88bfd3.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i187fdd445cd15d72d49a6258cc263c479ab0ef074c16c41f06832236319aeda6.MultiValueExtendedPropertiesRequestBuilder) {
    return i187fdd445cd15d72d49a6258cc263c479ab0ef074c16c41f06832236319aeda6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i407fa14099f8f21355cc4393cbbf7ba504cf2369dedf20f03995a6ab44e9f2fd.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i407fa14099f8f21355cc4393cbbf7ba504cf2369dedf20f03995a6ab44e9f2fd.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property calendarView in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ie86c1a49c7e108efe09ea7e41f8d9fdb50ddf77614df2d4ebba5986a99658451.SingleValueExtendedPropertiesRequestBuilder) {
    return ie86c1a49c7e108efe09ea7e41f8d9fdb50ddf77614df2d4ebba5986a99658451.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7b7e805ab81cde8a98356318b401680f3a8979a4b70d5264cfadbb4dc16ee816.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i7b7e805ab81cde8a98356318b401680f3a8979a4b70d5264cfadbb4dc16ee816.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie18c38c268443ce8e13c854391fffd06b1cb9ebd1893e0fe5a38148d0224945c.SnoozeReminderRequestBuilder) {
    return ie18c38c268443ce8e13c854391fffd06b1cb9ebd1893e0fe5a38148d0224945c.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i9736266d4318336b5c22cc28b3cd10ab3880d0e5e193c8310e680bc3cbffc9ff.TentativelyAcceptRequestBuilder) {
    return i9736266d4318336b5c22cc28b3cd10ab3880d0e5e193c8310e680bc3cbffc9ff.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

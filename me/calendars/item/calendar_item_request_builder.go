package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1e3d1947b13aae3f970d296ef71a4efc9e040742b37371ed6c138c0e253b1bf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarpermissions"
    i472232d154eda28ab8efd7fd9d2cc15ad8f3f68df3d90d6a2763b48e0a0fb5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/getschedule"
    i583738b29cef42ce350c32e40b7980bf07bc18ac8bde87fb209297b7485e7533 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview"
    ia34d5d2f8afdb98de85287585fbda3c10ce8763f60cf66a9c81a18f05aa37a17 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/singlevalueextendedproperties"
    ieb8d5dbdf577b51a4d75e40f7dc370bdfc82658f15e893321d1709a963ed78e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/multivalueextendedproperties"
    ied8df54815613d8a9050657a30d013ad8f4a43920a2fecbb829a53e69f7ef01a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events"
    ife78aa1675e428b35acac9c3d51963dd3aa40c64cd0973840b3ba47dc64b5df6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/allowedcalendarsharingroleswithuser"
    i00da9d0223281912e3c155074179db42ccbaa07b1ce54f8636bc10641c22665f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item"
    i285827c984f3b6bedeb07253a041706200e08a011dabe8ea34a88c14953200bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/singlevalueextendedproperties/item"
    i4e566b1ce953f539e8901716dfe774ae3102813087a55a1aa73d94bf3d010178 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item"
    i5a21f9744b199c1e3416254b970e7ad0211ea0e94c465251fdcf1915a650f97e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/multivalueextendedproperties/item"
    id86c8a2a7201c8af159055ea250d6a414ff1975b2f7042526a50a2700d86486d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarpermissions/item"
)

// CalendarItemRequestBuilder provides operations to manage the calendars property of the microsoft.graph.user entity.
type CalendarItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CalendarItemRequestBuilderGetQueryParameters the user's calendars. Read-only. Nullable.
type CalendarItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarItemRequestBuilderGetQueryParameters
}
// CalendarItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *CalendarItemRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ife78aa1675e428b35acac9c3d51963dd3aa40c64cd0973840b3ba47dc64b5df6.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ife78aa1675e428b35acac9c3d51963dd3aa40c64cd0973840b3ba47dc64b5df6.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
// CalendarPermissions the calendarPermissions property
func (m *CalendarItemRequestBuilder) CalendarPermissions()(*i1e3d1947b13aae3f970d296ef71a4efc9e040742b37371ed6c138c0e253b1bf4.CalendarPermissionsRequestBuilder) {
    return i1e3d1947b13aae3f970d296ef71a4efc9e040742b37371ed6c138c0e253b1bf4.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarPermissions.item collection
func (m *CalendarItemRequestBuilder) CalendarPermissionsById(id string)(*id86c8a2a7201c8af159055ea250d6a414ff1975b2f7042526a50a2700d86486d.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission%2Did"] = id
    }
    return id86c8a2a7201c8af159055ea250d6a414ff1975b2f7042526a50a2700d86486d.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *CalendarItemRequestBuilder) CalendarView()(*i583738b29cef42ce350c32e40b7980bf07bc18ac8bde87fb209297b7485e7533.CalendarViewRequestBuilder) {
    return i583738b29cef42ce350c32e40b7980bf07bc18ac8bde87fb209297b7485e7533.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item collection
func (m *CalendarItemRequestBuilder) CalendarViewById(id string)(*i4e566b1ce953f539e8901716dfe774ae3102813087a55a1aa73d94bf3d010178.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i4e566b1ce953f539e8901716dfe774ae3102813087a55a1aa73d94bf3d010178.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarItemRequestBuilderInternal instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarItemRequestBuilder) {
    m := &CalendarItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarItemRequestBuilder instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *CalendarItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the user's calendars. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the user's calendars. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CalendarItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, requestConfiguration *CalendarItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *CalendarItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property calendars for me
func (m *CalendarItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *CalendarItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Events the events property
func (m *CalendarItemRequestBuilder) Events()(*ied8df54815613d8a9050657a30d013ad8f4a43920a2fecbb829a53e69f7ef01a.EventsRequestBuilder) {
    return ied8df54815613d8a9050657a30d013ad8f4a43920a2fecbb829a53e69f7ef01a.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item collection
func (m *CalendarItemRequestBuilder) EventsById(id string)(*i00da9d0223281912e3c155074179db42ccbaa07b1ce54f8636bc10641c22665f.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i00da9d0223281912e3c155074179db42ccbaa07b1ce54f8636bc10641c22665f.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GetSchedule the getSchedule property
func (m *CalendarItemRequestBuilder) GetSchedule()(*i472232d154eda28ab8efd7fd9d2cc15ad8f3f68df3d90d6a2763b48e0a0fb5d2.GetScheduleRequestBuilder) {
    return i472232d154eda28ab8efd7fd9d2cc15ad8f3f68df3d90d6a2763b48e0a0fb5d2.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler the user's calendars. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) GetWithResponseHandler(requestConfiguration *CalendarItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the user's calendars. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) GetWithResponseHandler(requestConfiguration *CalendarItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCalendarFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *CalendarItemRequestBuilder) MultiValueExtendedProperties()(*ieb8d5dbdf577b51a4d75e40f7dc370bdfc82658f15e893321d1709a963ed78e2.MultiValueExtendedPropertiesRequestBuilder) {
    return ieb8d5dbdf577b51a4d75e40f7dc370bdfc82658f15e893321d1709a963ed78e2.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5a21f9744b199c1e3416254b970e7ad0211ea0e94c465251fdcf1915a650f97e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i5a21f9744b199c1e3416254b970e7ad0211ea0e94c465251fdcf1915a650f97e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, requestConfiguration *CalendarItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property calendars in me
func (m *CalendarItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, requestConfiguration *CalendarItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *CalendarItemRequestBuilder) SingleValueExtendedProperties()(*ia34d5d2f8afdb98de85287585fbda3c10ce8763f60cf66a9c81a18f05aa37a17.SingleValueExtendedPropertiesRequestBuilder) {
    return ia34d5d2f8afdb98de85287585fbda3c10ce8763f60cf66a9c81a18f05aa37a17.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i285827c984f3b6bedeb07253a041706200e08a011dabe8ea34a88c14953200bc.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i285827c984f3b6bedeb07253a041706200e08a011dabe8ea34a88c14953200bc.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package bookingbusinesses

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BookingBusinessesBookingBusinessItemRequestBuilder provides operations to manage the collection of bookingBusiness entities.
type BookingBusinessesBookingBusinessItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BookingBusinessesBookingBusinessItemRequestBuilderGetQueryParameters get the properties and relationships of a bookingBusiness object.
type BookingBusinessesBookingBusinessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookingBusinessesBookingBusinessItemRequestBuilderGetQueryParameters
}
// BookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Appointments provides operations to manage the appointments property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Appointments()(*BookingBusinessesItemAppointmentsRequestBuilder) {
    return NewBookingBusinessesItemAppointmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppointmentsById provides operations to manage the appointments property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) AppointmentsById(id string)(*BookingBusinessesItemAppointmentsBookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return NewBookingBusinessesItemAppointmentsBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CalendarView()(*BookingBusinessesItemCalendarViewRequestBuilder) {
    return NewBookingBusinessesItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CalendarViewById(id string)(*BookingBusinessesItemCalendarViewBookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return NewBookingBusinessesItemCalendarViewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBookingBusinessesBookingBusinessItemRequestBuilderInternal instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessesBookingBusinessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesBookingBusinessItemRequestBuilder) {
    m := &BookingBusinessesBookingBusinessItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/bookingBusinesses/{bookingBusiness%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBookingBusinessesBookingBusinessItemRequestBuilder instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessesBookingBusinessItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesBookingBusinessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessesBookingBusinessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete a bookingBusiness object.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get the properties and relationships of a bookingBusiness object.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of a bookingBusiness object.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Customers provides operations to manage the customers property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Customers()(*BookingBusinessesItemCustomersRequestBuilder) {
    return NewBookingBusinessesItemCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomersById provides operations to manage the customers property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CustomersById(id string)(*BookingBusinessesItemCustomersBookingCustomerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomer%2Did"] = id
    }
    return NewBookingBusinessesItemCustomersBookingCustomerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CustomQuestions provides operations to manage the customQuestions property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CustomQuestions()(*BookingBusinessesItemCustomQuestionsRequestBuilder) {
    return NewBookingBusinessesItemCustomQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomQuestionsById provides operations to manage the customQuestions property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CustomQuestionsById(id string)(*BookingBusinessesItemCustomQuestionsBookingCustomQuestionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomQuestion%2Did"] = id
    }
    return NewBookingBusinessesItemCustomQuestionsBookingCustomQuestionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete a bookingBusiness object.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get get the properties and relationships of a bookingBusiness object.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBookingBusinessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable), nil
}
// GetStaffAvailability provides operations to call the getStaffAvailability method.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) GetStaffAvailability()(*BookingBusinessesItemGetStaffAvailabilityRequestBuilder) {
    return NewBookingBusinessesItemGetStaffAvailabilityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the properties of a bookingBusiness object.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBookingBusinessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable), nil
}
// Publish provides operations to call the publish method.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Publish()(*BookingBusinessesItemPublishRequestBuilder) {
    return NewBookingBusinessesItemPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Services provides operations to manage the services property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Services()(*BookingBusinessesItemServicesRequestBuilder) {
    return NewBookingBusinessesItemServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicesById provides operations to manage the services property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) ServicesById(id string)(*BookingBusinessesItemServicesBookingServiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingService%2Did"] = id
    }
    return NewBookingBusinessesItemServicesBookingServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// StaffMembers provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) StaffMembers()(*BookingBusinessesItemStaffMembersRequestBuilder) {
    return NewBookingBusinessesItemStaffMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StaffMembersById provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) StaffMembersById(id string)(*BookingBusinessesItemStaffMembersBookingStaffMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingStaffMember%2Did"] = id
    }
    return NewBookingBusinessesItemStaffMembersBookingStaffMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unpublish provides operations to call the unpublish method.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Unpublish()(*BookingBusinessesItemUnpublishRequestBuilder) {
    return NewBookingBusinessesItemUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

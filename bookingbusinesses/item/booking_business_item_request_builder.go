package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0a7cdda26cb222932b3d70abaab89b217f2509a29415fc8ea0dc313dc6cb47a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/customers"
    i131bc9ff8c2fc67a2f7430a039bbbd0a92b1ff652e2e5f0a116686e51dc82b17 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/staffmembers"
    i282b5f21d91d31c3e5c81bcea750a4bf77b691c6d3b1b55d9359d4c1665f8f46 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/customquestions"
    i4d80f01302d2f12cfc7856faaf46602d45a2e40cc7854a1ede6303697bc2d4a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/services"
    i4e2baa0d10e2585111c8f7d52c8fcfc31b186c9cf6904f0475b65f6b86432585 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/calendarview"
    ia9496c7dee807d8e9cd4fb056272de5cbdd62a3c5eafa9377589cb50e4cb3f72 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/unpublish"
    iab9cfaa5004260332caf3903526a9ebdb267173b4a21a7fabc9278975a7053fe "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/getstaffavailability"
    iaf3ac87b764e164c3d5655c1a72252a5099d53e9aa3e9fa4fe32a0d44e6edf92 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/appointments"
    ibff8216433da6211e79401e82d7468eed7b0bf469f245d3915e84b5cadc35981 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/publish"
    i0fae5949907818b2ae894763546198c15bc34196d1bbc68fff8d5e35bf9e62bf "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/appointments/item"
    i79400fbdb20c4e1b3bc4a91d12f5b6146511a68e789f1456369fc0bfe4434560 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/staffmembers/item"
    i8dfad3dfd772d23aa0ae1b04ac3fa7b45824e6084187c192928925f7bcb86a70 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/services/item"
    i9cec441a87d2feb52d21615818f6b762a3fd05b35a9b73951208540f5eb80c77 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/customquestions/item"
    ia4fe11073e9da7c50ece70bf500822d1a07d9f1965f083164f821c9bca857112 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/customers/item"
    ieb67a4e973c38f9a77d4f81d39f0b7685c2e72c00379ef53269c356ae983bd0d "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/calendarview/item"
)

// BookingBusinessItemRequestBuilder provides operations to manage the collection of bookingBusiness entities.
type BookingBusinessItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BookingBusinessItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BookingBusinessItemRequestBuilderGetQueryParameters represents a Microsot Bookings Business.
type BookingBusinessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BookingBusinessItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookingBusinessItemRequestBuilderGetQueryParameters
}
// BookingBusinessItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Appointments the appointments property
func (m *BookingBusinessItemRequestBuilder) Appointments()(*iaf3ac87b764e164c3d5655c1a72252a5099d53e9aa3e9fa4fe32a0d44e6edf92.AppointmentsRequestBuilder) {
    return iaf3ac87b764e164c3d5655c1a72252a5099d53e9aa3e9fa4fe32a0d44e6edf92.NewAppointmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppointmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.bookingBusinesses.item.appointments.item collection
func (m *BookingBusinessItemRequestBuilder) AppointmentsById(id string)(*i0fae5949907818b2ae894763546198c15bc34196d1bbc68fff8d5e35bf9e62bf.BookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return i0fae5949907818b2ae894763546198c15bc34196d1bbc68fff8d5e35bf9e62bf.NewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *BookingBusinessItemRequestBuilder) CalendarView()(*i4e2baa0d10e2585111c8f7d52c8fcfc31b186c9cf6904f0475b65f6b86432585.CalendarViewRequestBuilder) {
    return i4e2baa0d10e2585111c8f7d52c8fcfc31b186c9cf6904f0475b65f6b86432585.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.bookingBusinesses.item.calendarView.item collection
func (m *BookingBusinessItemRequestBuilder) CalendarViewById(id string)(*ieb67a4e973c38f9a77d4f81d39f0b7685c2e72c00379ef53269c356ae983bd0d.BookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return ieb67a4e973c38f9a77d4f81d39f0b7685c2e72c00379ef53269c356ae983bd0d.NewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBookingBusinessItemRequestBuilderInternal instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessItemRequestBuilder) {
    m := &BookingBusinessItemRequestBuilder{
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
// NewBookingBusinessItemRequestBuilder instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete entity from bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *BookingBusinessItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents a Microsot Bookings Business.
func (m *BookingBusinessItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration represents a Microsot Bookings Business.
func (m *BookingBusinessItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *BookingBusinessItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update entity in bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, requestConfiguration *BookingBusinessItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Customers the customers property
func (m *BookingBusinessItemRequestBuilder) Customers()(*i0a7cdda26cb222932b3d70abaab89b217f2509a29415fc8ea0dc313dc6cb47a5.CustomersRequestBuilder) {
    return i0a7cdda26cb222932b3d70abaab89b217f2509a29415fc8ea0dc313dc6cb47a5.NewCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.bookingBusinesses.item.customers.item collection
func (m *BookingBusinessItemRequestBuilder) CustomersById(id string)(*ia4fe11073e9da7c50ece70bf500822d1a07d9f1965f083164f821c9bca857112.BookingCustomerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomer%2Did"] = id
    }
    return ia4fe11073e9da7c50ece70bf500822d1a07d9f1965f083164f821c9bca857112.NewBookingCustomerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CustomQuestions the customQuestions property
func (m *BookingBusinessItemRequestBuilder) CustomQuestions()(*i282b5f21d91d31c3e5c81bcea750a4bf77b691c6d3b1b55d9359d4c1665f8f46.CustomQuestionsRequestBuilder) {
    return i282b5f21d91d31c3e5c81bcea750a4bf77b691c6d3b1b55d9359d4c1665f8f46.NewCustomQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomQuestionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.bookingBusinesses.item.customQuestions.item collection
func (m *BookingBusinessItemRequestBuilder) CustomQuestionsById(id string)(*i9cec441a87d2feb52d21615818f6b762a3fd05b35a9b73951208540f5eb80c77.BookingCustomQuestionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomQuestion%2Did"] = id
    }
    return i9cec441a87d2feb52d21615818f6b762a3fd05b35a9b73951208540f5eb80c77.NewBookingCustomQuestionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete entity from bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete entity from bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *BookingBusinessItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get represents a Microsot Bookings Business.
func (m *BookingBusinessItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetStaffAvailability the getStaffAvailability property
func (m *BookingBusinessItemRequestBuilder) GetStaffAvailability()(*iab9cfaa5004260332caf3903526a9ebdb267173b4a21a7fabc9278975a7053fe.GetStaffAvailabilityRequestBuilder) {
    return iab9cfaa5004260332caf3903526a9ebdb267173b4a21a7fabc9278975a7053fe.NewGetStaffAvailabilityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler represents a Microsot Bookings Business.
func (m *BookingBusinessItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *BookingBusinessItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBookingBusinessFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable), nil
}
// Patch update entity in bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update entity in bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingBusinessable, requestConfiguration *BookingBusinessItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Publish the publish property
func (m *BookingBusinessItemRequestBuilder) Publish()(*ibff8216433da6211e79401e82d7468eed7b0bf469f245d3915e84b5cadc35981.PublishRequestBuilder) {
    return ibff8216433da6211e79401e82d7468eed7b0bf469f245d3915e84b5cadc35981.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Services the services property
func (m *BookingBusinessItemRequestBuilder) Services()(*i4d80f01302d2f12cfc7856faaf46602d45a2e40cc7854a1ede6303697bc2d4a3.ServicesRequestBuilder) {
    return i4d80f01302d2f12cfc7856faaf46602d45a2e40cc7854a1ede6303697bc2d4a3.NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.bookingBusinesses.item.services.item collection
func (m *BookingBusinessItemRequestBuilder) ServicesById(id string)(*i8dfad3dfd772d23aa0ae1b04ac3fa7b45824e6084187c192928925f7bcb86a70.BookingServiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingService%2Did"] = id
    }
    return i8dfad3dfd772d23aa0ae1b04ac3fa7b45824e6084187c192928925f7bcb86a70.NewBookingServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// StaffMembers the staffMembers property
func (m *BookingBusinessItemRequestBuilder) StaffMembers()(*i131bc9ff8c2fc67a2f7430a039bbbd0a92b1ff652e2e5f0a116686e51dc82b17.StaffMembersRequestBuilder) {
    return i131bc9ff8c2fc67a2f7430a039bbbd0a92b1ff652e2e5f0a116686e51dc82b17.NewStaffMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StaffMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.bookingBusinesses.item.staffMembers.item collection
func (m *BookingBusinessItemRequestBuilder) StaffMembersById(id string)(*i79400fbdb20c4e1b3bc4a91d12f5b6146511a68e789f1456369fc0bfe4434560.BookingStaffMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingStaffMember%2Did"] = id
    }
    return i79400fbdb20c4e1b3bc4a91d12f5b6146511a68e789f1456369fc0bfe4434560.NewBookingStaffMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *BookingBusinessItemRequestBuilder) Unpublish()(*ia9496c7dee807d8e9cd4fb056272de5cbdd62a3c5eafa9377589cb50e4cb3f72.UnpublishRequestBuilder) {
    return ia9496c7dee807d8e9cd4fb056272de5cbdd62a3c5eafa9377589cb50e4cb3f72.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

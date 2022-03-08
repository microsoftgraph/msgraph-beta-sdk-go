package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0a7cdda26cb222932b3d70abaab89b217f2509a29415fc8ea0dc313dc6cb47a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/customers"
    i131bc9ff8c2fc67a2f7430a039bbbd0a92b1ff652e2e5f0a116686e51dc82b17 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/staffmembers"
    i282b5f21d91d31c3e5c81bcea750a4bf77b691c6d3b1b55d9359d4c1665f8f46 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/customquestions"
    i4d80f01302d2f12cfc7856faaf46602d45a2e40cc7854a1ede6303697bc2d4a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/services"
    i4e2baa0d10e2585111c8f7d52c8fcfc31b186c9cf6904f0475b65f6b86432585 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/calendarview"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia9496c7dee807d8e9cd4fb056272de5cbdd62a3c5eafa9377589cb50e4cb3f72 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/unpublish"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BookingBusinessItemRequestBuilderDeleteOptions options for Delete
type BookingBusinessItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BookingBusinessItemRequestBuilderGetOptions options for Get
type BookingBusinessItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BookingBusinessItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BookingBusinessItemRequestBuilderGetQueryParameters represents a Microsot Bookings Business.
type BookingBusinessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BookingBusinessItemRequestBuilderPatchOptions options for Patch
type BookingBusinessItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BookingBusinessable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["bookingAppointment_id"] = id
    }
    return i0fae5949907818b2ae894763546198c15bc34196d1bbc68fff8d5e35bf9e62bf.NewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["bookingAppointment_id"] = id
    }
    return ieb67a4e973c38f9a77d4f81d39f0b7685c2e72c00379ef53269c356ae983bd0d.NewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBookingBusinessItemRequestBuilderInternal instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookingBusinessItemRequestBuilder) {
    m := &BookingBusinessItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/bookingBusinesses/{bookingBusiness_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBookingBusinessItemRequestBuilder instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookingBusinessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) CreateDeleteRequestInformation(options *BookingBusinessItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation represents a Microsot Bookings Business.
func (m *BookingBusinessItemRequestBuilder) CreateGetRequestInformation(options *BookingBusinessItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) CreatePatchRequestInformation(options *BookingBusinessItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
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
        urlTplParams["bookingCustomer_id"] = id
    }
    return ia4fe11073e9da7c50ece70bf500822d1a07d9f1965f083164f821c9bca857112.NewBookingCustomerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["bookingCustomQuestion_id"] = id
    }
    return i9cec441a87d2feb52d21615818f6b762a3fd05b35a9b73951208540f5eb80c77.NewBookingCustomQuestionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete entity from bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) Delete(options *BookingBusinessItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get represents a Microsot Bookings Business.
func (m *BookingBusinessItemRequestBuilder) Get(options *BookingBusinessItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BookingBusinessable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateBookingBusinessFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BookingBusinessable), nil
}
// Patch update entity in bookingBusinesses
func (m *BookingBusinessItemRequestBuilder) Patch(options *BookingBusinessItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *BookingBusinessItemRequestBuilder) Publish()(*ibff8216433da6211e79401e82d7468eed7b0bf469f245d3915e84b5cadc35981.PublishRequestBuilder) {
    return ibff8216433da6211e79401e82d7468eed7b0bf469f245d3915e84b5cadc35981.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["bookingService_id"] = id
    }
    return i8dfad3dfd772d23aa0ae1b04ac3fa7b45824e6084187c192928925f7bcb86a70.NewBookingServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["bookingStaffMember_id"] = id
    }
    return i79400fbdb20c4e1b3bc4a91d12f5b6146511a68e789f1456369fc0bfe4434560.NewBookingStaffMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessItemRequestBuilder) Unpublish()(*ia9496c7dee807d8e9cd4fb056272de5cbdd62a3c5eafa9377589cb50e4cb3f72.UnpublishRequestBuilder) {
    return ia9496c7dee807d8e9cd4fb056272de5cbdd62a3c5eafa9377589cb50e4cb3f72.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

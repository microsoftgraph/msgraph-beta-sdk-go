package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0a7cdda26cb222932b3d70abaab89b217f2509a29415fc8ea0dc313dc6cb47a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/customers"
    i131bc9ff8c2fc67a2f7430a039bbbd0a92b1ff652e2e5f0a116686e51dc82b17 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/staffmembers"
    i4d80f01302d2f12cfc7856faaf46602d45a2e40cc7854a1ede6303697bc2d4a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/services"
    i4e2baa0d10e2585111c8f7d52c8fcfc31b186c9cf6904f0475b65f6b86432585 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/calendarview"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia9496c7dee807d8e9cd4fb056272de5cbdd62a3c5eafa9377589cb50e4cb3f72 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/unpublish"
    iaf3ac87b764e164c3d5655c1a72252a5099d53e9aa3e9fa4fe32a0d44e6edf92 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/appointments"
    ibff8216433da6211e79401e82d7468eed7b0bf469f245d3915e84b5cadc35981 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/publish"
    i0fae5949907818b2ae894763546198c15bc34196d1bbc68fff8d5e35bf9e62bf "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/appointments/item"
    i79400fbdb20c4e1b3bc4a91d12f5b6146511a68e789f1456369fc0bfe4434560 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/staffmembers/item"
    i8dfad3dfd772d23aa0ae1b04ac3fa7b45824e6084187c192928925f7bcb86a70 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/services/item"
    ia4fe11073e9da7c50ece70bf500822d1a07d9f1965f083164f821c9bca857112 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/customers/item"
    ieb67a4e973c38f9a77d4f81d39f0b7685c2e72c00379ef53269c356ae983bd0d "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item/calendarview/item"
)

type BookingBusinessRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type BookingBusinessRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *BookingBusinessRequestBuilder) Appointments()(*iaf3ac87b764e164c3d5655c1a72252a5099d53e9aa3e9fa4fe32a0d44e6edf92.AppointmentsRequestBuilder) {
    return iaf3ac87b764e164c3d5655c1a72252a5099d53e9aa3e9fa4fe32a0d44e6edf92.NewAppointmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) AppointmentsById(id string)(*i0fae5949907818b2ae894763546198c15bc34196d1bbc68fff8d5e35bf9e62bf.BookingAppointmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["bookingAppointment_id"] = id
    }
    return i0fae5949907818b2ae894763546198c15bc34196d1bbc68fff8d5e35bf9e62bf.NewBookingAppointmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) CalendarView()(*i4e2baa0d10e2585111c8f7d52c8fcfc31b186c9cf6904f0475b65f6b86432585.CalendarViewRequestBuilder) {
    return i4e2baa0d10e2585111c8f7d52c8fcfc31b186c9cf6904f0475b65f6b86432585.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) CalendarViewById(id string)(*ieb67a4e973c38f9a77d4f81d39f0b7685c2e72c00379ef53269c356ae983bd0d.BookingAppointmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["bookingAppointment_id"] = id
    }
    return ieb67a4e973c38f9a77d4f81d39f0b7685c2e72c00379ef53269c356ae983bd0d.NewBookingAppointmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewBookingBusinessRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookingBusinessRequestBuilder) {
    m := &BookingBusinessRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/bookingBusinesses/{bookingBusiness_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewBookingBusinessRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookingBusinessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *BookingBusinessRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *BookingBusinessRequestBuilder) CreateGetRequestInformation(q func (value *BookingBusinessRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(BookingBusinessRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *BookingBusinessRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BookingBusiness, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *BookingBusinessRequestBuilder) Customers()(*i0a7cdda26cb222932b3d70abaab89b217f2509a29415fc8ea0dc313dc6cb47a5.CustomersRequestBuilder) {
    return i0a7cdda26cb222932b3d70abaab89b217f2509a29415fc8ea0dc313dc6cb47a5.NewCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) CustomersById(id string)(*ia4fe11073e9da7c50ece70bf500822d1a07d9f1965f083164f821c9bca857112.BookingCustomerRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["bookingCustomer_id"] = id
    }
    return ia4fe11073e9da7c50ece70bf500822d1a07d9f1965f083164f821c9bca857112.NewBookingCustomerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *BookingBusinessRequestBuilder) Get(q func (value *BookingBusinessRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BookingBusiness, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewBookingBusiness() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BookingBusiness), nil
}
func (m *BookingBusinessRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BookingBusiness, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *BookingBusinessRequestBuilder) Publish()(*ibff8216433da6211e79401e82d7468eed7b0bf469f245d3915e84b5cadc35981.PublishRequestBuilder) {
    return ibff8216433da6211e79401e82d7468eed7b0bf469f245d3915e84b5cadc35981.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) Services()(*i4d80f01302d2f12cfc7856faaf46602d45a2e40cc7854a1ede6303697bc2d4a3.ServicesRequestBuilder) {
    return i4d80f01302d2f12cfc7856faaf46602d45a2e40cc7854a1ede6303697bc2d4a3.NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) ServicesById(id string)(*i8dfad3dfd772d23aa0ae1b04ac3fa7b45824e6084187c192928925f7bcb86a70.BookingServiceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["bookingService_id"] = id
    }
    return i8dfad3dfd772d23aa0ae1b04ac3fa7b45824e6084187c192928925f7bcb86a70.NewBookingServiceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) StaffMembers()(*i131bc9ff8c2fc67a2f7430a039bbbd0a92b1ff652e2e5f0a116686e51dc82b17.StaffMembersRequestBuilder) {
    return i131bc9ff8c2fc67a2f7430a039bbbd0a92b1ff652e2e5f0a116686e51dc82b17.NewStaffMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) StaffMembersById(id string)(*i79400fbdb20c4e1b3bc4a91d12f5b6146511a68e789f1456369fc0bfe4434560.BookingStaffMemberRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["bookingStaffMember_id"] = id
    }
    return i79400fbdb20c4e1b3bc4a91d12f5b6146511a68e789f1456369fc0bfe4434560.NewBookingStaffMemberRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BookingBusinessRequestBuilder) Unpublish()(*ia9496c7dee807d8e9cd4fb056272de5cbdd62a3c5eafa9377589cb50e4cb3f72.UnpublishRequestBuilder) {
    return ia9496c7dee807d8e9cd4fb056272de5cbdd62a3c5eafa9377589cb50e4cb3f72.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

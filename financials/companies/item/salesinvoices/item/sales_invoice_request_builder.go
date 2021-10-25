package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i02155747beea3539ff57236a7d1a9ced210439349101163afadd0ca4dddc11ff "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/shipmentmethod"
    i0650e7911918a65a74f67fe1c197bdc18b90494b3361ef73af6e6fed5d0e8e47 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/currency"
    i34ab7468987d894218eb3c0a57bf71468ae3b457fdc7c9a1e1e06dac5693aad7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/salesinvoicelines"
    i36d06f0f74d1b6fcb1f7c52c368d55e3ae1c78fea244620a809d74b4050c21f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/cancel"
    i37e5adacee5361adfd45947751eeada85d4fb063399b1012befbf77338d1551b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/post"
    i4029970bec644a3ffd8071814f20a9d8eb5762118c76c4b80b23161fb86c2fee "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/cancelandsend"
    i73fdc6873e10630216ab19990248e2068cf89ebe3b996094b3b49c3edfb216d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/send"
    i93777e17bc2372af2aae81772eb24d80da0156546c0cdfca56a28e9618f8d889 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/customer"
    i973a00f44707ce3fe81e75ce782fefb6c04714e28ca331b622c712a4b29cb6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/paymentterm"
    iff094cb38c7f93c080b831ba288f86fafe7cb758d1a92b96280ccf5a97b91609 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/postandsend"
    if8797c053079d45780da07f4c01fb650fc8015394e895a3fdf2d677f2013e9ae "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/salesinvoicelines/item"
)

type SalesInvoiceRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SalesInvoiceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *SalesInvoiceRequestBuilder) Cancel()(*i36d06f0f74d1b6fcb1f7c52c368d55e3ae1c78fea244620a809d74b4050c21f2.CancelRequestBuilder) {
    return i36d06f0f74d1b6fcb1f7c52c368d55e3ae1c78fea244620a809d74b4050c21f2.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) CancelAndSend()(*i4029970bec644a3ffd8071814f20a9d8eb5762118c76c4b80b23161fb86c2fee.CancelAndSendRequestBuilder) {
    return i4029970bec644a3ffd8071814f20a9d8eb5762118c76c4b80b23161fb86c2fee.NewCancelAndSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewSalesInvoiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesInvoiceRequestBuilder) {
    m := &SalesInvoiceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/financials/companies/{company_id}/salesInvoices/{salesInvoice_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewSalesInvoiceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesInvoiceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesInvoiceRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SalesInvoiceRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesInvoiceRequestBuilder) CreateGetRequestInformation(q func (value *SalesInvoiceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SalesInvoiceRequestBuilderGetQueryParameters)
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
func (m *SalesInvoiceRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesInvoice, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesInvoiceRequestBuilder) Currency()(*i0650e7911918a65a74f67fe1c197bdc18b90494b3361ef73af6e6fed5d0e8e47.CurrencyRequestBuilder) {
    return i0650e7911918a65a74f67fe1c197bdc18b90494b3361ef73af6e6fed5d0e8e47.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) Customer()(*i93777e17bc2372af2aae81772eb24d80da0156546c0cdfca56a28e9618f8d889.CustomerRequestBuilder) {
    return i93777e17bc2372af2aae81772eb24d80da0156546c0cdfca56a28e9618f8d889.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SalesInvoiceRequestBuilder) Get(q func (value *SalesInvoiceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesInvoice, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSalesInvoice() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesInvoice), nil
}
func (m *SalesInvoiceRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesInvoice, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SalesInvoiceRequestBuilder) PaymentTerm()(*i973a00f44707ce3fe81e75ce782fefb6c04714e28ca331b622c712a4b29cb6cb.PaymentTermRequestBuilder) {
    return i973a00f44707ce3fe81e75ce782fefb6c04714e28ca331b622c712a4b29cb6cb.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) Post()(*i37e5adacee5361adfd45947751eeada85d4fb063399b1012befbf77338d1551b.PostRequestBuilder) {
    return i37e5adacee5361adfd45947751eeada85d4fb063399b1012befbf77338d1551b.NewPostRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) PostAndSend()(*iff094cb38c7f93c080b831ba288f86fafe7cb758d1a92b96280ccf5a97b91609.PostAndSendRequestBuilder) {
    return iff094cb38c7f93c080b831ba288f86fafe7cb758d1a92b96280ccf5a97b91609.NewPostAndSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) SalesInvoiceLines()(*i34ab7468987d894218eb3c0a57bf71468ae3b457fdc7c9a1e1e06dac5693aad7.SalesInvoiceLinesRequestBuilder) {
    return i34ab7468987d894218eb3c0a57bf71468ae3b457fdc7c9a1e1e06dac5693aad7.NewSalesInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) SalesInvoiceLinesById(id string)(*if8797c053079d45780da07f4c01fb650fc8015394e895a3fdf2d677f2013e9ae.SalesInvoiceLineRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoiceLine_id"] = id
    }
    return if8797c053079d45780da07f4c01fb650fc8015394e895a3fdf2d677f2013e9ae.NewSalesInvoiceLineRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) Send()(*i73fdc6873e10630216ab19990248e2068cf89ebe3b996094b3b49c3edfb216d3.SendRequestBuilder) {
    return i73fdc6873e10630216ab19990248e2068cf89ebe3b996094b3b49c3edfb216d3.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) ShipmentMethod()(*i02155747beea3539ff57236a7d1a9ced210439349101163afadd0ca4dddc11ff.ShipmentMethodRequestBuilder) {
    return i02155747beea3539ff57236a7d1a9ced210439349101163afadd0ca4dddc11ff.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

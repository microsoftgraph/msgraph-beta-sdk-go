package vendor_escaped

import (
	ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
	i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
	i492fcd884507cf5a1f3d87f066b6c00cda10f6321ff116df2682eb108c2cd692 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped/currency"
	icba4a323f9088474207a977dc2dc628e98e963693c938d5ce5c30e950d85ae89 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped/paymentmethod"
	id98c49e997a726a9f7b79d44019ff48dbae0a247d027461613188c31b69f6f7a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped/paymentterm"
	i93bc4bbcd1fa7551af807f53dc2f9b7c68950a789993a759c58cdb0ce1c6a637 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped/picture"
	i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type VendorRequestBuilder struct {
	pathParameters map[string]string
	requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter
	urlTemplate    string
}
type VendorRequestBuilderGetQueryParameters struct {
	ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
	Expand          []string
	Select_escpaped []string
}

func NewVendorRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter) *VendorRequestBuilder {
	m := &VendorRequestBuilder{}
	m.urlTemplate = "https://graph.microsoft.com/beta/financials/companies/{company_id}/purchaseInvoices/{purchaseInvoice_id}/vendor{?select,expand}"
	urlTplParams := make(map[string]string)
	if pathParameters != nil {
		for idx, item := range pathParameters {
			urlTplParams[idx] = item
		}
	}
	m.pathParameters = pathParameters
	m.requestAdapter = requestAdapter
	return m
}
func NewVendorRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter) *VendorRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewVendorRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *VendorRequestBuilder) CreateDeleteRequestInformation(h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption) (*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *VendorRequestBuilder) CreateGetRequestInformation(q func(value *VendorRequestBuilderGetQueryParameters) (err error), h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption) (*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
	requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
	if q != nil {
		qParams := new(VendorRequestBuilderGetQueryParameters)
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
func (m *VendorRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor, h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption) (*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *VendorRequestBuilder) Currency() *i492fcd884507cf5a1f3d87f066b6c00cda10f6321ff116df2682eb108c2cd692.CurrencyRequestBuilder {
	return i492fcd884507cf5a1f3d87f066b6c00cda10f6321ff116df2682eb108c2cd692.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
func (m *VendorRequestBuilder) Delete(h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler) error {
	requestInfo, err := m.CreateDeleteRequestInformation(h, o)
	if err != nil {
		return err
	}
	err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
	if err != nil {
		return err
	}
	return nil
}
func (m *VendorRequestBuilder) Get(q func(value *VendorRequestBuilderGetQueryParameters) (err error), h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler) (*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor, error) {
	requestInfo, err := m.CreateGetRequestInformation(q, h, o)
	if err != nil {
		return nil, err
	}
	res, err := m.requestAdapter.SendAsync(*requestInfo, func() i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable {
		return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewVendor()
	}, responseHandler)
	if err != nil {
		return nil, err
	}
	return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor), nil
}
func (m *VendorRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor, h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler) error {
	requestInfo, err := m.CreatePatchRequestInformation(body, h, o)
	if err != nil {
		return err
	}
	err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
	if err != nil {
		return err
	}
	return nil
}
func (m *VendorRequestBuilder) PaymentMethod() *icba4a323f9088474207a977dc2dc628e98e963693c938d5ce5c30e950d85ae89.PaymentMethodRequestBuilder {
	return icba4a323f9088474207a977dc2dc628e98e963693c938d5ce5c30e950d85ae89.NewPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
func (m *VendorRequestBuilder) PaymentTerm() *id98c49e997a726a9f7b79d44019ff48dbae0a247d027461613188c31b69f6f7a.PaymentTermRequestBuilder {
	return id98c49e997a726a9f7b79d44019ff48dbae0a247d027461613188c31b69f6f7a.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
func (m *VendorRequestBuilder) Picture() *i93bc4bbcd1fa7551af807f53dc2f9b7c68950a789993a759c58cdb0ce1c6a637.PictureRequestBuilder {
	return i93bc4bbcd1fa7551af807f53dc2f9b7c68950a789993a759c58cdb0ce1c6a637.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
func (m *VendorRequestBuilder) PictureById(id string) *i93bc4bbcd1fa7551af807f53dc2f9b7c68950a789993a759c58cdb0ce1c6a637.PictureRequestBuilder {
	urlTplParams := make(map[string]string)
	if m.pathParameters != nil {
		for idx, item := range m.pathParameters {
			urlTplParams[idx] = item
		}
	}
	if id != "" {
		urlTplParams["picture_id"] = id
	}
	return i93bc4bbcd1fa7551af807f53dc2f9b7c68950a789993a759c58cdb0ce1c6a637.NewPictureRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

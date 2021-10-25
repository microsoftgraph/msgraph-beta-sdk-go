package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/tasktriggers"
    i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/share"
    i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/restorefactorydefaults"
    i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/getcapabilities"
    ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/shares"
    icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/connectors"
    id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/resetdefaults"
    i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/tasktriggers/item"
)

type PrinterRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PrinterRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *PrinterRequestBuilder) Connectors()(*icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a.ConnectorsRequestBuilder) {
    return icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewPrinterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterRequestBuilder) {
    m := &PrinterRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/print/printers/{printer_id}{?select,expand}";
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
func NewPrinterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrinterRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrinterRequestBuilder) CreateGetRequestInformation(q func (value *PrinterRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PrinterRequestBuilderGetQueryParameters)
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
func (m *PrinterRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printer, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrinterRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrinterRequestBuilder) Get(q func (value *PrinterRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printer, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrinter() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printer), nil
}
func (m *PrinterRequestBuilder) GetCapabilities()(*i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.GetCapabilitiesRequestBuilder) {
    return i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.NewGetCapabilitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printer, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrinterRequestBuilder) ResetDefaults()(*id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.ResetDefaultsRequestBuilder) {
    return id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.NewResetDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) RestoreFactoryDefaults()(*i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.RestoreFactoryDefaultsRequestBuilder) {
    return i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.NewRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) Share()(*i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.ShareRequestBuilder) {
    return i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) Shares()(*ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5.SharesRequestBuilder) {
    return ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) TaskTriggers()(*i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb.TaskTriggersRequestBuilder) {
    return i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb.NewTaskTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) TaskTriggersById(id string)(*i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d.PrintTaskTriggerRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printTaskTrigger_id"] = id
    }
    return i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d.NewPrintTaskTriggerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

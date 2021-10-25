package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i01f2133292f0b08e5c7ad62ad1a919908e250568ad80f9fa7990e01a1a4675da "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reopen"
    i2ba9e3b5dde4370f9d4ee21032c5d1cf837c8172fdd1c29a522a55e73f42e955 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/operations"
    i5a4461258b51ae222704978e5dc6e8b0f90af958cbbfe0291aa23989e511468b "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections"
    i7e8c62f91304400acff09ffefec4fe5fc66f1640b9a230f33487eb271fea144f "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds"
    i7ee9cd47ece1636db5f23d48461a1f54657a90bc6e0ee26ce62b3d5e84e35d9a "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians"
    ib38135aba802e63a95783cb40679431dfd8b0b006b136ce2cf392d13cbcf39b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/settings"
    ibce968213312dcf70b6cebf357574ebaff57289bb8d2d5537d0fd2e4c8f604ea "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources"
    ibf27493a3efa3e33589ec8953bae196df0f5b31cd7f0010efe7b3550bdd8d64b "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/tags"
    icbe808dca9f215e46a201507bb9fd3d6682ffd22c4228bd2934fa6d092ec2013 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/close"
    ie37f7b6846339f3be72955eebac7b1a5f9a9b0e6768d7276f5894199b0fc4e64 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets"
    i2d904c4ba0f64303f7ec694166a350f070305a6d3282664966c3c7ee03bee849 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item"
    i6c4f9c432182d1a8383e210ea7282204df1fd76b96287e8d3e0ac99ac798c705 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item"
    i6f12c740d8ad4a07e1c35e99570b5c006a24cd6710fc2378825d45f2cebe46e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item"
    i700c8523a708bdda9b2320cf511ec749704101f5da23584bcca53db2919ad351 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item"
    i71c6b856224a9a329639e731310314e9aaa3c89985c5bf77e3adab5540d605e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/tags/item"
    i85cba7b9ace1e93810a0e9c5210f57fed67ca0ac127d57c09a74eadcbfe393df "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/operations/item"
    ia8b23b6ab36c2b10209f024c3f8396df9a37f4b7c7dbf0cdcc02383064607d36 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item"
)

type CaseRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type CaseRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *CaseRequestBuilder) Close()(*icbe808dca9f215e46a201507bb9fd3d6682ffd22c4228bd2934fa6d092ec2013.CloseRequestBuilder) {
    return icbe808dca9f215e46a201507bb9fd3d6682ffd22c4228bd2934fa6d092ec2013.NewCloseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewCaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseRequestBuilder) {
    m := &CaseRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/compliance/ediscovery/cases/{case_id}{?select,expand}";
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
func NewCaseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCaseRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *CaseRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CaseRequestBuilder) CreateGetRequestInformation(q func (value *CaseRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(CaseRequestBuilderGetQueryParameters)
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
func (m *CaseRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Case_escpaped, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CaseRequestBuilder) Custodians()(*i7ee9cd47ece1636db5f23d48461a1f54657a90bc6e0ee26ce62b3d5e84e35d9a.CustodiansRequestBuilder) {
    return i7ee9cd47ece1636db5f23d48461a1f54657a90bc6e0ee26ce62b3d5e84e35d9a.NewCustodiansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) CustodiansById(id string)(*i6f12c740d8ad4a07e1c35e99570b5c006a24cd6710fc2378825d45f2cebe46e6.CustodianRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["custodian_id"] = id
    }
    return i6f12c740d8ad4a07e1c35e99570b5c006a24cd6710fc2378825d45f2cebe46e6.NewCustodianRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CaseRequestBuilder) Get(q func (value *CaseRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Case_escpaped, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCase_escpaped() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Case_escpaped), nil
}
func (m *CaseRequestBuilder) LegalHolds()(*i7e8c62f91304400acff09ffefec4fe5fc66f1640b9a230f33487eb271fea144f.LegalHoldsRequestBuilder) {
    return i7e8c62f91304400acff09ffefec4fe5fc66f1640b9a230f33487eb271fea144f.NewLegalHoldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) LegalHoldsById(id string)(*ia8b23b6ab36c2b10209f024c3f8396df9a37f4b7c7dbf0cdcc02383064607d36.LegalHoldRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["legalHold_id"] = id
    }
    return ia8b23b6ab36c2b10209f024c3f8396df9a37f4b7c7dbf0cdcc02383064607d36.NewLegalHoldRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) NoncustodialDataSources()(*ibce968213312dcf70b6cebf357574ebaff57289bb8d2d5537d0fd2e4c8f604ea.NoncustodialDataSourcesRequestBuilder) {
    return ibce968213312dcf70b6cebf357574ebaff57289bb8d2d5537d0fd2e4c8f604ea.NewNoncustodialDataSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) NoncustodialDataSourcesById(id string)(*i2d904c4ba0f64303f7ec694166a350f070305a6d3282664966c3c7ee03bee849.NoncustodialDataSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["noncustodialDataSource_id"] = id
    }
    return i2d904c4ba0f64303f7ec694166a350f070305a6d3282664966c3c7ee03bee849.NewNoncustodialDataSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) Operations()(*i2ba9e3b5dde4370f9d4ee21032c5d1cf837c8172fdd1c29a522a55e73f42e955.OperationsRequestBuilder) {
    return i2ba9e3b5dde4370f9d4ee21032c5d1cf837c8172fdd1c29a522a55e73f42e955.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) OperationsById(id string)(*i85cba7b9ace1e93810a0e9c5210f57fed67ca0ac127d57c09a74eadcbfe393df.CaseOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["caseOperation_id"] = id
    }
    return i85cba7b9ace1e93810a0e9c5210f57fed67ca0ac127d57c09a74eadcbfe393df.NewCaseOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Case_escpaped, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CaseRequestBuilder) Reopen()(*i01f2133292f0b08e5c7ad62ad1a919908e250568ad80f9fa7990e01a1a4675da.ReopenRequestBuilder) {
    return i01f2133292f0b08e5c7ad62ad1a919908e250568ad80f9fa7990e01a1a4675da.NewReopenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) ReviewSets()(*ie37f7b6846339f3be72955eebac7b1a5f9a9b0e6768d7276f5894199b0fc4e64.ReviewSetsRequestBuilder) {
    return ie37f7b6846339f3be72955eebac7b1a5f9a9b0e6768d7276f5894199b0fc4e64.NewReviewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) ReviewSetsById(id string)(*i6c4f9c432182d1a8383e210ea7282204df1fd76b96287e8d3e0ac99ac798c705.ReviewSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["reviewSet_id"] = id
    }
    return i6c4f9c432182d1a8383e210ea7282204df1fd76b96287e8d3e0ac99ac798c705.NewReviewSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) Settings()(*ib38135aba802e63a95783cb40679431dfd8b0b006b136ce2cf392d13cbcf39b3.SettingsRequestBuilder) {
    return ib38135aba802e63a95783cb40679431dfd8b0b006b136ce2cf392d13cbcf39b3.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) SourceCollections()(*i5a4461258b51ae222704978e5dc6e8b0f90af958cbbfe0291aa23989e511468b.SourceCollectionsRequestBuilder) {
    return i5a4461258b51ae222704978e5dc6e8b0f90af958cbbfe0291aa23989e511468b.NewSourceCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) SourceCollectionsById(id string)(*i700c8523a708bdda9b2320cf511ec749704101f5da23584bcca53db2919ad351.SourceCollectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["sourceCollection_id"] = id
    }
    return i700c8523a708bdda9b2320cf511ec749704101f5da23584bcca53db2919ad351.NewSourceCollectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseRequestBuilder) Tags()(*ibf27493a3efa3e33589ec8953bae196df0f5b31cd7f0010efe7b3550bdd8d64b.TagsRequestBuilder) {
    return ibf27493a3efa3e33589ec8953bae196df0f5b31cd7f0010efe7b3550bdd8d64b.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseRequestBuilder) TagsById(id string)(*i71c6b856224a9a329639e731310314e9aaa3c89985c5bf77e3adab5540d605e7.TagRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["tag_id"] = id
    }
    return i71c6b856224a9a329639e731310314e9aaa3c89985c5bf77e3adab5540d605e7.NewTagRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

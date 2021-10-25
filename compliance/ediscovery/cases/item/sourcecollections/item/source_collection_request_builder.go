package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0597b69327ef47612760d3a06e5168d4357386c9da9233bf1fe1f95375bde771 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/custodiansources"
    i2c744ed9b76601b529ed40d577c487cc9b9cdba275f49afd6d038d45b275a597 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/additionalsources"
    i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/lastestimatestatisticsoperation"
    i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/addtoreviewsetoperation"
    i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/estimatestatistics"
    ie844c3a4d7bb8cc79a58d398b260c44e966bd08e45f89594aa2bccaa7e8180e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/noncustodialsources"
    id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/additionalsources/item"
)

type SourceCollectionRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SourceCollectionRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *SourceCollectionRequestBuilder) AdditionalSources()(*i2c744ed9b76601b529ed40d577c487cc9b9cdba275f49afd6d038d45b275a597.AdditionalSourcesRequestBuilder) {
    return i2c744ed9b76601b529ed40d577c487cc9b9cdba275f49afd6d038d45b275a597.NewAdditionalSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) AdditionalSourcesById(id string)(*id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3.DataSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["dataSource_id"] = id
    }
    return id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3.NewDataSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) AddToReviewSetOperation()(*i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409.AddToReviewSetOperationRequestBuilder) {
    return i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409.NewAddToReviewSetOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewSourceCollectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SourceCollectionRequestBuilder) {
    m := &SourceCollectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/compliance/ediscovery/cases/{case_id}/sourceCollections/{sourceCollection_id}{?select,expand}";
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
func NewSourceCollectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SourceCollectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSourceCollectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SourceCollectionRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SourceCollectionRequestBuilder) CreateGetRequestInformation(q func (value *SourceCollectionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SourceCollectionRequestBuilderGetQueryParameters)
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
func (m *SourceCollectionRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SourceCollectionRequestBuilder) CustodianSources()(*i0597b69327ef47612760d3a06e5168d4357386c9da9233bf1fe1f95375bde771.CustodianSourcesRequestBuilder) {
    return i0597b69327ef47612760d3a06e5168d4357386c9da9233bf1fe1f95375bde771.NewCustodianSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SourceCollectionRequestBuilder) EstimateStatistics()(*i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86.EstimateStatisticsRequestBuilder) {
    return i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86.NewEstimateStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) Get(q func (value *SourceCollectionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSourceCollection() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection), nil
}
func (m *SourceCollectionRequestBuilder) LastEstimateStatisticsOperation()(*i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3.LastEstimateStatisticsOperationRequestBuilder) {
    return i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3.NewLastEstimateStatisticsOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) NoncustodialSources()(*ie844c3a4d7bb8cc79a58d398b260c44e966bd08e45f89594aa2bccaa7e8180e9.NoncustodialSourcesRequestBuilder) {
    return ie844c3a4d7bb8cc79a58d398b260c44e966bd08e45f89594aa2bccaa7e8180e9.NewNoncustodialSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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

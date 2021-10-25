package schema

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2875f990bc810a0e416397fd14a01a4725fa7f0a20108aba18cb3e18e820df4a "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/directories"
    i36fdc7569d3f0809e9eefdd0c8c16c5e0395c922bdaace8196665e44a09adb27 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/functions"
    i9303411a20178a126012bdf802c85819c1aa86d05b037afdb540ae9ec7518178 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/filteroperators"
    idd8cf83633c367895d6f418067b0f38663f134d6ecc5ab7c35bc0987d1b4ac9e "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/parseexpression"
    i7fa3917741c65999dc9b62767c91ce9eefb3d6af40f071ce0f23f64af751b350 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/directories/item"
)

type SchemaRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SchemaRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewSchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchemaRequestBuilder) {
    m := &SchemaRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/servicePrincipals/{servicePrincipal_id}/synchronization/templates/{synchronizationTemplate_id}/schema{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewSchemaRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SchemaRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SchemaRequestBuilder) CreateGetRequestInformation(q func (value *SchemaRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SchemaRequestBuilderGetQueryParameters)
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
func (m *SchemaRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSchema, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SchemaRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SchemaRequestBuilder) Directories()(*i2875f990bc810a0e416397fd14a01a4725fa7f0a20108aba18cb3e18e820df4a.DirectoriesRequestBuilder) {
    return i2875f990bc810a0e416397fd14a01a4725fa7f0a20108aba18cb3e18e820df4a.NewDirectoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SchemaRequestBuilder) DirectoriesById(id string)(*i7fa3917741c65999dc9b62767c91ce9eefb3d6af40f071ce0f23f64af751b350.DirectoryDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryDefinition_id"] = id
    }
    return i7fa3917741c65999dc9b62767c91ce9eefb3d6af40f071ce0f23f64af751b350.NewDirectoryDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SchemaRequestBuilder) FilterOperators()(*i9303411a20178a126012bdf802c85819c1aa86d05b037afdb540ae9ec7518178.FilterOperatorsRequestBuilder) {
    return i9303411a20178a126012bdf802c85819c1aa86d05b037afdb540ae9ec7518178.NewFilterOperatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SchemaRequestBuilder) Functions()(*i36fdc7569d3f0809e9eefdd0c8c16c5e0395c922bdaace8196665e44a09adb27.FunctionsRequestBuilder) {
    return i36fdc7569d3f0809e9eefdd0c8c16c5e0395c922bdaace8196665e44a09adb27.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SchemaRequestBuilder) Get(q func (value *SchemaRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSchema, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSynchronizationSchema() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSchema), nil
}
func (m *SchemaRequestBuilder) ParseExpression()(*idd8cf83633c367895d6f418067b0f38663f134d6ecc5ab7c35bc0987d1b4ac9e.ParseExpressionRequestBuilder) {
    return idd8cf83633c367895d6f418067b0f38663f134d6ecc5ab7c35bc0987d1b4ac9e.NewParseExpressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SchemaRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSchema, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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

package schema

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3f8397838cdcbf421c948cec12bb6c8ae29443adfbfffafafcb0800257df0249 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/directories"
    i5b8f0ae1eb8c05092bdc285776ceda0a64242dbf7a6bd4a2dea538cb5922a020 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/parseexpression"
    i60bb4f0c9eb8dde47c85b8b4f412b0bca85a017fe9fc37056e7b1b55f179d459 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/filteroperators"
    ic3d9837872d528f66b24cce114afdbd85a16641424441ba4d77e26ac34ff3528 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/functions"
    iffe98a2db6c23deb1db6d3a5137091a707788c003479b60e45e086d4e05d1658 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item/schema/directories/item"
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
    m.urlTemplate = "https://graph.microsoft.com/beta/applications/{application_id}/synchronization/templates/{synchronizationTemplate_id}/schema{?select,expand}";
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
func (m *SchemaRequestBuilder) Directories()(*i3f8397838cdcbf421c948cec12bb6c8ae29443adfbfffafafcb0800257df0249.DirectoriesRequestBuilder) {
    return i3f8397838cdcbf421c948cec12bb6c8ae29443adfbfffafafcb0800257df0249.NewDirectoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SchemaRequestBuilder) DirectoriesById(id string)(*iffe98a2db6c23deb1db6d3a5137091a707788c003479b60e45e086d4e05d1658.DirectoryDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryDefinition_id"] = id
    }
    return iffe98a2db6c23deb1db6d3a5137091a707788c003479b60e45e086d4e05d1658.NewDirectoryDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SchemaRequestBuilder) FilterOperators()(*i60bb4f0c9eb8dde47c85b8b4f412b0bca85a017fe9fc37056e7b1b55f179d459.FilterOperatorsRequestBuilder) {
    return i60bb4f0c9eb8dde47c85b8b4f412b0bca85a017fe9fc37056e7b1b55f179d459.NewFilterOperatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SchemaRequestBuilder) Functions()(*ic3d9837872d528f66b24cce114afdbd85a16641424441ba4d77e26ac34ff3528.FunctionsRequestBuilder) {
    return ic3d9837872d528f66b24cce114afdbd85a16641424441ba4d77e26ac34ff3528.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *SchemaRequestBuilder) ParseExpression()(*i5b8f0ae1eb8c05092bdc285776ceda0a64242dbf7a6bd4a2dea538cb5922a020.ParseExpressionRequestBuilder) {
    return i5b8f0ae1eb8c05092bdc285776ceda0a64242dbf7a6bd4a2dea538cb5922a020.NewParseExpressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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

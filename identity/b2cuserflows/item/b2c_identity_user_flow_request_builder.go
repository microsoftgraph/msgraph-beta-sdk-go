package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3c30639732fa96ae7b2270ea94ac4d733ab4ce3489526a5e5cd1c27cbb90b0df "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/identityproviders"
    ia8501a142e35d07776280d07596b55002a93c3239073e18d8e5ae9479c0a6576 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userattributeassignments"
    ib718f8165ffdc29693dd68518ac7ea324def0c6882130bda31e56b4f400e55ed "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userflowidentityproviders"
    ic8b27e7bcdb899e4a47d9ea30cb80e33646e9ca77dcb752f0bec144f3c6411be "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages"
    i4f9895372ea85cee1ee4bf9a8b414752e4270d4b2216846b499db783b5644d6c "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item"
    ib4590aff7f09f446c7ae6a4a70dfc243aac07fb4522106f724124d4539ac2888 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userattributeassignments/item"
)

type B2cIdentityUserFlowRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type B2cIdentityUserFlowRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewB2cIdentityUserFlowRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2cIdentityUserFlowRequestBuilder) {
    m := &B2cIdentityUserFlowRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identity/b2cUserFlows/{b2cIdentityUserFlow_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewB2cIdentityUserFlowRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2cIdentityUserFlowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cIdentityUserFlowRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *B2cIdentityUserFlowRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *B2cIdentityUserFlowRequestBuilder) CreateGetRequestInformation(q func (value *B2cIdentityUserFlowRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(B2cIdentityUserFlowRequestBuilderGetQueryParameters)
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
func (m *B2cIdentityUserFlowRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2cIdentityUserFlow, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *B2cIdentityUserFlowRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *B2cIdentityUserFlowRequestBuilder) Get(q func (value *B2cIdentityUserFlowRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2cIdentityUserFlow, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewB2cIdentityUserFlow() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2cIdentityUserFlow), nil
}
func (m *B2cIdentityUserFlowRequestBuilder) IdentityProviders()(*i3c30639732fa96ae7b2270ea94ac4d733ab4ce3489526a5e5cd1c27cbb90b0df.IdentityProvidersRequestBuilder) {
    return i3c30639732fa96ae7b2270ea94ac4d733ab4ce3489526a5e5cd1c27cbb90b0df.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *B2cIdentityUserFlowRequestBuilder) Languages()(*ic8b27e7bcdb899e4a47d9ea30cb80e33646e9ca77dcb752f0bec144f3c6411be.LanguagesRequestBuilder) {
    return ic8b27e7bcdb899e4a47d9ea30cb80e33646e9ca77dcb752f0bec144f3c6411be.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *B2cIdentityUserFlowRequestBuilder) LanguagesById(id string)(*i4f9895372ea85cee1ee4bf9a8b414752e4270d4b2216846b499db783b5644d6c.UserFlowLanguageConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration_id"] = id
    }
    return i4f9895372ea85cee1ee4bf9a8b414752e4270d4b2216846b499db783b5644d6c.NewUserFlowLanguageConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *B2cIdentityUserFlowRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2cIdentityUserFlow, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *B2cIdentityUserFlowRequestBuilder) UserAttributeAssignments()(*ia8501a142e35d07776280d07596b55002a93c3239073e18d8e5ae9479c0a6576.UserAttributeAssignmentsRequestBuilder) {
    return ia8501a142e35d07776280d07596b55002a93c3239073e18d8e5ae9479c0a6576.NewUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *B2cIdentityUserFlowRequestBuilder) UserAttributeAssignmentsById(id string)(*ib4590aff7f09f446c7ae6a4a70dfc243aac07fb4522106f724124d4539ac2888.IdentityUserFlowAttributeAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment_id"] = id
    }
    return ib4590aff7f09f446c7ae6a4a70dfc243aac07fb4522106f724124d4539ac2888.NewIdentityUserFlowAttributeAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *B2cIdentityUserFlowRequestBuilder) UserFlowIdentityProviders()(*ib718f8165ffdc29693dd68518ac7ea324def0c6882130bda31e56b4f400e55ed.UserFlowIdentityProvidersRequestBuilder) {
    return ib718f8165ffdc29693dd68518ac7ea324def0c6882130bda31e56b4f400e55ed.NewUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

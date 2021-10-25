package directorysettingtemplates

import (
    i03094d9d1021b32e492d609c092d84ed4cfb558c32f0c8ade3ad08ceb1eead66 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/validateproperties"
    i8b3159fe7053462b6a5b7687099a8e653ffaaa3e49ac1c52748ad001ac05d4d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/getbyids"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    if2eec762f4a60cfa480122020666674cae3dd0c0dd67417b7da42d6d7d9efc9f "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/getuserownedobjects"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type DirectorySettingTemplatesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DirectorySettingTemplatesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escaped []string;
    Skip *int32;
    Top *int32;
}
func NewDirectorySettingTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplatesRequestBuilder) {
    m := &DirectorySettingTemplatesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/directorySettingTemplates{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDirectorySettingTemplatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectorySettingTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DirectorySettingTemplatesRequestBuilder) CreateGetRequestInformation(q func (value *DirectorySettingTemplatesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DirectorySettingTemplatesRequestBuilderGetQueryParameters)
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
func (m *DirectorySettingTemplatesRequestBuilder) CreatePostRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
func (m *DirectorySettingTemplatesRequestBuilder) Get(q func (value *DirectorySettingTemplatesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*DirectorySettingTemplatesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectorySettingTemplatesResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*DirectorySettingTemplatesResponse), nil
}
func (m *DirectorySettingTemplatesRequestBuilder) GetByIds()(*i8b3159fe7053462b6a5b7687099a8e653ffaaa3e49ac1c52748ad001ac05d4d9.GetByIdsRequestBuilder) {
    return i8b3159fe7053462b6a5b7687099a8e653ffaaa3e49ac1c52748ad001ac05d4d9.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectorySettingTemplatesRequestBuilder) GetUserOwnedObjects()(*if2eec762f4a60cfa480122020666674cae3dd0c0dd67417b7da42d6d7d9efc9f.GetUserOwnedObjectsRequestBuilder) {
    return if2eec762f4a60cfa480122020666674cae3dd0c0dd67417b7da42d6d7d9efc9f.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectorySettingTemplatesRequestBuilder) Post(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectorySettingTemplate() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate), nil
}
func (m *DirectorySettingTemplatesRequestBuilder) ValidateProperties()(*i03094d9d1021b32e492d609c092d84ed4cfb558c32f0c8ade3ad08ceb1eead66.ValidatePropertiesRequestBuilder) {
    return i03094d9d1021b32e492d609c092d84ed4cfb558c32f0c8ade3ad08ceb1eead66.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i09e0fb5ad1ad570bae5ac63f3e01cd3eeb4d487f97efe0766919eec487abbd14 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/getmembergroups"
    i350af3bac14d3575d18b98f3718632a63e909339a7e1d311aeed01b007c6326e "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/getmemberobjects"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/checkmembergroups"
    i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/checkmemberobjects"
    i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/restore"
)

type DirectorySettingTemplateRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DirectorySettingTemplateRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *DirectorySettingTemplateRequestBuilder) CheckMemberGroups()(*i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e.CheckMemberGroupsRequestBuilder) {
    return i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectorySettingTemplateRequestBuilder) CheckMemberObjects()(*i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24.CheckMemberObjectsRequestBuilder) {
    return i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewDirectorySettingTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplateRequestBuilder) {
    m := &DirectorySettingTemplateRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/directorySettingTemplates/{directorySettingTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDirectorySettingTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectorySettingTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DirectorySettingTemplateRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectorySettingTemplateRequestBuilder) CreateGetRequestInformation(q func (value *DirectorySettingTemplateRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DirectorySettingTemplateRequestBuilderGetQueryParameters)
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
func (m *DirectorySettingTemplateRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectorySettingTemplateRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DirectorySettingTemplateRequestBuilder) Get(q func (value *DirectorySettingTemplateRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectorySettingTemplate() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate), nil
}
func (m *DirectorySettingTemplateRequestBuilder) GetMemberGroups()(*i09e0fb5ad1ad570bae5ac63f3e01cd3eeb4d487f97efe0766919eec487abbd14.GetMemberGroupsRequestBuilder) {
    return i09e0fb5ad1ad570bae5ac63f3e01cd3eeb4d487f97efe0766919eec487abbd14.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectorySettingTemplateRequestBuilder) GetMemberObjects()(*i350af3bac14d3575d18b98f3718632a63e909339a7e1d311aeed01b007c6326e.GetMemberObjectsRequestBuilder) {
    return i350af3bac14d3575d18b98f3718632a63e909339a7e1d311aeed01b007c6326e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectorySettingTemplateRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DirectorySettingTemplateRequestBuilder) Restore()(*i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee.RestoreRequestBuilder) {
    return i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i08aa694b5efe17e6cfb40a0b57c60c903682f60b25cedf6b77b700e69456e371 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/addlanguagefiles"
    i30872138ba0a6007708c6d50246b61e9dbb1c6693196e74b8fe1b81dbab24eed "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/remove"
    i573a2f06fa9d3a5f4c3539c0c567663d335bba6395fe3de82bf7752385eacc91 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/grouppolicyoperations"
    ib8f5596b1f4d2b0563c2e8fec9b75de399dfc5c22c6593b08984b3c43081a2f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/uploadnewversion"
    ie6c6ba293495002c0854fd282065a1b58a7aa28885d4f70b197baa6c4d5221e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/updatelanguagefiles"
    ie97617c59d880c34963563417867a7e1047adb153e1210b59f965efbd77e2b11 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/removelanguagefiles"
    ib268e3b08fe1fad55fddc2f3fcee0a20255a98eba2bf68e561d4e06a67c8ce93 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/grouppolicyoperations/item"
)

type GroupPolicyUploadedDefinitionFileRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type GroupPolicyUploadedDefinitionFileRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) AddLanguageFiles()(*i08aa694b5efe17e6cfb40a0b57c60c903682f60b25cedf6b77b700e69456e371.AddLanguageFilesRequestBuilder) {
    return i08aa694b5efe17e6cfb40a0b57c60c903682f60b25cedf6b77b700e69456e371.NewAddLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewGroupPolicyUploadedDefinitionFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyUploadedDefinitionFileRequestBuilder) {
    m := &GroupPolicyUploadedDefinitionFileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewGroupPolicyUploadedDefinitionFileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyUploadedDefinitionFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyUploadedDefinitionFileRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) CreateGetRequestInformation(q func (value *GroupPolicyUploadedDefinitionFileRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(GroupPolicyUploadedDefinitionFileRequestBuilderGetQueryParameters)
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
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyUploadedDefinitionFile, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) Get(q func (value *GroupPolicyUploadedDefinitionFileRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyUploadedDefinitionFile, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyUploadedDefinitionFile() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyUploadedDefinitionFile), nil
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) GroupPolicyOperations()(*i573a2f06fa9d3a5f4c3539c0c567663d335bba6395fe3de82bf7752385eacc91.GroupPolicyOperationsRequestBuilder) {
    return i573a2f06fa9d3a5f4c3539c0c567663d335bba6395fe3de82bf7752385eacc91.NewGroupPolicyOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) GroupPolicyOperationsById(id string)(*ib268e3b08fe1fad55fddc2f3fcee0a20255a98eba2bf68e561d4e06a67c8ce93.GroupPolicyOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyOperation_id"] = id
    }
    return ib268e3b08fe1fad55fddc2f3fcee0a20255a98eba2bf68e561d4e06a67c8ce93.NewGroupPolicyOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyUploadedDefinitionFile, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) Remove()(*i30872138ba0a6007708c6d50246b61e9dbb1c6693196e74b8fe1b81dbab24eed.RemoveRequestBuilder) {
    return i30872138ba0a6007708c6d50246b61e9dbb1c6693196e74b8fe1b81dbab24eed.NewRemoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) RemoveLanguageFiles()(*ie97617c59d880c34963563417867a7e1047adb153e1210b59f965efbd77e2b11.RemoveLanguageFilesRequestBuilder) {
    return ie97617c59d880c34963563417867a7e1047adb153e1210b59f965efbd77e2b11.NewRemoveLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) UpdateLanguageFiles()(*ie6c6ba293495002c0854fd282065a1b58a7aa28885d4f70b197baa6c4d5221e4.UpdateLanguageFilesRequestBuilder) {
    return ie6c6ba293495002c0854fd282065a1b58a7aa28885d4f70b197baa6c4d5221e4.NewUpdateLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) UploadNewVersion()(*ib8f5596b1f4d2b0563c2e8fec9b75de399dfc5c22c6593b08984b3c43081a2f1.UploadNewVersionRequestBuilder) {
    return ib8f5596b1f4d2b0563c2e8fec9b75de399dfc5c22c6593b08984b3c43081a2f1.NewUploadNewVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

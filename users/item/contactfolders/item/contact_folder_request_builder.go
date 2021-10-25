package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3a83c58b232b381cc2ac12b3f11210de504ea48c5af551848ecd877ae1c766d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/singlevalueextendedproperties"
    ibda5f39767e403563993d59e2b18b8dcfe3d7c36ddd838d47b66f2a86342a85e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/multivalueextendedproperties"
    id4d4dad6b141e4f3af220acfdabfa9caf138cb05be16a53c9acd99720e8b78b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders"
    if41f8eaad27f2c9af0ca44df54e1acce03cf14a2cbde5af7e2af9a93abc757dd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts"
    i0137c5e81c3575ccf39564d07df5d5d12b3baf9db8099260cdee9dc017beb2b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/multivalueextendedproperties/item"
    i92daaea591536ee1438537d5a544b203a04f572b3ae5ff8711b9ddf8556888ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/singlevalueextendedproperties/item"
    ia5d8eb4c45f9ae822ce66db9645f27c43d4c66be10b3e37ff4e067424f9efd39 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/childfolders/item"
    icddb3f84902cf1910a8d5c630745aff38f5a668f2ea8de5510f40532344b9f4b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item/contacts/item"
)

type ContactFolderRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ContactFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ContactFolderRequestBuilder) ChildFolders()(*id4d4dad6b141e4f3af220acfdabfa9caf138cb05be16a53c9acd99720e8b78b7.ChildFoldersRequestBuilder) {
    return id4d4dad6b141e4f3af220acfdabfa9caf138cb05be16a53c9acd99720e8b78b7.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) ChildFoldersById(id string)(*ia5d8eb4c45f9ae822ce66db9645f27c43d4c66be10b3e37ff4e067424f9efd39.ContactFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["contactFolder_id1"] = id
    }
    return ia5d8eb4c45f9ae822ce66db9645f27c43d4c66be10b3e37ff4e067424f9efd39.NewContactFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewContactFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    m := &ContactFolderRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/contactFolders/{contactFolder_id}{?select,expand}";
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
func NewContactFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContactFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContactFolderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContactFolderRequestBuilder) Contacts()(*if41f8eaad27f2c9af0ca44df54e1acce03cf14a2cbde5af7e2af9a93abc757dd.ContactsRequestBuilder) {
    return if41f8eaad27f2c9af0ca44df54e1acce03cf14a2cbde5af7e2af9a93abc757dd.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) ContactsById(id string)(*icddb3f84902cf1910a8d5c630745aff38f5a668f2ea8de5510f40532344b9f4b.ContactRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return icddb3f84902cf1910a8d5c630745aff38f5a668f2ea8de5510f40532344b9f4b.NewContactRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ContactFolderRequestBuilder) CreateGetRequestInformation(q func (value *ContactFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ContactFolderRequestBuilderGetQueryParameters)
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
func (m *ContactFolderRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ContactFolderRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ContactFolderRequestBuilder) Get(q func (value *ContactFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContactFolder() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolder), nil
}
func (m *ContactFolderRequestBuilder) MultiValueExtendedProperties()(*ibda5f39767e403563993d59e2b18b8dcfe3d7c36ddd838d47b66f2a86342a85e.MultiValueExtendedPropertiesRequestBuilder) {
    return ibda5f39767e403563993d59e2b18b8dcfe3d7c36ddd838d47b66f2a86342a85e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0137c5e81c3575ccf39564d07df5d5d12b3baf9db8099260cdee9dc017beb2b5.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0137c5e81c3575ccf39564d07df5d5d12b3baf9db8099260cdee9dc017beb2b5.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContactFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ContactFolderRequestBuilder) SingleValueExtendedProperties()(*i3a83c58b232b381cc2ac12b3f11210de504ea48c5af551848ecd877ae1c766d6.SingleValueExtendedPropertiesRequestBuilder) {
    return i3a83c58b232b381cc2ac12b3f11210de504ea48c5af551848ecd877ae1c766d6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContactFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i92daaea591536ee1438537d5a544b203a04f572b3ae5ff8711b9ddf8556888ed.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i92daaea591536ee1438537d5a544b203a04f572b3ae5ff8711b9ddf8556888ed.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

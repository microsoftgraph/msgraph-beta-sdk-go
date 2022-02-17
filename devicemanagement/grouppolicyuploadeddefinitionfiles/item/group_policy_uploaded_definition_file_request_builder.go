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

// GroupPolicyUploadedDefinitionFileRequestBuilder builds and executes requests for operations under \deviceManagement\groupPolicyUploadedDefinitionFiles\{groupPolicyUploadedDefinitionFile-id}
type GroupPolicyUploadedDefinitionFileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupPolicyUploadedDefinitionFileRequestBuilderDeleteOptions options for Delete
type GroupPolicyUploadedDefinitionFileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyUploadedDefinitionFileRequestBuilderGetOptions options for Get
type GroupPolicyUploadedDefinitionFileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyUploadedDefinitionFileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyUploadedDefinitionFileRequestBuilderGetQueryParameters the available group policy uploaded definition files for this account.
type GroupPolicyUploadedDefinitionFileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GroupPolicyUploadedDefinitionFileRequestBuilderPatchOptions options for Patch
type GroupPolicyUploadedDefinitionFileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyUploadedDefinitionFile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) AddLanguageFiles()(*i08aa694b5efe17e6cfb40a0b57c60c903682f60b25cedf6b77b700e69456e371.AddLanguageFilesRequestBuilder) {
    return i08aa694b5efe17e6cfb40a0b57c60c903682f60b25cedf6b77b700e69456e371.NewAddLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupPolicyUploadedDefinitionFileRequestBuilderInternal instantiates a new GroupPolicyUploadedDefinitionFileRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyUploadedDefinitionFileRequestBuilder) {
    m := &GroupPolicyUploadedDefinitionFileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyUploadedDefinitionFileRequestBuilder instantiates a new GroupPolicyUploadedDefinitionFileRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyUploadedDefinitionFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyUploadedDefinitionFileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyUploadedDefinitionFileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) CreateGetRequestInformation(options *GroupPolicyUploadedDefinitionFileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyUploadedDefinitionFileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) Delete(options *GroupPolicyUploadedDefinitionFileRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) Get(options *GroupPolicyUploadedDefinitionFileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyUploadedDefinitionFile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyUploadedDefinitionFile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyUploadedDefinitionFile), nil
}
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) GroupPolicyOperations()(*i573a2f06fa9d3a5f4c3539c0c567663d335bba6395fe3de82bf7752385eacc91.GroupPolicyOperationsRequestBuilder) {
    return i573a2f06fa9d3a5f4c3539c0c567663d335bba6395fe3de82bf7752385eacc91.NewGroupPolicyOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyOperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyUploadedDefinitionFiles.item.groupPolicyOperations.item collection
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
// Patch the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileRequestBuilder) Patch(options *GroupPolicyUploadedDefinitionFileRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
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

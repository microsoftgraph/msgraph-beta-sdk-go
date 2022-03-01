package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i124c1a9dce598f2005f17f9c0aac4d9fb6304152df8c29b1d9e949847c6ee284 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/item/checkmembergroups"
    i1ff63b415e315d04703f6c2c7e54fc5e3002e5a1b4d0cab74e10f9ac6d90bf6e "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/item/getmemberobjects"
    i385093f6933c3d82d593f6649995a7bea55ac8e10ec0d1bf2e617b7aa87564d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/item/checkmemberobjects"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i794ec9bbf31af9f4ef213a94300299f42c65fec1d605adcb17effb7dec394f16 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/item/restore"
    ib41b1a1b649cee7ce263c38052e18a122c5cfd842f4909256bbff18420458682 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/item/getmembergroups"
)

// DirectoryRoleTemplateItemRequestBuilder builds and executes requests for operations under \directoryRoleTemplates\{directoryRoleTemplate-id}
type DirectoryRoleTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRoleTemplateItemRequestBuilderDeleteOptions options for Delete
type DirectoryRoleTemplateItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleTemplateItemRequestBuilderGetOptions options for Get
type DirectoryRoleTemplateItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRoleTemplateItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleTemplateItemRequestBuilderGetQueryParameters get entity from directoryRoleTemplates by key
type DirectoryRoleTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRoleTemplateItemRequestBuilderPatchOptions options for Patch
type DirectoryRoleTemplateItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectoryRoleTemplateItemRequestBuilder) CheckMemberGroups()(*i124c1a9dce598f2005f17f9c0aac4d9fb6304152df8c29b1d9e949847c6ee284.CheckMemberGroupsRequestBuilder) {
    return i124c1a9dce598f2005f17f9c0aac4d9fb6304152df8c29b1d9e949847c6ee284.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplateItemRequestBuilder) CheckMemberObjects()(*i385093f6933c3d82d593f6649995a7bea55ac8e10ec0d1bf2e617b7aa87564d9.CheckMemberObjectsRequestBuilder) {
    return i385093f6933c3d82d593f6649995a7bea55ac8e10ec0d1bf2e617b7aa87564d9.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRoleTemplateItemRequestBuilderInternal instantiates a new DirectoryRoleTemplateItemRequestBuilder and sets the default values.
func NewDirectoryRoleTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplateItemRequestBuilder) {
    m := &DirectoryRoleTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoleTemplates/{directoryRoleTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRoleTemplateItemRequestBuilder instantiates a new DirectoryRoleTemplateItemRequestBuilder and sets the default values.
func NewDirectoryRoleTemplateItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directoryRoleTemplates
func (m *DirectoryRoleTemplateItemRequestBuilder) CreateDeleteRequestInformation(options *DirectoryRoleTemplateItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from directoryRoleTemplates by key
func (m *DirectoryRoleTemplateItemRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleTemplateItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in directoryRoleTemplates
func (m *DirectoryRoleTemplateItemRequestBuilder) CreatePatchRequestInformation(options *DirectoryRoleTemplateItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from directoryRoleTemplates
func (m *DirectoryRoleTemplateItemRequestBuilder) Delete(options *DirectoryRoleTemplateItemRequestBuilderDeleteOptions)(error) {
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
// Get get entity from directoryRoleTemplates by key
func (m *DirectoryRoleTemplateItemRequestBuilder) Get(options *DirectoryRoleTemplateItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectoryRoleTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplate), nil
}
func (m *DirectoryRoleTemplateItemRequestBuilder) GetMemberGroups()(*ib41b1a1b649cee7ce263c38052e18a122c5cfd842f4909256bbff18420458682.GetMemberGroupsRequestBuilder) {
    return ib41b1a1b649cee7ce263c38052e18a122c5cfd842f4909256bbff18420458682.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplateItemRequestBuilder) GetMemberObjects()(*i1ff63b415e315d04703f6c2c7e54fc5e3002e5a1b4d0cab74e10f9ac6d90bf6e.GetMemberObjectsRequestBuilder) {
    return i1ff63b415e315d04703f6c2c7e54fc5e3002e5a1b4d0cab74e10f9ac6d90bf6e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in directoryRoleTemplates
func (m *DirectoryRoleTemplateItemRequestBuilder) Patch(options *DirectoryRoleTemplateItemRequestBuilderPatchOptions)(error) {
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
func (m *DirectoryRoleTemplateItemRequestBuilder) Restore()(*i794ec9bbf31af9f4ef213a94300299f42c65fec1d605adcb17effb7dec394f16.RestoreRequestBuilder) {
    return i794ec9bbf31af9f4ef213a94300299f42c65fec1d605adcb17effb7dec394f16.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

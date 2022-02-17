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

// DirectoryRoleTemplateRequestBuilder builds and executes requests for operations under \directoryRoleTemplates\{directoryRoleTemplate-id}
type DirectoryRoleTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRoleTemplateRequestBuilderDeleteOptions options for Delete
type DirectoryRoleTemplateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleTemplateRequestBuilderGetOptions options for Get
type DirectoryRoleTemplateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRoleTemplateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleTemplateRequestBuilderGetQueryParameters get entity from directoryRoleTemplates by key
type DirectoryRoleTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRoleTemplateRequestBuilderPatchOptions options for Patch
type DirectoryRoleTemplateRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectoryRoleTemplateRequestBuilder) CheckMemberGroups()(*i124c1a9dce598f2005f17f9c0aac4d9fb6304152df8c29b1d9e949847c6ee284.CheckMemberGroupsRequestBuilder) {
    return i124c1a9dce598f2005f17f9c0aac4d9fb6304152df8c29b1d9e949847c6ee284.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplateRequestBuilder) CheckMemberObjects()(*i385093f6933c3d82d593f6649995a7bea55ac8e10ec0d1bf2e617b7aa87564d9.CheckMemberObjectsRequestBuilder) {
    return i385093f6933c3d82d593f6649995a7bea55ac8e10ec0d1bf2e617b7aa87564d9.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRoleTemplateRequestBuilderInternal instantiates a new DirectoryRoleTemplateRequestBuilder and sets the default values.
func NewDirectoryRoleTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplateRequestBuilder) {
    m := &DirectoryRoleTemplateRequestBuilder{
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
// NewDirectoryRoleTemplateRequestBuilder instantiates a new DirectoryRoleTemplateRequestBuilder and sets the default values.
func NewDirectoryRoleTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directoryRoleTemplates
func (m *DirectoryRoleTemplateRequestBuilder) CreateDeleteRequestInformation(options *DirectoryRoleTemplateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectoryRoleTemplateRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleTemplateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectoryRoleTemplateRequestBuilder) CreatePatchRequestInformation(options *DirectoryRoleTemplateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectoryRoleTemplateRequestBuilder) Delete(options *DirectoryRoleTemplateRequestBuilderDeleteOptions)(error) {
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
func (m *DirectoryRoleTemplateRequestBuilder) Get(options *DirectoryRoleTemplateRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplate, error) {
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
func (m *DirectoryRoleTemplateRequestBuilder) GetMemberGroups()(*ib41b1a1b649cee7ce263c38052e18a122c5cfd842f4909256bbff18420458682.GetMemberGroupsRequestBuilder) {
    return ib41b1a1b649cee7ce263c38052e18a122c5cfd842f4909256bbff18420458682.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplateRequestBuilder) GetMemberObjects()(*i1ff63b415e315d04703f6c2c7e54fc5e3002e5a1b4d0cab74e10f9ac6d90bf6e.GetMemberObjectsRequestBuilder) {
    return i1ff63b415e315d04703f6c2c7e54fc5e3002e5a1b4d0cab74e10f9ac6d90bf6e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in directoryRoleTemplates
func (m *DirectoryRoleTemplateRequestBuilder) Patch(options *DirectoryRoleTemplateRequestBuilderPatchOptions)(error) {
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
func (m *DirectoryRoleTemplateRequestBuilder) Restore()(*i794ec9bbf31af9f4ef213a94300299f42c65fec1d605adcb17effb7dec394f16.RestoreRequestBuilder) {
    return i794ec9bbf31af9f4ef213a94300299f42c65fec1d605adcb17effb7dec394f16.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

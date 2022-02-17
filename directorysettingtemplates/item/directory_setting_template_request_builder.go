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

// DirectorySettingTemplateRequestBuilder builds and executes requests for operations under \directorySettingTemplates\{directorySettingTemplate-id}
type DirectorySettingTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectorySettingTemplateRequestBuilderDeleteOptions options for Delete
type DirectorySettingTemplateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectorySettingTemplateRequestBuilderGetOptions options for Get
type DirectorySettingTemplateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectorySettingTemplateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectorySettingTemplateRequestBuilderGetQueryParameters get entity from directorySettingTemplates by key
type DirectorySettingTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectorySettingTemplateRequestBuilderPatchOptions options for Patch
type DirectorySettingTemplateRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectorySettingTemplateRequestBuilder) CheckMemberGroups()(*i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e.CheckMemberGroupsRequestBuilder) {
    return i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectorySettingTemplateRequestBuilder) CheckMemberObjects()(*i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24.CheckMemberObjectsRequestBuilder) {
    return i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectorySettingTemplateRequestBuilderInternal instantiates a new DirectorySettingTemplateRequestBuilder and sets the default values.
func NewDirectorySettingTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplateRequestBuilder) {
    m := &DirectorySettingTemplateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directorySettingTemplates/{directorySettingTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectorySettingTemplateRequestBuilder instantiates a new DirectorySettingTemplateRequestBuilder and sets the default values.
func NewDirectorySettingTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectorySettingTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directorySettingTemplates
func (m *DirectorySettingTemplateRequestBuilder) CreateDeleteRequestInformation(options *DirectorySettingTemplateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from directorySettingTemplates by key
func (m *DirectorySettingTemplateRequestBuilder) CreateGetRequestInformation(options *DirectorySettingTemplateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in directorySettingTemplates
func (m *DirectorySettingTemplateRequestBuilder) CreatePatchRequestInformation(options *DirectorySettingTemplateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from directorySettingTemplates
func (m *DirectorySettingTemplateRequestBuilder) Delete(options *DirectorySettingTemplateRequestBuilderDeleteOptions)(error) {
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
// Get get entity from directorySettingTemplates by key
func (m *DirectorySettingTemplateRequestBuilder) Get(options *DirectorySettingTemplateRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectorySettingTemplate() }, nil, nil)
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
// Patch update entity in directorySettingTemplates
func (m *DirectorySettingTemplateRequestBuilder) Patch(options *DirectorySettingTemplateRequestBuilderPatchOptions)(error) {
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
func (m *DirectorySettingTemplateRequestBuilder) Restore()(*i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee.RestoreRequestBuilder) {
    return i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

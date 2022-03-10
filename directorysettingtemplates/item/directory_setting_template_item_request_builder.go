package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i09e0fb5ad1ad570bae5ac63f3e01cd3eeb4d487f97efe0766919eec487abbd14 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/getmembergroups"
    i350af3bac14d3575d18b98f3718632a63e909339a7e1d311aeed01b007c6326e "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/getmemberobjects"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/checkmembergroups"
    i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/checkmemberobjects"
    i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/restore"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// DirectorySettingTemplateItemRequestBuilder provides operations to manage the collection of directorySettingTemplate entities.
type DirectorySettingTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectorySettingTemplateItemRequestBuilderDeleteOptions options for Delete
type DirectorySettingTemplateItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectorySettingTemplateItemRequestBuilderGetOptions options for Get
type DirectorySettingTemplateItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectorySettingTemplateItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectorySettingTemplateItemRequestBuilderGetQueryParameters get entity from directorySettingTemplates by key
type DirectorySettingTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectorySettingTemplateItemRequestBuilderPatchOptions options for Patch
type DirectorySettingTemplateItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplateable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectorySettingTemplateItemRequestBuilder) CheckMemberGroups()(*i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e.CheckMemberGroupsRequestBuilder) {
    return i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectorySettingTemplateItemRequestBuilder) CheckMemberObjects()(*i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24.CheckMemberObjectsRequestBuilder) {
    return i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectorySettingTemplateItemRequestBuilderInternal instantiates a new DirectorySettingTemplateItemRequestBuilder and sets the default values.
func NewDirectorySettingTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplateItemRequestBuilder) {
    m := &DirectorySettingTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directorySettingTemplates/{directorySettingTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectorySettingTemplateItemRequestBuilder instantiates a new DirectorySettingTemplateItemRequestBuilder and sets the default values.
func NewDirectorySettingTemplateItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectorySettingTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directorySettingTemplates
func (m *DirectorySettingTemplateItemRequestBuilder) CreateDeleteRequestInformation(options *DirectorySettingTemplateItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectorySettingTemplateItemRequestBuilder) CreateGetRequestInformation(options *DirectorySettingTemplateItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectorySettingTemplateItemRequestBuilder) CreatePatchRequestInformation(options *DirectorySettingTemplateItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectorySettingTemplateItemRequestBuilder) Delete(options *DirectorySettingTemplateItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from directorySettingTemplates by key
func (m *DirectorySettingTemplateItemRequestBuilder) Get(options *DirectorySettingTemplateItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDirectorySettingTemplateFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplateable), nil
}
func (m *DirectorySettingTemplateItemRequestBuilder) GetMemberGroups()(*i09e0fb5ad1ad570bae5ac63f3e01cd3eeb4d487f97efe0766919eec487abbd14.GetMemberGroupsRequestBuilder) {
    return i09e0fb5ad1ad570bae5ac63f3e01cd3eeb4d487f97efe0766919eec487abbd14.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectorySettingTemplateItemRequestBuilder) GetMemberObjects()(*i350af3bac14d3575d18b98f3718632a63e909339a7e1d311aeed01b007c6326e.GetMemberObjectsRequestBuilder) {
    return i350af3bac14d3575d18b98f3718632a63e909339a7e1d311aeed01b007c6326e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in directorySettingTemplates
func (m *DirectorySettingTemplateItemRequestBuilder) Patch(options *DirectorySettingTemplateItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DirectorySettingTemplateItemRequestBuilder) Restore()(*i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee.RestoreRequestBuilder) {
    return i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

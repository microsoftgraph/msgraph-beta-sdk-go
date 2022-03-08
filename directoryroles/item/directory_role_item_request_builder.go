package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4e1a4185ef1644f3671e08895ba2e1b95a54a1c95e280bf43511b3c9807e0daa "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/getmemberobjects"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5485bb9c491434e8f03f3f928d1e9aae35354496db691be6f7158b8c396ccd9b "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/scopedmembers"
    i97065f3b11569623118c02d59611043962f19e93102e7f5dfe72d976fbf4a73e "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/getmembergroups"
    i9935b468ea0b8a63a043fea5dd4c071f39943fdd1a429815ff6943d7bf1da7c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/checkmembergroups"
    ibe15d8c336cbfe09bcf7d7a61546d490ce45e00077f924a095523cbf08039c76 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/restore"
    iea9aebdcf95599ab8429a0d32f34af8848c1aaf957a501ef2affa3de9d6c4949 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members"
    if07c44821ea10029e12a10833d0eb1edcf1c37d67dd91ede2fd5a2dad86555f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/checkmemberobjects"
    ia2a3af813d19ede6948e6cd8a4fdc1da146bb512d7519e6200eb556b45d1b28e "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/scopedmembers/item"
    iedf3077aa9b00be73cba43b6523f0fed72ea50c39d60527d1af9a3c75c7c4728 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/item"
)

// DirectoryRoleItemRequestBuilder provides operations to manage the collection of directoryRole entities.
type DirectoryRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRoleItemRequestBuilderDeleteOptions options for Delete
type DirectoryRoleItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleItemRequestBuilderGetOptions options for Get
type DirectoryRoleItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRoleItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleItemRequestBuilderGetQueryParameters get entity from directoryRoles by key
type DirectoryRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRoleItemRequestBuilderPatchOptions options for Patch
type DirectoryRoleItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectoryRoleItemRequestBuilder) CheckMemberGroups()(*i9935b468ea0b8a63a043fea5dd4c071f39943fdd1a429815ff6943d7bf1da7c4.CheckMemberGroupsRequestBuilder) {
    return i9935b468ea0b8a63a043fea5dd4c071f39943fdd1a429815ff6943d7bf1da7c4.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleItemRequestBuilder) CheckMemberObjects()(*if07c44821ea10029e12a10833d0eb1edcf1c37d67dd91ede2fd5a2dad86555f7.CheckMemberObjectsRequestBuilder) {
    return if07c44821ea10029e12a10833d0eb1edcf1c37d67dd91ede2fd5a2dad86555f7.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRoleItemRequestBuilderInternal instantiates a new DirectoryRoleItemRequestBuilder and sets the default values.
func NewDirectoryRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleItemRequestBuilder) {
    m := &DirectoryRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles/{directoryRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRoleItemRequestBuilder instantiates a new DirectoryRoleItemRequestBuilder and sets the default values.
func NewDirectoryRoleItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directoryRoles
func (m *DirectoryRoleItemRequestBuilder) CreateDeleteRequestInformation(options *DirectoryRoleItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from directoryRoles by key
func (m *DirectoryRoleItemRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in directoryRoles
func (m *DirectoryRoleItemRequestBuilder) CreatePatchRequestInformation(options *DirectoryRoleItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from directoryRoles
func (m *DirectoryRoleItemRequestBuilder) Delete(options *DirectoryRoleItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from directoryRoles by key
func (m *DirectoryRoleItemRequestBuilder) Get(options *DirectoryRoleItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDirectoryRoleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleable), nil
}
func (m *DirectoryRoleItemRequestBuilder) GetMemberGroups()(*i97065f3b11569623118c02d59611043962f19e93102e7f5dfe72d976fbf4a73e.GetMemberGroupsRequestBuilder) {
    return i97065f3b11569623118c02d59611043962f19e93102e7f5dfe72d976fbf4a73e.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleItemRequestBuilder) GetMemberObjects()(*i4e1a4185ef1644f3671e08895ba2e1b95a54a1c95e280bf43511b3c9807e0daa.GetMemberObjectsRequestBuilder) {
    return i4e1a4185ef1644f3671e08895ba2e1b95a54a1c95e280bf43511b3c9807e0daa.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleItemRequestBuilder) Members()(*iea9aebdcf95599ab8429a0d32f34af8848c1aaf957a501ef2affa3de9d6c4949.MembersRequestBuilder) {
    return iea9aebdcf95599ab8429a0d32f34af8848c1aaf957a501ef2affa3de9d6c4949.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directoryRoles.item.members.item collection
func (m *DirectoryRoleItemRequestBuilder) MembersById(id string)(*iedf3077aa9b00be73cba43b6523f0fed72ea50c39d60527d1af9a3c75c7c4728.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return iedf3077aa9b00be73cba43b6523f0fed72ea50c39d60527d1af9a3c75c7c4728.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in directoryRoles
func (m *DirectoryRoleItemRequestBuilder) Patch(options *DirectoryRoleItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DirectoryRoleItemRequestBuilder) Restore()(*ibe15d8c336cbfe09bcf7d7a61546d490ce45e00077f924a095523cbf08039c76.RestoreRequestBuilder) {
    return ibe15d8c336cbfe09bcf7d7a61546d490ce45e00077f924a095523cbf08039c76.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleItemRequestBuilder) ScopedMembers()(*i5485bb9c491434e8f03f3f928d1e9aae35354496db691be6f7158b8c396ccd9b.ScopedMembersRequestBuilder) {
    return i5485bb9c491434e8f03f3f928d1e9aae35354496db691be6f7158b8c396ccd9b.NewScopedMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directoryRoles.item.scopedMembers.item collection
func (m *DirectoryRoleItemRequestBuilder) ScopedMembersById(id string)(*ia2a3af813d19ede6948e6cd8a4fdc1da146bb512d7519e6200eb556b45d1b28e.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ia2a3af813d19ede6948e6cd8a4fdc1da146bb512d7519e6200eb556b45d1b28e.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

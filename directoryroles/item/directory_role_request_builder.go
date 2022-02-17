package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4e1a4185ef1644f3671e08895ba2e1b95a54a1c95e280bf43511b3c9807e0daa "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/getmemberobjects"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5485bb9c491434e8f03f3f928d1e9aae35354496db691be6f7158b8c396ccd9b "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/scopedmembers"
    i97065f3b11569623118c02d59611043962f19e93102e7f5dfe72d976fbf4a73e "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/getmembergroups"
    i9935b468ea0b8a63a043fea5dd4c071f39943fdd1a429815ff6943d7bf1da7c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/checkmembergroups"
    ibe15d8c336cbfe09bcf7d7a61546d490ce45e00077f924a095523cbf08039c76 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/restore"
    iea9aebdcf95599ab8429a0d32f34af8848c1aaf957a501ef2affa3de9d6c4949 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members"
    if07c44821ea10029e12a10833d0eb1edcf1c37d67dd91ede2fd5a2dad86555f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/checkmemberobjects"
    ia2a3af813d19ede6948e6cd8a4fdc1da146bb512d7519e6200eb556b45d1b28e "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/scopedmembers/item"
)

// DirectoryRoleRequestBuilder builds and executes requests for operations under \directoryRoles\{directoryRole-id}
type DirectoryRoleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRoleRequestBuilderDeleteOptions options for Delete
type DirectoryRoleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleRequestBuilderGetOptions options for Get
type DirectoryRoleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRoleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleRequestBuilderGetQueryParameters get entity from directoryRoles by key
type DirectoryRoleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryRoleRequestBuilderPatchOptions options for Patch
type DirectoryRoleRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRole;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectoryRoleRequestBuilder) CheckMemberGroups()(*i9935b468ea0b8a63a043fea5dd4c071f39943fdd1a429815ff6943d7bf1da7c4.CheckMemberGroupsRequestBuilder) {
    return i9935b468ea0b8a63a043fea5dd4c071f39943fdd1a429815ff6943d7bf1da7c4.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleRequestBuilder) CheckMemberObjects()(*if07c44821ea10029e12a10833d0eb1edcf1c37d67dd91ede2fd5a2dad86555f7.CheckMemberObjectsRequestBuilder) {
    return if07c44821ea10029e12a10833d0eb1edcf1c37d67dd91ede2fd5a2dad86555f7.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRoleRequestBuilderInternal instantiates a new DirectoryRoleRequestBuilder and sets the default values.
func NewDirectoryRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleRequestBuilder) {
    m := &DirectoryRoleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles/{directoryRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRoleRequestBuilder instantiates a new DirectoryRoleRequestBuilder and sets the default values.
func NewDirectoryRoleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directoryRoles
func (m *DirectoryRoleRequestBuilder) CreateDeleteRequestInformation(options *DirectoryRoleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectoryRoleRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectoryRoleRequestBuilder) CreatePatchRequestInformation(options *DirectoryRoleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DirectoryRoleRequestBuilder) Delete(options *DirectoryRoleRequestBuilderDeleteOptions)(error) {
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
// Get get entity from directoryRoles by key
func (m *DirectoryRoleRequestBuilder) Get(options *DirectoryRoleRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRole, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectoryRole() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRole), nil
}
func (m *DirectoryRoleRequestBuilder) GetMemberGroups()(*i97065f3b11569623118c02d59611043962f19e93102e7f5dfe72d976fbf4a73e.GetMemberGroupsRequestBuilder) {
    return i97065f3b11569623118c02d59611043962f19e93102e7f5dfe72d976fbf4a73e.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleRequestBuilder) GetMemberObjects()(*i4e1a4185ef1644f3671e08895ba2e1b95a54a1c95e280bf43511b3c9807e0daa.GetMemberObjectsRequestBuilder) {
    return i4e1a4185ef1644f3671e08895ba2e1b95a54a1c95e280bf43511b3c9807e0daa.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleRequestBuilder) Members()(*iea9aebdcf95599ab8429a0d32f34af8848c1aaf957a501ef2affa3de9d6c4949.MembersRequestBuilder) {
    return iea9aebdcf95599ab8429a0d32f34af8848c1aaf957a501ef2affa3de9d6c4949.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in directoryRoles
func (m *DirectoryRoleRequestBuilder) Patch(options *DirectoryRoleRequestBuilderPatchOptions)(error) {
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
func (m *DirectoryRoleRequestBuilder) Restore()(*ibe15d8c336cbfe09bcf7d7a61546d490ce45e00077f924a095523cbf08039c76.RestoreRequestBuilder) {
    return ibe15d8c336cbfe09bcf7d7a61546d490ce45e00077f924a095523cbf08039c76.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleRequestBuilder) ScopedMembers()(*i5485bb9c491434e8f03f3f928d1e9aae35354496db691be6f7158b8c396ccd9b.ScopedMembersRequestBuilder) {
    return i5485bb9c491434e8f03f3f928d1e9aae35354496db691be6f7158b8c396ccd9b.NewScopedMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directoryRoles.item.scopedMembers.item collection
func (m *DirectoryRoleRequestBuilder) ScopedMembersById(id string)(*ia2a3af813d19ede6948e6cd8a4fdc1da146bb512d7519e6200eb556b45d1b28e.ScopedRoleMembershipRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ia2a3af813d19ede6948e6cd8a4fdc1da146bb512d7519e6200eb556b45d1b28e.NewScopedRoleMembershipRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4e1a4185ef1644f3671e08895ba2e1b95a54a1c95e280bf43511b3c9807e0daa "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/getmemberobjects"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryRoleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleItemRequestBuilderGetQueryParameters retrieve the properties of a directoryRole object. You can use both the object ID and template ID of the **directoryRole** with this API. The template ID of a built-in role is immutable and can be seen in the role description on the Azure portal. For details, see Role template IDs.
type DirectoryRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleItemRequestBuilderGetQueryParameters
}
// DirectoryRoleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *DirectoryRoleItemRequestBuilder) CheckMemberGroups()(*i9935b468ea0b8a63a043fea5dd4c071f39943fdd1a429815ff6943d7bf1da7c4.CheckMemberGroupsRequestBuilder) {
    return i9935b468ea0b8a63a043fea5dd4c071f39943fdd1a429815ff6943d7bf1da7c4.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *DirectoryRoleItemRequestBuilder) CheckMemberObjects()(*if07c44821ea10029e12a10833d0eb1edcf1c37d67dd91ede2fd5a2dad86555f7.CheckMemberObjectsRequestBuilder) {
    return if07c44821ea10029e12a10833d0eb1edcf1c37d67dd91ede2fd5a2dad86555f7.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRoleItemRequestBuilderInternal instantiates a new DirectoryRoleItemRequestBuilder and sets the default values.
func NewDirectoryRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleItemRequestBuilder) {
    m := &DirectoryRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles/{directoryRole%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRoleItemRequestBuilder instantiates a new DirectoryRoleItemRequestBuilder and sets the default values.
func NewDirectoryRoleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directoryRoles
func (m *DirectoryRoleItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation retrieve the properties of a directoryRole object. You can use both the object ID and template ID of the **directoryRole** with this API. The template ID of a built-in role is immutable and can be seen in the role description on the Azure portal. For details, see Role template IDs.
func (m *DirectoryRoleItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in directoryRoles
func (m *DirectoryRoleItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable, requestConfiguration *DirectoryRoleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete entity from directoryRoles
func (m *DirectoryRoleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve the properties of a directoryRole object. You can use both the object ID and template ID of the **directoryRole** with this API. The template ID of a built-in role is immutable and can be seen in the role description on the Azure portal. For details, see Role template IDs.
func (m *DirectoryRoleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *DirectoryRoleItemRequestBuilder) GetMemberGroups()(*i97065f3b11569623118c02d59611043962f19e93102e7f5dfe72d976fbf4a73e.GetMemberGroupsRequestBuilder) {
    return i97065f3b11569623118c02d59611043962f19e93102e7f5dfe72d976fbf4a73e.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *DirectoryRoleItemRequestBuilder) GetMemberObjects()(*i4e1a4185ef1644f3671e08895ba2e1b95a54a1c95e280bf43511b3c9807e0daa.GetMemberObjectsRequestBuilder) {
    return i4e1a4185ef1644f3671e08895ba2e1b95a54a1c95e280bf43511b3c9807e0daa.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.directoryRole entity.
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
        urlTplParams["directoryObject%2Did"] = id
    }
    return iedf3077aa9b00be73cba43b6523f0fed72ea50c39d60527d1af9a3c75c7c4728.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in directoryRoles
func (m *DirectoryRoleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable, requestConfiguration *DirectoryRoleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable), nil
}
// Restore provides operations to call the restore method.
func (m *DirectoryRoleItemRequestBuilder) Restore()(*ibe15d8c336cbfe09bcf7d7a61546d490ce45e00077f924a095523cbf08039c76.RestoreRequestBuilder) {
    return ibe15d8c336cbfe09bcf7d7a61546d490ce45e00077f924a095523cbf08039c76.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedMembers provides operations to manage the scopedMembers property of the microsoft.graph.directoryRole entity.
func (m *DirectoryRoleItemRequestBuilder) ScopedMembers()(*i5485bb9c491434e8f03f3f928d1e9aae35354496db691be6f7158b8c396ccd9b.ScopedMembersRequestBuilder) {
    return i5485bb9c491434e8f03f3f928d1e9aae35354496db691be6f7158b8c396ccd9b.NewScopedMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedMembersById provides operations to manage the scopedMembers property of the microsoft.graph.directoryRole entity.
func (m *DirectoryRoleItemRequestBuilder) ScopedMembersById(id string)(*ia2a3af813d19ede6948e6cd8a4fdc1da146bb512d7519e6200eb556b45d1b28e.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return ia2a3af813d19ede6948e6cd8a4fdc1da146bb512d7519e6200eb556b45d1b28e.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}

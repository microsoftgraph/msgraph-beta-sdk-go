package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i150aab97f1f2524e3b90b58049a38acb3ea567fd7b50e2041780889dfdc331f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item/checkmembergroups"
    i6dae04482805de37a04879751a17cc7796dc3d76cf9eaaf38d9c9786d2af4e42 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item/checkmemberobjects"
    i7fcd18b152b0bf6522aa91647cd77f6b35cf0b1ed498a9e97e6248e5339052f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item/restore"
    ia3fa7372a45c09ac1b7b34eba9e20dc374b9817cffc1e34eea6afa985a7355cc "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item/group"
    ibb9768f289fbadec5f87cde24e363e0f2c3556f5ff9d970941d716894c9e54b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item/application"
    iecc50f59642d3ded03c6cdda66cc48a3fd034beb1a00070995aa7a5d9202ec63 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item/getmemberobjects"
    ifa902ae6cdb0d6db6dfb06b83126cf372c36769e3a03f598113ca40094f4a8a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item/getmembergroups"
    ifca3b2702985940d02d9e6addc2729a9cd9203e3486ed30cd6ac6c7b8fe62a10 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/item/user"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the deletedItems property of the microsoft.graph.directory entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryObjectItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryObjectItemRequestBuilderGetQueryParameters get deletedItems from directory
type DirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryObjectItemRequestBuilderGetQueryParameters
}
// DirectoryObjectItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryObjectItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Application the application property
func (m *DirectoryObjectItemRequestBuilder) Application()(*ibb9768f289fbadec5f87cde24e363e0f2c3556f5ff9d970941d716894c9e54b3.ApplicationRequestBuilder) {
    return ibb9768f289fbadec5f87cde24e363e0f2c3556f5ff9d970941d716894c9e54b3.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *DirectoryObjectItemRequestBuilder) CheckMemberGroups()(*i150aab97f1f2524e3b90b58049a38acb3ea567fd7b50e2041780889dfdc331f5.CheckMemberGroupsRequestBuilder) {
    return i150aab97f1f2524e3b90b58049a38acb3ea567fd7b50e2041780889dfdc331f5.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *DirectoryObjectItemRequestBuilder) CheckMemberObjects()(*i6dae04482805de37a04879751a17cc7796dc3d76cf9eaaf38d9c9786d2af4e42.CheckMemberObjectsRequestBuilder) {
    return i6dae04482805de37a04879751a17cc7796dc3d76cf9eaaf38d9c9786d2af4e42.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/deletedItems/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deletedItems for directory
func (m *DirectoryObjectItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property deletedItems for directory
func (m *DirectoryObjectItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryObjectItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get deletedItems from directory
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get deletedItems from directory
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deletedItems in directory
func (m *DirectoryObjectItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property deletedItems in directory
func (m *DirectoryObjectItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *DirectoryObjectItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property deletedItems for directory
func (m *DirectoryObjectItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryObjectItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get get deletedItems from directory
func (m *DirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// GetMemberGroups the getMemberGroups property
func (m *DirectoryObjectItemRequestBuilder) GetMemberGroups()(*ifa902ae6cdb0d6db6dfb06b83126cf372c36769e3a03f598113ca40094f4a8a5.GetMemberGroupsRequestBuilder) {
    return ifa902ae6cdb0d6db6dfb06b83126cf372c36769e3a03f598113ca40094f4a8a5.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *DirectoryObjectItemRequestBuilder) GetMemberObjects()(*iecc50f59642d3ded03c6cdda66cc48a3fd034beb1a00070995aa7a5d9202ec63.GetMemberObjectsRequestBuilder) {
    return iecc50f59642d3ded03c6cdda66cc48a3fd034beb1a00070995aa7a5d9202ec63.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group the group property
func (m *DirectoryObjectItemRequestBuilder) Group()(*ia3fa7372a45c09ac1b7b34eba9e20dc374b9817cffc1e34eea6afa985a7355cc.GroupRequestBuilder) {
    return ia3fa7372a45c09ac1b7b34eba9e20dc374b9817cffc1e34eea6afa985a7355cc.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property deletedItems in directory
func (m *DirectoryObjectItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *DirectoryObjectItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Restore the restore property
func (m *DirectoryObjectItemRequestBuilder) Restore()(*i7fcd18b152b0bf6522aa91647cd77f6b35cf0b1ed498a9e97e6248e5339052f2.RestoreRequestBuilder) {
    return i7fcd18b152b0bf6522aa91647cd77f6b35cf0b1ed498a9e97e6248e5339052f2.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*ifca3b2702985940d02d9e6addc2729a9cd9203e3486ed30cd6ac6c7b8fe62a10.UserRequestBuilder) {
    return ifca3b2702985940d02d9e6addc2729a9cd9203e3486ed30cd6ac6c7b8fe62a10.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

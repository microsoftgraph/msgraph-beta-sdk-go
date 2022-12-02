package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder provides operations to manage the permissionGrants property of the microsoft.graph.chat entity.
type UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters a collection of permissions granted to apps for the chat.
type UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters
}
// UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberGroups()(*UsersItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder) {
    return NewUsersItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberObjects()(*UsersItemChatsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder) {
    return NewUsersItemChatsItemPermissionGrantsItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewUsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    m := &UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/permissionGrants/{resourceSpecificPermissionGrant%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewUsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property permissionGrants for users
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of permissions granted to apps for the chat.
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property permissionGrants in users
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property permissionGrants for users
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of permissions granted to apps for the chat.
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) GetMemberGroups()(*UsersItemChatsItemPermissionGrantsItemGetMemberGroupsRequestBuilder) {
    return NewUsersItemChatsItemPermissionGrantsItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) GetMemberObjects()(*UsersItemChatsItemPermissionGrantsItemGetMemberObjectsRequestBuilder) {
    return NewUsersItemChatsItemPermissionGrantsItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property permissionGrants in users
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable), nil
}
// Restore provides operations to call the restore method.
func (m *UsersItemChatsItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Restore()(*UsersItemChatsItemPermissionGrantsItemRestoreRequestBuilder) {
    return NewUsersItemChatsItemPermissionGrantsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

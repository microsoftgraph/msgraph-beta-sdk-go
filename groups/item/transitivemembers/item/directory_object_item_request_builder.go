package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b8fa3b14cd8d9e4cbccaae3f2a86e1254e84ba0390728ce44644002bb819078 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/orgcontact"
    i2fb901df3b5906e3721f8321201356cac6bfbd696d2681a2a287c6f6a008b65a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/user"
    i42287541979c5d1b3785cfd503f5d0cdcd56ccef2727b31ab03aff058325da8b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/serviceprincipal"
    i44b62aae1ae9d8af12e30a87d8a61b906e15e67787ce130a8b3a4b5cd4d1f7fd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group"
    i54bef4a77ade54d07f646d6d1f880255049667b0ec08215be63f2eb472076925 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/application"
    i662d4f7ee7a89e38f452309e6041c6ceb88e921531be901af2a981900ae36131 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/device"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderGetQueryParameters the direct and transitive members of a group. Nullable.
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
// Application casts the previous resource to application.
func (m *DirectoryObjectItemRequestBuilder) Application()(*i54bef4a77ade54d07f646d6d1f880255049667b0ec08215be63f2eb472076925.ApplicationRequestBuilder) {
    return i54bef4a77ade54d07f646d6d1f880255049667b0ec08215be63f2eb472076925.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/transitiveMembers/{directoryObject%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation the direct and transitive members of a group. Nullable.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Device casts the previous resource to device.
func (m *DirectoryObjectItemRequestBuilder) Device()(*i662d4f7ee7a89e38f452309e6041c6ceb88e921531be901af2a981900ae36131.DeviceRequestBuilder) {
    return i662d4f7ee7a89e38f452309e6041c6ceb88e921531be901af2a981900ae36131.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the direct and transitive members of a group. Nullable.
func (m *DirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
// Group casts the previous resource to group.
func (m *DirectoryObjectItemRequestBuilder) Group()(*i44b62aae1ae9d8af12e30a87d8a61b906e15e67787ce130a8b3a4b5cd4d1f7fd.GroupRequestBuilder) {
    return i44b62aae1ae9d8af12e30a87d8a61b906e15e67787ce130a8b3a4b5cd4d1f7fd.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*i0b8fa3b14cd8d9e4cbccaae3f2a86e1254e84ba0390728ce44644002bb819078.OrgContactRequestBuilder) {
    return i0b8fa3b14cd8d9e4cbccaae3f2a86e1254e84ba0390728ce44644002bb819078.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i42287541979c5d1b3785cfd503f5d0cdcd56ccef2727b31ab03aff058325da8b.ServicePrincipalRequestBuilder) {
    return i42287541979c5d1b3785cfd503f5d0cdcd56ccef2727b31ab03aff058325da8b.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryObjectItemRequestBuilder) User()(*i2fb901df3b5906e3721f8321201356cac6bfbd696d2681a2a287c6f6a008b65a.UserRequestBuilder) {
    return i2fb901df3b5906e3721f8321201356cac6bfbd696d2681a2a287c6f6a008b65a.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

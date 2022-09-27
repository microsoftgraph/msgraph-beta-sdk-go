package rolescopetags

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i33808f02abca01e1267c66ae1677f591fa5a4d6e6c99557066e7ad9ab5bdeb72 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/rolescopetags/getrolescopetagsbyid"
    i52487895386ac5b9b67d47fdade3472c4d5e919e1c2bc876ca92efd882afe9ae "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/rolescopetags/hascustomrolescopetag"
    i87e24dbe25af599aa1e532c37a185b0fc3836b1f20d55a440891f8edc915a488 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/rolescopetags/count"
)

// RoleScopeTagsRequestBuilder provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
type RoleScopeTagsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleScopeTagsRequestBuilderGetQueryParameters the Role Scope Tags.
type RoleScopeTagsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// RoleScopeTagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleScopeTagsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleScopeTagsRequestBuilderGetQueryParameters
}
// RoleScopeTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleScopeTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleScopeTagsRequestBuilderInternal instantiates a new RoleScopeTagsRequestBuilder and sets the default values.
func NewRoleScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleScopeTagsRequestBuilder) {
    m := &RoleScopeTagsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/roleScopeTags{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleScopeTagsRequestBuilder instantiates a new RoleScopeTagsRequestBuilder and sets the default values.
func NewRoleScopeTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *RoleScopeTagsRequestBuilder) Count()(*i87e24dbe25af599aa1e532c37a185b0fc3836b1f20d55a440891f8edc915a488.CountRequestBuilder) {
    return i87e24dbe25af599aa1e532c37a185b0fc3836b1f20d55a440891f8edc915a488.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the Role Scope Tags.
func (m *RoleScopeTagsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the Role Scope Tags.
func (m *RoleScopeTagsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RoleScopeTagsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to roleScopeTags for deviceManagement
func (m *RoleScopeTagsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to roleScopeTags for deviceManagement
func (m *RoleScopeTagsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable, requestConfiguration *RoleScopeTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the Role Scope Tags.
func (m *RoleScopeTagsRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleScopeTagsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoleScopeTagCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagCollectionResponseable), nil
}
// GetRoleScopeTagsById the getRoleScopeTagsById property
func (m *RoleScopeTagsRequestBuilder) GetRoleScopeTagsById()(*i33808f02abca01e1267c66ae1677f591fa5a4d6e6c99557066e7ad9ab5bdeb72.GetRoleScopeTagsByIdRequestBuilder) {
    return i33808f02abca01e1267c66ae1677f591fa5a4d6e6c99557066e7ad9ab5bdeb72.NewGetRoleScopeTagsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HasCustomRoleScopeTag provides operations to call the hasCustomRoleScopeTag method.
func (m *RoleScopeTagsRequestBuilder) HasCustomRoleScopeTag()(*i52487895386ac5b9b67d47fdade3472c4d5e919e1c2bc876ca92efd882afe9ae.HasCustomRoleScopeTagRequestBuilder) {
    return i52487895386ac5b9b67d47fdade3472c4d5e919e1c2bc876ca92efd882afe9ae.NewHasCustomRoleScopeTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to roleScopeTags for deviceManagement
func (m *RoleScopeTagsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable, requestConfiguration *RoleScopeTagsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoleScopeTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable), nil
}

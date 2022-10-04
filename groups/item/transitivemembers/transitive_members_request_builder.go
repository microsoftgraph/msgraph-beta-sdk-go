package transitivemembers

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09c8ec187be02eb8c582ed81e1e61003595821ab2bbae63fbe3cc6f03af340cb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/count"
    i0bfb2f159df03e765a70621592a9bac6bea666cafcacbe7360925a51c4aa4035 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/orgcontact"
    i253d4eefafb34c1d668f6ffcb8eb7b397cff8f52d64cd1018c52c4208fe7f416 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/user"
    i6c2186b52272f4e15e16e8e4ed70b613d91134b9d1c00d170d8c9c5667e23cb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/group"
    i798ff338a423619b5a9f4a247c449578aae41c4ba79b20951f1390abfb3fae67 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/device"
    i98ed2293d1a8873eab4be0f1d527226547fd88f6bead39b1ad64c8ccc0ddffa1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/serviceprincipal"
    if86ba66b7b16edf432a08f35e77e5db696d3bab06d5d99c69d0d5b251d310f3a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/application"
)

// TransitiveMembersRequestBuilder provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
type TransitiveMembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TransitiveMembersRequestBuilderGetQueryParameters the direct and transitive members of a group. Nullable.
type TransitiveMembersRequestBuilderGetQueryParameters struct {
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
// TransitiveMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransitiveMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TransitiveMembersRequestBuilderGetQueryParameters
}
// Application the application property
func (m *TransitiveMembersRequestBuilder) Application()(*if86ba66b7b16edf432a08f35e77e5db696d3bab06d5d99c69d0d5b251d310f3a.ApplicationRequestBuilder) {
    return if86ba66b7b16edf432a08f35e77e5db696d3bab06d5d99c69d0d5b251d310f3a.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTransitiveMembersRequestBuilderInternal instantiates a new TransitiveMembersRequestBuilder and sets the default values.
func NewTransitiveMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMembersRequestBuilder) {
    m := &TransitiveMembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/transitiveMembers{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTransitiveMembersRequestBuilder instantiates a new TransitiveMembersRequestBuilder and sets the default values.
func NewTransitiveMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransitiveMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *TransitiveMembersRequestBuilder) Count()(*i09c8ec187be02eb8c582ed81e1e61003595821ab2bbae63fbe3cc6f03af340cb.CountRequestBuilder) {
    return i09c8ec187be02eb8c582ed81e1e61003595821ab2bbae63fbe3cc6f03af340cb.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the direct and transitive members of a group. Nullable.
func (m *TransitiveMembersRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TransitiveMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Device the device property
func (m *TransitiveMembersRequestBuilder) Device()(*i798ff338a423619b5a9f4a247c449578aae41c4ba79b20951f1390abfb3fae67.DeviceRequestBuilder) {
    return i798ff338a423619b5a9f4a247c449578aae41c4ba79b20951f1390abfb3fae67.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the direct and transitive members of a group. Nullable.
func (m *TransitiveMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *TransitiveMembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Group the group property
func (m *TransitiveMembersRequestBuilder) Group()(*i6c2186b52272f4e15e16e8e4ed70b613d91134b9d1c00d170d8c9c5667e23cb2.GroupRequestBuilder) {
    return i6c2186b52272f4e15e16e8e4ed70b613d91134b9d1c00d170d8c9c5667e23cb2.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *TransitiveMembersRequestBuilder) OrgContact()(*i0bfb2f159df03e765a70621592a9bac6bea666cafcacbe7360925a51c4aa4035.OrgContactRequestBuilder) {
    return i0bfb2f159df03e765a70621592a9bac6bea666cafcacbe7360925a51c4aa4035.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *TransitiveMembersRequestBuilder) ServicePrincipal()(*i98ed2293d1a8873eab4be0f1d527226547fd88f6bead39b1ad64c8ccc0ddffa1.ServicePrincipalRequestBuilder) {
    return i98ed2293d1a8873eab4be0f1d527226547fd88f6bead39b1ad64c8ccc0ddffa1.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *TransitiveMembersRequestBuilder) User()(*i253d4eefafb34c1d668f6ffcb8eb7b397cff8f52d64cd1018c52c4208fe7f416.UserRequestBuilder) {
    return i253d4eefafb34c1d668f6ffcb8eb7b397cff8f52d64cd1018c52c4208fe7f416.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

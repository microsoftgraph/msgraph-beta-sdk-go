package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43f566393f76360ac3ca287c870ef2baf5787c60c6f468f416c835d802571cf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/device"
    i4431c64c4893640baf9ae4bd923f21f993b2c5a91a48a0ebdc2777c5c5df5818 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal"
    i69827f845da09907f3b68791eb586821c7951d1a1c2521a0264005f3c3a5d2d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/orgcontact"
    i9f634b6a60b2cb11ab9275f9284ef2a197be590be48a4c07b49a82c72bfdb38c "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/user"
    ia9fe0981135f75615d85ee5338b1a7b57df19457f228494d1bcd8ba24bd6e9ea "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/group"
    ibd79a3ed447e1cf322e69ff53ac0105e979ac3381b2d12dc8df9f7574d5e2fc9 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/application"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderGetQueryParameters get transitiveMemberOf from servicePrincipals
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
// Application the application property
func (m *DirectoryObjectItemRequestBuilder) Application()(*ibd79a3ed447e1cf322e69ff53ac0105e979ac3381b2d12dc8df9f7574d5e2fc9.ApplicationRequestBuilder) {
    return ibd79a3ed447e1cf322e69ff53ac0105e979ac3381b2d12dc8df9f7574d5e2fc9.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/transitiveMemberOf/{directoryObject%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation get transitiveMemberOf from servicePrincipals
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get transitiveMemberOf from servicePrincipals
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
// Device the device property
func (m *DirectoryObjectItemRequestBuilder) Device()(*i43f566393f76360ac3ca287c870ef2baf5787c60c6f468f416c835d802571cf7.DeviceRequestBuilder) {
    return i43f566393f76360ac3ca287c870ef2baf5787c60c6f468f416c835d802571cf7.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get transitiveMemberOf from servicePrincipals
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
// Group the group property
func (m *DirectoryObjectItemRequestBuilder) Group()(*ia9fe0981135f75615d85ee5338b1a7b57df19457f228494d1bcd8ba24bd6e9ea.GroupRequestBuilder) {
    return ia9fe0981135f75615d85ee5338b1a7b57df19457f228494d1bcd8ba24bd6e9ea.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*i69827f845da09907f3b68791eb586821c7951d1a1c2521a0264005f3c3a5d2d7.OrgContactRequestBuilder) {
    return i69827f845da09907f3b68791eb586821c7951d1a1c2521a0264005f3c3a5d2d7.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i4431c64c4893640baf9ae4bd923f21f993b2c5a91a48a0ebdc2777c5c5df5818.ServicePrincipalRequestBuilder) {
    return i4431c64c4893640baf9ae4bd923f21f993b2c5a91a48a0ebdc2777c5c5df5818.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i9f634b6a60b2cb11ab9275f9284ef2a197be590be48a4c07b49a82c72bfdb38c.UserRequestBuilder) {
    return i9f634b6a60b2cb11ab9275f9284ef2a197be590be48a4c07b49a82c72bfdb38c.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

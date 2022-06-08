package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0963c5970c78866023651991cee5002ef406e2c64de627cb0683f5436086a03f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/application"
    i55b9076cdd5c40ba0f1f6be0681996eff040223a12d5dd7c1e2a70b1f4a5d491 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group"
    i9931190ada5754ea3041db4c9a2290d28cdba4ad809a3434515616fc63fd1445 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/device"
    icec4b12626b2f1a274d430e4a4b2037a6faa73fc2d9b18e5240d64732685d946 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/user"
    ie8b1ae18725068bfc936408064d4de65b52ef6c67f5209d8b03e6e00a4a2a8c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/orgcontact"
    if6f8523e578c4ce23b8a7c58e1f6eb28e772e77d324752c940daa1d4134c4d09 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/serviceprincipal"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.group entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderGetQueryParameters get transitiveMemberOf from groups
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
func (m *DirectoryObjectItemRequestBuilder) Application()(*i0963c5970c78866023651991cee5002ef406e2c64de627cb0683f5436086a03f.ApplicationRequestBuilder) {
    return i0963c5970c78866023651991cee5002ef406e2c64de627cb0683f5436086a03f.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/transitiveMemberOf/{directoryObject%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation get transitiveMemberOf from groups
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get transitiveMemberOf from groups
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
func (m *DirectoryObjectItemRequestBuilder) Device()(*i9931190ada5754ea3041db4c9a2290d28cdba4ad809a3434515616fc63fd1445.DeviceRequestBuilder) {
    return i9931190ada5754ea3041db4c9a2290d28cdba4ad809a3434515616fc63fd1445.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get transitiveMemberOf from groups
func (m *DirectoryObjectItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get transitiveMemberOf from groups
func (m *DirectoryObjectItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Group the group property
func (m *DirectoryObjectItemRequestBuilder) Group()(*i55b9076cdd5c40ba0f1f6be0681996eff040223a12d5dd7c1e2a70b1f4a5d491.GroupRequestBuilder) {
    return i55b9076cdd5c40ba0f1f6be0681996eff040223a12d5dd7c1e2a70b1f4a5d491.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*ie8b1ae18725068bfc936408064d4de65b52ef6c67f5209d8b03e6e00a4a2a8c5.OrgContactRequestBuilder) {
    return ie8b1ae18725068bfc936408064d4de65b52ef6c67f5209d8b03e6e00a4a2a8c5.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*if6f8523e578c4ce23b8a7c58e1f6eb28e772e77d324752c940daa1d4134c4d09.ServicePrincipalRequestBuilder) {
    return if6f8523e578c4ce23b8a7c58e1f6eb28e772e77d324752c940daa1d4134c4d09.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*icec4b12626b2f1a274d430e4a4b2037a6faa73fc2d9b18e5240d64732685d946.UserRequestBuilder) {
    return icec4b12626b2f1a274d430e4a4b2037a6faa73fc2d9b18e5240d64732685d946.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

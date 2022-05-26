package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07ad8c96151bd2d31d2a0f0b6b533e2a4373be7e802e519d23e6d5155f50a542 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/orgcontact"
    i08face1df25225fe35a15fab2bd5299df1e5400f1cb53516290120c10462fff8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/user"
    i29477516fb8946699fd4c8b23cae038591b8ae2c1dbd18e62f579f8efb03e18d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/application"
    i587651bf33bf8f7d0e6d38887d6ec9865ebf9ffd01583543e61946e7fc35ffab "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/serviceprincipal"
    i7f11a0bab32915a6fe1b5198246a7ffa94bb1da190905b37a237f34c5013a4f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group"
    ie3c23faa9e8ae0b05ebca4c7d827c1c674325bda02bac619984a70c03127e26d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/device"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderGetQueryParameters a list of group members with license errors from this group-based license assignment. Read-only.
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
func (m *DirectoryObjectItemRequestBuilder) Application()(*i29477516fb8946699fd4c8b23cae038591b8ae2c1dbd18e62f579f8efb03e18d.ApplicationRequestBuilder) {
    return i29477516fb8946699fd4c8b23cae038591b8ae2c1dbd18e62f579f8efb03e18d.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation a list of group members with license errors from this group-based license assignment. Read-only.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration a list of group members with license errors from this group-based license assignment. Read-only.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
func (m *DirectoryObjectItemRequestBuilder) Device()(*ie3c23faa9e8ae0b05ebca4c7d827c1c674325bda02bac619984a70c03127e26d.DeviceRequestBuilder) {
    return ie3c23faa9e8ae0b05ebca4c7d827c1c674325bda02bac619984a70c03127e26d.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a list of group members with license errors from this group-based license assignment. Read-only.
func (m *DirectoryObjectItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler a list of group members with license errors from this group-based license assignment. Read-only.
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
func (m *DirectoryObjectItemRequestBuilder) Group()(*i7f11a0bab32915a6fe1b5198246a7ffa94bb1da190905b37a237f34c5013a4f4.GroupRequestBuilder) {
    return i7f11a0bab32915a6fe1b5198246a7ffa94bb1da190905b37a237f34c5013a4f4.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*i07ad8c96151bd2d31d2a0f0b6b533e2a4373be7e802e519d23e6d5155f50a542.OrgContactRequestBuilder) {
    return i07ad8c96151bd2d31d2a0f0b6b533e2a4373be7e802e519d23e6d5155f50a542.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i587651bf33bf8f7d0e6d38887d6ec9865ebf9ffd01583543e61946e7fc35ffab.ServicePrincipalRequestBuilder) {
    return i587651bf33bf8f7d0e6d38887d6ec9865ebf9ffd01583543e61946e7fc35ffab.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i08face1df25225fe35a15fab2bd5299df1e5400f1cb53516290120c10462fff8.UserRequestBuilder) {
    return i08face1df25225fe35a15fab2bd5299df1e5400f1cb53516290120c10462fff8.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

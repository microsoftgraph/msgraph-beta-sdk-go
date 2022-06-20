package roleassignments

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i334c41831d8aa39d2eb97c99a39e2e5d732b6dcb80b9bca493ed07074555542a "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignments/count"
    ib43f9c9c96e8803e0a6dcc7f505e321e14b49389f81434fc0940b1e55fe32164 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignments/export"
)

// RoleAssignmentsRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.privilegedAccess entity.
type RoleAssignmentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleAssignmentsRequestBuilderGetQueryParameters a collection of role assignments for the provider.
type RoleAssignmentsRequestBuilderGetQueryParameters struct {
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
// RoleAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleAssignmentsRequestBuilderGetQueryParameters
}
// RoleAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleAssignmentsRequestBuilderInternal instantiates a new RoleAssignmentsRequestBuilder and sets the default values.
func NewRoleAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleAssignmentsRequestBuilder) {
    m := &RoleAssignmentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/roleAssignments{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleAssignmentsRequestBuilder instantiates a new RoleAssignmentsRequestBuilder and sets the default values.
func NewRoleAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *RoleAssignmentsRequestBuilder) Count()(*i334c41831d8aa39d2eb97c99a39e2e5d732b6dcb80b9bca493ed07074555542a.CountRequestBuilder) {
    return i334c41831d8aa39d2eb97c99a39e2e5d732b6dcb80b9bca493ed07074555542a.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation a collection of role assignments for the provider.
func (m *RoleAssignmentsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration a collection of role assignments for the provider.
func (m *RoleAssignmentsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RoleAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to roleAssignments for privilegedAccess
func (m *RoleAssignmentsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to roleAssignments for privilegedAccess
func (m *RoleAssignmentsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, requestConfiguration *RoleAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Export provides operations to call the export method.
func (m *RoleAssignmentsRequestBuilder) Export()(*ib43f9c9c96e8803e0a6dcc7f505e321e14b49389f81434fc0940b1e55fe32164.ExportRequestBuilder) {
    return ib43f9c9c96e8803e0a6dcc7f505e321e14b49389f81434fc0940b1e55fe32164.NewExportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a collection of role assignments for the provider.
func (m *RoleAssignmentsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler a collection of role assignments for the provider.
func (m *RoleAssignmentsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RoleAssignmentsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentCollectionResponseable), nil
}
// Post create new navigation property to roleAssignments for privilegedAccess
func (m *RoleAssignmentsRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to roleAssignments for privilegedAccess
func (m *RoleAssignmentsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, requestConfiguration *RoleAssignmentsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentable), nil
}

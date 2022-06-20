package accesspackageresource

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i897c2fba433b27857b779c277bc3368d266f4ff52b1c90a243dfdee2b48155bb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageresourcescope/accesspackageresource/accesspackageresourcescopes"
    ic61459b7f69db0087bd52939220c7d5841e85378bef774f58268405210aaeee9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageresourcescope/accesspackageresource/accesspackageresourceenvironment"
    ic8cec4a9da711f4e1bbb3ab6394870d2f3a28d878b2d43dc7065cf71750ba481 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageresourcescope/accesspackageresource/accesspackageresourceroles"
    i7c7430fabee6617f9ee7baebdda9a390f1aadf600f3f871ba48cf6655d673476 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageresourcescope/accesspackageresource/accesspackageresourcescopes/item"
    i7e57e899523d78cfbdb5fdd38865a512f62da0c51bfa42f069a0903560447a46 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageresourcescope/accesspackageresource/accesspackageresourceroles/item"
)

// AccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceScope entity.
type AccessPackageResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessPackageResourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageResourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceRequestBuilderGetQueryParameters get accessPackageResource from identityGovernance
type AccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessPackageResourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageResourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessPackageResourceRequestBuilderGetQueryParameters
}
// AccessPackageResourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageResourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageResourceEnvironment the accessPackageResourceEnvironment property
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*ic61459b7f69db0087bd52939220c7d5841e85378bef774f58268405210aaeee9.AccessPackageResourceEnvironmentRequestBuilder) {
    return ic61459b7f69db0087bd52939220c7d5841e85378bef774f58268405210aaeee9.NewAccessPackageResourceEnvironmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoles the accessPackageResourceRoles property
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRoles()(*ic8cec4a9da711f4e1bbb3ab6394870d2f3a28d878b2d43dc7065cf71750ba481.AccessPackageResourceRolesRequestBuilder) {
    return ic8cec4a9da711f4e1bbb3ab6394870d2f3a28d878b2d43dc7065cf71750ba481.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageResourceScope.accessPackageResource.accessPackageResourceRoles.item collection
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRolesById(id string)(*i7e57e899523d78cfbdb5fdd38865a512f62da0c51bfa42f069a0903560447a46.AccessPackageResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole%2Did"] = id
    }
    return i7e57e899523d78cfbdb5fdd38865a512f62da0c51bfa42f069a0903560447a46.NewAccessPackageResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceScopes the accessPackageResourceScopes property
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*i897c2fba433b27857b779c277bc3368d266f4ff52b1c90a243dfdee2b48155bb.AccessPackageResourceScopesRequestBuilder) {
    return i897c2fba433b27857b779c277bc3368d266f4ff52b1c90a243dfdee2b48155bb.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageResourceScope.accessPackageResource.accessPackageResourceScopes.item collection
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopesById(id string)(*i7c7430fabee6617f9ee7baebdda9a390f1aadf600f3f871ba48cf6655d673476.AccessPackageResourceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope%2Did"] = id
    }
    return i7c7430fabee6617f9ee7baebdda9a390f1aadf600f3f871ba48cf6655d673476.NewAccessPackageResourceScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageResourceRequestBuilderInternal instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    m := &AccessPackageResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageResourceScope/accessPackageResource{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageResourceRequestBuilder instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageResource for identityGovernance
func (m *AccessPackageResourceRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property accessPackageResource for identityGovernance
func (m *AccessPackageResourceRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AccessPackageResourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get accessPackageResource from identityGovernance
func (m *AccessPackageResourceRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get accessPackageResource from identityGovernance
func (m *AccessPackageResourceRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AccessPackageResourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageResource in identityGovernance
func (m *AccessPackageResourceRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property accessPackageResource in identityGovernance
func (m *AccessPackageResourceRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *AccessPackageResourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property accessPackageResource for identityGovernance
func (m *AccessPackageResourceRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property accessPackageResource for identityGovernance
func (m *AccessPackageResourceRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *AccessPackageResourceRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get accessPackageResource from identityGovernance
func (m *AccessPackageResourceRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get accessPackageResource from identityGovernance
func (m *AccessPackageResourceRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AccessPackageResourceRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable), nil
}
// Patch update the navigation property accessPackageResource in identityGovernance
func (m *AccessPackageResourceRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property accessPackageResource in identityGovernance
func (m *AccessPackageResourceRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, requestConfiguration *AccessPackageResourceRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}

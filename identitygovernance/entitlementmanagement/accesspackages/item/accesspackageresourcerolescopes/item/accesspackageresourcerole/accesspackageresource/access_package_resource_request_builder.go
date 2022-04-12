package accesspackageresource

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i7cdde6386682a7b30d1897e2f5b10b1dba49e35085058d46b9e97030e955dc08 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageresourcerolescopes/item/accesspackageresourcerole/accesspackageresource/accesspackageresourcescopes"
    i8d0f38c14939e12933c77e7ed82bae427a73a08d276a308a0399c054513d3ce3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageresourcerolescopes/item/accesspackageresourcerole/accesspackageresource/accesspackageresourceroles"
    ife7445bfdce7d42f048f3698b718854bc535e4daa00bba4bc176b533f45a7077 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageresourcerolescopes/item/accesspackageresourcerole/accesspackageresource/accesspackageresourceenvironment"
    i95bd789c6b6d638c6d9e88a9ce6120e1cc2d851f19b94543af06cdffa7aab1f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageresourcerolescopes/item/accesspackageresourcerole/accesspackageresource/accesspackageresourceroles/item"
    i967b53f70b934e90a12f7597314710840dfde37afa761e574abb6ed700e05925 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageresourcerolescopes/item/accesspackageresourcerole/accesspackageresource/accesspackageresourcescopes/item"
)

// AccessPackageResourceRequestBuilder provides operations to manage the accessPackageResource property of the microsoft.graph.accessPackageResourceRole entity.
type AccessPackageResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessPackageResourceRequestBuilderDeleteOptions options for Delete
type AccessPackageResourceRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AccessPackageResourceRequestBuilderGetOptions options for Get
type AccessPackageResourceRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessPackageResourceRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AccessPackageResourceRequestBuilderGetQueryParameters read-only. Nullable.
type AccessPackageResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessPackageResourceRequestBuilderPatchOptions options for Patch
type AccessPackageResourceRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AccessPackageResourceEnvironment the accessPackageResourceEnvironment property
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*ife7445bfdce7d42f048f3698b718854bc535e4daa00bba4bc176b533f45a7077.AccessPackageResourceEnvironmentRequestBuilder) {
    return ife7445bfdce7d42f048f3698b718854bc535e4daa00bba4bc176b533f45a7077.NewAccessPackageResourceEnvironmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoles the accessPackageResourceRoles property
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRoles()(*i8d0f38c14939e12933c77e7ed82bae427a73a08d276a308a0399c054513d3ce3.AccessPackageResourceRolesRequestBuilder) {
    return i8d0f38c14939e12933c77e7ed82bae427a73a08d276a308a0399c054513d3ce3.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.accessPackageResourceRoleScopes.item.accessPackageResourceRole.accessPackageResource.accessPackageResourceRoles.item collection
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRolesById(id string)(*i95bd789c6b6d638c6d9e88a9ce6120e1cc2d851f19b94543af06cdffa7aab1f4.AccessPackageResourceRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole%2Did"] = id
    }
    return i95bd789c6b6d638c6d9e88a9ce6120e1cc2d851f19b94543af06cdffa7aab1f4.NewAccessPackageResourceRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceScopes the accessPackageResourceScopes property
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*i7cdde6386682a7b30d1897e2f5b10b1dba49e35085058d46b9e97030e955dc08.AccessPackageResourceScopesRequestBuilder) {
    return i7cdde6386682a7b30d1897e2f5b10b1dba49e35085058d46b9e97030e955dc08.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.accessPackageResourceRoleScopes.item.accessPackageResourceRole.accessPackageResource.accessPackageResourceScopes.item collection
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopesById(id string)(*i967b53f70b934e90a12f7597314710840dfde37afa761e574abb6ed700e05925.AccessPackageResourceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope%2Did"] = id
    }
    return i967b53f70b934e90a12f7597314710840dfde37afa761e574abb6ed700e05925.NewAccessPackageResourceScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageResourceRequestBuilderInternal instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
func NewAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    m := &AccessPackageResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope%2Did}/accessPackageResourceRole/accessPackageResource{?%24select,%24expand}";
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
func (m *AccessPackageResourceRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageResourceRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation read-only. Nullable.
func (m *AccessPackageResourceRequestBuilder) CreateGetRequestInformation(options *AccessPackageResourceRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property accessPackageResource in identityGovernance
func (m *AccessPackageResourceRequestBuilder) CreatePatchRequestInformation(options *AccessPackageResourceRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property accessPackageResource for identityGovernance
func (m *AccessPackageResourceRequestBuilder) Delete(options *AccessPackageResourceRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable.
func (m *AccessPackageResourceRequestBuilder) Get(options *AccessPackageResourceRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceable), nil
}
// Patch update the navigation property accessPackageResource in identityGovernance
func (m *AccessPackageResourceRequestBuilder) Patch(options *AccessPackageResourceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}

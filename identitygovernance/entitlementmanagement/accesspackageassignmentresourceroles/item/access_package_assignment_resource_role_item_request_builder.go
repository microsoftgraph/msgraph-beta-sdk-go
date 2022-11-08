package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i7d890f65851793ed799c10549d6fd6c6c2494d30953973e4207fcc7acf22a83b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments"
    iaa99fd9136f8f73ab1315e849913867001cc91a1d284f0e03cee25e796bbae38 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageresourcescope"
    ic9d50657d48a7afc267e2cecedcbf53309e95d9f9bb8c499a5ff668c121cdc13 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageresourcerole"
    ife3670cfcf012fea44d391fd12342b004d1d1ec7881d957f768ace00e28c468f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackagesubject"
    i959b19586acf09d3de28ace7531945eb1c09ede616185b0017775aa599fd0304 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item"
)

// AccessPackageAssignmentResourceRoleItemRequestBuilder provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.entitlementManagement entity.
type AccessPackageAssignmentResourceRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters represents the resource-specific role which a subject has been assigned through an access package assignment.
type AccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters
}
// AccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignments provides operations to manage the accessPackageAssignments property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageAssignments()(*i7d890f65851793ed799c10549d6fd6c6c2494d30953973e4207fcc7acf22a83b.AccessPackageAssignmentsRequestBuilder) {
    return i7d890f65851793ed799c10549d6fd6c6c2494d30953973e4207fcc7acf22a83b.NewAccessPackageAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentsById provides operations to manage the accessPackageAssignments property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageAssignmentsById(id string)(*i959b19586acf09d3de28ace7531945eb1c09ede616185b0017775aa599fd0304.AccessPackageAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignment%2Did"] = id
    }
    return i959b19586acf09d3de28ace7531945eb1c09ede616185b0017775aa599fd0304.NewAccessPackageAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackageResourceRole provides operations to manage the accessPackageResourceRole property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageResourceRole()(*ic9d50657d48a7afc267e2cecedcbf53309e95d9f9bb8c499a5ff668c121cdc13.AccessPackageResourceRoleRequestBuilder) {
    return ic9d50657d48a7afc267e2cecedcbf53309e95d9f9bb8c499a5ff668c121cdc13.NewAccessPackageResourceRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScope provides operations to manage the accessPackageResourceScope property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageResourceScope()(*iaa99fd9136f8f73ab1315e849913867001cc91a1d284f0e03cee25e796bbae38.AccessPackageResourceScopeRequestBuilder) {
    return iaa99fd9136f8f73ab1315e849913867001cc91a1d284f0e03cee25e796bbae38.NewAccessPackageResourceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageSubject provides operations to manage the accessPackageSubject property of the microsoft.graph.accessPackageAssignmentResourceRole entity.
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageSubject()(*ife3670cfcf012fea44d391fd12342b004d1d1ec7881d957f768ace00e28c468f.AccessPackageSubjectRequestBuilder) {
    return ife3670cfcf012fea44d391fd12342b004d1d1ec7881d957f768ace00e28c468f.NewAccessPackageSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageAssignmentResourceRoleItemRequestBuilderInternal instantiates a new AccessPackageAssignmentResourceRoleItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentResourceRoleItemRequestBuilder) {
    m := &AccessPackageAssignmentResourceRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentResourceRoleItemRequestBuilder instantiates a new AccessPackageAssignmentResourceRoleItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentResourceRoleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessPackageAssignmentResourceRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageAssignmentResourceRoles for identityGovernance
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageAssignmentResourceRoles in identityGovernance
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, requestConfiguration *AccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property accessPackageAssignmentResourceRoles for identityGovernance
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessPackageAssignmentResourceRoleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the resource-specific role which a subject has been assigned through an access package assignment.
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessPackageAssignmentResourceRoleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable), nil
}
// Patch update the navigation property accessPackageAssignmentResourceRoles in identityGovernance
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, requestConfiguration *AccessPackageAssignmentResourceRoleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentResourceRoleable), nil
}

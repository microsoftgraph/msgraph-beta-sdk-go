package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder provides operations to manage the accessPackageAssignmentRequests property of the microsoft.graph.accessPackageAssignment entity.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters get accessPackageAssignmentRequests from identityGovernance
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentRequest entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) AccessPackage()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignment provides operations to manage the accessPackageAssignment property of the microsoft.graph.accessPackageAssignmentRequest entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) AccessPackageAssignment()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageAssignmentRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemAccessPackageAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Cancel()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemCancelRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal instantiates a new AccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentRequests/{accessPackageAssignmentRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder instantiates a new AccessPackageAssignmentRequestItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageAssignmentRequests for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get accessPackageAssignmentRequests from identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageAssignmentRequests in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property accessPackageAssignmentRequests for identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get accessPackageAssignmentRequests from identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable), nil
}
// Patch update the navigation property accessPackageAssignmentRequests in identityGovernance
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, requestConfiguration *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentRequestable), nil
}
// Reprocess provides operations to call the reprocess method.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Reprocess()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemReprocessRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Requestor provides operations to manage the requestor property of the microsoft.graph.accessPackageAssignmentRequest entity.
func (m *IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) Requestor()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

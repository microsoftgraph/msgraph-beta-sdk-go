package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder provides operations to manage the connectedOrganization property of the microsoft.graph.accessPackageSubject entity.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderGetQueryParameters the connected organization of the subject. Read-only. Nullable.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderGetQueryParameters
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderInternal instantiates a new ConnectedOrganizationRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentRequests/{accessPackageAssignmentRequest%2Did}/requestor/connectedOrganization{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder instantiates a new ConnectedOrganizationRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the connected organization of the subject. Read-only. Nullable.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions on 2023-03-01 and will be removed 2023-12-31
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConnectedOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable), nil
}
// ToGetRequestInformation the connected organization of the subject. Read-only. Nullable.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions on 2023-03-01 and will be removed 2023-12-31
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions on 2023-03-01 and will be removed 2023-12-31
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentRequestsItemRequestorConnectedOrganizationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

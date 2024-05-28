package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder provides operations to manage the accessPackageAssignment property of the microsoft.graph.accessPackageAssignmentRequest entity.
type EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderGetQueryParameters for a requestType of UserAdd or AdminAdd, an access package assignment requested to be created. For a requestType of UserRemove, AdminRemove, or SystemRemove, this property has the id property of an existing assignment to be removed. Supports $expand.
type EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder) {
    m := &EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentRequests/{accessPackageAssignmentRequest%2Did}/accessPackageAssignment{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder instantiates a new EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get for a requestType of UserAdd or AdminAdd, an access package assignment requested to be created. For a requestType of UserRemove, AdminRemove, or SystemRemove, this property has the id property of an existing assignment to be removed. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a AccessPackageAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageAssignmentable), nil
}
// ToGetRequestInformation for a requestType of UserAdd or AdminAdd, an access package assignment requested to be created. For a requestType of UserRemove, AdminRemove, or SystemRemove, this property has the id property of an existing assignment to be removed. Supports $expand.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemAccesspackageassignmentAccessPackageAssignmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

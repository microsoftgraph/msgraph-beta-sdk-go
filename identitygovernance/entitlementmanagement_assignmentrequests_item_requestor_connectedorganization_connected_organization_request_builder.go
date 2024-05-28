package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder provides operations to manage the connectedOrganization property of the microsoft.graph.accessPackageSubject entity.
type EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderGetQueryParameters the connected organization of the subject. Read-only. Nullable.
type EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder) {
    m := &EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentRequests/{accessPackageAssignmentRequest%2Did}/requestor/connectedOrganization{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder instantiates a new EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the connected organization of the subject. Read-only. Nullable.
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a ConnectedOrganizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Deprecated:  as of 2022-10/PrivatePreview:MicrosofEntitlementManagementCustomextensions
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemRequestorConnectedorganizationConnectedOrganizationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest entity.
type PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderGetQueryParameters references the group that is the scope of the membership or ownership eligibility request through PIM for groups. Supports $expand.
type PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderGetQueryParameters
}
// NewPrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder) {
    m := &PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/{privilegedAccessGroupEligibilityScheduleRequest%2Did}/group{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get references the group that is the scope of the membership or ownership eligibility request through PIM for groups. Supports $expand.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder) ServiceProvisioningErrors()(*PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupServiceProvisioningErrorsRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsItemGroupServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation references the group that is the scope of the membership or ownership eligibility request through PIM for groups. Supports $expand.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

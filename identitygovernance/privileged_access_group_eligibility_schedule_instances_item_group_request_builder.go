// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupEligibilityScheduleInstance entity.
type PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderGetQueryParameters references the group that is the scope of the membership or ownership eligibility through PIM for Groups. Supports $expand.
type PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderGetQueryParameters
}
// NewPrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderInternal instantiates a new PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder) {
    m := &PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances/{privilegedAccessGroupEligibilityScheduleInstance%2Did}/group{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder instantiates a new PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get references the group that is the scope of the membership or ownership eligibility through PIM for Groups. Supports $expand.
// returns a Groupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupServiceProvisioningErrorsRequestBuilder when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder) ServiceProvisioningErrors()(*PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupServiceProvisioningErrorsRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleInstancesItemGroupServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation references the group that is the scope of the membership or ownership eligibility through PIM for Groups. Supports $expand.
// returns a *RequestInformation when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder when successful
func (m *PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleInstancesItemGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

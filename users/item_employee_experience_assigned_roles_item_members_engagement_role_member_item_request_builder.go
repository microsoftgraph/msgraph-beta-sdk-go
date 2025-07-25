// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder provides operations to manage the members property of the microsoft.graph.engagementRole entity.
type ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderGetQueryParameters users who have been assigned this role.
type ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderGetQueryParameters
}
// ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderInternal instantiates a new ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder and sets the default values.
func NewItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) {
    m := &ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/employeeExperience/assignedRoles/{engagementRole%2Did}/members/{engagementRoleMember%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder instantiates a new ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder and sets the default values.
func NewItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property members for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get users who have been assigned this role.
// returns a EngagementRoleMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementRoleMemberable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEngagementRoleMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementRoleMemberable), nil
}
// Patch update the navigation property members in users
// returns a EngagementRoleMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementRoleMemberable, requestConfiguration *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementRoleMemberable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEngagementRoleMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementRoleMemberable), nil
}
// ToDeleteRequestInformation delete navigation property members for users
// returns a *RequestInformation when successful
func (m *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation users who have been assigned this role.
// returns a *RequestInformation when successful
func (m *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property members in users
// returns a *RequestInformation when successful
func (m *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementRoleMemberable, requestConfiguration *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// User provides operations to manage the user property of the microsoft.graph.engagementRoleMember entity.
// returns a *ItemEmployeeExperienceAssignedRolesItemMembersItemUserRequestBuilder when successful
func (m *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) User()(*ItemEmployeeExperienceAssignedRolesItemMembersItemUserRequestBuilder) {
    return NewItemEmployeeExperienceAssignedRolesItemMembersItemUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder when successful
func (m *ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) WithUrl(rawUrl string)(*ItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder) {
    return NewItemEmployeeExperienceAssignedRolesItemMembersEngagementRoleMemberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

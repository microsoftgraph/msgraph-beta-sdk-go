package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.appleUserInitiatedEnrollmentProfile entity.
type AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetQueryParameters the list of assignments for this profile.
type AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetQueryParameters
}
// AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderInternal instantiates a new AppleEnrollmentProfileAssignmentItemRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) {
    m := &AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile%2Did}/assignments/{appleEnrollmentProfileAssignment%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder instantiates a new AppleEnrollmentProfileAssignmentItemRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of assignments for this profile.
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable), nil
}
// Patch update the navigation property assignments in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property assignments for deviceManagement
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of assignments for this profile.
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in deviceManagement
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) {
    return NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

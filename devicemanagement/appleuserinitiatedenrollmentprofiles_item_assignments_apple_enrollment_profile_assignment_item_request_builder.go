package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.appleUserInitiatedEnrollmentProfile entity.
type AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetQueryParameters the list of assignments for this profile.
type AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetQueryParameters
}
// AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderInternal instantiates a new AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder and sets the default values.
func NewAppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) {
    m := &AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile%2Did}/assignments/{appleEnrollmentProfileAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder instantiates a new AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder and sets the default values.
func NewAppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of assignments for this profile.
// returns a AppleEnrollmentProfileAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a AppleEnrollmentProfileAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, requestConfiguration *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of assignments for this profile.
// returns a *RequestInformation when successful
func (m *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RequestInformation when successful
func (m *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, requestConfiguration *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder when successful
func (m *AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*AppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) {
    return NewAppleuserinitiatedenrollmentprofilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

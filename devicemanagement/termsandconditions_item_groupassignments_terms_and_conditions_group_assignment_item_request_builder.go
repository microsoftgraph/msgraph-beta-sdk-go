package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder provides operations to manage the groupAssignments property of the microsoft.graph.termsAndConditions entity.
type TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderGetQueryParameters the list of group assignments for this T&C policy.
type TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderGetQueryParameters
}
// TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderInternal instantiates a new TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder and sets the default values.
func NewTermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) {
    m := &TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions%2Did}/groupAssignments/{termsAndConditionsGroupAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder instantiates a new TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder and sets the default values.
func NewTermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupAssignments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of group assignments for this T&C policy.
// returns a TermsAndConditionsGroupAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsGroupAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTermsAndConditionsGroupAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsGroupAssignmentable), nil
}
// Patch update the navigation property groupAssignments in deviceManagement
// returns a TermsAndConditionsGroupAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsGroupAssignmentable, requestConfiguration *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsGroupAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTermsAndConditionsGroupAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsGroupAssignmentable), nil
}
// TermsAndConditions provides operations to manage the termsAndConditions property of the microsoft.graph.termsAndConditionsGroupAssignment entity.
// returns a *TermsandconditionsItemGroupassignmentsItemTermsandconditionsTermsAndConditionsRequestBuilder when successful
func (m *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) TermsAndConditions()(*TermsandconditionsItemGroupassignmentsItemTermsandconditionsTermsAndConditionsRequestBuilder) {
    return NewTermsandconditionsItemGroupassignmentsItemTermsandconditionsTermsAndConditionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property groupAssignments for deviceManagement
// returns a *RequestInformation when successful
func (m *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of group assignments for this T&C policy.
// returns a *RequestInformation when successful
func (m *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupAssignments in deviceManagement
// returns a *RequestInformation when successful
func (m *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsGroupAssignmentable, requestConfiguration *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder when successful
func (m *TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*TermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) {
    return NewTermsandconditionsItemGroupassignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

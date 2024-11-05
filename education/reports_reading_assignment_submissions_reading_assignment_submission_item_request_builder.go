package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder provides operations to manage the readingAssignmentSubmissions property of the microsoft.graph.reportsRoot entity.
type ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderGetQueryParameters get readingAssignmentSubmissions from education
type ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderGetQueryParameters
}
// ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderInternal instantiates a new ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder and sets the default values.
func NewReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) {
    m := &ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/reports/readingAssignmentSubmissions/{readingAssignmentSubmission%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder instantiates a new ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder and sets the default values.
func NewReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property readingAssignmentSubmissions for education
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get readingAssignmentSubmissions from education
// returns a ReadingAssignmentSubmissionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReadingAssignmentSubmissionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReadingAssignmentSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReadingAssignmentSubmissionable), nil
}
// Patch update the navigation property readingAssignmentSubmissions in education
// returns a ReadingAssignmentSubmissionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReadingAssignmentSubmissionable, requestConfiguration *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReadingAssignmentSubmissionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReadingAssignmentSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReadingAssignmentSubmissionable), nil
}
// ToDeleteRequestInformation delete navigation property readingAssignmentSubmissions for education
// returns a *RequestInformation when successful
func (m *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get readingAssignmentSubmissions from education
// returns a *RequestInformation when successful
func (m *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property readingAssignmentSubmissions in education
// returns a *RequestInformation when successful
func (m *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReadingAssignmentSubmissionable, requestConfiguration *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder when successful
func (m *ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) WithUrl(rawUrl string)(*ReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder) {
    return NewReportsReadingAssignmentSubmissionsReadingAssignmentSubmissionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

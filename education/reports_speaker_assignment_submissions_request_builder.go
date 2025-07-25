// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsSpeakerAssignmentSubmissionsRequestBuilder provides operations to manage the speakerAssignmentSubmissions property of the microsoft.graph.reportsRoot entity.
type ReportsSpeakerAssignmentSubmissionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsSpeakerAssignmentSubmissionsRequestBuilderGetQueryParameters get a list of speaker assignments that were submitted by a student.
type ReportsSpeakerAssignmentSubmissionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ReportsSpeakerAssignmentSubmissionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsSpeakerAssignmentSubmissionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsSpeakerAssignmentSubmissionsRequestBuilderGetQueryParameters
}
// ReportsSpeakerAssignmentSubmissionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsSpeakerAssignmentSubmissionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySpeakerAssignmentSubmissionId provides operations to manage the speakerAssignmentSubmissions property of the microsoft.graph.reportsRoot entity.
// returns a *ReportsSpeakerAssignmentSubmissionsSpeakerAssignmentSubmissionItemRequestBuilder when successful
func (m *ReportsSpeakerAssignmentSubmissionsRequestBuilder) BySpeakerAssignmentSubmissionId(speakerAssignmentSubmissionId string)(*ReportsSpeakerAssignmentSubmissionsSpeakerAssignmentSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if speakerAssignmentSubmissionId != "" {
        urlTplParams["speakerAssignmentSubmission%2Did"] = speakerAssignmentSubmissionId
    }
    return NewReportsSpeakerAssignmentSubmissionsSpeakerAssignmentSubmissionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewReportsSpeakerAssignmentSubmissionsRequestBuilderInternal instantiates a new ReportsSpeakerAssignmentSubmissionsRequestBuilder and sets the default values.
func NewReportsSpeakerAssignmentSubmissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsSpeakerAssignmentSubmissionsRequestBuilder) {
    m := &ReportsSpeakerAssignmentSubmissionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/reports/speakerAssignmentSubmissions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewReportsSpeakerAssignmentSubmissionsRequestBuilder instantiates a new ReportsSpeakerAssignmentSubmissionsRequestBuilder and sets the default values.
func NewReportsSpeakerAssignmentSubmissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsSpeakerAssignmentSubmissionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsSpeakerAssignmentSubmissionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ReportsSpeakerAssignmentSubmissionsCountRequestBuilder when successful
func (m *ReportsSpeakerAssignmentSubmissionsRequestBuilder) Count()(*ReportsSpeakerAssignmentSubmissionsCountRequestBuilder) {
    return NewReportsSpeakerAssignmentSubmissionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of speaker assignments that were submitted by a student.
// returns a SpeakerAssignmentSubmissionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportsroot-list-speakerassignmentsubmissions?view=graph-rest-beta
func (m *ReportsSpeakerAssignmentSubmissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsSpeakerAssignmentSubmissionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SpeakerAssignmentSubmissionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSpeakerAssignmentSubmissionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SpeakerAssignmentSubmissionCollectionResponseable), nil
}
// Post create new navigation property to speakerAssignmentSubmissions for education
// returns a SpeakerAssignmentSubmissionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsSpeakerAssignmentSubmissionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SpeakerAssignmentSubmissionable, requestConfiguration *ReportsSpeakerAssignmentSubmissionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SpeakerAssignmentSubmissionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSpeakerAssignmentSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SpeakerAssignmentSubmissionable), nil
}
// ToGetRequestInformation get a list of speaker assignments that were submitted by a student.
// returns a *RequestInformation when successful
func (m *ReportsSpeakerAssignmentSubmissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsSpeakerAssignmentSubmissionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to speakerAssignmentSubmissions for education
// returns a *RequestInformation when successful
func (m *ReportsSpeakerAssignmentSubmissionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SpeakerAssignmentSubmissionable, requestConfiguration *ReportsSpeakerAssignmentSubmissionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ReportsSpeakerAssignmentSubmissionsRequestBuilder when successful
func (m *ReportsSpeakerAssignmentSubmissionsRequestBuilder) WithUrl(rawUrl string)(*ReportsSpeakerAssignmentSubmissionsRequestBuilder) {
    return NewReportsSpeakerAssignmentSubmissionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

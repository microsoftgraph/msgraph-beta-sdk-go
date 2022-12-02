package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder provides operations to manage the dispositionReviewStages property of the microsoft.graph.security.retentionLabel entity.
type SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderGetQueryParameters review stages during which reviewers are notified to determine whether a document must be deleted or retained.
type SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderGetQueryParameters
}
// SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderInternal instantiates a new DispositionReviewStageItemRequestBuilder and sets the default values.
func NewSecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder) {
    m := &SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/labels/retentionLabels/{retentionLabel%2Did}/dispositionReviewStages/{dispositionReviewStage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder instantiates a new DispositionReviewStageItemRequestBuilder and sets the default values.
func NewSecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property dispositionReviewStages for security
func (m *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation review stages during which reviewers are notified to determine whether a document must be deleted or retained.
func (m *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property dispositionReviewStages in security
func (m *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DispositionReviewStageable, requestConfiguration *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property dispositionReviewStages for security
func (m *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get review stages during which reviewers are notified to determine whether a document must be deleted or retained.
func (m *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DispositionReviewStageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateDispositionReviewStageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DispositionReviewStageable), nil
}
// Patch update the navigation property dispositionReviewStages in security
func (m *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DispositionReviewStageable, requestConfiguration *SecurityLabelsRetentionLabelsItemDispositionReviewStagesDispositionReviewStageItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DispositionReviewStageable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateDispositionReviewStageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DispositionReviewStageable), nil
}

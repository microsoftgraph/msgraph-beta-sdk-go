package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder provides operations to manage the results property of the microsoft.graph.threatAssessmentRequest entity.
type MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetQueryParameters a collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
type MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetQueryParameters
}
// MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderInternal instantiates a new ThreatAssessmentResultItemRequestBuilder and sets the default values.
func NewMeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) {
    m := &MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/informationProtection/threatAssessmentRequests/{threatAssessmentRequest%2Did}/results/{threatAssessmentResult%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder instantiates a new ThreatAssessmentResultItemRequestBuilder and sets the default values.
func NewMeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property results for me
func (m *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
func (m *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property results in me
func (m *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, requestConfiguration *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property results for me
func (m *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
func (m *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateThreatAssessmentResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable), nil
}
// Patch update the navigation property results in me
func (m *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, requestConfiguration *MeInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateThreatAssessmentResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable), nil
}

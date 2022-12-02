package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder provides operations to count the resources in the collection.
type MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewMeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder) {
    m := &MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/informationProtection/threatAssessmentRequests/{threatAssessmentRequest%2Did}/results/$count";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewMeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the number of the resource
func (m *MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "text/plain"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the number of the resource
func (m *MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *MeInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}

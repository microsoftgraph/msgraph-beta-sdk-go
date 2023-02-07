package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder provides operations to call the getFinalReport method.
type SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilderInternal instantiates a new MicrosoftGraphGetFinalReportRequestBuilder and sets the default values.
func NewSubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder) {
    m := &SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/subjectRightsRequests/{subjectRightsRequest%2Did}/microsoft.graph.getFinalReport()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder instantiates a new MicrosoftGraphGetFinalReportRequestBuilder and sets the default values.
func NewSubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getFinalReport
func (m *SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getFinalReport
func (m *SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectRightsRequestsItemMicrosoftGraphGetFinalReportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

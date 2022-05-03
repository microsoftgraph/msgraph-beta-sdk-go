package getsuggestedenrollmentlimitwithenrollmenttype

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder provides operations to call the getSuggestedEnrollmentLimit method.
type GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal instantiates a new GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder and sets the default values.
func NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, enrollmentType *string)(*GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    m := &GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getSuggestedEnrollmentLimit(enrollmentType='{enrollmentType}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if enrollmentType != nil {
        urlTplParams[""] = *enrollmentType
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder instantiates a new GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder and sets the default values.
func NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getSuggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getSuggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler invoke function getSuggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) GetWithResponseHandler(requestConfiguration *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SuggestedEnrollmentLimitable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getSuggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) GetWithResponseHandler(requestConfiguration *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SuggestedEnrollmentLimitable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSuggestedEnrollmentLimitFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SuggestedEnrollmentLimitable), nil
}

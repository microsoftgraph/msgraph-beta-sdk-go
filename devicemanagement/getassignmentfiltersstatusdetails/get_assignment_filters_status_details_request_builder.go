package getassignmentfiltersstatusdetails

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// GetAssignmentFiltersStatusDetailsRequestBuilder provides operations to call the getAssignmentFiltersStatusDetails method.
type GetAssignmentFiltersStatusDetailsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetAssignmentFiltersStatusDetailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetAssignmentFiltersStatusDetailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetAssignmentFiltersStatusDetailsRequestBuilderInternal instantiates a new GetAssignmentFiltersStatusDetailsRequestBuilder and sets the default values.
func NewGetAssignmentFiltersStatusDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAssignmentFiltersStatusDetailsRequestBuilder) {
    m := &GetAssignmentFiltersStatusDetailsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getAssignmentFiltersStatusDetails";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetAssignmentFiltersStatusDetailsRequestBuilder instantiates a new GetAssignmentFiltersStatusDetailsRequestBuilder and sets the default values.
func NewGetAssignmentFiltersStatusDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAssignmentFiltersStatusDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAssignmentFiltersStatusDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getAssignmentFiltersStatusDetails
func (m *GetAssignmentFiltersStatusDetailsRequestBuilder) CreatePostRequestInformation(body GetAssignmentFiltersStatusDetailsRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getAssignmentFiltersStatusDetails
func (m *GetAssignmentFiltersStatusDetailsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetAssignmentFiltersStatusDetailsRequestBodyable, requestConfiguration *GetAssignmentFiltersStatusDetailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getAssignmentFiltersStatusDetails
func (m *GetAssignmentFiltersStatusDetailsRequestBuilder) Post(body GetAssignmentFiltersStatusDetailsRequestBodyable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterStatusDetailsable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getAssignmentFiltersStatusDetails
func (m *GetAssignmentFiltersStatusDetailsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetAssignmentFiltersStatusDetailsRequestBodyable, requestConfiguration *GetAssignmentFiltersStatusDetailsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterStatusDetailsable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentFilterStatusDetailsFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterStatusDetailsable), nil
}

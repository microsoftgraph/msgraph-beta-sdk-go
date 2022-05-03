package getstaffavailability

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetStaffAvailabilityRequestBuilder provides operations to call the getStaffAvailability method.
type GetStaffAvailabilityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetStaffAvailabilityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetStaffAvailabilityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetStaffAvailabilityRequestBuilderInternal instantiates a new GetStaffAvailabilityRequestBuilder and sets the default values.
func NewGetStaffAvailabilityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetStaffAvailabilityRequestBuilder) {
    m := &GetStaffAvailabilityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/bookingBusinesses/{bookingBusiness%2Did}/microsoft.graph.getStaffAvailability";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetStaffAvailabilityRequestBuilder instantiates a new GetStaffAvailabilityRequestBuilder and sets the default values.
func NewGetStaffAvailabilityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetStaffAvailabilityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetStaffAvailabilityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getStaffAvailability
func (m *GetStaffAvailabilityRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetStaffAvailabilityRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getStaffAvailability
func (m *GetStaffAvailabilityRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetStaffAvailabilityRequestBodyable, requestConfiguration *GetStaffAvailabilityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action getStaffAvailability
func (m *GetStaffAvailabilityRequestBuilder) PostWithResponseHandler(body GetStaffAvailabilityRequestBodyable, requestConfiguration *GetStaffAvailabilityRequestBuilderPostRequestConfiguration)(GetStaffAvailabilityResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getStaffAvailability
func (m *GetStaffAvailabilityRequestBuilder) PostWithResponseHandler(body GetStaffAvailabilityRequestBodyable, requestConfiguration *GetStaffAvailabilityRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetStaffAvailabilityResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetStaffAvailabilityResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetStaffAvailabilityResponseable), nil
}

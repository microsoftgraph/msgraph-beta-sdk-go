package getstaffavailability

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
// CreatePostRequestInformation get the availability information of staff members of a Microsoft Bookings calendar.
func (m *GetStaffAvailabilityRequestBuilder) CreatePostRequestInformation(body GetStaffAvailabilityPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration get the availability information of staff members of a Microsoft Bookings calendar.
func (m *GetStaffAvailabilityRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetStaffAvailabilityPostRequestBodyable, requestConfiguration *GetStaffAvailabilityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post get the availability information of staff members of a Microsoft Bookings calendar.
func (m *GetStaffAvailabilityRequestBuilder) Post(ctx context.Context, body GetStaffAvailabilityPostRequestBodyable, requestConfiguration *GetStaffAvailabilityRequestBuilderPostRequestConfiguration)(GetStaffAvailabilityResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateGetStaffAvailabilityResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetStaffAvailabilityResponseable), nil
}

package bookingbusinesses

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilder provides operations to call the getStaffAvailability method.
type ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilderInternal instantiates a new GetStaffAvailabilityRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilder) {
    m := &ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/bookingBusinesses/{bookingBusiness%2Did}/microsoft.graph.getStaffAvailability";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilder instantiates a new GetStaffAvailabilityRequestBuilder and sets the default values.
func NewItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the availability information of staff members of a Microsoft Bookings calendar.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/bookingbusiness-getstaffavailability?view=graph-rest-1.0
func (m *ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilder) Post(ctx context.Context, body ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilderPostRequestConfiguration)(ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityResponseable), nil
}
// ToPostRequestInformation get the availability information of staff members of a Microsoft Bookings calendar.
func (m *ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphGetStaffAvailabilityGetStaffAvailabilityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}

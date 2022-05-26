package searchexistingidentities

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SearchExistingIdentitiesRequestBuilder provides operations to call the searchExistingIdentities method.
type SearchExistingIdentitiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SearchExistingIdentitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SearchExistingIdentitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSearchExistingIdentitiesRequestBuilderInternal instantiates a new SearchExistingIdentitiesRequestBuilder and sets the default values.
func NewSearchExistingIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchExistingIdentitiesRequestBuilder) {
    m := &SearchExistingIdentitiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/importedDeviceIdentities/microsoft.graph.searchExistingIdentities";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSearchExistingIdentitiesRequestBuilder instantiates a new SearchExistingIdentitiesRequestBuilder and sets the default values.
func NewSearchExistingIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchExistingIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSearchExistingIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action searchExistingIdentities
func (m *SearchExistingIdentitiesRequestBuilder) CreatePostRequestInformation(body SearchExistingIdentitiesPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action searchExistingIdentities
func (m *SearchExistingIdentitiesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SearchExistingIdentitiesPostRequestBodyable, requestConfiguration *SearchExistingIdentitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action searchExistingIdentities
func (m *SearchExistingIdentitiesRequestBuilder) Post(body SearchExistingIdentitiesPostRequestBodyable)(SearchExistingIdentitiesResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action searchExistingIdentities
func (m *SearchExistingIdentitiesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body SearchExistingIdentitiesPostRequestBodyable, requestConfiguration *SearchExistingIdentitiesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(SearchExistingIdentitiesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSearchExistingIdentitiesResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(SearchExistingIdentitiesResponseable), nil
}

package creategoogleplaywebtoken

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CreateGooglePlayWebTokenRequestBuilder provides operations to call the createGooglePlayWebToken method.
type CreateGooglePlayWebTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CreateGooglePlayWebTokenRequestBuilderPostOptions options for Post
type CreateGooglePlayWebTokenRequestBuilderPostOptions struct {
    // 
    Body CreateGooglePlayWebTokenRequestBodyable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewCreateGooglePlayWebTokenRequestBuilderInternal instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewCreateGooglePlayWebTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateGooglePlayWebTokenRequestBuilder) {
    m := &CreateGooglePlayWebTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/microsoft.graph.createGooglePlayWebToken";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreateGooglePlayWebTokenRequestBuilder instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewCreateGooglePlayWebTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateGooglePlayWebTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateGooglePlayWebTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation generates a web token that is used in an embeddable component.
func (m *CreateGooglePlayWebTokenRequestBuilder) CreatePostRequestInformation(options *CreateGooglePlayWebTokenRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post generates a web token that is used in an embeddable component.
func (m *CreateGooglePlayWebTokenRequestBuilder) Post(options *CreateGooglePlayWebTokenRequestBuilderPostOptions)(CreateGooglePlayWebTokenResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateCreateGooglePlayWebTokenResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(CreateGooglePlayWebTokenResponseable), nil
}

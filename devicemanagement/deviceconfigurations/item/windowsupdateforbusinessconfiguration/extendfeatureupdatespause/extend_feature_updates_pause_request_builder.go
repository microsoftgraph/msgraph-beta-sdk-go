package extendfeatureupdatespause

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExtendFeatureUpdatesPauseRequestBuilder provides operations to call the extendFeatureUpdatesPause method.
type ExtendFeatureUpdatesPauseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExtendFeatureUpdatesPauseRequestBuilderPostOptions options for Post
type ExtendFeatureUpdatesPauseRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewExtendFeatureUpdatesPauseRequestBuilderInternal instantiates a new ExtendFeatureUpdatesPauseRequestBuilder and sets the default values.
func NewExtendFeatureUpdatesPauseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExtendFeatureUpdatesPauseRequestBuilder) {
    m := &ExtendFeatureUpdatesPauseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/microsoft.graph.windowsUpdateForBusinessConfiguration/microsoft.graph.extendFeatureUpdatesPause";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExtendFeatureUpdatesPauseRequestBuilder instantiates a new ExtendFeatureUpdatesPauseRequestBuilder and sets the default values.
func NewExtendFeatureUpdatesPauseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExtendFeatureUpdatesPauseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExtendFeatureUpdatesPauseRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation extend Feature Updates Pause for a Windows Update for Business ring.
func (m *ExtendFeatureUpdatesPauseRequestBuilder) CreatePostRequestInformation(options *ExtendFeatureUpdatesPauseRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Post extend Feature Updates Pause for a Windows Update for Business ring.
func (m *ExtendFeatureUpdatesPauseRequestBuilder) Post(options *ExtendFeatureUpdatesPauseRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}

package getoffice365groupsactivitystoragewithperiod

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetOffice365GroupsActivityStorageWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityStorage method.
type GetOffice365GroupsActivityStorageWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetOptions options for Get
type GetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal instantiates a new GetOffice365GroupsActivityStorageWithPeriodRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    m := &GetOffice365GroupsActivityStorageWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getOffice365GroupsActivityStorage(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams[""] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilder instantiates a new GetOffice365GroupsActivityStorageWithPeriodRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getOffice365GroupsActivityStorage
func (m *GetOffice365GroupsActivityStorageWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
// Get invoke function getOffice365GroupsActivityStorage
func (m *GetOffice365GroupsActivityStorageWithPeriodRequestBuilder) Get(options *GetOffice365GroupsActivityStorageWithPeriodRequestBuilderGetOptions)(GetOffice365GroupsActivityStorageWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetOffice365GroupsActivityStorageWithPeriodResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetOffice365GroupsActivityStorageWithPeriodResponseable), nil
}

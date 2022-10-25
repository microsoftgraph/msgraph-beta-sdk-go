package manageddeviceenrollmentabandonmentdetailswithskipwithtopwithfilterwithskiptoken

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder provides operations to call the managedDeviceEnrollmentAbandonmentDetails method.
type ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal instantiates a new ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder and sets the default values.
func NewManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, filter *string, skip *int32, skipToken *string, top *int32)(*ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    m := &ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports/microsoft.graph.managedDeviceEnrollmentAbandonmentDetails(skip={skip},top={top},filter='{filter}',skipToken='{skipToken}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if filter != nil {
        urlTplParams["filter"] = *filter
    }
    if skip != nil {
        urlTplParams["skip"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*skip), 10)
    }
    if skipToken != nil {
        urlTplParams["skipToken"] = *skipToken
    }
    if top != nil {
        urlTplParams["top"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*top), 10)
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder instantiates a new ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder and sets the default values.
func NewManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil, nil)
}
// CreateGetRequestInformation metadata for Enrollment abandonment details report
func (m *ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get metadata for Enrollment abandonment details report
func (m *ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}

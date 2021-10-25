package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i058f5de0abd53df643c0b1d4535f14317cbfcc0a5ef23108987eaeb903f7eb21 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/apphealthmetrics"
    i4f5cd14f349081a4fb0b52a2ea5ec7b6c90bc3d3cc26fbbd0891609c4b1b0d07 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/workfromanywheremetrics"
    i69e441d79d9ffb70d727d5891fb03cb7264fe1ac972103efbef3cdeff428347c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/devicebootperformancemetrics"
    icd0719662a86b52cac5fcea4da793c6735d376b486c1450eed1b5ddc8ff8e6e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/rebootanalyticsmetrics"
    iec2f756b9554ce46110f6b3f733d0617bf3ef06dc480e86cb09cbf2d74628de8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/bestpracticesmetrics"
    iecb3f73f1e42490bd6888019100a7693b3f982d354ec76f90bf067fc6fabe7de "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/resourceperformancemetrics"
)

type UserExperienceAnalyticsBaselineRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type UserExperienceAnalyticsBaselineRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) AppHealthMetrics()(*i058f5de0abd53df643c0b1d4535f14317cbfcc0a5ef23108987eaeb903f7eb21.AppHealthMetricsRequestBuilder) {
    return i058f5de0abd53df643c0b1d4535f14317cbfcc0a5ef23108987eaeb903f7eb21.NewAppHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) BestPracticesMetrics()(*iec2f756b9554ce46110f6b3f733d0617bf3ef06dc480e86cb09cbf2d74628de8.BestPracticesMetricsRequestBuilder) {
    return iec2f756b9554ce46110f6b3f733d0617bf3ef06dc480e86cb09cbf2d74628de8.NewBestPracticesMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewUserExperienceAnalyticsBaselineRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsBaselineRequestBuilder) {
    m := &UserExperienceAnalyticsBaselineRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewUserExperienceAnalyticsBaselineRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsBaselineRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsBaselineRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) CreateGetRequestInformation(q func (value *UserExperienceAnalyticsBaselineRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(UserExperienceAnalyticsBaselineRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsBaseline, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) DeviceBootPerformanceMetrics()(*i69e441d79d9ffb70d727d5891fb03cb7264fe1ac972103efbef3cdeff428347c.DeviceBootPerformanceMetricsRequestBuilder) {
    return i69e441d79d9ffb70d727d5891fb03cb7264fe1ac972103efbef3cdeff428347c.NewDeviceBootPerformanceMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) Get(q func (value *UserExperienceAnalyticsBaselineRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsBaseline, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsBaseline() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsBaseline), nil
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsBaseline, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) RebootAnalyticsMetrics()(*icd0719662a86b52cac5fcea4da793c6735d376b486c1450eed1b5ddc8ff8e6e2.RebootAnalyticsMetricsRequestBuilder) {
    return icd0719662a86b52cac5fcea4da793c6735d376b486c1450eed1b5ddc8ff8e6e2.NewRebootAnalyticsMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) ResourcePerformanceMetrics()(*iecb3f73f1e42490bd6888019100a7693b3f982d354ec76f90bf067fc6fabe7de.ResourcePerformanceMetricsRequestBuilder) {
    return iecb3f73f1e42490bd6888019100a7693b3f982d354ec76f90bf067fc6fabe7de.NewResourcePerformanceMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) WorkFromAnywhereMetrics()(*i4f5cd14f349081a4fb0b52a2ea5ec7b6c90bc3d3cc26fbbd0891609c4b1b0d07.WorkFromAnywhereMetricsRequestBuilder) {
    return i4f5cd14f349081a4fb0b52a2ea5ec7b6c90bc3d3cc26fbbd0891609c4b1b0d07.NewWorkFromAnywhereMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

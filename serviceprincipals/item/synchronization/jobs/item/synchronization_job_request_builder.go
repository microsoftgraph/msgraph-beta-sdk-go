package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i072f978c5e76c8f6f7f146e9adbc77e18bc617a145954334b9d760baaf124b24 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/schema"
    i1037a2ef85bc7c104d93f4c5a1e5a28ff559376fe6cb1bb568e1d1c84273c5d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/stop"
    i17256af27a6172f385302adc8f2924bddbaa8027a02a4e613619e8304ec94fe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/provisionondemand"
    i37054552034248ca68fec46390fca0923f7a2e6c10539e7a67204b742ef8e9fb "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/restart"
    ia80ae8b8dfc2c4ddfe68e3d7f5fb7158905c34e938b0cc1afbae2798e432005b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/validatecredentials"
    id1feb9558d42229603da95cf0e6bf17dfe11c8a034277b3ec93c1b87e7a947a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/pause"
    ie7dd06601f7be4a75e499cc220f5c87d45dca4695a9b092b00f512bf0edb0643 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/start"
)

type SynchronizationJobRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SynchronizationJobRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewSynchronizationJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationJobRequestBuilder) {
    m := &SynchronizationJobRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/servicePrincipals/{servicePrincipal_id}/synchronization/jobs/{synchronizationJob_id}{?select,expand}";
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
func NewSynchronizationJobRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationJobRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationJobRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SynchronizationJobRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SynchronizationJobRequestBuilder) CreateGetRequestInformation(q func (value *SynchronizationJobRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SynchronizationJobRequestBuilderGetQueryParameters)
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
func (m *SynchronizationJobRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJob, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SynchronizationJobRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SynchronizationJobRequestBuilder) Get(q func (value *SynchronizationJobRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJob, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSynchronizationJob() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJob), nil
}
func (m *SynchronizationJobRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJob, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SynchronizationJobRequestBuilder) Pause()(*id1feb9558d42229603da95cf0e6bf17dfe11c8a034277b3ec93c1b87e7a947a7.PauseRequestBuilder) {
    return id1feb9558d42229603da95cf0e6bf17dfe11c8a034277b3ec93c1b87e7a947a7.NewPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) ProvisionOnDemand()(*i17256af27a6172f385302adc8f2924bddbaa8027a02a4e613619e8304ec94fe5.ProvisionOnDemandRequestBuilder) {
    return i17256af27a6172f385302adc8f2924bddbaa8027a02a4e613619e8304ec94fe5.NewProvisionOnDemandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) Restart()(*i37054552034248ca68fec46390fca0923f7a2e6c10539e7a67204b742ef8e9fb.RestartRequestBuilder) {
    return i37054552034248ca68fec46390fca0923f7a2e6c10539e7a67204b742ef8e9fb.NewRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) Schema()(*i072f978c5e76c8f6f7f146e9adbc77e18bc617a145954334b9d760baaf124b24.SchemaRequestBuilder) {
    return i072f978c5e76c8f6f7f146e9adbc77e18bc617a145954334b9d760baaf124b24.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) Start()(*ie7dd06601f7be4a75e499cc220f5c87d45dca4695a9b092b00f512bf0edb0643.StartRequestBuilder) {
    return ie7dd06601f7be4a75e499cc220f5c87d45dca4695a9b092b00f512bf0edb0643.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) Stop()(*i1037a2ef85bc7c104d93f4c5a1e5a28ff559376fe6cb1bb568e1d1c84273c5d7.StopRequestBuilder) {
    return i1037a2ef85bc7c104d93f4c5a1e5a28ff559376fe6cb1bb568e1d1c84273c5d7.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) ValidateCredentials()(*ia80ae8b8dfc2c4ddfe68e3d7f5fb7158905c34e938b0cc1afbae2798e432005b.ValidateCredentialsRequestBuilder) {
    return ia80ae8b8dfc2c4ddfe68e3d7f5fb7158905c34e938b0cc1afbae2798e432005b.NewValidateCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}

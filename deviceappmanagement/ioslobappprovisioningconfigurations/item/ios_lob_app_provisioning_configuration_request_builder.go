package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/userstatuses"
    i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/groupassignments"
    ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assign"
    ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assignments"
    if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/devicestatuses"
    i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/devicestatuses/item"
    i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/groupassignments/item"
    i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/userstatuses/item"
    ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assignments/item"
)

type IosLobAppProvisioningConfigurationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type IosLobAppProvisioningConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Assign()(*ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85.AssignRequestBuilder) {
    return ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Assignments()(*ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567.AssignmentsRequestBuilder) {
    return ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) AssignmentsById(id string)(*ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0.IosLobAppProvisioningConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosLobAppProvisioningConfigurationAssignment_id"] = id
    }
    return ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0.NewIosLobAppProvisioningConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewIosLobAppProvisioningConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosLobAppProvisioningConfigurationRequestBuilder) {
    m := &IosLobAppProvisioningConfigurationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewIosLobAppProvisioningConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosLobAppProvisioningConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosLobAppProvisioningConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *IosLobAppProvisioningConfigurationRequestBuilder) CreateGetRequestInformation(q func (value *IosLobAppProvisioningConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(IosLobAppProvisioningConfigurationRequestBuilderGetQueryParameters)
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
func (m *IosLobAppProvisioningConfigurationRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *IosLobAppProvisioningConfigurationRequestBuilder) DeviceStatuses()(*if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5.DeviceStatusesRequestBuilder) {
    return if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) DeviceStatusesById(id string)(*i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8.ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus_id"] = id
    }
    return i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8.NewManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Get(q func (value *IosLobAppProvisioningConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewIosLobAppProvisioningConfiguration() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfiguration), nil
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) GroupAssignments()(*i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0.GroupAssignmentsRequestBuilder) {
    return i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) GroupAssignmentsById(id string)(*i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19.MobileAppProvisioningConfigGroupAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppProvisioningConfigGroupAssignment_id"] = id
    }
    return i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19.NewMobileAppProvisioningConfigGroupAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *IosLobAppProvisioningConfigurationRequestBuilder) UserStatuses()(*i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929.UserStatusesRequestBuilder) {
    return i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) UserStatusesById(id string)(*i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d.ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus_id"] = id
    }
    return i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d.NewManagedDeviceMobileAppConfigurationUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
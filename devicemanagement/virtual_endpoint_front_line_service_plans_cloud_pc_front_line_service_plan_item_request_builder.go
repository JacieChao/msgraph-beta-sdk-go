package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder provides operations to manage the frontLineServicePlans property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetQueryParameters get the properties and relationships of a cloudPcFrontLineServicePlan object. This API is available in the following national cloud deployments.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetQueryParameters
}
// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderInternal instantiates a new CloudPcFrontLineServicePlanItemRequestBuilder and sets the default values.
func NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) {
    m := &VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/frontLineServicePlans/{cloudPcFrontLineServicePlan%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder instantiates a new CloudPcFrontLineServicePlanItemRequestBuilder and sets the default values.
func NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property frontLineServicePlans for deviceManagement
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get the properties and relationships of a cloudPcFrontLineServicePlan object. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcfrontlineserviceplan-get?view=graph-rest-1.0
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcFrontLineServicePlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable), nil
}
// Patch update the navigation property frontLineServicePlans in deviceManagement
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcFrontLineServicePlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable), nil
}
// ToDeleteRequestInformation delete navigation property frontLineServicePlans for deviceManagement
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of a cloudPcFrontLineServicePlan object. This API is available in the following national cloud deployments.
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property frontLineServicePlans in deviceManagement
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) {
    return NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

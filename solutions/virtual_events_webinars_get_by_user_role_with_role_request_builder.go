package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder provides operations to call the getByUserRole method.
type VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderGetQueryParameters get a list of virtualEventWebinar objects where the signed-in user is either the organizer or a coorganizer.
type VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderGetQueryParameters
}
// NewVirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderInternal instantiates a new VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, role *string)(*VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder) {
    m := &VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/getByUserRole(role='{role}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if role != nil {
        m.BaseRequestBuilder.PathParameters["role"] = *role
    }
    return m
}
// NewVirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder instantiates a new VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get a list of virtualEventWebinar objects where the signed-in user is either the organizer or a coorganizer.
// Deprecated: This method is obsolete. Use GetAsGetByUserRoleWithRoleGetResponse instead.
// returns a VirtualEventsWebinarsGetByUserRoleWithRoleResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventwebinar-getbyuserrole?view=graph-rest-beta
func (m *VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration)(VirtualEventsWebinarsGetByUserRoleWithRoleResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEventsWebinarsGetByUserRoleWithRoleResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEventsWebinarsGetByUserRoleWithRoleResponseable), nil
}
// GetAsGetByUserRoleWithRoleGetResponse get a list of virtualEventWebinar objects where the signed-in user is either the organizer or a coorganizer.
// returns a VirtualEventsWebinarsGetByUserRoleWithRoleGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventwebinar-getbyuserrole?view=graph-rest-beta
func (m *VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder) GetAsGetByUserRoleWithRoleGetResponse(ctx context.Context, requestConfiguration *VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration)(VirtualEventsWebinarsGetByUserRoleWithRoleGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEventsWebinarsGetByUserRoleWithRoleGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEventsWebinarsGetByUserRoleWithRoleGetResponseable), nil
}
// ToGetRequestInformation get a list of virtualEventWebinar objects where the signed-in user is either the organizer or a coorganizer.
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder when successful
func (m *VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder) {
    return NewVirtualEventsWebinarsGetByUserRoleWithRoleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}

package svc_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	userprovider "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	typesv1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	ctxpkg "github.com/cs3org/reva/v2/pkg/ctx"
	"github.com/cs3org/reva/v2/pkg/rgrpc/status"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	libregraph "github.com/owncloud/libre-graph-api-go"
	ogrpc "github.com/owncloud/ocis/v2/ocis-pkg/service/grpc"
	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
	v0 "github.com/owncloud/ocis/v2/protogen/gen/ocis/messages/settings/v0"
	settingssvc "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/settings/v0"
	"github.com/owncloud/ocis/v2/services/graph/mocks"
	"github.com/owncloud/ocis/v2/services/graph/pkg/config"
	"github.com/owncloud/ocis/v2/services/graph/pkg/config/defaults"
	service "github.com/owncloud/ocis/v2/services/graph/pkg/service/v0"
	"github.com/owncloud/ocis/v2/services/graph/pkg/service/v0/errorcode"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Graph", func() {
	var (
		svc               service.Service
		gatewayClient     *mocks.GatewayClient
		eventsPublisher   mocks.Publisher
		permissionService mocks.Permissions
		ctx               context.Context
		cfg               *config.Config
	)

	JustBeforeEach(func() {
		ctx = ctxpkg.ContextSetUser(context.Background(), &userprovider.User{Id: &userprovider.UserId{Type: userprovider.UserType_USER_TYPE_PRIMARY, OpaqueId: "testuser"}, Username: "testuser"})
		cfg = defaults.FullDefaultConfig()
		cfg.Identity.LDAP.CACert = "" // skip the startup checks, we don't use LDAP at all in this tests
		cfg.TokenManager.JWTSecret = "loremipsum"
		cfg.Commons = &shared.Commons{}
		cfg.GRPCClientTLS = &shared.GRPCClientTLS{}

		_ = ogrpc.Configure(ogrpc.GetClientOptions(cfg.GRPCClientTLS)...)
		gatewayClient = &mocks.GatewayClient{}
		eventsPublisher = mocks.Publisher{}
		permissionService = mocks.Permissions{}
		svc = service.NewService(
			service.Config(cfg),
			service.WithGatewayClient(gatewayClient),
			service.EventsPublisher(&eventsPublisher),
			service.PermissionService(&permissionService),
		)
	})

	Describe("NewService", func() {
		It("returns a service", func() {
			Expect(svc).ToNot(BeNil())
		})
	})

	Describe("List drives", func() {
		It("can list an empty list of spaces", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status:        status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{},
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)
			Expect(rr.Code).To(Equal(http.StatusOK))
		})
		It("can list an empty list of all spaces", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Times(1).Return(&provider.ListStorageSpacesResponse{
				Status:        status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{},
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/drives", nil)
			rr := httptest.NewRecorder()
			svc.GetAllDrives(rr, r)
			Expect(rr.Code).To(Equal(http.StatusOK))
		})

		It("can list a space without owner", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Times(1).Return(&provider.ListStorageSpacesResponse{
				Status: status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{
					{
						Id:        &provider.StorageSpaceId{OpaqueId: "sameID"},
						SpaceType: "aspacetype",
						Root: &provider.ResourceId{
							StorageId: "pro-1",
							SpaceId:   "sameID",
							OpaqueId:  "sameID",
						},
						Name: "aspacename",
					},
				},
			}, nil)
			gatewayClient.On("InitiateFileDownload", mock.Anything, mock.Anything).Return(&gateway.InitiateFileDownloadResponse{
				Status: status.NewNotFound(ctx, "not found"),
			}, nil)
			gatewayClient.On("GetQuota", mock.Anything, mock.Anything).Return(&provider.GetQuotaResponse{
				Status: status.NewUnimplemented(ctx, fmt.Errorf("not supported"), "not supported"),
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))

			body, _ := io.ReadAll(rr.Body)
			Expect(body).To(MatchJSON(`
			{
				"value":[
					{
						"driveType":"aspacetype",
						"id":"pro-1$sameID",
						"name":"aspacename",
						"root":{
							"id":"pro-1$sameID",
							"webDavUrl":"https://localhost:9200/dav/spaces/pro-1$sameID"
						},
						"webUrl": "https://localhost:9200/f/pro-1$sameID"
					}
				]
			}
			`))
		})
		It("can list a spaces with sort", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status: status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{
					{
						Id:        &provider.StorageSpaceId{OpaqueId: "bsameID"},
						SpaceType: "bspacetype",
						Root: &provider.ResourceId{
							StorageId: "pro-1",
							SpaceId:   "bsameID",
							OpaqueId:  "bsameID",
						},
						Name: "bspacename",
						Opaque: &typesv1beta1.Opaque{
							Map: map[string]*typesv1beta1.OpaqueEntry{
								"spaceAlias": {Decoder: "plain", Value: []byte("bspacetype/bspacename")},
								"etag":       {Decoder: "plain", Value: []byte("123456789")},
							},
						},
					},
					{
						Id:        &provider.StorageSpaceId{OpaqueId: "asameID"},
						SpaceType: "aspacetype",
						Root: &provider.ResourceId{
							StorageId: "pro-1",
							SpaceId:   "asameID",
							OpaqueId:  "asameID",
						},
						Name: "aspacename",
						Opaque: &typesv1beta1.Opaque{
							Map: map[string]*typesv1beta1.OpaqueEntry{
								"spaceAlias": {Decoder: "plain", Value: []byte("aspacetype/aspacename")},
								"etag":       {Decoder: "plain", Value: []byte("101112131415")},
							},
						},
					},
				},
			}, nil)
			gatewayClient.On("InitiateFileDownload", mock.Anything, mock.Anything).Return(&gateway.InitiateFileDownloadResponse{
				Status: status.NewNotFound(ctx, "not found"),
			}, nil)
			gatewayClient.On("GetQuota", mock.Anything, mock.Anything).Return(&provider.GetQuotaResponse{
				Status: status.NewUnimplemented(ctx, fmt.Errorf("not supported"), "not supported"),
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives?$orderby=name%20asc", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))

			body, _ := io.ReadAll(rr.Body)
			Expect(body).To(MatchJSON(`
			{
				"value":[
					{
						"driveAlias":"aspacetype/aspacename",
						"driveType":"aspacetype",
						"id":"pro-1$asameID",
						"name":"aspacename",
						"root":{
							"eTag":"101112131415",
							"id":"pro-1$asameID",
							"webDavUrl":"https://localhost:9200/dav/spaces/pro-1$asameID"
						},
						"webUrl": "https://localhost:9200/f/pro-1$asameID"
					},
					{
						"driveAlias":"bspacetype/bspacename",
						"driveType":"bspacetype",
						"id":"pro-1$bsameID",
						"name":"bspacename",
						"root":{
							"eTag":"123456789",
							"id":"pro-1$bsameID",
							"webDavUrl":"https://localhost:9200/dav/spaces/pro-1$bsameID"
						},
						"webUrl": "https://localhost:9200/f/pro-1$bsameID"
					}
				]
			}
			`))
		})
		It("can list a spaces type mountpoint", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status: status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{
					{
						Id:        &provider.StorageSpaceId{OpaqueId: "prID$aID!differentID"},
						SpaceType: "mountpoint",
						Root: &provider.ResourceId{
							StorageId: "prID",
							SpaceId:   "aID",
							OpaqueId:  "differentID",
						},
						Name: "New Folder",
						Opaque: &typesv1beta1.Opaque{
							Map: map[string]*typesv1beta1.OpaqueEntry{
								"spaceAlias":     {Decoder: "plain", Value: []byte("mountpoint/new-folder")},
								"etag":           {Decoder: "plain", Value: []byte("101112131415")},
								"grantStorageID": {Decoder: "plain", Value: []byte("ownerStorageID")},
								"grantSpaceID":   {Decoder: "plain", Value: []byte("spaceID")},
								"grantOpaqueID":  {Decoder: "plain", Value: []byte("opaqueID")},
							},
						},
					},
				},
			}, nil)
			gatewayClient.On("Stat", mock.Anything, mock.Anything).Return(&provider.StatResponse{
				Status: status.NewOK(ctx),
				Info: &provider.ResourceInfo{
					Etag:  "123456789",
					Type:  provider.ResourceType_RESOURCE_TYPE_CONTAINER,
					Id:    &provider.ResourceId{StorageId: "ownerStorageID", SpaceId: "spaceID", OpaqueId: "opaqueID"},
					Path:  "New Folder",
					Mtime: &typesv1beta1.Timestamp{Seconds: 1648327606, Nanos: 0},
					Size:  uint64(1234),
				},
			}, nil)
			gatewayClient.On("GetQuota", mock.Anything, mock.Anything).Return(&provider.GetQuotaResponse{
				Status: status.NewUnimplemented(ctx, fmt.Errorf("not supported"), "not supported"),
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))

			body, _ := io.ReadAll(rr.Body)

			var response map[string][]libregraph.Drive
			err := json.Unmarshal(body, &response)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(response["value"])).To(Equal(1))
			value := response["value"][0]
			Expect(*value.DriveAlias).To(Equal("mountpoint/new-folder"))
			Expect(*value.DriveType).To(Equal("mountpoint"))
			Expect(*value.Id).To(Equal("prID$aID!differentID"))
			Expect(*value.Name).To(Equal("New Folder"))
			Expect(*value.Root.WebDavUrl).To(Equal("https://localhost:9200/dav/spaces/prID$aID%21differentID"))
			Expect(*value.Root.ETag).To(Equal("101112131415"))
			Expect(*value.Root.Id).To(Equal("prID$aID!differentID"))
			Expect(*value.Root.RemoteItem.ETag).To(Equal("123456789"))
			Expect(*value.Root.RemoteItem.Id).To(Equal("ownerStorageID$spaceID!opaqueID"))
			Expect(value.Root.RemoteItem.LastModifiedDateTime.UTC()).To(Equal(time.Unix(1648327606, 0).UTC()))
			Expect(*value.Root.RemoteItem.Folder).To(Equal(libregraph.Folder{}))
			Expect(*value.Root.RemoteItem.Name).To(Equal("New Folder"))
			Expect(*value.Root.RemoteItem.Size).To(Equal(int64(1234)))
			Expect(*value.Root.RemoteItem.WebDavUrl).To(Equal("https://localhost:9200/dav/spaces/ownerStorageID$spaceID%21opaqueID"))
		})
		It("can not list spaces with wrong sort parameter", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status:        status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{}}, nil)
			gatewayClient.On("InitiateFileDownload", mock.Anything, mock.Anything).Return(&gateway.InitiateFileDownloadResponse{
				Status: status.NewNotFound(ctx, "not found"),
			}, nil)
			gatewayClient.On("GetQuota", mock.Anything, mock.Anything).Return(&provider.GetQuotaResponse{
				Status: status.NewUnimplemented(ctx, fmt.Errorf("not supported"), "not supported"),
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives?$orderby=owner%20asc", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)
			Expect(rr.Code).To(Equal(http.StatusBadRequest))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("we do not support <owner> as a order parameter"))
			Expect(libreError.Error.Code).To(Equal(errorcode.InvalidRequest.String()))
		})
		It("can list a spaces with invalid query parameter", func() {
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives?§orderby=owner%20asc", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)
			Expect(rr.Code).To(Equal(http.StatusBadRequest))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("Query parameter '§orderby' is not supported. Cause: Query parameter '§orderby' is not supported"))
			Expect(libreError.Error.Code).To(Equal(errorcode.InvalidRequest.String()))
		})
		It("can list a spaces with an unsupported operand", func() {
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives?$filter=contains(driveType,personal)", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)
			Expect(rr.Code).To(Equal(http.StatusNotImplemented))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("unsupported filter operand: contains"))
			Expect(libreError.Error.Code).To(Equal(errorcode.NotSupported.String()))
		})
		It("transport error", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(nil, errors.New("transport error"))

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives)", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)
			Expect(rr.Code).To(Equal(http.StatusInternalServerError))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("transport error"))
			Expect(libreError.Error.Code).To(Equal(errorcode.GeneralException.String()))
		})
		It("grpc error", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status:        status.NewInternal(ctx, "internal error"),
				StorageSpaces: []*provider.StorageSpace{}}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives)", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)
			Expect(rr.Code).To(Equal(http.StatusInternalServerError))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("internal error"))
			Expect(libreError.Error.Code).To(Equal(errorcode.GeneralException.String()))
		})
		It("grpc not found", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status:        status.NewNotFound(ctx, "no spaces found"),
				StorageSpaces: []*provider.StorageSpace{}}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives)", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)
			Expect(rr.Code).To(Equal(http.StatusOK))

			body, _ := io.ReadAll(rr.Body)

			var response map[string][]libregraph.Drive
			err := json.Unmarshal(body, &response)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(response)).To(Equal(0))
		})
		It("quota error", func() {
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status: status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{
					{
						Id:        &provider.StorageSpaceId{OpaqueId: "sameID"},
						SpaceType: "aspacetype",
						Root: &provider.ResourceId{
							StorageId: "pro-1",
							SpaceId:   "sameID",
							OpaqueId:  "sameID",
						},
						Name: "aspacename",
					},
				},
			}, nil)
			gatewayClient.On("InitiateFileDownload", mock.Anything, mock.Anything).Return(&gateway.InitiateFileDownloadResponse{
				Status: status.NewNotFound(ctx, "not found"),
			}, nil)
			gatewayClient.On("GetQuota", mock.Anything, mock.Anything).Return(&provider.GetQuotaResponse{
				Status: status.NewInternal(ctx, "internal quota error"),
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me/drives", nil)
			rr := httptest.NewRecorder()
			svc.GetDrives(rr, r)

			Expect(rr.Code).To(Equal(http.StatusInternalServerError))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("internal quota error"))
			Expect(libreError.Error.Code).To(Equal(errorcode.GeneralException.String()))
		})
	})
	Describe("Create Drive", func() {
		It("cannot create a space without valid user in context", func() {
			jsonBody := []byte(`{"Name": "Test Space"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody))
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusUnauthorized))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("invalid user"))
			Expect(libreError.Error.Code).To(Equal(errorcode.NotAllowed.String()))
		})
		It("cannot create a space without permission", func() {
			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_UNKNOWN,
					Constraint: v0.Permission_CONSTRAINT_OWN,
				},
			}, nil)
			jsonBody := []byte(`{"Name": "Test Space"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusUnauthorized))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("insufficient permissions to create a space."))
			Expect(libreError.Error.Code).To(Equal(errorcode.NotAllowed.String()))
		})
		It("cannot create a space with wrong request body", func() {
			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			jsonBody := []byte(`{"Name": "Test Space"`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusBadRequest))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("invalid body schema definition"))
			Expect(libreError.Error.Code).To(Equal(errorcode.InvalidRequest.String()))
		})
		It("transport error", func() {
			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			gatewayClient.On("CreateStorageSpace", mock.Anything, mock.Anything).Return(&provider.CreateStorageSpaceResponse{}, errors.New("transport error"))
			jsonBody := []byte(`{"Name": "Test Space"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusInternalServerError))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("transport error"))
			Expect(libreError.Error.Code).To(Equal(errorcode.GeneralException.String()))
		})
		It("grpc permission denied error", func() {
			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			gatewayClient.On("CreateStorageSpace", mock.Anything, mock.Anything).Return(&provider.CreateStorageSpaceResponse{
				Status: status.NewPermissionDenied(ctx, nil, "grpc permission denied"),
			}, nil)

			jsonBody := []byte(`{"Name": "Test Space"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusForbidden))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("permission denied"))
			Expect(libreError.Error.Code).To(Equal(errorcode.NotAllowed.String()))
		})
		It("grpc general error", func() {
			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			gatewayClient.On("CreateStorageSpace", mock.Anything, mock.Anything).Return(&provider.CreateStorageSpaceResponse{
				Status: status.NewInternal(ctx, "grpc error"),
			}, nil)

			jsonBody := []byte(`{"Name": "Test Space"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusInternalServerError))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("grpc error"))
			Expect(libreError.Error.Code).To(Equal(errorcode.GeneralException.String()))
		})
		It("cannot create a space with empty name", func() {
			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			jsonBody := []byte(`{"Name": ""}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusBadRequest))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("invalid spacename: spacename must not be empty"))
			Expect(libreError.Error.Code).To(Equal(errorcode.InvalidRequest.String()))
		})
		It("cannot create a space with a name that exceeds 255 chars", func() {
			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			jsonBody := []byte(`{"Name": "uufZ2MEUjUMJa84RkPsjJ1zf4XXRTdVMxRsJGfevwHuUBojo5JEdNU22O1FGgzXXTi9tl5ZKWaluIef8pPmEAxn9lHGIjyDVYeRQPiX5PCAZ7rVszrpLJryY5x1p6fFGQ6WQsPpNaqnKnfMliJDsbkAwMf7rCpzo0GUuadgHY9s2mfoXHDnpxqEmDsheucqVAFcNlFZNbNHoZAebHfv78KYc8C0WnhWvqvSPGBkNPQbZUkFCOAIlqpQ2Q3MubgI2"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusBadRequest))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("invalid spacename: spacename must be smaller than 255"))
			Expect(libreError.Error.Code).To(Equal(errorcode.InvalidRequest.String()))
		})
		It("cannot create a space with a wrong type", func() {
			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			jsonBody := []byte(`{"Name": "Test", "DriveType": "media"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusBadRequest))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("drives of this type cannot be created via this api"))
			Expect(libreError.Error.Code).To(Equal(errorcode.InvalidRequest.String()))
		})
		It("cannot create a space with a name that contains invalid chars", func() {
			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			jsonBody := []byte(`{"Name": "Space / Name"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusBadRequest))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("invalid spacename: spacenames must not contain [/ \\ . : ? * \" > < |]"))
			Expect(libreError.Error.Code).To(Equal(errorcode.InvalidRequest.String()))
		})
		It("can create a project space", func() {
			gatewayClient.On("CreateStorageSpace", mock.Anything, mock.Anything).Return(&provider.CreateStorageSpaceResponse{
				Status: status.NewOK(ctx),
				StorageSpace: &provider.StorageSpace{
					Id:        &provider.StorageSpaceId{OpaqueId: "newID"},
					Name:      "Test Space",
					SpaceType: "project",
					Root: &provider.ResourceId{
						StorageId: "pro-1",
						SpaceId:   "newID",
						OpaqueId:  "newID",
					},
					Opaque: &typesv1beta1.Opaque{
						Map: map[string]*typesv1beta1.OpaqueEntry{
							"description": {Decoder: "plain", Value: []byte("This space is for testing")},
							"spaceAlias":  {Decoder: "plain", Value: []byte("project/testspace")},
						},
					},
				},
			}, nil)

			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			jsonBody := []byte(`{"Name": "Test Space", "DriveType": "project", "Description": "This space is for testing", "DriveAlias": "project/testspace"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusCreated))

			body, _ := io.ReadAll(rr.Body)
			var response libregraph.Drive
			err := json.Unmarshal(body, &response)
			Expect(err).ToNot(HaveOccurred())
			Expect(*response.Name).To(Equal("Test Space"))
			Expect(*response.DriveType).To(Equal("project"))
			Expect(*response.DriveAlias).To(Equal("project/testspace"))
			Expect(*response.Description).To(Equal("This space is for testing"))
		})
		It("Incomplete space", func() {
			gatewayClient.On("CreateStorageSpace", mock.Anything, mock.Anything).Return(&provider.CreateStorageSpaceResponse{
				Status: status.NewOK(ctx),
				StorageSpace: &provider.StorageSpace{
					Id:        &provider.StorageSpaceId{OpaqueId: "newID"},
					Name:      "Test Space",
					SpaceType: "project",
				},
			}, nil)

			permissionService.On("GetPermissionByID", mock.Anything, mock.Anything).Return(&settingssvc.GetPermissionByIDResponse{
				Permission: &v0.Permission{
					Operation:  v0.Permission_OPERATION_READWRITE,
					Constraint: v0.Permission_CONSTRAINT_ALL,
				},
			}, nil)
			jsonBody := []byte(`{"Name": "Test Space"}`)
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/drives", bytes.NewBuffer(jsonBody)).WithContext(ctx)
			rr := httptest.NewRecorder()
			svc.CreateDrive(rr, r)
			Expect(rr.Code).To(Equal(http.StatusInternalServerError))

			body, _ := io.ReadAll(rr.Body)
			var libreError libregraph.OdataError
			err := json.Unmarshal(body, &libreError)
			Expect(err).To(Not(HaveOccurred()))
			Expect(libreError.Error.Message).To(Equal("space has no root"))
			Expect(libreError.Error.Code).To(Equal(errorcode.GeneralException.String()))
		})
	})
})

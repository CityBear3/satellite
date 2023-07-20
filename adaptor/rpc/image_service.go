package rpc

import (
	"github.com/CityBear3/satellite/adaptor/repository/mysql"
	"github.com/CityBear3/satellite/adaptor/rpc/convertors"
	"github.com/CityBear3/satellite/adaptor/rpc/validations"
	"github.com/CityBear3/satellite/logic/dto"
	archiveLogic "github.com/CityBear3/satellite/logic/interactor/archive"
	"github.com/CityBear3/satellite/pb/image/v1"
	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net/http"
)

type ImageService struct {
	uploadArchiveInteractor *archiveLogic.UploadArchiveInteractor
	image.UnimplementedImageServiceServer
}

func NewImageService() *ImageService {
	return &ImageService{
		uploadArchiveInteractor: archiveLogic.NewArchiveImageInteractor(mysql.NewArchiveRepository(), mysql.NewTxManger()),
	}
}

func (s ImageService) UploadImage(server image.ImageService_UploadImageServer) error {
	ctx := server.Context()
	var meta *image.Meta
	var data []byte
	for {
		request, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return convertors.ConvertError(err)
		}

		if m := request.GetMeta(); m != nil {
			meta = m
		}
		if c := request.GetChunk(); c != nil {
			data = append(data, c...)
		}
	}

	contentType := http.DetectContentType(data)
	if err := validations.ValidateUploadImageStream(meta, data, contentType); err != nil {
		return convertors.ConvertError(err)
	}

	id, err := ulid.Parse(meta.Id)
	if err != nil {
		return convertors.ConvertError(err)
	}

	if err := s.uploadArchiveInteractor.Handle(ctx, dto.NewUploadArchiveRequest(id, contentType, data)); err != nil {
		return convertors.ConvertError(err)
	}

	if err := server.SendAndClose(&emptypb.Empty{}); err != nil {
		return convertors.ConvertError(err)
	}

	return nil
}

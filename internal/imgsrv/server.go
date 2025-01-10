package imgsrv

import (
    "context"
)

type ImageProcessingServiceInterface interface {
    Validate(ctx context.Context, img *Image) error
    Process(ctx context.Context, img *Image) (*ImgProcessResponse, error)
}

package imgsrv

import (
    "log"
    "context"
    "time"
)

type ImgProcessResponse struct {
    Result string
}

type Image struct {
    ID           string 
    ResourcePath string
}

// ImageProcessingService implements ImageProcessingServiceInterface
type ImageProcessingService struct {}

func NewImageProcessingService() ImageProcessingServiceInterface {
    return NewImageProcessingServiceProxy(56) // Proxy is now returned as ImageProcessingServiceInterface
}

func (s *ImageProcessingService) Validate(ctx context.Context, img *Image) error {
    log.Print("Validating image")
    select {
    case <-time.After(10 * time.Second):
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func (s *ImageProcessingService) Process(ctx context.Context, img *Image) (*ImgProcessResponse, error) {
    log.Print("Processing image")
    select {
    case <-time.After(10 * time.Second):
        return &ImgProcessResponse{Result: "Success"}, nil
    case <-ctx.Done():
        return nil, ctx.Err()
    }
}

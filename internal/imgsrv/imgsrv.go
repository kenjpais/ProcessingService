package imgsrv

import (
    "log"
    "time"
    "context"
)

type ImgProcessResponse struct {
    Result string `json:"result"`
}

type Image struct {
    ID           string `json:"id"`
    ResourcePath string `json:"resource_path"`
}

type ImageProcessingService struct {}

func NewImageProcessingService() *ImageProcessingService {
    return &ImageProcessingService{}
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

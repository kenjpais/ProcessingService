package imgsrv

import (
    "context"
    "log"
    "fmt"
)

type ImageProcessingServiceProxy struct {
    mainInstance      *ImageProcessingService
    maxAllowedRequest uint32
    rateLimiter       map[string]uint32
}

// NewImageProcessingServiceProxy creates a new proxy with rate limiting.
func NewImageProcessingServiceProxy(maxAllowedRequest uint32) *ImageProcessingServiceProxy {
    return &ImageProcessingServiceProxy{
        mainInstance:      &ImageProcessingService{},
        maxAllowedRequest: maxAllowedRequest,
        rateLimiter:       make(map[string]uint32),
    }
}

// Decorator function to apply rate-limiting check before executing the actual function
func (s *ImageProcessingServiceProxy) rateLimitDecorator(id string, f func()) bool {
    log.Print("Checking rate limit for id:", id)
    if limit, exists := s.rateLimiter[id]; exists && s.maxAllowedRequest <= limit {
        log.Print("Rate limit exceeded for id:", id)
        return false // Rate limit exceeded
    }
    f() // Execute the function
    return true
}

// Implement the Validate method (required by ImageProcessingServiceInterface)
func (s *ImageProcessingServiceProxy) Validate(ctx context.Context, img *Image) error {
    allowed := s.rateLimitDecorator("validate", func() {
        s.mainInstance.Validate(ctx, img)
    })
    if !allowed {
        return fmt.Errorf("rate limit exceeded for Validate")
    }
    return nil
}

// Implement the Process method (required by ImageProcessingServiceInterface)
func (s *ImageProcessingServiceProxy) Process(ctx context.Context, img *Image) (*ImgProcessResponse, error) {
    var result *ImgProcessResponse
    allowed := s.rateLimitDecorator("process", func() {
        result, _ = s.mainInstance.Process(ctx, img)
    })
    if !allowed {
        return nil, fmt.Errorf("rate limit exceeded for Process")
    }
    return result, nil
}

// Increment the request count for a specific id
func (s *ImageProcessingServiceProxy) IncrementRequestCount(id string) {
    if _, exists := s.rateLimiter[id]; !exists {
        s.rateLimiter[id] = 0
    }
    s.rateLimiter[id]++
}

// Reset the request count for a specific id
func (s *ImageProcessingServiceProxy) ResetRequestCount(id string) {
    delete(s.rateLimiter, id)
}

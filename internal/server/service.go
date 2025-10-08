package server

import (
	"context"
	"log"
	"sync"

	pb "test_task/pb/storage"
)

// StorageService реализует pb.StorageServiceServer
type StorageService struct {
	pb.UnimplementedStorageServiceServer
	data map[string]string
	mu   sync.RWMutex
}

// NewStorageService создаёт новый экземпляр сервиса
func NewStorageService() *StorageService {
	return &StorageService{
		data: make(map[string]string),
	}
}

// Set сохраняет данные по ключу
func (s *StorageService) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[req.Key] = req.Value
	log.Printf("[SERVER] Set key=%s, value=%s\n", req.Key, req.Value)
	return &pb.SetResponse{Success: true, Message: "Stored successfully"}, nil
}

// Get возвращает данные по ключу
func (s *StorageService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[req.Key]
	log.Printf("[SERVER] Get key=%s -> %v\n", req.Key, value)
	if !ok {
		return &pb.GetResponse{Key: req.Key, Found: false}, nil
	}
	return &pb.GetResponse{Key: req.Key, Value: value, Found: true}, nil
}

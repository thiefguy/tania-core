package inmemory

import (
	"github.com/Tanibox/tania-server/src/assets/query"
	"github.com/Tanibox/tania-server/src/assets/storage"
	uuid "github.com/satori/go.uuid"
)

type AreaReadQueryInMemory struct {
	Storage *storage.AreaReadStorage
}

func NewAreaReadQueryInMemory(s *storage.AreaReadStorage) query.AreaReadQuery {
	return AreaReadQueryInMemory{Storage: s}
}

func (s AreaReadQueryInMemory) FindByID(uid uuid.UUID) <-chan query.QueryResult {
	result := make(chan query.QueryResult)

	go func() {
		s.Storage.Lock.RLock()
		defer s.Storage.Lock.RUnlock()

		area := storage.AreaRead{}
		for _, val := range s.Storage.AreaReadMap {
			if val.UID == uid {
				area = val
			}
		}

		result <- query.QueryResult{Result: area}

		close(result)
	}()

	return result
}

func (s AreaReadQueryInMemory) FindAllByFarm(farmUID uuid.UUID) <-chan query.QueryResult {
	result := make(chan query.QueryResult)

	go func() {
		s.Storage.Lock.RLock()
		defer s.Storage.Lock.RUnlock()

		areas := []storage.AreaRead{}
		for _, val := range s.Storage.AreaReadMap {
			if val.Farm.UID == farmUID {
				areas = append(areas, val)
			}
		}

		result <- query.QueryResult{Result: areas}

		close(result)
	}()

	return result
}
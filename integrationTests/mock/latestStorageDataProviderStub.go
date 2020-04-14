package mock

import "github.com/ElrondNetwork/elrond-go/storage/factory"

// LatestStorageDataProviderStub -
type LatestStorageDataProviderStub struct {
	GetCalled                      func() (factory.LatestDataFromStorage, error)
	GetParentDirAndLastEpochCalled func() (string, uint32, error)
}

// GetParentDirAndLastEpoch -
func (l *LatestStorageDataProviderStub) GetParentDirAndLastEpoch() (string, uint32, error) {
	if l.GetParentDirAndLastEpochCalled != nil {
		return l.GetParentDirAndLastEpochCalled()
	}

	return "", 0, nil
}

// Get -
func (l *LatestStorageDataProviderStub) Get() (factory.LatestDataFromStorage, error) {
	if l.GetCalled != nil {
		return l.GetCalled()
	}

	return factory.LatestDataFromStorage{}, nil
}
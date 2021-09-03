package main

import "context"

type M map[string]interface{}

type Store interface {
	GetCountries(ctx context.Context) ([]M, error)
	GetStates(ctx context.Context, country string) ([]M, error)
	GetCities(ctx context.Context, country, state string) ([]M, error)
}

type mockStore struct {
}

func (m *mockStore) GetCountries(ctx context.Context) ([]M, error) {
	return []M{
		{"country": "Andorra"},
		{"country": "Argentina"},
		{"country": "Australia"},
		{"country": "Austria"},
		{"country": "Bahamas"},
		{"country": "Bahrain"},
		{"country": "Bangladesh"},
		{"country": "Belgium"},
		{"country": "Bosnia Herzegovina"},
		{"country": "Brazil"},
		{"country": "Cambodia"},
		{"country": "Canada"},
		{"country": "Chile"},
		{"country": "China"},
		{"country": "Colombia"},
		{"country": "Croatia"},
		{"country": "Cyprus"},
		{"country": "Czech Republic"},
		{"country": "Denmark"},
		{"country": "Ethiopia"},
		{"country": "Finland"},
		{"country": "France"},
		{"country": "Germany"},
		{"country": "Hungary"},
		{"country": "India"},
		{"country": "Indonesia"},
		{"country": "Iran"},
		{"country": "Ireland"},
		{"country": "Israel"},
		{"country": "Italy"},
		{"country": "Vietnam"},
	}, nil
}

func (m *mockStore) GetStates(ctx context.Context, country string) ([]M, error) {
	return []M{
		{"state": "Anhui"},
		{"state": "Beijing"},
		{"state": "Chongqing"},
		{"state": "Fujian"},
		{"state": "Gansu"},
		{"state": "Guangdong"},
		{"state": "Guangxi"},
		{"state": "Guizhou"},
		{"state": "Hainan"},
		{"state": "Hebei"},
		{"state": "Heilongjiang"},
	}, nil
}

func (m *mockStore) GetCities(ctx context.Context, country, state string) ([]M, error) {
	return []M{
		{"city": "Addison"},
		{"city": "Albany"},
		{"city": "Buffalo"},
		{"city": "Carmel"},
		{"city": "Dunkirk"},
		{"city": "East Syracuse"},
		{"city": "Farmingdale"},
		{"city": "Holtsville"},
		{"city": "Ithaca"},
	}, nil
}

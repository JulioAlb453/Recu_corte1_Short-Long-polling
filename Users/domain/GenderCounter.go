package domain

type GenderCounter struct{
	Counts map[string] int `json: counts`
}
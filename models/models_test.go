package models

import "testing"

func TestProductStruct(t *testing.T) {
	p := Product{
		ID:         1,
		Name:       "iPhone",
		Price:      500000,
		CategoryID: 1,
	}

	if p.Name != "iPhone" {
		t.Errorf("ожидали iPhone, получили %s", p.Name)
	}

	if p.Price != 500000 {
		t.Errorf("ожидали 500000, получили %d", p.Price)
	}
}

func TestCategoryStruct(t *testing.T) {
	c := Category{
		ID:   1,
		Name: "Телефоны",
	}

	if c.Name != "Телефоны" {
		t.Errorf("ожидали Телефоны, получили %s", c.Name)
	}
}
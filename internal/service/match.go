package service

import schema "github.com/denislyubo/matchmaker"

type match struct {
	groupSize uint
}

var _ schema.MatchService = (*match)(nil)

func New(groupSize uint) *match {
	return &match{groupSize: groupSize}
}

func (m *match) AddUser()  {}
func (m *match) GetMatch() {}

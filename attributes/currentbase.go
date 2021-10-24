package attributes

import (
	"github.com/vitalick/go-d2editor/utils"
	"io"
)

type CurrentBase struct {
	Current utils.FloatD2sGo `json:"current"`
	Base    utils.FloatD2sGo `json:"base"`
}

func NewCurrentBase(r io.Reader, bs []bool, i *int) (*CurrentBase, error) {
	nowIter := *i
	*i += 2
	var err error
	c := utils.FloatD2sGo(0)
	b := utils.FloatD2sGo(0)
	if bs[nowIter] || true {
		c, err = utils.NewFloatD2sGo(r)
		if err != nil {
			return nil, err
		}
	}
	if bs[nowIter+1] || true {
		b, err = utils.NewFloatD2sGo(r)
		if err != nil {
			return nil, err
		}
	}
	return &CurrentBase{
		Current: c,
		Base:    b,
	}, nil
}

func (cb *CurrentBase) GetPacked() ([]byte, error) {
	var outB []byte
	cp, err := cb.Current.GetPacked()
	if err != nil {
		return nil, err
	}
	outB = append(outB, cp[:]...)
	bp, err := cb.Base.GetPacked()
	if err != nil {
		return nil, err
	}
	outB = append(outB, bp[:]...)
	return outB, nil
}

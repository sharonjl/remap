package remap

import (
	"testing"
	"time"
)

type StudentDetail struct {
	Id          int
	Name        string
	Address     string
	Grade       int
	DOB         *time.Time
	ModifiedOn  *time.Time
	CreatedTime time.Time
}

type Student struct {
	Name        string
	Grade       int
	DOB         *time.Time
	ModifiedOn  time.Time
	CreatedTime *time.Time
}

const twelveYears = time.Hour * 24 * 365 * 12

var studentADob = time.Now()
var studentACreated = studentADob.Add(twelveYears)
var studentAModified = studentADob.Add(twelveYears).Add(time.Hour * 24)

var studentA = &StudentDetail{
	Id:          1,
	Name:        "My string",
	Address:     "12 My Street, My State, Country",
	Grade:       int(6),
	DOB:         &studentADob,
	CreatedTime: studentACreated,
	ModifiedOn:  &studentAModified,
}

func TestOnFields(t *testing.T) {
	b := &Student{}

	OnFields(studentA, b)
	if b.Name != studentA.Name ||
		b.Grade != studentA.Grade {
		t.Errorf("failed copying primitive values")
	}

	if b.DOB != studentA.DOB {
		t.Errorf("failed copying ptr to ptr")
	}

	if b.CreatedTime.Sub(studentA.CreatedTime) != 0 {
		t.Errorf("failed copying val to ptr")
	}

	if studentA.ModifiedOn.Sub(b.ModifiedOn) != 0 {
		t.Errorf("failed copying ptr to val")
	}
}

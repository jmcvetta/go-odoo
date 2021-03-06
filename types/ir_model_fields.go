package types

import (
	"time"
)

type IrModelFields struct {
	Column1          string      `xmlrpc:"column1"`
	Column2          string      `xmlrpc:"column2"`
	CompleteName     string      `xmlrpc:"complete_name"`
	Compute          string      `xmlrpc:"compute"`
	Copy             bool        `xmlrpc:"copy"`
	CreateDate       time.Time   `xmlrpc:"create_date"`
	CreateUid        Many2One    `xmlrpc:"create_uid"`
	Depends          string      `xmlrpc:"depends"`
	DisplayName      string      `xmlrpc:"display_name"`
	Domain           string      `xmlrpc:"domain"`
	FieldDescription string      `xmlrpc:"field_description"`
	Groups           []int64     `xmlrpc:"groups"`
	Help             string      `xmlrpc:"help"`
	Id               int64       `xmlrpc:"id"`
	Index            bool        `xmlrpc:"index"`
	LastUpdate       time.Time   `xmlrpc:"__last_update"`
	Model            string      `xmlrpc:"model"`
	ModelId          Many2One    `xmlrpc:"model_id"`
	Modules          string      `xmlrpc:"modules"`
	Name             string      `xmlrpc:"name"`
	OnDelete         interface{} `xmlrpc:"on_delete"`
	Readonly         bool        `xmlrpc:"readonly"`
	Related          string      `xmlrpc:"related"`
	Relation         string      `xmlrpc:"relation"`
	RelationField    string      `xmlrpc:"relation_field"`
	RelationTable    string      `xmlrpc:"relation_table"`
	Required         bool        `xmlrpc:"required"`
	Selectable       bool        `xmlrpc:"selectable"`
	Selection        string      `xmlrpc:"selection"`
	Size             int64       `xmlrpc:"size"`
	State            interface{} `xmlrpc:"state"`
	Store            bool        `xmlrpc:"store"`
	TrackVisibility  interface{} `xmlrpc:"track_visibility"`
	Translate        bool        `xmlrpc:"translate"`
	Ttype            interface{} `xmlrpc:"ttype"`
	WriteDate        time.Time   `xmlrpc:"write_date"`
	WriteUid         Many2One    `xmlrpc:"write_uid"`
}

type IrModelFieldsNil struct {
	Column1          interface{} `xmlrpc:"column1"`
	Column2          interface{} `xmlrpc:"column2"`
	CompleteName     interface{} `xmlrpc:"complete_name"`
	Compute          interface{} `xmlrpc:"compute"`
	Copy             bool        `xmlrpc:"copy"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	Depends          interface{} `xmlrpc:"depends"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	Domain           interface{} `xmlrpc:"domain"`
	FieldDescription interface{} `xmlrpc:"field_description"`
	Groups           interface{} `xmlrpc:"groups"`
	Help             interface{} `xmlrpc:"help"`
	Id               interface{} `xmlrpc:"id"`
	Index            bool        `xmlrpc:"index"`
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	Model            interface{} `xmlrpc:"model"`
	ModelId          interface{} `xmlrpc:"model_id"`
	Modules          interface{} `xmlrpc:"modules"`
	Name             interface{} `xmlrpc:"name"`
	OnDelete         interface{} `xmlrpc:"on_delete"`
	Readonly         bool        `xmlrpc:"readonly"`
	Related          interface{} `xmlrpc:"related"`
	Relation         interface{} `xmlrpc:"relation"`
	RelationField    interface{} `xmlrpc:"relation_field"`
	RelationTable    interface{} `xmlrpc:"relation_table"`
	Required         bool        `xmlrpc:"required"`
	Selectable       bool        `xmlrpc:"selectable"`
	Selection        interface{} `xmlrpc:"selection"`
	Size             interface{} `xmlrpc:"size"`
	State            interface{} `xmlrpc:"state"`
	Store            bool        `xmlrpc:"store"`
	TrackVisibility  interface{} `xmlrpc:"track_visibility"`
	Translate        bool        `xmlrpc:"translate"`
	Ttype            interface{} `xmlrpc:"ttype"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var IrModelFieldsModel string = "ir.model.fields"

type IrModelFieldss []IrModelFields

type IrModelFieldssNil []IrModelFieldsNil

func (s *IrModelFields) NilableType_() interface{} {
	return &IrModelFieldsNil{}
}

func (ns *IrModelFieldsNil) Type_() interface{} {
	s := &IrModelFields{}
	return load(ns, s)
}

func (s *IrModelFieldss) NilableType_() interface{} {
	return &IrModelFieldssNil{}
}

func (ns *IrModelFieldssNil) Type_() interface{} {
	s := &IrModelFieldss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModelFields))
	}
	return s
}

package xsnowflake

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/bwmarrin/snowflake"
)

type ID snowflake.ID

func (id ID) IsZero() bool {
	return snowflake.ID(id) == 0
}

func (id *ID) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	parsedID, err := Parse(str)
	if err != nil {
		return err
	}

	*id = parsedID
	return nil
}

func (id *ID) UnmarshalText(data []byte) error {
	parsedID, err := Parse(string(data))
	if err != nil {
		return err
	}

	*id = parsedID
	return nil
}

func (id ID) String() string {
	return snowflake.ID(id).String()
}

func (id *ID) Scan(src interface{}) error {
	switch s := src.(type) {
	case string:
		parsedID, err := Parse(s)
		if err != nil {
			return err
		}
		*id = parsedID
	case int64:
		*id = ID(s)
	default:
		return errors.New("Unsupported scan source")
	}
	return nil
}

func (id ID) Value() (driver.Value, error) {
	return int64(id), nil
}

type Node snowflake.Node

func (n *Node) Generate() ID {
	return ID((*snowflake.Node)(n).Generate())
}

func Parse(s string) (ID, error) {
	id, err := snowflake.ParseString(s)
	if err != nil {
		return ID(0), err
	}
	return ID(id), nil
}

func NewGenerator(nodeId int64, t time.Time) (*Node, error) {
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		return nil, err
	}

	snowflake.Epoch = t.UnixMilli()
	return (*Node)(node), nil
}

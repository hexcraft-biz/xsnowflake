package xsnowflake

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

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

func New(nodeId int64) ID {
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		panic(err)
	}
	return ID(node.Generate())
}

func Parse(s string) (ID, error) {
	id, err := snowflake.ParseString(s)
	if err != nil {
		return ID(0), err
	}
	return ID(id), nil
}

package db

import (
	"database/sql"
	"fmt"
	"strings"
)

type Kit struct {
	Id          int            `db:"id"`
	Uid         string         `db:"uid"`
	Name        string         `db:"name"`
	IsCustom    int            `db:"iscustom"`
	Description sql.NullString `db:"description"`
	Copyright   sql.NullString `db:"copyright"`
	Licence     sql.NullString `db:"licence"`
	Credits     sql.NullString `db:"credits"`
	Url         sql.NullString `db:"url"`
	Tags        sql.NullString `db:"tags"`
	TagList     []string
}

type KitTag struct {
	Id   int    `db:"id"`
	Kit  int    `db:"kit"`
	Name string `db:"name"`
}

// TODO: optional filter by
//   - like name
//   - isCustom
//   - in (tags)
func (d *Sqlite) ListKits() (*[]Kit, error) {
	kits := []Kit{}

	rows, err := d.db.Queryx(`select k.*, string_agg(t.name) as tags
	from kit k left join kit_tag t on t.kit = k.id
	group by k.id, k.uid, k.name, k.iscustom, k.description, k.copyright, k.licence, k.credits, k.url
	order by k.name, k.id
	`)
	if err != nil {
		return &kits, fmt.Errorf("failed sql: %w", err)
	}
	for rows.Next() {
		kit := Kit{}
		err := rows.StructScan(kit)
		if err != nil {
			return &kits, fmt.Errorf("failed sql: %w", err)
		}
		if kit.Tags.Valid {
			kit.TagList = strings.Split(kit.Tags.String, ",")
		}
		kits = append(kits, kit)
	}
	return &kits, nil
}

package example

import (
	"context"
	"database/sql"

	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	sqlCreateTable = `
	create table if not exists item (
		id serial primary key,
		Name text
	);
	`
	sqlGetOne = `
	select id, Name
	from item
	where id=$1
	;
	`
	sqlGetAll = `
	select id, Name
	from item
	;
	`
	sqlCreateOne = `
	insert into item
	(Name)
	values
	($1)
	returning id, Name
	;
	`
	sqlUpdateOne = `
	update item
	set 
		Name=$2
	where id=$1
	returning id, Name
	;
	`
	sqlDeleteOne = `
	delete from item
	where id=$1
	;
	`
)

type postgresStorage struct {
	UnimplementedItemStorageServer

	db *sql.DB
}

func NewPostgresItemStorage(db *sql.DB) ItemStorageServer {
	_, err := db.Exec(sqlCreateTable)
	if err != nil {
		panic(err)
	}
	return &postgresStorage{db: db}
}

func (s *postgresStorage) GetItem(ctx context.Context, req *GetItemRequest) (*GetItemResponse, error) {
	row := s.db.QueryRowContext(ctx, sqlGetOne, req.Id)
	item, err := ScanRowToItem(row)
	if err != nil {
		return nil, err
	}
	return &GetItemResponse{
		Item: item,
	}, nil
}

func (s *postgresStorage) GetAllItems(ctx context.Context, req *emptypb.Empty) (*GetAllItemsResponse, error) {
	rows, err := s.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return nil, err
	}

	items := make([]*Item, 0)
	for rows.Next() {
		i, err := ScanRowsToItem(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return &GetAllItemsResponse{
		Items: items,
	}, nil
}

func (s *postgresStorage) CreateItem(ctx context.Context, req *CreateItemRequest) (*CreateItemResponse, error) {
	resp := &CreateItemResponse{
		Item: &Item{},
	}
	row := s.db.QueryRowContext(ctx, sqlCreateOne, req.Item.Name)
	if err := row.Scan(&resp.Item.Id, &resp.Item.Name); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *postgresStorage) UpdateItem(ctx context.Context, req *UpdateItemRequest) (*UpdateItemResponse, error) {
	row := s.db.QueryRowContext(ctx, sqlUpdateOne, req.Id, req.Item.Name)
	item, err := ScanRowToItem(row)
	if err != nil {
		return nil, err
	}

	return &UpdateItemResponse{
		Item: item,
	}, nil
}

func (s *postgresStorage) DeleteItem(ctx context.Context, req *DeleteItemRequest) (*emptypb.Empty, error) {
	_, err := s.db.ExecContext(ctx, sqlDeleteOne, req.Id)
	return &emptypb.Empty{}, err
}

func ScanRowToItem(r *sql.Row) (*Item, error) {
	i := &Item{}
	if err := r.Scan(&i.Id, &i.Name); err != nil {
		return nil, err
	}
	return i, nil
}

func ScanRowsToItem(r *sql.Rows) (*Item, error) {
	i := &Item{}
	if err := r.Scan(&i.Id, &i.Name); err != nil {
		return nil, err
	}
	return i, nil
}
